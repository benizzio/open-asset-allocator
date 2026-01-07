import { Notification, NotificationType } from "../infra/infra-types";
import * as handlebars from "handlebars";
import * as bootstrap from "bootstrap";
import DomInfra from "../infra/dom";
import { APIErrorResponse } from "../api/api";

const NotificationTypeBootstrapClasses = {
    info: "text-bg-primary",
    warning: "text-bg-warning",
    error: "text-bg-danger",
    success: "text-bg-success",
};

type BootstrapNotification = Notification & { contextClasses: string; };

function buildBootstrapNotification(notification: Notification): BootstrapNotification {
    const contextClasses =
        NotificationTypeBootstrapClasses[notification.type] || NotificationTypeBootstrapClasses.info;
    return { ...notification, contextClasses };
}

function buildBootstrapErrorNotification(errorResponse: APIErrorResponse): BootstrapNotification {

    const title = "Error";
    let content = DomInfra.DomUtils.escapeHtml(errorResponse.errorMessage ?? "");

    if(errorResponse.details && errorResponse.details.length > 0) {

        const detailsList = errorResponse.details.map(
            detail => `<li>${ DomInfra.DomUtils.escapeHtml(detail ?? "") }</li>`,
        ).join("");

        content += `<ul>${ detailsList }</ul>`;
    }

    return {
        title,
        content,
        type: NotificationType.ERROR,
        contextClasses: NotificationTypeBootstrapClasses.error,
    };
}

const notifications = {

    toastTemplate: null as HandlebarsTemplateDelegate,
    notificationsContainer: null as HTMLDivElement,

    init() {
        this.notificationsContainer = window["toast-notification-container"] as HTMLDivElement;
        const toastTemplateElement = window["template-toast-notification"];
        this.toastTemplate = handlebars.compile(toastTemplateElement.innerHTML);
    },

    notify(notification: Notification) {
        const bootstrapNotification = buildBootstrapNotification(notification);
        this._notifyAsToast(bootstrapNotification);
    },

    notifyErrorResponse(errorResponse: APIErrorResponse) {
        const bootstrapNotification = buildBootstrapErrorNotification(errorResponse);
        this._notifyAsToast(bootstrapNotification);
    },

    notifyError(error: Error) {
        const bootstrapNotification = buildBootstrapNotification({
            title: "Error",
            content: DomInfra.DomUtils.escapeHtml(error.message),
            type: NotificationType.ERROR,
        });
        this._notifyAsToast(bootstrapNotification);
    },

    _notifyAsToast(bootstrapNotification: Notification & { contextClasses: string }) {

        const notificationHTML = this.toastTemplate(bootstrapNotification);

        const notificationElement = document.createElement("div");
        notificationElement.innerHTML = notificationHTML;
        const toastElement = notificationElement.firstElementChild as HTMLElement;

        this.notificationsContainer.appendChild(toastElement);

        const toast = new bootstrap.Toast(toastElement, { delay: 5000 });
        toast.show();

        toastElement.addEventListener("hidden.bs.toast", () => {
            toastElement.remove();
        });
    },
};

export default notifications;