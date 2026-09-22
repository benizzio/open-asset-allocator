/**
 * Coordinates the shared asset list and create/edit form page.
 *
 * The controller changes one form between create and edit operations based on the active route.
 * HTMX remains responsible for all asset HTTP requests, while this module validates responses,
 * populates the shared form, and preserves deferred external asset data during updates.
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
type AssetOperation = "create" | "edit";

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

/** Returns the single form shared by create and edit operations. Authored by: OpenCode. */
function getAssetForm(): HTMLFormElement | null {
    return document.querySelector("#asset-form") as HTMLFormElement | null;
}

/** Extracts and decodes the asset identifier from the current detail route. Authored by: OpenCode. */
function getAssetIdentifierFromLocation(): string | null {

    const match = globalThis.location.pathname.match(ASSET_IDENTIFIER_PATH_PATTERN);
    return match ? decodeURIComponent(match[1]) : null;
}

/** Reads the active create or edit operation from the shared form. Authored by: OpenCode. */
function getAssetOperation(form: HTMLFormElement): AssetOperation {
    return form.dataset.operation === "edit" ? "edit" : "create";
}

/** Reports whether a loaded asset carries external data that an update must preserve. Authored by: OpenCode. */
function hasExternalData(asset: AssetDTO): boolean {
    return Object.prototype.hasOwnProperty.call(asset, "externalData")
        && asset.externalData !== undefined
        && asset.externalData !== null;
}

/** Updates both the current and reset baseline value of one form input. Authored by: OpenCode. */
function setInputValue(form: HTMLFormElement, name: string, value: string): void {

    const input = form.elements.namedItem(name) as HTMLInputElement | null;

    if(input) {
        input.value = value;
        input.defaultValue = value;
    }
}

/** Stores the loaded asset in the page for later update payload preparation. Authored by: OpenCode. */
function setAssetData(asset: AssetDTO | null): void {

    const assetDataElement = document.querySelector("#asset-data");

    if(assetDataElement) {
        assetDataElement.textContent = asset ? JSON.stringify(asset) : "";
    }
}

/** Reads and validates the asset stored in the page's JSON state element. Authored by: OpenCode. */
function readAssetData(): AssetDTO | null {
    const assetDataElement = document.querySelector("#asset-data");
    return normalizeAsset(parseJSON(assetDataElement?.textContent));
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

/** Resets and configures the shared form for a create or edit operation. Authored by: OpenCode. */
function configureAssetForm(operation: AssetOperation, asset: AssetDTO | null = null): void {

    const form = getAssetForm();
    const heading = document.querySelector("#asset-form-heading");
    const submitButton = document.querySelector("#asset-submit") as HTMLButtonElement | null;
    const formCard = document.querySelector("#asset-form-content");
    const errorElement = document.querySelector("#asset-form-error") as HTMLElement | null;
    const idInput = form?.elements.namedItem("id") as HTMLInputElement | null;

    if(!form || !heading || !submitButton || !formCard || !idInput) {
        return;
    }

    form.reset();
    form.classList.remove("was-validated");
    form.dataset.operation = operation;
    form.removeAttribute("hx-vals");
    form.querySelector("[data-external-data-marker]")?.remove();

    const isEdit = operation === "edit";
    heading.textContent = isEdit ? "Edit asset" : "New asset";
    submitButton.textContent = isEdit ? "Save" : "Create";
    formCard.classList.toggle("border-primary", !isEdit);
    formCard.classList.toggle("border-dashed", !isEdit);
    idInput.disabled = !isEdit;

    if(errorElement) {
        errorElement.innerHTML = "";
        errorElement.style.display = "none";
    }

    setInputValue(form, "id", isEdit && asset ? String(asset.id) : "");
    setInputValue(form, "ticker", isEdit && asset ? asset.ticker : "");
    setInputValue(form, "name", isEdit && asset ? asset.name : "");
    setAssetData(isEdit ? asset : null);
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

/** Adds unchanged external data to an edit request so PUT does not clear it. Authored by: OpenCode. */
function prepareExternalData(form: HTMLFormElement, event: CustomEvent): void {

    const asset = readAssetData();
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
}

/**
 * Provides browser behavior for the shared asset list and create/edit form page.
 *
 * Use the exported instance through `globalThis.assetPage` from HTMX attributes. For example,
 * `assetPage.initializeAssetFormRoute()` configures the shared form after route navigation.
 *
 * Authored by: OpenCode
 */
const AssetPage = {

    /**
     * Navigates to the shared form in create mode.
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
     * Navigates to one asset and loads it into the shared edit form.
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
     * Configures the shared form for `/asset/new` or requests data for `/asset/:assetId`.
     *
     * Example: invoke from the route-trigger elements in `asset.html`.
     *
     * Authored by: OpenCode
     */
    initializeAssetFormRoute(): void {

        displayAssetForm();

        if(globalThis.location.pathname === NEW_ASSET_PATH) {
            configureAssetForm("create");
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
     * Validates a loaded asset against the route identifier and populates the shared edit form.
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

        configureAssetForm("edit", asset);
    },

    /**
     * Changes the shared HTMX form request between POST and PUT and retains external data on edits.
     *
     * Example: configure on the form's `htmx:configRequest` event.
     *
     * Authored by: OpenCode
     */
    prepareAssetRequest(event: CustomEvent): void {
        const form = event.currentTarget as HTMLFormElement | null;

        if(!form || getAssetOperation(form) === "create") {
            return;
        }

        event.detail.verb = "put";
        prepareExternalData(form, event);
    },

    /**
     * Handles successful POST and PUT responses according to the shared form's current operation.
     * Failed requests retain the user's form values for correction.
     *
     * Example: configure on the form's `htmx:afterRequest` event.
     *
     * Authored by: OpenCode
     */
    handleMutationAfterRequest(event: AssetRequestEvent): void {
        if(!event.detail.successful) {
            return;
        }

        const form = event.currentTarget as HTMLFormElement | null;
        const asset = normalizeAsset(parseJSON(event.detail.xhr.response));

        if(!form || !asset || asset.id <= 0) {
            notifyUnexpectedResponse("The server returned an invalid asset.");
            return;
        }

        if(getAssetOperation(form) === "create") {
            notifySuccess("Asset created", "The asset was created and is ready for editing.");
            this.navigateToAsset(String(asset.id));
            return;
        }

        const requestedId = (form.elements.namedItem("id") as HTMLInputElement | null)?.value;

        if(String(asset.id) !== requestedId) {
            notifyUnexpectedResponse("The server returned an invalid updated asset.");
            return;
        }

        configureAssetForm("edit", asset);
        notifySuccess("Asset saved", "The asset changes were saved.");
    },
};

export default AssetPage;
