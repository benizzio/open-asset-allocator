import { ErrorResponse } from "./infra-types";
import { logger, LogLevel } from "./logging";

const InfraTypesUtils = {

    toErrorResponse(errorResponseJson: string): ErrorResponse | undefined {

        try {

            const errorResponse = JSON.parse(errorResponseJson) as ErrorResponse;

            if(errorResponse && typeof errorResponse.error === "string") {
                return errorResponse;
            }
        } catch(jsonParseError) {
            logger(LogLevel.WARN, "Failed to parse error response as JSON", jsonParseError);
        }

        return undefined;
    },
};

export default InfraTypesUtils;