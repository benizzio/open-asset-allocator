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
    calculateAndFormatPercent(value: BigNumber, total: BigNumber): string {
        return this.formatPercent(value.div(total).toNumber());
    },
};

export default format;