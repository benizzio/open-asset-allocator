/**
 * ESLint flat configuration for frontend code, including module boundaries and code-quality rules.
 * @author OpenCode
 * @author benizzio
 * @author GPT-6 Luna
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */
import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import stylistic from "@stylistic/eslint-plugin";
import sonarjs from "eslint-plugin-sonarjs";

export default [

    { files: ["**/*.{js,jsx,mjs,cjs,ts,tsx}"] },

    { ignores: ["dist/**"] },

    { languageOptions: { globals: globals.browser } },

    pluginJs.configs.recommended,

    ...tseslint.configs.recommended,

    {
        plugins: { "@stylistic": stylistic, sonarjs },
        // ESLint rules: https://eslint.org/docs/latest/rules/
        // ESLint Stylistic rules: https://eslint.style/packages/default
        rules: {
            "sonarjs/cognitive-complexity": ["error", 15],
            "sonarjs/expression-complexity": ["error", { max: 3 }],
            "sonarjs/no-duplicate-string": ["warn", { threshold: 3 }],
            "sonarjs/no-nested-conditional": "error",
            "@typescript-eslint/prefer-optional-chain": "warn",
            "@typescript-eslint/prefer-nullish-coalescing": "warn",
            "max-len": ["warn", { code: 120, tabWidth: 4 }],
            "max-lines": ["warn", { max: 600 }],
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
            "@stylistic/no-trailing-spaces": ["error"],
            "@stylistic/keyword-spacing": [
                "warn",
                { overrides: { if: { after: false }, for: { after: false }, switch: { after: false } } },
            ],
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
                        // domain/asset public API only
                        // Authored by: OpenCode
                        {
                            group: [
                                "**/domain/asset/**",
                                "domain/asset/**",
                                "./domain/asset/**",
                                "../domain/asset/**",
                            ],
                            message:
                                "Import from the public API 'domain/asset' (index.ts) only; " +
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
                        // pages public API only
                        // Authored by: OpenCode
                        {
                            group: [
                                "**/pages/**",
                                "pages/**",
                                "./pages/**",
                                "../pages/**",
                            ],
                            message:
                                "Import from the public API 'pages' (index.ts) only; " +
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
                        // Block deep imports into the local 'infra/htmx' module; import only from its public API
                        // Authored by: GitHub Copilot
                        {
                            group: [
                                // generic patterns
                                "**/infra/htmx/*",
                                "**/infra/htmx/**",
                                // common relative forms
                                "infra/htmx/*",
                                "infra/htmx/**",
                                "./infra/htmx/*",
                                "./infra/htmx/**",
                                "../infra/htmx/*",
                                "../infra/htmx/**",
                            ],
                            message:
                                "Import from the public API 'infra/htmx' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into the local 'infra/routing' module; import only from its public API
                        // Authored by: GitHub Copilot
                        {
                            group: [
                                // generic patterns
                                "**/infra/routing/*",
                                "**/infra/routing/**",
                                // common relative forms
                                "infra/routing/*",
                                "infra/routing/**",
                                "./infra/routing/*",
                                "./infra/routing/**",
                                "../infra/routing/*",
                                "../infra/routing/**",
                            ],
                            message:
                                "Import from the public API 'infra/routing' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into the composed asset-column module while leaving its local imports alone
                        // Authored by: GPT-6 Luna
                        {
                            group: [
                                "**/asset-composed-columns-input/*",
                                "**/asset-composed-columns-input/**",
                                "asset-composed-columns-input/*",
                                "asset-composed-columns-input/**",
                                "./asset-composed-columns-input/*",
                                "./asset-composed-columns-input/**",
                                "../asset-composed-columns-input/*",
                                "../asset-composed-columns-input/**",
                            ],
                            message:
                                "Import from the public API 'components/asset-composed-columns-input' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into feature modules from outside their public APIs.
                        // Authored by: GPT-6.1 Terra
                        {
                            group: [
                                "**/components/allocation-plan/**", "components/allocation-plan/**",
                                "../components/allocation-plan/**", "../allocation-plan/**",
                            ],
                            message: "Import from the public API 'components/allocation-plan' (index.ts) only.",
                        },
                        {
                            group: [
                                "**/components/portfolio-history/**", "components/portfolio-history/**",
                                "../components/portfolio-history/**", "../portfolio-history/**",
                            ],
                            message: "Import from the public API 'components/portfolio-history' (index.ts) only.",
                        },
                        // Block deep imports into the portfolio-history-quote module; use its public entry point
                        // Authored by: GPT-6 Luna
                        {
                            group: [
                                "**/portfolio-history-quote/*",
                                "**/portfolio-history-quote/**",
                                "portfolio-history-quote/*",
                                "portfolio-history-quote/**",
                                "./portfolio-history-quote/*",
                                "./portfolio-history-quote/**",
                                "../portfolio-history-quote/*",
                                "../portfolio-history-quote/**",
                            ],
                            message:
                                "Import from the public API 'components/portfolio-history/portfolio-history-quote' " +
                                "(index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into the management metadata editor; use its public entry point
                        // Authored by: GPT-6.1 Sol
                        {
                            group: [
                                "**/management-metadata-editor/*",
                                "**/management-metadata-editor/**",
                                "management-metadata-editor/*",
                                "management-metadata-editor/**",
                                "./management-metadata-editor/*",
                                "./management-metadata-editor/**",
                                "../management-metadata-editor/*",
                                "../management-metadata-editor/**",
                            ],
                            message:
                                "Import from the public API 'components/management-metadata-editor' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into allocation plan management metadata; use its public entry point
                        // Authored by: GPT-6.1 Sol
                        {
                            group: [
                                "**/allocation-plan-management-metadata/*",
                                "**/allocation-plan-management-metadata/**",
                                "allocation-plan-management-metadata/*",
                                "allocation-plan-management-metadata/**",
                                "./allocation-plan-management-metadata/*",
                                "./allocation-plan-management-metadata/**",
                                "../allocation-plan-management-metadata/*",
                                "../allocation-plan-management-metadata/**",
                            ],
                            message:
                                "Import from the public API 'components/allocation-plan/allocation-plan-management-metadata' " +
                                "(index.ts) only; deep imports are not allowed.",
                        },
                        // Block deep imports into allocation-plan management response mapping; use its public API
                        // Authored by: GPT-6 Luna
                        {
                            group: [
                                "**/allocation-plan-management-response/*",
                                "**/allocation-plan-management-response/**",
                                "allocation-plan-management-response/*",
                                "allocation-plan-management-response/**",
                                "./allocation-plan-management-response/*",
                                "./allocation-plan-management-response/**",
                                "../allocation-plan-management-response/*",
                                "../allocation-plan-management-response/**",
                            ],
                            message:
                                "Import from the public API 'components/allocation-plan/allocation-plan-management-response' " +
                                "(index.ts) only; deep imports are not allowed.",
                        },
                        // Block deep imports into shared paginated-list behavior; use its public entry point
                        // Authored by: GPT-6 Luna
                        {
                            group: [
                                "**/paginated-list/*",
                                "**/paginated-list/**",
                                "paginated-list/*",
                                "paginated-list/**",
                                "./paginated-list/*",
                                "./paginated-list/**",
                                "../paginated-list/*",
                                "../paginated-list/**",
                            ],
                            message:
                                "Import from the public API 'components/paginated-list' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Block deep imports into chart infrastructure; use its public entry point
                        // Authored by: GPT-6 Luna
                        {
                            group: [
                                "**/infra/chart/chart",
                                "**/infra/chart/chart.ts",
                                "./chart/chart",
                                "./chart/chart.ts",
                            ],
                            message: "Import from the public API 'infra/chart' (index.ts) only; deep imports are not allowed.",
                        },
                        // Block deep imports into history observation details; use its public entry point
                        // Authored by: GPT-6.1 Sol
                        {
                            group: [
                                "**/portfolio-history-management-observation-details/*",
                                "**/portfolio-history-management-observation-details/**",
                                "portfolio-history-management-observation-details/*",
                                "portfolio-history-management-observation-details/**",
                                "./portfolio-history-management-observation-details/*",
                                "./portfolio-history-management-observation-details/**",
                                "../portfolio-history-management-observation-details/*",
                                "../portfolio-history-management-observation-details/**",
                            ],
                            message:
                                "Import from the public API " +
                                "'components/portfolio-history/portfolio-history-management-observation-details' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                        // Observation-clone internals are private to the clone module.
                        // Authored by: GPT-6.1 Terra
                        {
                            group: [
                                "**/portfolio-history-management-observation-clone/*",
                                "**/portfolio-history-management-observation-clone/**",
                                "./portfolio-history-management-observation-clone/*",
                                "../portfolio-history-management-observation-clone/*",
                            ],
                            message: "Import observation cloning through its index.ts public API only.",
                        },
                        // Form identifiers have a dependency-free public boundary.
                        // Authored by: GPT-6.1 Sol
                        {
                            group: ["**/portfolio-history-management-form/*", "**/portfolio-history-management-form/**"],
                            message: "Import history form identifiers through its index.ts public API only.",
                        },
                        // Block deep imports into history observation metadata; use its public entry point
                        // Authored by: GPT-6.1 Sol
                        {
                            group: [
                                "**/portfolio-history-management-observation-metadata/*",
                                "**/portfolio-history-management-observation-metadata/**",
                                "portfolio-history-management-observation-metadata/*",
                                "portfolio-history-management-observation-metadata/**",
                                "./portfolio-history-management-observation-metadata/*",
                                "./portfolio-history-management-observation-metadata/**",
                                "../portfolio-history-management-observation-metadata/*",
                                "../portfolio-history-management-observation-metadata/**",
                            ],
                            message:
                                "Import from the public API " +
                                "'components/portfolio-history/portfolio-history-management-observation-metadata' (index.ts) only; " +
                                "deep imports are not allowed.",
                        },
                    ],
                },
            ],

        },
    },

    // Restrict sibling feature internals reached from nested components without affecting domain imports.
    // Authored by: GPT-6.1 Terra
    {
        files: ["websrc/components/**/*.{js,jsx,mjs,cjs,ts,tsx}"],
        ignores: ["websrc/components/allocation-plan/**"],
        plugins: { allocationPlanBoundary: tseslint.plugin },
        rules: {
            "allocationPlanBoundary/no-restricted-imports": ["error", {
                patterns: [{
                    regex: "^(?:\\.\\./){2,}allocation-plan/",
                    message: "Import from the public API 'components/allocation-plan' (index.ts) only.",
                }],
            }],
        },
    },
    {
        files: ["websrc/components/**/*.{js,jsx,mjs,cjs,ts,tsx}"],
        ignores: ["websrc/components/portfolio-history/**"],
        rules: {
            "@typescript-eslint/no-restricted-imports": ["error", {
                patterns: [{
                    regex: "^(?:\\.\\./){2,}portfolio-history/",
                    message: "Import from the public API 'components/portfolio-history' (index.ts) only.",
                }],
            }],
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
        languageOptions: {
            parserOptions: {
                project: "./tsconfig.eslint.json",
                tsconfigRootDir: import.meta.dirname,
            },
        },
        rules: { "no-unused-vars": "off", "@typescript-eslint/no-unused-vars": ["warn"] },
    },
];
