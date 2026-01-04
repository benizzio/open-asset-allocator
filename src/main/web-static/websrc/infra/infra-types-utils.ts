import { ErrorResponse } from "./infra-types";
import { logger, LogLevel } from "./logging";

const InfraTypesUtils = {

    toErrorResponse(errorResponseJson: string): ErrorResponse | undefined {

        try {

            const errorResponse = JSON.parse(errorResponseJson) as ErrorResponse;

            const hasErrorProperty = errorResponse && typeof errorResponse.error === "string";

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