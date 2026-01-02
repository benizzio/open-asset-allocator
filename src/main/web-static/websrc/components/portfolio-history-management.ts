import htmx from "htmx.org";
import BigNumber from "bignumber.js";
import { AfterRequestEventDetail, HtmxInfra } from "../infra/htmx";
import { ObservationTimestamp } from "../domain/portfolio-allocation";
import router from "../infra/routing/router";
import AssetComposedColumnsInput from "./asset-composed-columns-input";
import { toInt } from "../utils/lang";
import type { TemplateDelegate } from "handlebars";

const PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER = "accordion-portfolio-history-management";
const PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX = "portfolio-history-management-form-";
const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-history-management-form-tbody-";

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

        const quantity = this.quantityInput.value || 0;
        const marketPrice = this.marketPriceInput.value || 0;

        if(quantity && marketPrice) {
            const totalMarketValue = new BigNumber(quantity)
                .times(marketPrice)
                .decimalPlaces(0, BigNumber.ROUND_HALF_UP);
            this.totalMarketValueInput.value = totalMarketValue.toString();
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

    init() {
        HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "addObservationZero",
            modifyObservationsResponse,
        );
    },

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

        focusOnNewLine(newRow);
        // Process the newly added row with htmx to enable bindings
        htmx.process(newRow);
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

    handleAfterPostObservationHistory(event: CustomEvent, observationTimestampId: number) {

        const eventDetail = event.detail as AfterRequestEventDetail;

        if(!eventDetail.successful) {
            return;
        }

        propagateRefreshDataAfterPost(observationTimestampId);
    },

    navigateToPortfolioAllocationViewing() {
        const globalPortfolioIdField = document.querySelector("[name='portfolioId']") as HTMLInputElement;
        const portfolioId = globalPortfolioIdField.value;
        router.navigateTo(`/portfolio/${ portfolioId }/history`);
    },
};

export default portfolioHistoryManagement;
