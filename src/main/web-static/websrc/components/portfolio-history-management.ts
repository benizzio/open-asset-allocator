/**
 * Manages portfolio-history observation forms and synchronizes each row's first external-asset association.
 *
 * @example `portfolioHistoryManagement.init()`
 * @author GPT-6 Luna
 * @author benizzio
 * @author GPT-6 Sol
 * @author GPT-6.1 Sol
 * @author GPT-6 Luna
 */

import htmx from "htmx.org";
import { BigNumber } from "bignumber.js";
import { AfterRequestEventDetail, HtmxInfra, type BeforeSwapEventDetail } from "../infra/htmx";
import { ObservationTimestamp } from "../domain/portfolio-allocation";
import Router from "../infra/routing";
import {
    createPortfolioHistoryQuoteAction,
    readExternalAssetKeys,
    EXTERNAL_ASSET_KEYS_SELECTOR,
    EXTERNAL_ASSET_SOURCE_SELECTOR,
    EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR,
    EXTERNAL_ASSET_TICKER_SELECTOR,
    PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX,
} from "./portfolio-history-quote";
import AssetComposedColumnsInput, {
    ASSET_ROW_SELECTION_CHANGE_EVENT,
    AssetRowSelectionChangeState,
    type AssetRowSelectionChangeEvent,
} from "./asset-composed-columns-input";
import { toInt } from "../utils/lang";
import type { TemplateDelegate } from "handlebars";
import notifications from "./notifications";
import { NotificationType } from "../infra/infra-types";
import { getBrowserLocalCalendarDate } from "../utils/calendar-date";
import {
    activateManagementMetadataEditor,
    bindReadOnlyManagementMetadataHeaderToggle,
    hideManagementMetadataEditor,
} from "./management-metadata-editor";
import {
    areObservationDetailsReady,
    bindObservationDetailsLoader,
    reloadObservationDetails,
    retryObservationDetails,
} from "./portfolio-history-management-observation-details";
import { bindObservationMetadataRequest } from "./portfolio-history-management-observation-metadata";
import paginatedList, { PaginationMode } from "./paginated-list";

const PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER = "accordion-portfolio-history-management";
const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-history-management-form-tbody-";
const NEW_OBSERVATION_DATE_INPUT_ID = "portfolio-history-management-observation-date";
const ACCORDION_ITEM_SELECTOR = ".accordion-item";

/** Creates the row-scoped quote action before its first use in row normalization.
 *
 * @author GPT-6 Sol
 */
const portfolioHistoryQuoteAction = createPortfolioHistoryQuoteAction(recalculatePortfolioHistoryAllocation);

/** Reuses the existing portfolio-history row calculation after a quote changes Market Price.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function recalculatePortfolioHistoryAllocation(allocationIndex: number, observationTimestampId: number): void {
    portfolioHistoryManagement.handleInputQuantityOrMarketPrice(allocationIndex, String(observationTimestampId));
}

/** Writes a complete key set atomically, or clears and disables the row's entire association group.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function synchronizeExternalAssetKeyGroup(
    fieldset: HTMLFieldSetElement,
    values: { source?: unknown; exchangeId?: unknown; ticker?: unknown } | undefined,
): void {
    const sourceInput = fieldset.querySelector<HTMLInputElement>(EXTERNAL_ASSET_SOURCE_SELECTOR);
    const exchangeIdInput = fieldset.querySelector<HTMLInputElement>(EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR);
    const tickerInput = fieldset.querySelector<HTMLInputElement>(EXTERNAL_ASSET_TICKER_SELECTOR);
    const keys = readExternalAssetKeys(values);

    if(!sourceInput || !exchangeIdInput || !tickerInput || !keys) {
        [sourceInput, exchangeIdInput, tickerInput].forEach(input => {
            if(input) {
                input.value = "";
            }
        });
        fieldset.disabled = true;
        return;
    }

    sourceInput.value = keys.source;
    exchangeIdInput.value = keys.exchangeId;
    tickerInput.value = keys.ticker;
    fieldset.disabled = false;
}

/** Normalizes one server-rendered row without replacing valid first-provider values.
 *
 * @author GPT-6 Luna
 */
function normalizePortfolioHistoryRow(row: HTMLTableRowElement): void {
    const generation = Number(row.dataset.assetSelectionGeneration);

    if(!Number.isSafeInteger(generation) || generation < 0) {
        row.dataset.assetSelectionGeneration = "0";
    }

    const fieldset = row.querySelector<HTMLFieldSetElement>(EXTERNAL_ASSET_KEYS_SELECTOR);

    if(fieldset) {
        synchronizeExternalAssetKeyGroup(fieldset, {
            source: fieldset.querySelector<HTMLInputElement>(EXTERNAL_ASSET_SOURCE_SELECTOR)?.value,
            exchangeId: fieldset.querySelector<HTMLInputElement>(EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR)?.value,
            ticker: fieldset.querySelector<HTMLInputElement>(EXTERNAL_ASSET_TICKER_SELECTOR)?.value,
        });
    }

    portfolioHistoryQuoteAction.bindRow(row);
    portfolioHistoryQuoteAction.updateButtonVisibility(row);
}

/** Applies one bubbling asset-selection event only to its owning portfolio-history row.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function handleAssetRowSelectionChange(event: Event): void {
    if(!(event instanceof CustomEvent)) {
        return;
    }

    const row = event.target;

    if(!(row instanceof HTMLTableRowElement)
        || !row.id.startsWith(PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX)) {
        return;
    }

    const fieldset = row.querySelector<HTMLFieldSetElement>(EXTERNAL_ASSET_KEYS_SELECTOR);
    const detail = (event as AssetRowSelectionChangeEvent).detail;

    if(!detail
        || !Number.isSafeInteger(detail.generation)
        || row.dataset.assetSelectionGeneration !== String(detail.generation)) {
        return;
    }

    if(detail.state === AssetRowSelectionChangeState.EXISTING) {
        if(fieldset) {
            synchronizeExternalAssetKeyGroup(fieldset, detail.asset.externalData?.data?.[0]);
        }
    }
    else if(detail.state === AssetRowSelectionChangeState.NEW || detail.state === AssetRowSelectionChangeState.SEARCH) {
        if(fieldset) {
            synchronizeExternalAssetKeyGroup(fieldset, undefined);
        }
    }

    portfolioHistoryQuoteAction.updateButtonVisibility(row);
}

/** Normalizes history rows within one HTMX swap target, including a target that is itself a row.
 *
 * @author GPT-6 Luna
 */
function normalizePortfolioHistoryRowsWithin(element: Element): void {
    const rowSelector = `tr[id^="${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }"]`;

    if(element.matches(rowSelector) && element instanceof HTMLTableRowElement) {
        normalizePortfolioHistoryRow(element);
    }

    element.querySelectorAll<HTMLTableRowElement>(rowSelector).forEach(normalizePortfolioHistoryRow);
}

class FormRowValueElements {

    quantityInput: HTMLInputElement;
    marketPriceInput: HTMLInputElement;
    totalMarketValueInput: HTMLInputElement;

    constructor(formUniqueId: string, formRowIndex: number) {

        const formRowId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ formUniqueId }-row-${ formRowIndex }`;
        const formRow = window[formRowId] as HTMLElement;

        this.quantityInput = formRow.querySelector("[name$='[assetQuantity]']");
        this.marketPriceInput = formRow.querySelector("[name$='[assetMarketPrice]']");
        this.totalMarketValueInput = formRow.querySelector("[name$='[totalMarketValue]']");
    }

    handleInputQuantityOrMarketPrice() {

        const quantity = this.quantityInput.value;
        const marketPrice = this.marketPriceInput.value;

        if(quantity && marketPrice) {
            try {
                const totalMarketValue = new BigNumber(quantity)
                    .times(new BigNumber(marketPrice))
                    .decimalPlaces(0, BigNumber.ROUND_HALF_UP);
                this.totalMarketValueInput.value = totalMarketValue.toString();
            } catch {
                this.totalMarketValueInput.value = "";
            }
        }
    }

    handleInputTotalMarketValue() {
        this.quantityInput.value = "";
        this.marketPriceInput.value = "";
    }
}

function focusOnNewLine(newRow: HTMLElement) {

    const firstFieldOfNewRow: HTMLInputElement = newRow.querySelector("input[type='text']");

    if(firstFieldOfNewRow) {
        firstFieldOfNewRow.scrollIntoView({ behavior: "smooth", block: "center" });
        firstFieldOfNewRow.focus();
    }
}

function getNextPortfolioHistoryManagementIndex(tbody: HTMLElement): number {
    const rows = tbody.querySelectorAll("tr");
    const lastRow = rows[rows.length - 1] as HTMLElement;
    const lastRowId = lastRow?.id;
    const lastRowIdIndex = lastRowId?.split("-").pop();
    return lastRowIdIndex ? toInt(lastRowIdIndex) + 1 : 0;
}

/** Maps a paginated observation response, including the new-observation form only on initial loads.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function modifyObservationsResponse(
    originalServerResponseJSON: string,
    eventDetail: BeforeSwapEventDetail,
): string {
    return paginatedList.transformPageResponse<ObservationTimestamp, ObservationTimestamp[]>(
        originalServerResponseJSON,
        eventDetail,
        observation => observation.id,
        (observations, mode) => mode === PaginationMode.Initial
            ? [{ id: 0, timeTag: "New Observation *" }, ...observations]
            : observations,
    );
}

HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
    "addObservationZero",
    modifyObservationsResponse,
);

/** Refreshes observation details or its ordered list after a successful form submission.
 * @author GPT-6 Luna
 */
function propagateRefreshDataAfterPost(observationTimestampId: number, metadataChanged: boolean) {

    const portfolioHistoryManagementContainerElement =
        document.getElementById(PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER);

    if(observationTimestampId !== 0 && !metadataChanged) {
        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = document.getElementById(formId);
        const header = document.getElementById(`portfolio-history-management-trigger-${ observationTimestampId }`);

        const editorFields = form?.closest<HTMLElement>(ACCORDION_ITEM_SELECTOR)
            ?.querySelector<HTMLElement>("[data-metadata-editor-active]");

        const tagInput = form?.closest<HTMLElement>(ACCORDION_ITEM_SELECTOR)
            ?.querySelector<HTMLInputElement>("[data-observation-editor-tag]");

        const dateInput = form?.closest<HTMLElement>(ACCORDION_ITEM_SELECTOR)
            ?.querySelector<HTMLInputElement>("[data-observation-editor-date]");

        if(form instanceof HTMLFormElement && header instanceof HTMLElement) {
            form.reset();

            if(editorFields && tagInput && dateInput) {
                hideManagementMetadataEditor(editorFields, tagInput, dateInput);
            }

            reloadObservationDetails(header, form);
        }
    }

    else if(portfolioHistoryManagementContainerElement) {
        htmx.trigger(portfolioHistoryManagementContainerElement, "reload-portfolio-history-management");
    }

    const portfolioHistoryViewContainerElement = window["accordion-portfolio-history"];

    if(portfolioHistoryViewContainerElement) {
        htmx.trigger(portfolioHistoryViewContainerElement, "reload-portfolio-history");
    }

    AssetComposedColumnsInput.loadDatalists();
}

const portfolioHistoryManagement = {

    handlebarsPortfolioHistoryManagementRowTemplate: null as TemplateDelegate,
    handlebarsPortfolioHistoryManagementContainerTemplate: null as TemplateDelegate,

    /** Initializes this component on its own settle and normalizes rows inserted by nested history requests.
     *
     * @param event - HTMX settle event from this component or a nested allocation-row swap.
     * @param element - The portfolio-history management container receiving the settle event.
     * @example `portfolioHistoryManagement.handleAfterSettle(event, this)`
     * @author GPT-6 Luna
     * @author benizzio
     */
    handleAfterSettle(event: CustomEvent, element: HTMLElement) {

        if(event.target !== element) {
            if(event.target instanceof Element && element.contains(event.target)) {
                normalizePortfolioHistoryRowsWithin(event.target);
            }

            return;
        }

        this.init();
    },

    /** Initializes templates, row synchronization, and guards for the current management forms.
     *
     * @example `portfolioHistoryManagement.init()`
     * @author GPT-6 Luna
     * @author benizzio
     * @author GPT-6 Sol
     */
    init() {

        // Repeated init calls keep one delegated asset-row listener on the document.
        document.addEventListener(ASSET_ROW_SELECTION_CHANGE_EVENT, handleAssetRowSelectionChange);
        const managementContainer = document.getElementById(PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER);

        if(managementContainer) {
            bindReadOnlyManagementMetadataHeaderToggle(managementContainer);
            bindObservationMetadataRequest(managementContainer);
        }

        const newObservationDateInput = managementContainer?.querySelector<HTMLInputElement>(
            `#${ NEW_OBSERVATION_DATE_INPUT_ID }`,
        );

        if(newObservationDateInput && newObservationDateInput.dataset.dateDefaultInitialized !== "true") {
            if(!newObservationDateInput.value) {
                newObservationDateInput.value = getBrowserLocalCalendarDate();
            }

            newObservationDateInput.dataset.dateDefaultInitialized = "true";
        }

        managementContainer
            ?.querySelectorAll<HTMLTableRowElement>(`tr[id^="${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }"]`)
            .forEach(normalizePortfolioHistoryRow);

        managementContainer
            ?.querySelectorAll<HTMLFormElement>(`form[id^="${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }"]`)
            .forEach(form => {
                AssetComposedColumnsInput.bindFormValidationGuards(form);
                this.bindObservationDetailsSubmissionGuard(form);
            });

        managementContainer?.querySelectorAll<HTMLElement>("[hx-trigger^='load-observation-details']")
            .forEach(header => {
                const item = header.closest<HTMLElement>(ACCORDION_ITEM_SELECTOR);
                const panel = item?.querySelector<HTMLElement>(".accordion-collapse");
                const form = item?.querySelector<HTMLFormElement>("form");

                if(panel && form) {
                    bindObservationDetailsLoader(header, panel, form);
                }
            });

        const managementFormRowTemplateElement = window["template-portfolio-history-management-form-tbody-row"];
        const managementTemplateElement = window["template-portfolio-history-management"];

        this.handlebarsPortfolioHistoryManagementRowTemplate = Handlebars.compile(
            managementFormRowTemplateElement.innerHTML,
        );

        this.handlebarsPortfolioHistoryManagementContainerTemplate = Handlebars.compile(
            managementTemplateElement.innerHTML,
        );
    },

    /** Adds one blank allocation row to the requested observation.
     *
     * @param observationTimestampId - Observation whose allocation table receives the row.
     * @example `portfolioHistoryManagement.addPortfolioHistoryManagementRow(12)`
     * @author GPT-6 Luna
     * @author benizzio
     */
    addPortfolioHistoryManagementRow(observationTimestampId: number) {

        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = document.getElementById(formId);

        if(!(form instanceof HTMLFormElement) || !areObservationDetailsReady(form)) {
            return;
        }

        const tbodyId = PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX + observationTimestampId;
        const tbody: HTMLElement = window[tbodyId];
        const nextIndex = getNextPortfolioHistoryManagementIndex(tbody);

        const newRowHtml = this.handlebarsPortfolioHistoryManagementRowTemplate({
            allocationIndex: nextIndex,
            observationTimestampId: observationTimestampId,
        });
        tbody.insertAdjacentHTML("beforeend", newRowHtml);

        const newRow = tbody.lastElementChild as HTMLElement;

        if(newRow instanceof HTMLTableRowElement) {
            normalizePortfolioHistoryRow(newRow);
        }

        focusOnNewLine(newRow);
        // Process the newly added row with htmx to enable bindings
        htmx.process(newRow);
        htmx.trigger(newRow, "htmx:afterSettle");
    },

    assetActionButtonClickHandler(formRowIndex: number, formUniqueId: string) {

        const formRowId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ formUniqueId }-row-${ formRowIndex }`;
        const assetIdHiddenFieldName = `allocations[${ formRowIndex }][assetId]`;
        const assetTickerFieldName = `allocations[${ formRowIndex }][assetTicker]`;
        const assetNameFieldName = `allocations[${ formRowIndex }][assetName]`;

        AssetComposedColumnsInput.assetActionButtonClickHandler(
            formRowId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
    },

    validateAssetElementsForPost(formRowIndex: number, formUniqueId: string) {

        const formRowId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ formUniqueId }-row-${ formRowIndex }`;
        const assetIdHiddenFieldName = `allocations[${ formRowIndex }][assetId]`;
        const assetTickerFieldName = `allocations[${ formRowIndex }][assetTicker]`;
        const assetNameFieldName = `allocations[${ formRowIndex }][assetName]`;

        AssetComposedColumnsInput.validateAssetElementsForPost(
            formRowId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
    },

    handleInputQuantityOrMarketPrice(formRowIndex: number, formUniqueId: string) {
        const rowValueElements = new FormRowValueElements(formUniqueId, formRowIndex);
        rowValueElements.handleInputQuantityOrMarketPrice();
    },

    handleInputTotalMarketValue(formRowIndex: number, formUniqueId: string) {
        const rowValueElements = new FormRowValueElements(formUniqueId, formRowIndex);
        rowValueElements.handleInputTotalMarketValue();
    },

    handleInputObservationTimeTag(newTimeTagInput: HTMLInputElement, observationTimestampId: number) {
        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = window[formId] as HTMLFormElement;
        const observationTimeTagInput = form.elements.namedItem("observationTimestamp.timeTag") as HTMLInputElement;
        observationTimeTagInput.value = newTimeTagInput.value;
    },

    /** Opens one persisted observation's metadata editor and focuses its tag field.
     *
     * @param button - Pencil button activated by pointer or keyboard.
     * @param observationTimestampId - Persisted observation whose fields should be edited.
     * @example `portfolioHistoryManagement.activateObservationEditor(button, 12)` opens observation 12's editor.
     * @author GPT-6 Luna
     * @author GPT-6.1 Sol
     */
    activateObservationEditor(button: HTMLButtonElement, observationTimestampId: number) {
        activateManagementMetadataEditor(
            document.getElementById(`portfolio-history-management-container-${ observationTimestampId }`),
            "[data-observation-editor-tag]",
            "[data-observation-editor-date]",
            button,
        );
    },

    /** Retries loading the requested observation's allocation rows after a failed lazy request.
     *
     * @param observationTimestampId - Persisted observation whose detail request should be retried.
     * @example `portfolioHistoryManagement.retryObservationDetails(12)` retries observation 12's load.
     * @author GPT-6 Luna
     */
    retryObservationDetails(observationTimestampId: number) {
        const header = document.getElementById(`portfolio-history-management-trigger-${ observationTimestampId }`);

        if(header) {
            retryObservationDetails(header);
        }
    },

    /** Prevents form writes until the persisted observation's allocation details have loaded.
     *
     * @param form - Observation form receiving native and HTMX submission guards.
     * @example `portfolioHistoryManagement.bindObservationDetailsSubmissionGuard(form)` protects a persisted form.
     * @author GPT-6 Luna
     */
    bindObservationDetailsSubmissionGuard(form: HTMLFormElement) {
        if(form.dataset.observationDetailsGuardBound === "true") {
            return;
        }

        form.dataset.observationDetailsGuardBound = "true";

        form.addEventListener("submit", event => {
            if(areObservationDetailsReady(form)) {
                return;
            }

            event.preventDefault();
            event.stopImmediatePropagation();
        }, true);
    },

    /** Reloads the observation only after a successful POST dispatched by that exact observation form.
     *
     * @param event - HTMX after-request event bubbled through the observation form.
     * @param observationTimestampId - Observation whose exact form issued the POST.
     * @example `portfolioHistoryManagement.handleAfterPostObservationHistory(event, 12)`
     * @author GPT-6 Luna
     * @author benizzio
     */
    handleAfterPostObservationHistory(event: CustomEvent, observationTimestampId: number) {

        const eventDetail = event.detail as AfterRequestEventDetail;
        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = document.getElementById(formId);

        if(!eventDetail.successful
            || eventDetail.requestConfig.verb.toLowerCase() !== "post"
            || !(form instanceof HTMLFormElement)
            || eventDetail.requestConfig.elt !== form) {
            return;
        }

        const metadataChanged = form.dataset.submittedMetadataChanged === "true";
        propagateRefreshDataAfterPost(observationTimestampId, metadataChanged);

        notifications.notify({
            title: "Success",
            content: "Portfolio observation data saved successfully.",
            type: NotificationType.SUCCESS,
        });
    },

    navigateToPortfolioAllocationViewing() {
        const globalPortfolioIdField = document.querySelector("[name='portfolioId']") as HTMLInputElement;
        const portfolioId = globalPortfolioIdField.value;
        Router.navigateTo(`/portfolio/${ portfolioId }/history`);
    },
};

export default portfolioHistoryManagement;
