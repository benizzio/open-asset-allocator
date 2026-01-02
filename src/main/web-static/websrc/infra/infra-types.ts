export type CustomEventHandler = (event: CustomEvent) => void;

export enum NotificationType {
    INFO = "info",
    WARNING = "warning",
    ERROR = "error",
    SUCCESS = "success",
}

export type ErrorResponse = {
    error: string;
    details?: string[];
};

export type Notification = {
    title: string;
    content: string;
    type: NotificationType;
};