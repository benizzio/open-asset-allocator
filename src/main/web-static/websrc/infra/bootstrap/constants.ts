/**
 * Bootstrap CSS class constants to avoid hard-coded strings throughout the application.
 *
 * This module provides centralized access to Bootstrap class names, making the code
 * more maintainable and reducing the risk of typos.
 *
 * Authored by: GitHub Copilot
 */
export const BootstrapClasses = {
    // Form validation classes
    WAS_VALIDATED: "was-validated",
    NEEDS_VALIDATION: "needs-validation",

    // Button classes
    BUTTON_PRIMARY: "btn btn-primary",
    BUTTON_DANGER: "btn btn-danger",
} as const;

export const BootstrapIconClasses = {
    SEARCH: "bi bi-search",
    RESET: "bi bi-x-circle",
} as const;