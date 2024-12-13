import { navigoRouter } from "./routing-navigo";
import { bindHTMXEventOnRouteInDescendants } from "./binding-htmx";
import { bindNavigateToInDescendants } from "./binding-navigation";
import { bindAttributeOnRouteInDescendants, bindDisplayOnRouteInDescendants } from "./binding-dom";

const router = {
    init() {
        router.bindDocumentToRouting();
        router.initNavigoOnBrowserRoute();
    },
    initNavigoOnBrowserRoute() {
        navigoRouter.resolve();
    },
    bindDescendantsToRouting(element: HTMLElement) {
        bindHTMXEventOnRouteInDescendants(element);
        bindNavigateToInDescendants(element);
        bindAttributeOnRouteInDescendants(element);
        bindDisplayOnRouteInDescendants(element);
    },
    bindDocumentToRouting() {
        this.bindDescendantsToRouting(document.body);
    },
};

export default router;