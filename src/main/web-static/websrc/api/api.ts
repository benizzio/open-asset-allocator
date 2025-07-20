import { Asset } from "../domain/asset";

export type APIErrorResponse = {
    errorMessage: string;
    details?: string[];
};

class APIError extends Error {
    response: Response;
}

const api = {
    /**
     * Type guard for APIErrorResponse
     * @param obj
     *
     * Authored by: GitHub Copilot
     */
    isAPIErrorResponse(obj: unknown): obj is APIErrorResponse {
        return typeof obj === "object"
            && obj !== null
            && "errorMessage" in obj
            && Array.isArray((obj as APIErrorResponse).details);
    },
    getAsset: async(uniqueIdentifier: string): Promise<Asset | APIErrorResponse> => {

        const response = await fetch("/api/asset/" + uniqueIdentifier);

        if(!response.ok) {

            const error = new APIError("Network response was not ok");
            const contentType = response.headers.get("content-type");

            if(contentType && contentType.includes("application/json")) {
                return await response.json() as Promise<APIErrorResponse>;
            }

            throw error;
        }

        return await response.json() as Promise<Asset>;
    },
};

export default api;