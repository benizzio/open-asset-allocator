import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import stylisticTs from "@stylistic/eslint-plugin-ts";

export default [

    { files: ["**/*.{js,mjs,cjs,ts}"] },

    { languageOptions: { globals: globals.browser } },

    pluginJs.configs.recommended,

    ...tseslint.configs.recommended,

    {
        plugins: { "@stylistic/ts": stylisticTs },
        // ESLint rules: https://eslint.org/docs/latest/rules/
        // ESLint Stylistic rules for typescript: https://eslint.style/packages/ts
        rules: {
            "max-len": ["warn", { code: 120, tabWidth: 4 }],
            "no-unused-vars": "warn",
            "@stylistic/ts/indent": ["warn", 4],
            "@stylistic/ts/semi": ["error", "always"],
            "@stylistic/ts/quotes": ["warn", "double"],
            "@stylistic/ts/quote-props": ["warn", "consistent-as-needed"],
            "@stylistic/ts/object-curly-newline": ["warn", { multiline: true }],
            "@stylistic/ts/object-curly-spacing": ["warn", "always"],
            "@stylistic/ts/comma-dangle": ["warn", "always-multiline"],
            "@stylistic/ts/comma-spacing": ["warn", { before: false, after: true }],
            "@stylistic/ts/key-spacing": ["warn", { beforeColon: false, afterColon: true }],
            "@stylistic/ts/padding-line-between-statements": [
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
];