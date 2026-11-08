/**
 * Public routing infrastructure for route-bound DOM behavior and route lifecycle subscriptions.
 * @module infra/routing
 * @author benizzio
 * @author GPT-6.1 Terra
 */
import { bindHTMXTriggerOnRouteInDescendants } from "./binding-htmx-trigger-on-route";
import { bindNavigateToInDescendants } from "./binding-dom-navigate-to";
import { bindDisplayOnRouteInDescendants } from "./binding-dom-display-on-route";
import { bindAttributeOnRouteInDescendants } from "./binding-dom-attribute-on-route";
import {
    bootNavigoRouter,
    buildParameterizedDestinationPathFromCurrentLocationContext,
    NAVIGO_PATH_PARAM_PREFIX,
    navigateToRoute,
    navigoRouter,
} from "./routing-navigo";

/** Subscribes to departure or a path-parameter change on a registered route without replacing its hooks.
 * @param route - Registered route pattern to observe.
 * @param onExit - Callback invoked when the route is left.
 * @returns A function that removes this subscription.
 * @example `const unsubscribe = subscribeRouteExit('/portfolio/:portfolioId/history/manage', cancel);`
 * @author GPT-6.1 Terra
 */
export function subscribeRouteExit(route: string, onExit: () => void): () => void {
    const initialPath = window.location.pathname;

    const removeLeave = navigoRouter.addLeaveHook(route, (done: (success?: boolean) => void) => {
        try {
            onExit();
        } finally {
            done();
        }
    }) as () => void;

    const removeAfter = navigoRouter.addAfterHook(route, () => {
        if(window.location.pathname !== initialPath) {
            onExit();
        }
    }) as () => void;

    return () => {
        removeLeave();
        removeAfter();
    };
}

const Router = {

    NAVIGO_PATH_PARAM_PREFIX,

    init() {
        Router.bindDocumentToRouting();
        globalThis["navigateTo"] = this.navigateTo;
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
        navigateToRoute(destinationPath);
    },

    buildParameterizedDestinationPathFromCurrentLocationContext,
};

export default Router;
