import { Infra } from "./infra/infra";
import Application from "./application";
import { AfterRequestEventDetail, HtmxInfra } from "./infra/htmx";
import notifications from "./components/notifications";

const AFTER_REQUEST_ERROR_HANDLER = (event: CustomEvent) => {

    const eventDetail = event.detail as AfterRequestEventDetail;

    if(!eventDetail.successful) {

        const errorResponse = HtmxInfra.toErrorResponse(eventDetail);

        if(errorResponse) {
            notifications.notifyErrorResponse(errorResponse);
        }
        else {
            const fallbackErrorMessage = "An unexpected error occurred while communicating with the server.";
            notifications.notifyErrorResponse({ errorMessage: fallbackErrorMessage });
        }

        return;
    }
};

const GENERAL_UNCAUGHT_ERROR_HANDLER = (error: Error) => {
    notifications.notifyError(error);
};

export default Infra.init(AFTER_REQUEST_ERROR_HANDLER, GENERAL_UNCAUGHT_ERROR_HANDLER);
Application.init();
