import { bindHTMXTriggerOnRouteInDescendants } from "./binding-htmx-trigger-on-route";
import { bindNavigateToInDescendants } from "./binding-dom-navigate-to";
import { bindDisplayOnRouteInDescendants } from "./binding-dom-display-on-route";
import { bindAttributeOnRouteInDescendants } from "./binding-dom-attribute-on-route";
import { bootNavigoRouter, navigoRouter } from "./routing-navigo";

const router = {
    init(browserGlobal: Window) {

        router.bindDocumentToRouting();

        browserGlobal["navigateTo"] = (destinationPath: string) => {
            navigoRouter.navigate(destinationPath);
        };
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
};

export default router;