import { logger, LogLevel } from "./logging";
import { APIErrorResponse } from "../api/api";

const InfraTypesUtils = {

    toErrorResponse(errorResponseJson: string): APIErrorResponse | undefined {

        try {

            const errorResponse = JSON.parse(errorResponseJson) as APIErrorResponse;

            const hasErrorProperty = errorResponse && typeof errorResponse.errorMessage === "string";

            const hasValidDetails = errorResponse
                && (
                    errorResponse.details === undefined
                    || (
                        Array.isArray(errorResponse.details)
                        && errorResponse.details.every((detail) => typeof detail === "string")
                    )
                );

            if(errorResponse && hasErrorProperty && hasValidDetails) {
                return errorResponse;
            }
        } catch(jsonParseError) {
            logger(LogLevel.WARN, "Failed to parse error response as JSON", jsonParseError);
        }

        return undefined;
    },
};

export default InfraTypesUtils;