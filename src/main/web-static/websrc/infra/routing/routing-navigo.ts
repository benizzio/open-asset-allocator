import Navigo from "navigo";
import { logger, LogLevel } from "../logging";

export type HookCleanupFunction = (success?: boolean) => void;

export const navigoRouter = new Navigo("/", { strategy: "ALL" });

export const NAVIGO_PATH_PARAM_PREFIX = ":";

function hasPathParam(path: string) {
    return path.includes(NAVIGO_PATH_PARAM_PREFIX);
}

function resolvePathParamsFromCurrentLocationContext(path: string) {

    let resolvedPath = path;
    const currentLocation = navigoRouter.getCurrentLocation();
    const currentMatch = navigoRouter.match(currentLocation.url)[0];

    if(currentMatch) {
        for(const key in currentMatch.data) {
            resolvedPath = resolvedPath.replace(`${ NAVIGO_PATH_PARAM_PREFIX }${ key }`, currentMatch.data[key]);
        }
    }

    if(hasPathParam(resolvedPath)) {
        const errorMessage = "Could not find path param in current route navigate to";
        logger(LogLevel.ERROR, errorMessage, resolvedPath);
        throw new Error(errorMessage);
    }

    return resolvedPath;
}

export function buildParameterizedDestinationPathFromCurrentLocationContext(destinationPath: string) {
    if(hasPathParam(destinationPath)) {
        return resolvePathParamsFromCurrentLocationContext(destinationPath);
    }
    return destinationPath;
}

let routerBooted = false;

export function bootNavigoRouter() {
    if(!routerBooted) {
        const currentLocation = navigoRouter.getCurrentLocation().url;
        navigoRouter.navigate(currentLocation);
        routerBooted = true;
    }
}
