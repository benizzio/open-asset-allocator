import { Portfolio, PortfolioDTO } from "../domain/portfolio";
import { DomainService } from "../domain/service";
import htmx from "htmx.org";
import Router from "../infra/routing/router";

const PortfolioPage = {

    getContextPortfolio(): Portfolio {
        const portfolioElement = window["portfolio"] as HTMLScriptElement;
        const portfolioJSON = portfolioElement.textContent;
        const portfolioDTO = JSON.parse(portfolioJSON) as PortfolioDTO;
        return DomainService.mapping.mapToPortfolio(portfolioDTO);
    },

    reloadPortfolio(event: CustomEvent) {

        if(!event.detail?.successful) {
            return;
        }

        const requestConfig = event.detail?.requestConfig;
        const form = event.target as HTMLFormElement;

        const portfolioId =
            (requestConfig?.parameters && (requestConfig.parameters.id ?? requestConfig.parameters["id"])) ??
            new FormData(form).get("id");

        if(form.checkValidity() && portfolioId) {

            form.reset();

            const portfolioElement = document.querySelector("#portfolio-context");

            htmx.trigger(portfolioElement, "reload-portfolio", { routerPathData: { portfolioId: portfolioId } });

            Router.navigateTo(`/portfolio/${ portfolioId }`);
        }
    },

    preFillEditPortfolioForm() {

        const portfolioDataElement = window["portfolio-context"] as HTMLDivElement;
        const portfolioIdField = portfolioDataElement.querySelector("input[name='portfolioId']") as HTMLInputElement;

        const portfolionNameField =
            portfolioDataElement.querySelector("input[name='portfolioName']") as HTMLInputElement;

        const portfolioId = portfolioIdField.value;
        const portfolionName = portfolionNameField.value;

        const editPortfolioForm = window["edit-portfolio-form"] as HTMLFormElement;
        const editPortfolioIdInput = editPortfolioForm.querySelector("input[name='id']") as HTMLInputElement;
        editPortfolioIdInput.value = portfolioId;
        const editPortfolioNameInput = editPortfolioForm.querySelector("input[name='name']") as HTMLInputElement;
        editPortfolioNameInput.value = portfolionName;
    },
};

export default PortfolioPage;