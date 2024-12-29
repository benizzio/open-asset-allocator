import Navigo from "navigo";

export type HookCleanupFunction = (success?: boolean) => void;

export const navigoRouter = new Navigo("/", { strategy: "ALL" });
