/**
 * Manages portfolio-history observation forms and synchronizes each row's first external-asset association.
 *
 * @example `portfolioHistoryManagement.init()`
 * @author GPT-6 Luna
 * @author benizzio
 * @author GPT-6 Sol
 */

import htmx from "htmx.org";
import { BigNumber } from "bignumber.js";
import { AfterRequestEventDetail, HtmxInfra } from "../infra/htmx";
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

const PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER = "accordion-portfolio-history-management";
const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-history-management-form-tbody-";

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

function modifyObservationsResponse(originalServerResponseJSON: string): string {

    const originalServerResponse = JSON.parse(originalServerResponseJSON) as ObservationTimestamp[];

    const modifiedServerResponse: ObservationTimestamp[] = [
        {
            id: 0,
            timeTag: "New Observation *",
        },
        ...originalServerResponse,
    ];

    return JSON.stringify(modifiedServerResponse);
}

function propagateRefreshDataAfterPost(observationTimestampId: number) {

    if(observationTimestampId !== 0) {

        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = window[formId] as HTMLFormElement;
        form.reset();

        const observationManagementTriggerElement =
            window[`portfolio-history-management-trigger-${ observationTimestampId }`] as HTMLElement;
        htmx.trigger(observationManagementTriggerElement, "reload-portfolio-history-management-observation");
    }
    else {
        const portfolioHistoryManagementContainerElement =
            window[PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER] as HTMLElement;
        htmx.trigger(portfolioHistoryManagementContainerElement, "reload-portfolio-history-management");
    }

    const portfolioHistoryViewContainerElement = window["accordion-portfolio-history"];
    htmx.trigger(portfolioHistoryViewContainerElement, "reload-portfolio-history");

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

        // Repeated init calls use the same callback, so the document keeps one delegated listener.
        document.addEventListener(ASSET_ROW_SELECTION_CHANGE_EVENT, handleAssetRowSelectionChange);
        const managementContainer = document.getElementById(PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER);

        managementContainer
            ?.querySelectorAll<HTMLTableRowElement>(`tr[id^="${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }"]`)
            .forEach(normalizePortfolioHistoryRow);

        managementContainer
            ?.querySelectorAll<HTMLFormElement>(`form[id^="${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }"]`)
            .forEach(form => AssetComposedColumnsInput.bindFormValidationGuards(form));

        HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "addObservationZero",
            modifyObservationsResponse,
        );

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

        propagateRefreshDataAfterPost(observationTimestampId);

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
