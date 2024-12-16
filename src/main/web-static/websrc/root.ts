import * as clientSideTemplates from "htmx-ext-client-side-templates";
import * as bootstrap from "bootstrap";
import router from "./infra/routing/router";
import chart from "./infra/chart/chart";
import { handlebars } from "./infra/handlebars/handlebars";

//eslint-disable-next-line
const browserGlobal = (window as any);

browserGlobal.router = router;
browserGlobal.Handlebars = handlebars.register();
browserGlobal.chart = chart;

const onPageLoad = () => {
    router.init();
};
document.addEventListener("DOMContentLoaded", onPageLoad);

export default { clientSideTemplates, bootstrap };
