/**
 * Manages per-form external-asset drafts, provider search, and ordered record rendering.
 * The main asset form remains the only writer; search requests use a separate HTMX target.
 *
 * @author OpenCode
 */
import notifications from "../components/notifications";
import type { Asset, ExternalAsset } from "../domain/asset";
import DomInfra from "../infra/dom";
import { handlebarsInfra } from "../infra/handlebars";
import { NotificationType } from "../infra/infra-types";

/**
 * Holds the ordered provider records and current search state for one asset form.
 * @author OpenCode
 */
type Draft = { records: ExternalAsset[]; results: ExternalAsset[]; sequence: number; };
/**
 * Carries the HTMX request and outcome used by provider-search event handlers.
 * @author OpenCode
 */
type SearchEvent = CustomEvent<{ xhr: XMLHttpRequest; successful: boolean; }>;

const drafts = new WeakMap<HTMLFormElement, Draft>();
const requestSequences = new WeakMap<XMLHttpRequest, number>();

/**
 * Returns only fields accepted and persisted by the asset API.
 * @author OpenCode
 */
function persistedRecord(record: ExternalAsset): ExternalAsset {
    return { source: record.source, ticker: record.ticker, exchangeId: record.exchangeId };
}

/**
 * Updates the form-json extension's values using the current ordered draft.
 * @author OpenCode
 */
function updatePayload(form: HTMLFormElement, draft: Draft): void {
    const externalData = draft.records.length ? { data: draft.records.map(persistedRecord) } : null;
    form.setAttribute("hx-vals", JSON.stringify({ externalData }));
}

/**
 * Refreshes the ordered table and form payload together.
 * @author OpenCode
 */
function renderRecords(form: HTMLFormElement, draft: Draft): void {
    const target = DomInfra.DomUtils.queryDescendant(form, "[data-external-records]");

    if(target) {
        target.innerHTML = handlebarsInfra.utils.renderTemplateByElementId("asset-external-records", draft.records);
    }
    updatePayload(form, draft);
}

/**
 * Updates the polite status shown during a search or when it returns no results.
 * @author OpenCode
 */
function showMessage(form: HTMLFormElement, message: string): void {
    const target = DomInfra.DomUtils.queryDescendant(form, "[data-external-message]");

    if(target) {
        target.textContent = message;
    }
}

/**
 * Displays provider search and operation failures through the shared error-toast component.
 * @author OpenCode
 */
function showError(message: string): void {
    notifications.notifyError(new Error(message));
}

/**
 * Displays duplicate and search-validation feedback through the shared warning-toast component.
 * @author OpenCode
 */
function showWarning(message: string): void {
    notifications.notify({ title: "Warning", content: message, type: NotificationType.WARNING });
}

/**
 * Displays non-error search feedback through the shared notification component.
 * @author OpenCode
 */
function showInformation(message: string): void {
    notifications.notify({ title: "External asset search", content: message, type: NotificationType.INFO });
}

/**
 * Clears the current query and results, invalidating requests started before this reset.
 * @author OpenCode
 */
function clearSearch(form: HTMLFormElement, draft: Draft): void {
    const input = form.querySelector<HTMLInputElement>("[data-external-query]");
    const search = DomInfra.DomUtils.queryDescendant(form, "[data-external-search]");

    draft.sequence++;
    draft.results = [];

    if(input) {
        input.value = "";
    }

    if(search) {
        search.innerHTML = "";
        search.removeAttribute("hx-vals");
    }
    showMessage(form, "");
}

/**
 * Fills only blank main asset fields from the selected provider result.
 * @author OpenCode
 */
function fillEmptyAssetFields(form: HTMLFormElement, selected: ExternalAsset): void {
    const ticker = form.querySelector<HTMLInputElement>("input[name=\"ticker\"]");
    const name = form.querySelector<HTMLInputElement>("input[name=\"name\"]");

    if(ticker && !ticker.value.trim()) {
        ticker.value = `${ selected.exchangeId }:${ selected.ticker }`;
    }

    if(name && !name.value.trim() && selected.name) {
        name.value = selected.name;
    }
}

/**
 * Validates a provider record before showing or storing it.
 * @author OpenCode
 */
function isExternalAsset(value: unknown): value is ExternalAsset {
    if(typeof value !== "object" || value === null) {
        return false;
    }
    const record = value as Record<string, unknown>;
    return [record.source, record.ticker, record.exchangeId].every(
        field => typeof field === "string" && field.length > 0,
    );
}

/**
 * Checks the shape of persisted external data before allowing an editable form to be displayed.
 *
 * @example
 * ```ts
 * if(hasValidExternalData(response.externalData)) {
 *     // The external data can be used safely by the edit form.
 * }
 * ```
 * @author OpenCode
 */
export function hasValidExternalData(value: unknown): value is Asset["externalData"] {
    return value === undefined || value === null || (typeof value === "object" && value !== null
        && Array.isArray((value as { data?: unknown }).data)
        && (value as { data: unknown[] }).data.every(isExternalAsset));
}

/**
 * Initializes a newly rendered create or edit form with independent copies of its ordered records.
 *
 * @example
 * ```ts
 * initializeExternalDraft(form, loadedAsset);
 * ```
 * @author OpenCode
 */
export function initializeExternalDraft(form: HTMLFormElement, asset?: Asset): void {
    const draft: Draft = {
        records: asset?.externalData?.data.map(persistedRecord) ?? [],
        results: [],
        sequence: 0,
    };
    drafts.set(form, draft);
    renderRecords(form, draft);
}

/**
 * Starts a provider lookup from the attached search button or the input's Enter key.
 *
 * @example
 * ```ts
 * searchExternalAssets(button.closest("form")!);
 * ```
 * @author OpenCode
 */
export function searchExternalAssets(form: HTMLFormElement): void {
    const draft = drafts.get(form);
    const input = form.querySelector<HTMLInputElement>("[data-external-query]");
    const search = DomInfra.DomUtils.queryDescendant(form, "[data-external-search]");

    if(!draft || !input || !search) {
        return;
    }
    const query = input.value.trim();

    draft.sequence++;
    draft.results = [];
    search.innerHTML = "";
    search.removeAttribute("hx-vals");
    showMessage(form, "");

    if(!query || query.length > 100) {
        showWarning("Enter a search term containing 1 to 100 characters.");
        return;
    }
    showMessage(form, "Searching external assets…");
    search.setAttribute("hx-vals", JSON.stringify({ query }));
    globalThis.htmx.trigger(search, "search-external-assets");
}

/**
 * Records the HTMX request generation so late responses cannot replace newer results.
 *
 * @example
 * ```ts
 * searchTarget.addEventListener("htmx:beforeRequest", handleExternalSearchBeforeRequest);
 * ```
 * @author OpenCode
 */
export function handleExternalSearchBeforeRequest(event: SearchEvent): void {
    const form = (event.currentTarget as HTMLElement).closest("form");
    const draft = form && drafts.get(form);

    if(draft) {
        requestSequences.set(event.detail.xhr, draft.sequence);
    }
}

/**
 * Renders the current provider response, capped at ten rows, and reports empty or failed searches.
 *
 * @example
 * ```ts
 * searchTarget.addEventListener("htmx:afterRequest", handleExternalSearchAfterRequest);
 * ```
 * @author OpenCode
 */
export function handleExternalSearchAfterRequest(event: SearchEvent): void {
    const search = event.currentTarget as HTMLElement;
    const form = search.closest("form");
    const draft = form && drafts.get(form);

    if(!form || !draft || !form.isConnected || requestSequences.get(event.detail.xhr) !== draft.sequence) {
        return;
    }

    if(!event.detail.successful) {
        search.innerHTML = "";
        showMessage(form, "");
        showError("External assets could not be searched. Try again.");
        return;
    }

    try {
        const data: unknown = JSON.parse(event.detail.xhr.responseText);

        if(!Array.isArray(data) || !data.every(isExternalAsset)) {
            throw new Error("Invalid external asset search response");
        }
        draft.results = data.slice(0, 10);

        search.innerHTML = draft.results.length
            ? handlebarsInfra.utils.renderTemplateByElementId("asset-external-results", draft.results) : "";
        showMessage(form, "");

        if(draft.results.length === 0) {
            showInformation("No external assets found.");
        }
    } catch {
        search.innerHTML = "";
        showMessage(form, "");
        showError("External assets could not be searched. Try again.");
    }
}

/**
 * Adds a selected result unless its source, ticker and exchange ID are already in this draft.
 *
 * @example
 * ```ts
 * addExternalAsset(form, 0);
 * ```
 * @author OpenCode
 */
export function addExternalAsset(form: HTMLFormElement, index: number): void {
    const draft = drafts.get(form);
    const selected = draft?.results[index];

    if(!draft || !selected) {
        return;
    }

    if(draft.records.some(record => record.source === selected.source && record.ticker === selected.ticker
        && record.exchangeId === selected.exchangeId)) {
        showWarning(
            `${ selected.ticker } from ${ selected.source } on exchange ${ selected.exchangeId } is already added.`,
        );
        return;
    }

    fillEmptyAssetFields(form, selected);
    draft.records.push(persistedRecord(selected));
    renderRecords(form, draft);
    clearSearch(form, draft);
}

/**
 * Removes one registered provider entry from the local draft without saving the asset.
 *
 * @example
 * ```ts
 * removeExternalAsset(form, 0);
 * ```
 * @author OpenCode
 */
export function removeExternalAsset(form: HTMLFormElement, index: number): void {
    const draft = drafts.get(form);

    if(draft && index >= 0 && index < draft.records.length) {
        draft.records.splice(index, 1);
        showMessage(form, "");
        renderRecords(form, draft);
    }
}

/**
 * Moves one entry up or down and serializes the updated provider priority order.
 *
 * @example
 * ```ts
 * moveExternalAsset(form, 1, -1);
 * ```
 * @author OpenCode
 */
export function moveExternalAsset(form: HTMLFormElement, index: number, direction: number): void {
    const draft = drafts.get(form);
    const destination = index + direction;

    if(draft && index >= 0 && destination >= 0 && destination < draft.records.length) {
        [draft.records[index], draft.records[destination]] = [draft.records[destination], draft.records[index]];
        renderRecords(form, draft);
    }
}
