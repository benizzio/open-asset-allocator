/**
 * Coordinates asset list, creation, and editing interactions for the assets pages.
 *
 * The page uses HTMX for HTTP requests and Handlebars for response rendering. This module only
 * handles state that cannot be represented by HTML attributes, such as interpreting mutation
 * responses and preserving external asset data while its editor is deferred.
 *
 * Authored by: OpenCode
 */
import htmx from "htmx.org";
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

const ASSETS_PATH = "/assets";
const ASSET_IDENTIFIER_PATH_PATTERN = /^\/asset\/([^/]+)\/?$/;

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

function readAssetJSON(root: HTMLElement): AssetDTO | null {

    const assetJSONElement = (root.querySelector("script#asset-data")
        ?? root.closest("[data-asset-view]")?.querySelector("script#asset-data")) as HTMLScriptElement | null;

    if(!assetJSONElement) {
        return null;
    }

    return normalizeAsset(parseJSON(assetJSONElement.textContent));
}

function getRequestTarget(event: AssetRequestEvent): HTMLElement {
    return (event.detail.target as HTMLElement | undefined) ?? event.target as HTMLElement;
}

function getAssetIdentifierFromLocation(): string | null {

    const match = globalThis.location.pathname.match(ASSET_IDENTIFIER_PATH_PATTERN);
    return match ? decodeURIComponent(match[1]) : null;
}

function getAssetEditorForm(event: CustomEvent): HTMLFormElement | null {
    const eventTarget = (event.currentTarget ?? event.target) as HTMLElement | null;
    return eventTarget?.tagName === "FORM" ? eventTarget as HTMLFormElement : null;
}

function hasExternalData(asset: AssetDTO): boolean {
    return Object.prototype.hasOwnProperty.call(asset, "externalData")
        && asset.externalData !== undefined
        && asset.externalData !== null;
}

function setInputValue(form: HTMLFormElement, name: string, value: string): void {

    const input = form.elements.namedItem(name) as HTMLInputElement | null;

    if(input) {
        input.value = value;
        input.defaultValue = value;
    }
}

function notifySuccess(title: string, content: string): void {
    notifications.notify({ title, content, type: NotificationType.SUCCESS });
}

function notifyUnexpectedResponse(message: string): void {
    notifications.notifyError(new Error(message));
}

function renderListError(target: HTMLElement): void {

    target.innerHTML = `
        <div class="row justify-content-center my-5">
            <div class="col-12 col-lg-8">
                <div class="alert alert-danger" role="alert">
                    <h2 class="h5">Assets could not be loaded</h2>
                    <p>There was a problem loading the asset list.</p>
                    <button type="button" class="btn btn-outline-danger" onclick="assetsPage.reloadAssets()">
                        Try again
                    </button>
                </div>
            </div>
        </div>
    `;
}

function renderAssetLoadError(target: HTMLElement): void {

    if(!document.body.contains(target)) {
        return;
    }

    target.innerHTML = `
        <div class="row justify-content-center my-5">
            <div class="col-12 col-lg-8">
                <div class="alert alert-danger" role="alert">
                    <h2 class="h5">Asset could not be loaded</h2>
                    <p>The requested asset is not available for editing.</p>
                    <button type="button" class="btn btn-outline-danger" onclick="assetsPage.navigateToAssets()">
                        Back to assets
                    </button>
                </div>
            </div>
        </div>
    `;
}

function renderUpdateError(target: HTMLElement): void {

    if(!document.body.contains(target)) {
        return;
    }

    const form = target.querySelector("form") as HTMLFormElement | null;

    if(form) {
        form.querySelectorAll("button, input").forEach(element => {
            (element as HTMLInputElement | HTMLButtonElement).disabled = false;
        });
    }
}

/**
 * Provides the browser-side behavior required by the assets list and forms.
 *
 * Use the exported instance through `globalThis.assetsPage` from HTMX attributes. For example,
 * an asset form can call `assetsPage.handleUpdateAfterRequest(event)` after a successful PUT.
 *
 * Authored by: OpenCode
 */
const AssetsPage = {

    /**
     * Navigates to the asset table without submitting or mutating the current form.
     *
     * Example: `<button type="button" onclick="assetsPage.navigateToAssets()">Back</button>`.
     *
     * Authored by: OpenCode
     */
    navigateToAssets(): void {
        Router.navigateTo(ASSETS_PATH);

        const assetsElement = document.querySelector("#assets") as HTMLElement | null;

        if(assetsElement) {
            htmx.trigger(assetsElement, "reload-assets");
        }
    },

    /**
     * Reissues the list request after a failed asset-list load.
     *
     * Example: `assetsPage.reloadAssets()` from a retry button.
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
     * Handles an asset-list request failure while leaving successful HTMX rendering untouched.
     *
     * Example: configure `hx-on::after-request="assetsPage.handleAssetListAfterRequest(event)"`
     * on the list target.
     *
     * Authored by: OpenCode
     */
    handleAssetListAfterRequest(event: AssetRequestEvent): void {
        if(event.detail.successful) {
            return;
        }

        renderListError(getRequestTarget(event));
    },

    /**
     * Processes a successful create response and opens the generated asset in edit mode.
     * Failed responses remain on the create form so the global error handler can show its details.
     *
     * Example: configure `hx-on::after-request="assetsPage.handleCreateAfterRequest(event)"`
     * on a form posting to `/api/asset`.
     *
     * Authored by: OpenCode
     */
    handleCreateAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            return;
        }

        const createdAsset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!createdAsset || createdAsset.id <= 0) {
            notifyUnexpectedResponse("The server returned an invalid created asset.");
            return;
        }

        notifySuccess("Asset created", "The asset was created and is ready for editing.");
        Router.navigateTo(`/asset/${ createdAsset.id }`);
    },

    /**
     * Displays a stable error view for an unsuccessful asset detail request.
     *
     * Example: configure `hx-on::after-request="assetsPage.handleAssetLoadAfterRequest(event)"`
     * on the detail request target.
     *
     * Authored by: OpenCode
     */
    handleAssetLoadAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            renderAssetLoadError(getRequestTarget(event));
        }
    },

    /**
     * Validates and initializes the asset editor after its Handlebars response settles.
     * The returned identifier must match the route identifier before the form can be used.
     *
     * Example: configure `hx-on::after-settle="assetsPage.initializeAssetEditor(this)"`
     * on the detail request target.
     *
     * Authored by: OpenCode
     */
    initializeAssetEditor(target: HTMLElement): void {
        if(!document.body.contains(target)) {
            return;
        }

        const asset = readAssetJSON(target);
        const requestedIdentifier = getAssetIdentifierFromLocation();

        if(!asset || requestedIdentifier === null || String(asset.id) !== requestedIdentifier) {
            renderAssetLoadError(target);
        }
    },

    /**
     * Adds retained external asset data to the form-json request without exposing editing controls.
     * The hidden marker ensures form-json includes the key, while hx-vals supplies the nested object.
     *
     * Example: configure `hx-on::config-request="assetsPage.prepareUpdateRequest(event)"` on the
     * asset update form.
     *
     * Authored by: OpenCode
     */
    prepareUpdateRequest(event: CustomEvent): void {
        const form = getAssetEditorForm(event);

        if(!form) {
            return;
        }

        const asset = readAssetJSON(form);
        const externalDataMarker = form.querySelector("[data-external-data-marker]");
        const requestFormData = (event.detail as { formData?: FormData }).formData;

        if(!asset || !hasExternalData(asset)) {
            form.removeAttribute("hx-vals");
            externalDataMarker?.remove();
            requestFormData?.delete("externalData");
            return;
        }

        if(!externalDataMarker) {
            const marker = document.createElement("input");
            marker.type = "hidden";
            marker.name = "externalData";
            marker.value = "";
            marker.setAttribute("data-external-data-marker", "true");
            form.appendChild(marker);
            requestFormData?.append("externalData", "");
        }

        form.setAttribute("hx-vals", JSON.stringify({ externalData: asset.externalData }));
    },

    /**
     * Handles a successful basic-field update while keeping the user on the edit route.
     * The response becomes the new saved baseline for subsequent updates.
     *
     * Example: configure `hx-on::after-request="assetsPage.handleUpdateAfterRequest(event)"`
     * on a form sending `PUT /api/asset`.
     *
     * Authored by: OpenCode
     */
    handleUpdateAfterRequest(event: AssetRequestEvent): void {
        const form = getAssetEditorForm(event);

        if(!form) {
            return;
        }

        if(!event.detail.successful) {
            renderUpdateError(getRequestTarget(event));
            return;
        }

        const updatedAsset = normalizeAsset(parseJSON(event.detail.xhr.response));
        const requestedId = (form.elements.namedItem("id") as HTMLInputElement | null)?.value;

        if(!updatedAsset || String(updatedAsset.id) !== requestedId) {
            notifyUnexpectedResponse("The server returned an invalid updated asset.");
            return;
        }

        setInputValue(form, "id", String(updatedAsset.id));
        setInputValue(form, "ticker", updatedAsset.ticker);
        setInputValue(form, "name", updatedAsset.name);

        const assetJSONElement = form.closest("[data-asset-view]")?.querySelector("script#asset-data");

        if(assetJSONElement) {
            assetJSONElement.textContent = JSON.stringify(updatedAsset);
        }

        notifySuccess("Asset saved", "The asset changes were saved.");
    },
};

export default AssetsPage;
