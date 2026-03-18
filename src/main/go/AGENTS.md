# Specific agent instructions for the back-end and Go language

Refer to the general instructions in the root `../../../AGENTS.md` for broader instructions.

## Go language standards

- when declaring a variable, give preference to `var` over `:=` as it is more explicit and more similar to other
  languages
    - **exception**: multiple variable declarations with reusage, e.g.:
        ```go
        err := doSomething()
        <...>
        result, err := doSomethingElse()
        ```
- do not follow Godoc convention of adding a comment for every function, type, variable, etc. Clean code has priority
- most of the project's generic, reusable code can be found in the following listed packages. New code should be, in
  general, attentive to those packages to be DRY.
    - `src/main/go/infra`: represents the DDD infrastructure layer, and includes a lot of stack and utility code;
    - `src/main/go/inttest`: integration tests
        - `src/main/go/inttest/infra`: represents the DDD infrastructure layer specific for integration tests, and
          includes a lot of stack and utility code;
    - `src/main/go/langext`: includes implementations that extend the Go language and are not available in the standard
      implementations at the time of writing.