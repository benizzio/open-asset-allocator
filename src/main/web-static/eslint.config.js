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

            // Restrict deep imports under domain/service and application; enforce public API imports only
            // Authored by: GitHub Copilot
            "no-restricted-imports": [
                "error",
                {
                    patterns: [
                        // domain/service public API only
                        {
                            group: [
                                "**/domain/service/**",
                                "domain/service/**",
                                "./domain/service/**",
                                "../domain/service/**",
                            ],
                            message:
                                "Import from the public API 'domain/service' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // application public API only (outside the module)
                        {
                            group: [
                                "**/application/**",
                                "application/**",
                                "./application/**",
                                "../application/**",
                            ],
                            message:
                                "Import from the public API 'application' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into the local 'infra/handlebars' module; import only from its public API
                        // Authored by: GitHub Copilot
                        {
                            group: [
                                // generic patterns
                                "**/infra/handlebars/*",
                                "**/infra/handlebars/**",
                                // common relative forms
                                "infra/handlebars/*",
                                "infra/handlebars/**",
                                "./infra/handlebars/*",
                                "./infra/handlebars/**",
                                "../infra/handlebars/*",
                                "../infra/handlebars/**",
                            ],
                            message:
                                "Import from the public API 'infra/handlebars' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into the local 'infra/dom' module; import only from its public API
                        // Authored by: GitHub Copilot
                        {
                            group: [
                                // generic patterns
                                "**/infra/dom/*",
                                "**/infra/dom/**",
                                // common relative forms
                                "infra/dom/*",
                                "infra/dom/**",
                                "./infra/dom/*",
                                "./infra/dom/**",
                                "../infra/dom/*",
                                "../infra/dom/**",
                            ],
                            message:
                                "Import from the public API 'infra/dom' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                    ],
                },
            ],

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