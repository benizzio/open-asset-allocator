export const PAUL_TOL_PALETTE = {
    qualitative: ["#4477AA", "#EE6677", "#228833", "#CCBB44", "#66CCEE", "#AA3377", "#BBBBBB"],
    qualitativeVibrant: ["#EE7733", "#0077BB", "#33BBEE", "#EE3377", "#CC3311", "#009988", "#BBBBBB"],
    qualitativeMuted: [
        "#CC6677",
        "#332288",
        "#DDCC77",
        "#117733",
        "#88CCEE",
        "#882255",
        "#44AA99",
        "#999933",
        "#AA4499",
    ],
    qualitativeMediumContrast: ["#6699CC", "#004488", "#EECC66", "#994455", "#997700", "#EE99AA"],
    qualitativePale: ["#BBCCEE", "#CCEEFF", "#CCDDAA", "#EEEEBB", "#FFCCCC", "#DDDDDD"],
    qualitativeDark: ["#222255", "#225555", "#225522", "#666633", "#663333", "#555555"],
    qualitativeLight: [
        "#77AADD",
        "#EE8866",
        "#EEDD88",
        "#FFAABB",
        "#99DDFF",
        "#44BB99",
        "#BBCC33",
        "#AAAA00",
        "#DDDDDD",
    ],
};

export const BOOTSTRAP_BODY_BACKGROUND_COLOR =
    getComputedStyle(document.documentElement).getPropertyValue("--bs-body-bg").trim();