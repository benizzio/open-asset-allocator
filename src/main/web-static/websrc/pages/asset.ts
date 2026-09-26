/**
 * Coordinates asset list, creation, and editing on one page.
 *
 * HTMX handles HTTP requests and each form declares its own operation in HTML. This controller
 * renders the create or edit template, checks API response identities, and navigates between views.
 *
 * @author OpenCode
 */
import htmx from "htmx.org";
import * as Handlebars from "handlebars";
import notifications from "../components/notifications";
import { NotificationType } from "../infra/infra-types";
import { AfterRequestEventDetail, BeforeSwapEventDetail } from "../infra/htmx";
import Router from "../infra/routing";
import type { Asset } from "../domain/asset";
import {
    addExternalAsset,
    handleExternalSearchAfterRequest,
    handleExternalSearchBeforeRequest,
    hasValidExternalData,
    initializeExternalDraft,
    moveExternalAsset,
    removeExternalAsset,
    searchExternalAssets,
} from "./asset-external";

/**
 * Carries an HTMX request completion for asset list and form handling.
 * @author OpenCode
 */
type AssetRequestEvent = CustomEvent<AfterRequestEventDetail>;
/**
 * Carries an HTMX swap decision for asset detail response validation.
 * @author OpenCode
 */
type AssetBeforeSwapEvent = CustomEvent<BeforeSwapEventDetail>;

const ASSETS_PATH = "/asset";
const NEW_ASSET_PATH = "/asset/new";
const ASSET_IDENTIFIER_PATH_PATTERN = /^\/asset\/([^/]+)\/?$/;
const ASSET_DETAIL_API_PATH_PATTERN = /^\/api\/asset\/[^/]+\/?$/;

/**
 * Limits mutation handling to the form's own asset API response, excluding nested search GETs.
 * @author OpenCode
 */
function isAssetFormResponse(event: AssetRequestEvent, method: "post" | "put"): boolean {
    return event.detail.requestConfig.elt === event.currentTarget
        && event.detail.requestConfig.verb === method
        && new URL(event.detail.xhr.responseURL).pathname === "/api/asset";
}

/**
 * Parses an unknown response value as JSON without propagating malformed-response errors.
 * @author OpenCode
 */
function parseJSON(value: unknown): unknown | null {

    if(typeof value !== "string") {
        return value ?? null;
    }

    try {
        return JSON.parse(value);
    } catch {
        return null;
    }
}

/**
 * Validates and normalizes an API response into the asset shape used by this page.
 * @author OpenCode
 */
function normalizeAsset(value: unknown): Asset | null {

    if(typeof value !== "object" || value === null) {
        return null;
    }

    const candidate = value as Record<string, unknown>;
    const numericId = typeof candidate.id === "number" ? candidate.id : Number(candidate.id);

    if(!Number.isSafeInteger(numericId) || typeof candidate.name !== "string" || typeof candidate.ticker !== "string") {
        return null;
    }

    const externalData = candidate.externalData;

    if(!hasValidExternalData(externalData)) {
        return null;
    }

    return {
        id: numericId,
        name: candidate.name,
        ticker: candidate.ticker,
        externalData: externalData as Asset["externalData"],
    };
}

/**
 * Extracts and decodes the asset identifier from the current detail route.
 * @author OpenCode
 */
function getAssetIdentifierFromLocation(): string | null {

    const match = globalThis.location.pathname.match(ASSET_IDENTIFIER_PATH_PATTERN);
    return match ? decodeURIComponent(match[1]) : null;
}

/**
 * Returns the active detail-route identifier only when the response belongs to that route.
 * @author OpenCode
 */
function getCurrentAssetIdentifierForResponse(responseURL: string): string | null {

    const currentIdentifier = getAssetIdentifierFromLocation();

    if(!currentIdentifier || currentIdentifier === "new") {
        return null;
    }

    const responsePath = new URL(responseURL, globalThis.location.origin).pathname;
    return responsePath === `/api/asset/${ encodeURIComponent(currentIdentifier) }` ? currentIdentifier : null;
}

/**
 * Renders the operation-specific form and binds HTMX to the inserted form.
 * @author OpenCode
 */
function renderAssetForm(
    templateId: "asset-create-form" | "asset-edit-form",
    contentElement: HTMLElement,
    asset?: Asset,
): void {

    const template = document.getElementById(templateId);

    if(!template) {
        return;
    }

    contentElement.innerHTML = Handlebars.compile(template.innerHTML)(asset ?? {});
    const form = contentElement.querySelector("form") as HTMLFormElement | null;

    if(form) {
        initializeExternalDraft(form, asset);
    }
    // The script-loaded HTMX instance owns the form-json extension used by these forms.
    globalThis.htmx.process(contentElement);
}

/**
 * Shows a success toast through the application notification component.
 * @author OpenCode
 */
function notifySuccess(title: string, content: string): void {
    notifications.notify({ title, content, type: NotificationType.SUCCESS });
}

/**
 * Reports an invalid successful API response through the notification component.
 * @author OpenCode
 */
function notifyUnexpectedResponse(message: string): void {
    notifications.notifyError(new Error(message));
}

/**
 * Replaces the collection target with a retryable load error.
 * @author OpenCode
 */
function renderListError(target: HTMLElement): void {

    target.innerHTML = `
        <div class="row justify-content-center my-5">
            <div class="col-12 col-lg-8">
                <div class="alert alert-danger" role="alert">
                    <h2 class="h5">Assets could not be loaded</h2>
                    <p>There was a problem loading the asset list.</p>
                    <button type="button" class="btn btn-outline-danger" onclick="assetPage.reloadAssets()">
                        Try again
                    </button>
                </div>
            </div>
        </div>
    `;
}

/**
 * Displays a recoverable detail-load error without destroying the edit route bindings.
 * @author OpenCode
 */
function renderAssetLoadError(): void {

    const contentElement = document.querySelector("#asset-edit-content") as HTMLElement | null;
    const errorElement = document.querySelector("#asset-edit-error") as HTMLElement | null;
    const loadingElement = document.querySelector("#asset-edit-loading") as HTMLElement | null;

    if(!contentElement || !errorElement) {
        return;
    }

    contentElement.innerHTML = "";

    if(loadingElement) {
        loadingElement.style.display = "none";
    }
    errorElement.style.display = null;

    errorElement.innerHTML = `
        <div class="alert alert-danger mb-0" role="alert">
            <h2 class="h5">Asset could not be loaded</h2>
            <p>The requested asset is not available for editing.</p>
            <button type="button"
                    class="btn btn-outline-danger"
                    onclick="assetPage.navigateToAssets()"
                    aria-label="Back to assets"
                    title="Back to assets"
            >
                <span class="bi bi-arrow-left h5 mb-0" aria-hidden="true"></span>
            </button>
        </div>
    `;
}

/**
 * Provides browser behavior for the asset list and operation-specific forms.
 *
 * Use the exported instance through `globalThis.assetPage` from HTMX attributes.
 *
 * @example
 * ```ts
 * assetPage.renderCreateForm(view);
 * ```
 *
 * @author OpenCode
 */
const AssetPage = {

    searchExternalAssets,
    handleExternalSearchBeforeRequest,
    handleExternalSearchAfterRequest,
    addExternalAsset,
    removeExternalAsset,
    moveExternalAsset,

    /**
     * Runs provider search on Enter without submitting the parent asset form.
     *
     * @example
     * ```ts
     * assetPage.handleExternalSearchKeydown(event);
     * ```
     * @author OpenCode
     */
    handleExternalSearchKeydown(event: KeyboardEvent): void {
        if(event.key === "Enter") {
            event.preventDefault();
            const form = (event.target as HTMLElement).closest("form");

            if(form) {
                searchExternalAssets(form);
            }
        }
    },

    /**
     * Navigates to the asset creation form.
     *
     * @example
     * ```html
     * <button onclick="assetPage.navigateToNewAsset()">New asset</button>
     * ```
     *
     * @author OpenCode
     */
    navigateToNewAsset(): void {
        Router.navigateTo(NEW_ASSET_PATH);
    },

    /**
     * Navigates to one asset and loads it into the edit form.
     *
     * @example
     * ```ts
     * assetPage.navigateToAsset("42");
     * ```
     *
     * @author OpenCode
     */
    navigateToAsset(assetId: string): void {
        Router.navigateTo(`/asset/${ encodeURIComponent(assetId) }`);
    },

    /**
     * Activates asset-row navigation from Enter or Space without duplicating pointer navigation.
     *
     * @example
     * ```ts
     * assetPage.handleAssetNavigationKeypress(event, "42");
     * ```
     *
     * @author OpenCode
     */
    handleAssetNavigationKeypress(event: KeyboardEvent, assetId: string): void {
        if(event.key === "Enter" || event.key === " ") {
            event.preventDefault();
            this.navigateToAsset(assetId);
        }
    },

    /**
     * Navigates to the asset table, whose route binding reloads its data.
     *
     * @example
     * ```html
     * <button type="button" onclick="assetPage.navigateToAssets()">Back</button>
     * ```
     *
     * @author OpenCode
     */
    navigateToAssets(): void {
        Router.navigateTo(ASSETS_PATH);
    },

    /**
     * Reissues the asset collection request from an error-state retry control.
     *
     * @example
     * ```html
     * <button onclick="assetPage.reloadAssets()">Try again</button>
     * ```
     *
     * @author OpenCode
     */
    reloadAssets(): void {
        const assetsElement = document.querySelector("#assets") as HTMLElement | null;

        if(assetsElement) {
            htmx.trigger(assetsElement, "reload-assets");
        }
    },

    /**
     * Leaves successful list rendering untouched and renders a stable list error on failure.
     *
     * @example
     * ```html
     * <div hx-on::after-request="assetPage.handleAssetListAfterRequest(event)"></div>
     * ```
     *
     * @author OpenCode
     */
    handleAssetListAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            renderListError(event.detail.target as HTMLElement);
        }
    },

    /**
     * Renders a blank creation form each time the create route is entered.
     *
     * @example
     * ```ts
     * assetPage.renderCreateForm(view);
     * ```
     *
     * @author OpenCode
     */
    renderCreateForm(view: HTMLElement): void {
        const contentElement = view.querySelector("#asset-create-content") as HTMLElement | null;

        if(contentElement) {
            renderAssetForm("asset-create-form", contentElement);
        }
    },

    /**
     * Resets the edit view to its loading placeholder before the route-owned GET begins.
     *
     * @example
     * ```ts
     * assetPage.prepareAssetEditView(view);
     * ```
     *
     * @author OpenCode
     */
    prepareAssetEditView(view: HTMLElement): void {
        const contentElement = view.querySelector("#asset-edit-content") as HTMLElement | null;
        const errorElement = view.querySelector("#asset-edit-error") as HTMLElement | null;
        const loadingElement = view.querySelector("#asset-edit-loading") as HTMLElement | null;

        if(contentElement) {
            contentElement.innerHTML = "";
        }

        if(errorElement) {
            errorElement.style.display = "none";
            errorElement.innerHTML = "";
        }

        if(loadingElement) {
            loadingElement.style.display = null;
        }
    },

    /**
     * Rejects an asset response whose ID does not match the requested edit route before HTMX renders it.
     *
     * @example
     * ```html
     * <div hx-on::before-swap="assetPage.validateAssetBeforeSwap(event)"></div>
     * ```
     *
     * @author OpenCode
     */
    validateAssetBeforeSwap(event: AssetBeforeSwapEvent): void {
        if(event.detail.isError || event.detail.requestConfig.verb !== "get") {
            return;
        }

        const requestPath = new URL(event.detail.xhr.responseURL).pathname;

        if(!ASSET_DETAIL_API_PATH_PATTERN.test(requestPath)) {
            return;
        }

        const requestedIdentifier = getCurrentAssetIdentifierForResponse(event.detail.xhr.responseURL);

        if(!requestedIdentifier) {
            event.detail.shouldSwap = false;
            return;
        }

        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!asset || String(asset.id) !== requestedIdentifier) {
            event.detail.shouldSwap = false;
            renderAssetLoadError();
        }
    },

    /**
     * Hides the edit loading placeholder on success or displays a retryable load error on failure.
     *
     * @example
     * ```html
     * <div hx-on::after-request="assetPage.handleAssetLoadAfterRequest(event)"></div>
     * ```
     *
     * @author OpenCode
     */
    handleAssetLoadAfterRequest(event: AssetRequestEvent): void {
        if(event.target !== event.currentTarget || event.detail.requestConfig.verb !== "get"
            || !getCurrentAssetIdentifierForResponse(event.detail.xhr.responseURL)) {
            return;
        }

        if(!event.detail.successful) {
            renderAssetLoadError();
            return;
        }

        const loadingElement = document.querySelector("#asset-edit-loading") as HTMLElement | null;

        if(loadingElement) {
            loadingElement.style.display = "none";
        }
        const form = document.querySelector<HTMLFormElement>("#edit-asset-form");
        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(form && asset) {
            initializeExternalDraft(form, asset);
        }
    },

    /**
     * Handles a successful POST response by navigating to the generated asset's edit route.
     * A failed request leaves the creation form and its input values in place.
     *
     * @example
     * ```html
     * <form hx-on::after-request="assetPage.handleCreateAfterRequest(event)"></form>
     * ```
     *
     * @author OpenCode
     */
    handleCreateAfterRequest(event: AssetRequestEvent): void {
        if(!isAssetFormResponse(event, "post") || !event.detail.successful) {
            return;
        }

        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!asset || typeof asset.id !== "number" || asset.id <= 0) {
            notifyUnexpectedResponse("The server returned an invalid asset.");
            return;
        }

        notifySuccess("Asset created", "The asset was created and is ready for editing.");
        this.navigateToAsset(String(asset.id));
    },

    /**
     * Validates a successful PUT response and refreshes the saved values on the edit route.
     * A failed request leaves the edit form and its input values in place.
     *
     * @example
     * ```html
     * <form hx-on::after-request="assetPage.handleSaveAfterRequest(event)"></form>
     * ```
     *
     * @author OpenCode
     */
    handleSaveAfterRequest(event: AssetRequestEvent): void {
        if(!isAssetFormResponse(event, "put") || !event.detail.successful) {
            return;
        }

        const form = event.currentTarget as HTMLFormElement | null;
        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!form || !asset || typeof asset.id !== "number" || asset.id <= 0) {
            notifyUnexpectedResponse("The server returned an invalid asset.");
            return;
        }

        const requestedId = (form.elements.namedItem("id") as HTMLInputElement | null)?.value;

        if(String(asset.id) !== requestedId) {
            notifyUnexpectedResponse("The server returned an invalid updated asset.");
            return;
        }

        const contentElement = document.querySelector("#asset-edit-content") as HTMLElement | null;

        if(contentElement) {
            renderAssetForm("asset-edit-form", contentElement, asset);
        }
        notifySuccess("Asset saved", "The asset changes were saved.");
    },
};

export default AssetPage;
