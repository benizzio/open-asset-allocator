import { bindHTMXTriggerOnRouteInDescendants } from "./binding-htmx-trigger-on-route";
import { bindNavigateToInDescendants } from "./binding-dom-navigate-to";
import { bindDisplayOnRouteInDescendants } from "./binding-dom-display-on-route";
import { bindAttributeOnRouteInDescendants } from "./binding-dom-attribute-on-route";
import {
    bootNavigoRouter,
    buildParameterizedDestinationPathFromCurrentLocationContext,
    NAVIGO_PATH_PARAM_PREFIX,
    navigoRouter,
} from "./routing-navigo";

const Router = {

    NAVIGO_PATH_PARAM_PREFIX,

    init(browserGlobal: Window) {
        Router.bindDocumentToRouting();
        browserGlobal["navigateTo"] = this.navigateTo;
    },

    bindDescendants(element: HTMLElement) {
        bindHTMXTriggerOnRouteInDescendants(element);
        bindNavigateToInDescendants(element);
        bindAttributeOnRouteInDescendants(element);
        bindDisplayOnRouteInDescendants(element);
    },

    bindDocumentToRouting() {
        this.bindDescendants(document.body);
    },

    boot() {
        bootNavigoRouter();
    },

    navigateTo(destinationPath: string) {
        navigoRouter.navigate(destinationPath);
    },

    buildParameterizedDestinationPathFromCurrentLocationContext,
};

export default Router;