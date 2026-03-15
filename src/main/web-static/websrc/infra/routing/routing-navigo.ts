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

/**
 * Navigates to the given route path using the Navigo router. All navigation calls should go through this function.
 *
 * @param path - The destination route path to navigate to.
 *
 * @author benizzio
 * @author GitHub Copilot
 */
export function navigateToRoute(path: string) {
    navigoRouter.navigate(path);
    routerBooted = true;
}

export function bootNavigoRouter() {
    if(!routerBooted) {
        const currentLocation = navigoRouter.getCurrentLocation().url;
        navigateToRoute(currentLocation);
    }
}
