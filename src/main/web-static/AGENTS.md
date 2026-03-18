# Specific agent instructions for the front-end and JavaScript, TypeScript, HTML and SCSS languages

Refer to the general instructions in the root `../../../AGENTS.md` for broader instructions.

## Front-end and browser code standards

- the code is structured in the following modules
    - `src/main/web-static/websrc`: the main code for the web application
        - `src/main/web-static/websrc/api`: code related to API calls. Almost all API calls should be handled by HTMX so
          this should only include exceptional cases that need to be handled by JavaScript or TypeScript code
        - `src/main/web-static/websrc/pages`: base templates for the application UI
        - `src/main/web-static/websrc/components`: segmented templates and controller code for parts of the pages that
          form the application UI
        - `src/main/web-static/websrc/application`: code related to the application logic that allow the usage of
          the domain logic tied to the SPA infrastructure and UI
        - `src/main/web-static/websrc/domain`: domain logic code, including domain entities and specific functionality
          to manipulate them
        - `src/main/web-static/websrc/infra`: infrastructure code that allows the usage of the front end stack
          components and libraries (browser APIs, HTMX, Handlebars, etc.) with their own APIs and HTML element
          bindings
        - `src/main/web-static/websrc/utils`: generic utility code that can be used accross the other modules, including
          the infra modules. Base JavaScript helper functions and API wrappers should be implemented here
        - `src/main/web-static/websrc/stylesheet`: SCSS code for the application stylesheets
- except for generic utilities in `src/main/web-static/websrc/utils`, modules should only expose their APIs through the
  `index.ts`, and direct access should be blocked. This should be enforced by
  `src/main/web-static/websrc/infra/eslint.config.mjs`.