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
 * characters (both uppercase and lowercase) and the special characters "-", "_", ":", and ".",
 * removing any other character.
 *
 * @param inputElement The HTML input element to apply the mask to
 *
 * Authored by: GitHub Copilot
 */
export function maskTickerInput(inputElement: HTMLInputElement): void {

    let value = inputElement.value;

    if(value) {
        // Remove any character that is not alphanumeric (uppercase or lowercase), "-", "_", ":", or "."
        value = value.replace(/[^A-Za-z0-9\-_:.]/g, "");
    }

    inputElement.value = value;
}

/**
 * Applies a mask to number inputs to limit decimal places to a maximum of 8.
 * Blocks typing additional decimal digits beyond the 8th decimal place.
 *
 * @param inputElement The HTML input element to apply the mask to
 * @param decimalPlaces The maximum number of decimal places allowed (default is 8)
 *
 * Co-authored by: GitHub Copilot
 */
export function maskNumberDecimalPlaces(inputElement: HTMLInputElement, decimalPlaces: number = 8): void {

    if(inputElement.value && inputElement.value.includes(".")) {
        const parts = inputElement.value.split(".");

        if(parts[1] && parts[1].length > decimalPlaces) {
            inputElement.value = parts[0] + "." + parts[1].substring(0, decimalPlaces);
        }
    }
}
