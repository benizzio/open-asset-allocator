<!--suppress HtmlUnknownTag -->

# Specific agent instructions for the back-end and Go language

Refer to the general instructions in the root `../../../AGENTS.md` for broader instructions.

## Go language standards

<CodingStandards>

- when declaring a variable, give preference to `var` over `:=` as it is more explicit and more similar to other
  languages
    - **exception**: multiple variable declarations with reusage, e.g.:
        ```go
        err := doSomething()
        <...>
        result, err := doSomethingElse()
        ```
- do not follow Godoc convention of adding a comment for every function, type, variable, etc. Clean code has priority
    - exception: AI generated code according to general instructions

</CodingStandards>

<CodeStructure>

- most of the project's generic, reusable code can be found in the following listed packages. New code should be, in
  general, attentive to those packages to be DRY.
    - `src/main/go/infra`: represents the DDD infrastructure layer, and includes a lot of stack and utility code;
    - `src/main/go/inttest`: integration tests
        - `src/main/go/inttest/infra`: represents the DDD infrastructure layer specific for integration tests, and
          includes a lot of stack and utility code;
    - `src/main/go/langext`: includes implementations that extend the Go language and are not available in the standard
      implementations at the time of writing.

</CodeStructure>

### Testing standards

- the project should always prioritize integration tests over unit tests. Unit tests should be written only for
  components where higher complexity justifies the need, or when specifically prompted for them
- integration tests should **ALWAYS**:
    - use JSON strings for input and output data assertion, to test parsing and improve readability
    - be written in a black-box style; exceptions if strictly necessary, should be explicitly confirmed with a new
      prompt
    - have no side effects, i.e. if the default state of any persistence is modified, it should be reverted to original
      in a cleanup step that runs regardless of the test result
    - be atomic, i.e. should never rely on the result of another test
    - try to rely on initial persistence data. If the existing initial data is not sufficient, then it should
      create its own, following the no side effects rule

#### Integration test structure

<CodeStructure>

- `src/main/go/inttest`: base integration test package
- `src/main/go/inttest/infra`: infrastructure needed for running integration tests, including initial db state
- `src/main/go/inttest/util`: general utilities for all integration tests

</CodeStructure>