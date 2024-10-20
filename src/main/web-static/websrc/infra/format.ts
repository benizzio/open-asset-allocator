import BigNumber from "bignumber.js";

const DEFAULT_LOCALE: Intl.LocalesArgument = "en-US";
const DEFAULT_CURRENCY= "USD";

const format = {
    formatCurrency(value: number, locale = DEFAULT_LOCALE, currency = DEFAULT_CURRENCY): string {
        return value.toLocaleString(locale, { style: "currency", currency: currency });
    },
    formatPercent(value: number, locale = DEFAULT_LOCALE): string {
        return value.toLocaleString(locale, { style: "percent", maximumFractionDigits: 2 });
    },
    calculateAndFormatPercent(value: number, total: number): string {
        const valueBignumber = BigNumber(value);
        const totalBignumber = BigNumber(total);
        return this.formatPercent(valueBignumber.div(totalBignumber).toNumber());
    },
};

export default format;