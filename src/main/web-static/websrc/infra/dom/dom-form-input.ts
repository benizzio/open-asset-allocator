/**
 * Applies a complete input mask that only allows uppercase alphanumeric
 * characters and the special characters "-" and "_", removing any other character.
 *
 * @param inputElement The HTML input element to apply the mask to
 *
 * Authored by: GitHub Copilot
 */
export function maskTagInput(inputElement: HTMLInputElement): void {

    let value = inputElement.value;

    if(value) {
        // Convert to uppercase and remove any character that is not alphanumeric, "-", or "_"
        value = value.toUpperCase().replace(/[^A-Z0-9\-_]/g, "");
    }

    inputElement.value = value;
}

/**
 * Applies a complete input mask that only allows alphanumeric
 * characters (both uppercase and lowercase) and the special characters "-", "_", and ":", removing any other character.
 *
 * @param inputElement The HTML input element to apply the mask to
 *
 * Authored by: GitHub Copilot
 */
export function maskTickerInput(inputElement: HTMLInputElement): void {

    let value = inputElement.value;

    if(value) {
        // Remove any character that is not alphanumeric (uppercase or lowercase), "-", "_", or ":"
        value = value.replace(/[^A-Za-z0-9\-_:]/g, "");
    }

    inputElement.value = value;
}
