import htmx from "htmx.org";

const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-allocation-management-rows-";

const portfolioHistoryManagement = {

    handlebarPortfolioHistoryManagementRowTemplate: null,

    addPortfolioHistoryManagementRow(observationTimestampId: string) {

        const tbodyId = PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX + observationTimestampId;
        const tbody: HTMLElement = window[tbodyId];
        const nextIndex = this.getNextPortfolioHistoryManagementIndex(tbody);

        const newRowHtml = this.handlebarPortfolioHistoryManagementRowTemplate({
            allocationIndex: nextIndex,
            observationTimestampId: observationTimestampId,
        });
        tbody.insertAdjacentHTML("beforeend", newRowHtml);

        const newRow = tbody.lastElementChild;

        this.focusOnNewLine(newRow);
        // Process the newly added row with htmx to enable bindings
        htmx.process(newRow);
    },

    focusOnNewLine(newRow: HTMLElement) {

        const firstFieldOfNewRow: HTMLInputElement = newRow.querySelector("input[type='text']");

        if(firstFieldOfNewRow) {
            firstFieldOfNewRow.scrollIntoView({ behavior: "smooth", block: "center" });
            firstFieldOfNewRow.focus();
        }
    },

    getNextPortfolioHistoryManagementIndex(tbody: HTMLElement): number {
        const rows = tbody.querySelectorAll("tr");
        return rows.length;
    },
};

export default portfolioHistoryManagement;