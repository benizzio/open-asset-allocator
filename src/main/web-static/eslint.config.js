import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import stylistic from "@stylistic/eslint-plugin";

export default [

    { files: ["**/*.{js,jsx,mjs,cjs,ts,tsx}"] },

    { ignores: ["dist/**"] },

    { languageOptions: { globals: globals.browser } },

    pluginJs.configs.recommended,

    ...tseslint.configs.recommended,

    {
        plugins: { "@stylistic": stylistic },
        // ESLint rules: https://eslint.org/docs/latest/rules/
        // ESLint Stylistic rules: https://eslint.style/packages/default
        rules: {
            "max-len": ["warn", { code: 120, tabWidth: 4 }],
            // no-unused-vars rules are scoped per-file-type below
            "@stylistic/indent": ["warn", 4],
            "@stylistic/semi": ["error", "always"],
            "@stylistic/quotes": ["warn", "double"],
            "@stylistic/quote-props": ["warn", "consistent-as-needed"],
            "@stylistic/object-curly-newline": ["warn", { multiline: true }],
            "@stylistic/object-curly-spacing": ["warn", "always"],
            "@stylistic/comma-dangle": ["warn", "always-multiline"],
            "@stylistic/comma-spacing": ["warn", { before: false, after: true }],
            "@stylistic/key-spacing": ["warn", { beforeColon: false, afterColon: true }],
            "@stylistic/padding-line-between-statements": [
                "warn",
                { blankLine: "always", prev: "import", next: "*" },
                { blankLine: "never", prev: "import", next: "import" },
                { blankLine: "always", prev: "*", next: "multiline-expression" },
                { blankLine: "always", prev: "*", next: "multiline-block-like" },
                { blankLine: "always", prev: "*", next: "multiline-const" },
                { blankLine: "always", prev: "*", next: "multiline-let" },
                { blankLine: "always", prev: "*", next: "multiline-var" }],

        },
    },

    // Apply core no-unused-vars to JavaScript files
    {
        files: ["**/*.{js,jsx,mjs,cjs}"],
        rules: { "no-unused-vars": "warn" },
    },

    // Use TypeScript-specific no-unused-vars and disable the core rule for TS files
    {
        files: ["**/*.{ts,tsx}"],
        rules: { "no-unused-vars": "off", "@typescript-eslint/no-unused-vars": ["warn"] },
    },
];