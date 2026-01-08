import BigNumber from "bignumber.js";

const DEFAULT_LOCALE: Intl.LocalesArgument = "en-US";
const DEFAULT_CURRENCY = "USD";
const FINANCIAL_DECIMAL_DIGITS = 2;

const Format = {

    formatCurrency(value: number, locale = DEFAULT_LOCALE, currency = DEFAULT_CURRENCY): string {
        return value.toLocaleString(locale, { style: "currency", currency: currency });
    },

    formatPercent(value: number, locale = DEFAULT_LOCALE): string {
        return value.toLocaleString(locale, { style: "percent", maximumFractionDigits: 2 });
    },

    calculateAndFormatPercent(value: BigNumber, total: BigNumber): string {
        return this.formatPercent(value.div(total).toNumber());
    },

    /**
     * Formats a number as a financial value with locale-specific grouping separators.
     * Does not include currency symbol, only the formatted number.
     *
     * @param value The numeric value to format
     * @param locale The locale to use for formatting (default: "en-US")
     * @param decimalDigits The number of decimal digits to display (default: 2)
     * @returns The formatted number string with grouping separators
     *
     * @author GitHub Copilot
     */
    formatFinancialNumber(
        value: number,
        locale = DEFAULT_LOCALE,
        decimalDigits = FINANCIAL_DECIMAL_DIGITS,
    ): string {

        return value.toLocaleString(locale, {
            style: "decimal",
            minimumFractionDigits: decimalDigits,
            maximumFractionDigits: decimalDigits,
        });
    },

    /**
     * Parses a locale-formatted financial string back to a number.
     * Handles locale-specific grouping and decimal separators.
     *
     * @param formattedValue The formatted string to parse
     * @param locale The locale used for formatting (default: "en-US")
     * @returns The parsed number, or NaN if parsing fails
     *
     * @author GitHub Copilot
     */
    parseFinancialNumber(formattedValue: string, locale = DEFAULT_LOCALE): number {

        if(!formattedValue || formattedValue.trim() === "") {
            return NaN;
        }

        // Get locale-specific separators
        const parts = new Intl.NumberFormat(locale).formatToParts(1234.5);
        const groupSeparator = parts.find(part => part.type === "group")?.value || ",";
        const decimalSeparator = parts.find(part => part.type === "decimal")?.value || ".";

        // Remove grouping separators and replace decimal separator with standard "."
        let sanitized = formattedValue.replace(new RegExp(`\\${groupSeparator}`, "g"), "");

        if(decimalSeparator !== ".") {
            sanitized = sanitized.replace(new RegExp(`\\${decimalSeparator}`, "g"), ".");
        }

        return parseFloat(sanitized);
    },
};

export default Format;