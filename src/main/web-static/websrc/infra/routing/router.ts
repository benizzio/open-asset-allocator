import { bindHTMXTriggerOnRouteInDescendants } from "./binding-htmx-trigger-on-route";
import { bindNavigateToInDescendants } from "./binding-dom-navigate-to";
import { bindDisplayOnRouteInDescendants } from "./binding-dom-display-on-route";
import { bindAttributeOnRouteInDescendants } from "./binding-dom-attribute-on-route";

const router = {
    init() {
        router.bindDocumentToRouting();
        // router.initNavigoOnBrowserRoute();
    },
    // initNavigoOnBrowserRoute() {
    //     navigoRouter.resolve();
    // },
    bindDescendants(element: HTMLElement) {
        bindHTMXTriggerOnRouteInDescendants(element);
        bindNavigateToInDescendants(element);
        bindAttributeOnRouteInDescendants(element);
        bindDisplayOnRouteInDescendants(element);
    },
    bindDocumentToRouting() {
        this.bindDescendants(document.body);
    },
};

export default router;