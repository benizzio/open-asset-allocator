/**
 * Coordinates asset list, creation, and editing on one page.
 *
 * HTMX handles HTTP requests and each form declares its own operation in HTML. This controller
 * renders the create or edit template, checks API response identities, and navigates between views.
 *
 * Authored by: OpenCode
 */
import htmx from "htmx.org";
import * as Handlebars from "handlebars";
import notifications from "../components/notifications";
import { NotificationType } from "../infra/infra-types";
import { AfterRequestEventDetail } from "../infra/htmx";
import Router from "../infra/routing";

type ExternalAssetData = Record<string, unknown>;

type AssetDTO = {
    id: number;
    name: string;
    ticker: string;
    externalData?: ExternalAssetData | null;
};

type AssetRequestEvent = CustomEvent<AfterRequestEventDetail>;

const ASSETS_PATH = "/asset";
const NEW_ASSET_PATH = "/asset/new";
const ASSET_IDENTIFIER_PATH_PATTERN = /^\/asset\/([^/]+)\/?$/;

/** Parses an unknown response value as JSON without propagating malformed-response errors. Authored by: OpenCode. */
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

/** Validates and normalizes an API response into the asset shape used by this page. Authored by: OpenCode. */
function normalizeAsset(value: unknown): AssetDTO | null {

    if(typeof value !== "object" || value === null) {
        return null;
    }

    const candidate = value as Record<string, unknown>;
    const numericId = typeof candidate.id === "number" ? candidate.id : Number(candidate.id);

    if(!Number.isSafeInteger(numericId) || typeof candidate.name !== "string" || typeof candidate.ticker !== "string") {
        return null;
    }

    const externalData = candidate.externalData;

    if(externalData !== undefined && externalData !== null && typeof externalData !== "object") {
        return null;
    }

    return {
        id: numericId,
        name: candidate.name,
        ticker: candidate.ticker,
        externalData: externalData as ExternalAssetData | null | undefined,
    };
}

/** Extracts and decodes the asset identifier from the current detail route. Authored by: OpenCode. */
function getAssetIdentifierFromLocation(): string | null {

    const match = globalThis.location.pathname.match(ASSET_IDENTIFIER_PATH_PATTERN);
    return match ? decodeURIComponent(match[1]) : null;
}

/** Switches the edit view between its loading placeholder and form content. Authored by: OpenCode. */
function setFormLoading(isLoading: boolean): void {

    const loadingElement = document.querySelector("#asset-form-loading") as HTMLElement | null;
    const contentElement = document.querySelector("#asset-form-content") as HTMLElement | null;

    if(loadingElement) {
        loadingElement.style.display = isLoading ? null : "none";
    }

    if(contentElement) {
        contentElement.style.display = isLoading ? "none" : null;
    }
}

/** Displays the shared form and hides the asset collection in the mounted page. Authored by: OpenCode. */
function displayAssetForm(): void {

    const assetsElement = document.querySelector("#assets") as HTMLElement | null;
    const formViewElement = document.querySelector("#asset-form-view") as HTMLElement | null;

    if(assetsElement) {
        assetsElement.style.display = "none";
    }

    if(formViewElement) {
        formViewElement.style.display = null;
    }
}

/** Displays the asset collection and hides the shared form in the mounted page. Authored by: OpenCode. */
function displayAssetList(): void {

    const assetsElement = document.querySelector("#assets") as HTMLElement | null;
    const formViewElement = document.querySelector("#asset-form-view") as HTMLElement | null;

    if(formViewElement) {
        formViewElement.style.display = "none";
    }

    if(assetsElement) {
        assetsElement.style.display = null;
    }
}

/** Renders the operation-specific form and binds HTMX to the inserted form. Authored by: OpenCode. */
function renderAssetForm(templateId: "asset-create-form" | "asset-edit-form", asset?: AssetDTO): void {

    const template = document.getElementById(templateId);
    const contentElement = document.querySelector("#asset-form-content") as HTMLElement | null;
    const errorElement = document.querySelector("#asset-form-error") as HTMLElement | null;

    if(!template || !contentElement) {
        return;
    }

    if(errorElement) {
        errorElement.innerHTML = "";
        errorElement.style.display = "none";
    }

    contentElement.innerHTML = Handlebars.compile(template.innerHTML)(asset ?? {});
    // The script-loaded HTMX instance owns the form-json extension used by these forms.
    globalThis.htmx.process(contentElement);
    setFormLoading(false);
}

/** Shows a success toast through the application notification component. Authored by: OpenCode. */
function notifySuccess(title: string, content: string): void {
    notifications.notify({ title, content, type: NotificationType.SUCCESS });
}

/** Reports an invalid successful API response through the notification component. Authored by: OpenCode. */
function notifyUnexpectedResponse(message: string): void {
    notifications.notifyError(new Error(message));
}

/** Replaces the collection target with a retryable load error. Authored by: OpenCode. */
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

/** Displays a recoverable detail-load error without destroying the shared form. Authored by: OpenCode. */
function renderAssetLoadError(): void {

    const contentElement = document.querySelector("#asset-form-content") as HTMLElement | null;
    const errorElement = document.querySelector("#asset-form-error") as HTMLElement | null;

    if(!contentElement || !errorElement) {
        return;
    }

    setFormLoading(false);
    contentElement.style.display = "none";
    errorElement.style.display = null;

    errorElement.innerHTML = `
        <div class="alert alert-danger mb-0" role="alert">
            <h2 class="h5">Asset could not be loaded</h2>
            <p>The requested asset is not available for editing.</p>
            <button type="button" class="btn btn-outline-danger" onclick="assetPage.navigateToAssets()">
                Back to assets
            </button>
        </div>
    `;
}

/**
 * Provides browser behavior for the asset list and operation-specific forms.
 *
 * Use the exported instance through `globalThis.assetPage` from HTMX attributes. For example,
 * `assetPage.initializeAssetFormRoute()` renders the appropriate form after route navigation.
 *
 * Authored by: OpenCode
 */
const AssetPage = {

    /**
     * Navigates to the asset creation form.
     *
     * Example: `<button onclick="assetPage.navigateToNewAsset()">New asset</button>`.
     *
     * Authored by: OpenCode
     */
    navigateToNewAsset(): void {
        Router.navigateTo(NEW_ASSET_PATH);
        this.initializeAssetFormRoute();
    },

    /**
     * Navigates to one asset and loads it into the edit form.
     *
     * Example: `assetPage.navigateToAsset("42")` opens `/asset/42`.
     *
     * Authored by: OpenCode
     */
    navigateToAsset(assetId: string): void {
        Router.navigateTo(`/asset/${ encodeURIComponent(assetId) }`);
        this.initializeAssetFormRoute();
    },

    /**
     * Activates asset-row navigation from Enter or Space without duplicating pointer navigation.
     *
     * Example: `assetPage.handleAssetNavigationKeypress(event, "42")` from a table row.
     *
     * Authored by: OpenCode
     */
    handleAssetNavigationKeypress(event: KeyboardEvent, assetId: string): void {
        if(event.key === "Enter" || event.key === " ") {
            this.navigateToAsset(assetId);
        }
    },

    /**
     * Navigates to the asset table and reloads its data when the page is already mounted.
     *
     * Example: `<button type="button" onclick="assetPage.navigateToAssets()">Back</button>`.
     *
     * Authored by: OpenCode
     */
    navigateToAssets(): void {
        Router.navigateTo(ASSETS_PATH);
        displayAssetList();

        const assetsElement = document.querySelector("#assets") as HTMLElement | null;

        if(assetsElement) {
            htmx.trigger(assetsElement, "reload-assets");
        }
    },

    /**
     * Reissues the asset collection request from an error-state retry control.
     *
     * Example: `<button onclick="assetPage.reloadAssets()">Try again</button>`.
     *
     * Authored by: OpenCode
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
     * Example: configure on the collection target's `htmx:afterRequest` event.
     *
     * Authored by: OpenCode
     */
    handleAssetListAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            renderListError(event.detail.target as HTMLElement);
        }
    },

    /**
     * Renders the creation form for `/asset/new` or requests data for `/asset/:assetId`.
     *
     * Example: invoke from the route-trigger elements in `asset.html`.
     *
     * Authored by: OpenCode
     */
    initializeAssetFormRoute(): void {

        displayAssetForm();

        if(globalThis.location.pathname === NEW_ASSET_PATH) {
            renderAssetForm("asset-create-form");
            return;
        }

        const requestedIdentifier = getAssetIdentifierFromLocation();
        const detailLoader = document.querySelector("#asset-detail-loader") as HTMLElement | null;
        const identifierInput = detailLoader?.querySelector("input[name='assetId']") as HTMLInputElement | null;

        if(!requestedIdentifier || !detailLoader || !identifierInput) {
            renderAssetLoadError();
            return;
        }

        identifierInput.value = requestedIdentifier;
        setFormLoading(true);
        htmx.trigger(detailLoader, "load-asset");
    },

    /**
     * Validates a loaded asset against the route identifier and renders the edit form.
     *
     * Example: configure on the detail loader's `htmx:afterRequest` event.
     *
     * Authored by: OpenCode
     */
    handleAssetLoadAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            renderAssetLoadError();
            return;
        }

        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));
        const requestedIdentifier = getAssetIdentifierFromLocation();

        if(!asset || requestedIdentifier === null || String(asset.id) !== requestedIdentifier) {
            renderAssetLoadError();
            return;
        }

        renderAssetForm("asset-edit-form", asset);
    },

    /**
     * Handles a successful POST response by navigating to the generated asset's edit route.
     * A failed request leaves the creation form and its input values in place.
     *
     * Example: configure on the creation form's `htmx:afterRequest` event.
     *
     * Authored by: OpenCode
     */
    handleCreateAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            return;
        }

        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!asset || asset.id <= 0) {
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
     * Example: configure on the edit form's `htmx:afterRequest` event.
     *
     * Authored by: OpenCode
     */
    handleSaveAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            return;
        }

        const form = event.currentTarget as HTMLFormElement | null;
        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!form || !asset || asset.id <= 0) {
            notifyUnexpectedResponse("The server returned an invalid asset.");
            return;
        }

        const requestedId = (form.elements.namedItem("id") as HTMLInputElement | null)?.value;

        if(String(asset.id) !== requestedId) {
            notifyUnexpectedResponse("The server returned an invalid updated asset.");
            return;
        }

        renderAssetForm("asset-edit-form", asset);
        notifySuccess("Asset saved", "The asset changes were saved.");
    },
};

export default AssetPage;
