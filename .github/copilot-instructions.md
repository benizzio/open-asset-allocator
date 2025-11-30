# Code syntax preferences

> [!IMPORTANT]
> If the prompt does not mention the intent to change the code files, do not generate any changes and just print example
> code blocks.

## Language agnostic standards

### Base standards and principles

- use [clean code principles](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29)
    - give special attention to decomposing code into smaller functions
    - always use descriptive names for functions, variables, and types
    - when uncertain, check Context7 with "Clean Code <name of the language>" as prompt
- be [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)
    - before generation a function, look for existing functions that can be reused (more references in below language
      sections)
- follow [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
- follow [Domain Driven Design](https://www.infoq.com/minibooks/domain-driven-design-quickly/) principles
    - when uncertain, check Context7 with "Evans DDD Sample" as prompt

### Code style standards beyond linters

- when a block of code gets too big (not a hard rule, but around more than 3 instructions with multiple lines), the
  human eye perceives it better if it is divided into smaller contextual blocks separated by blank lines
    - when dividing a code unit into blocks, if it is a function, a blank line after the function declaration and before
      the first instruction is preferred, as it makes it more readable
    - Examples: check `EXAMPLE REF: CODE TOO BIG` in language specific instructions below
- commented code can exist for transitioning code betyween PRs, but it is a red flag and should be pointed in the review
  comments to make sure it is necessary;
- **all AI generated code**:
    - must contain proper minimal code comment documentation according to the language standards,
    - public API code (as in usable in other packages or modules) must contain very detailed usage instructions. It must
      also contain authoring documentation.
    - if the laguage has authoring documentation standards, it must be followed
    - Examples: check `EXAMPLE REF: CODE DOCS` in language specific instructions below
- when implementing tests:
    - follow standard names `expected` and `actual` for variables used in the appropriate context

> [!IMPORTANT]
> Code reviews must evaluate source code in all laguages cited below.

## Go standards

- when handling go language code, read [go copilot instructions](copilot_specific_instructions/go.instructions.md)

## Browser code standards

- when handling code for browser environments (Javascript, Typescript, HTML, CSS and similar), read
  [browser copilot instructions](copilot_specific_instructions/browser.instructions.md)

## DB related code standards

- when handling database related code (SQL, PL/pgSQL, database migration files, ORM related code), read
  [db copilot instructions](copilot_specific_instructions/db.copilot.instructions.md)