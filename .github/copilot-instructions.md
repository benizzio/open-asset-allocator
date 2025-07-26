# Code syntax preferences

> [!IMPORTANT]
> If the prompt does not mention the intent to change the code files, do not generate any changes and just print an
> example code.

## Language agnostic standards

- use [clean code principles](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29)
    - give special attention to decomposing code into smaller functions
- follow [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
    - before generation a function, look for existing functions that can be reused (more references in below language
      sections)
- be [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)
    - when a block of code gets too big (not a hard rule, but around more than 3 instructions with multiple lines), the
      human eye perceives it better if it is divided into smaller contextual blocks separated by blank lines
    - when dividing a code unit into blocks, if it is a function, a blank line after the function declaration and before
      the first instruction is preferred, as it makes it more readable
    - Examples:
      ```go 
      
      // bad example
      func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {
          var portfolioIdParam = context.Param(portfolioIdParam)
          portfolioId, err := strconv.Atoi(portfolioIdParam)
          var planType = allocation.AssetAllocationPlan
          allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(
              portfolioId,
              &planType,
          )
          if infra.HandleAPIError(context, "Error getting allocation plans", err) {
              return
          }
          var allocationPlansDTS = model.MapToAllocationPlanDTSs(allocationPlans)
          context.JSON(http.StatusOK, allocationPlansDTS)
      }
      
      // good example
      func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {
      
          var portfolioIdParam = context.Param(portfolioIdParam)
          portfolioId, err := strconv.Atoi(portfolioIdParam)
      
          var planType = allocation.AssetAllocationPlan
          allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(
              portfolioId,
              &planType,
          )
          if infra.HandleAPIError(context, "Error getting allocation plans", err) {
              return
          }
      
          var allocationPlansDTS = model.MapToAllocationPlanDTSs(allocationPlans)
      
          context.JSON(http.StatusOK, allocationPlansDTS)
      }
      ```
- follow [Domain Driven Design](https://www.infoq.com/minibooks/domain-driven-design-quickly/) principles
- **all AI generated code**:
    - must contain proper minimal code comment documentation according to the language standards,
    - public API code (as in usable in other packages or modules) must contain very detailed usage instructions. It must
      also contain authoring documentation.
    - Examples:
      ```go
      // private function:

      // sendValidationErrorResponse sends a standardized HTTP response for validation validationErrors.
      //
      // Authored by: GitHub Copilot
      func sendValidationErrorResponse(context *gin.Context, errorMessages []string) {
      <...>
      }
    
      // public function:
      
      // GetStructNamespaceDescription extracts field name and builds the full namespace
      //
      // from a struct and field namespace string.
      // Parameters:
      //   - targetStruct: The struct or pointer to struct
      //   - fieldNamespace: The field namespace (can be a simple field name or a dot-separated path)
      //
      // Returns:
      //   - namespace: The full namespace including the struct name
      //   - fieldName: The simple field name (last part of the namespace)
      //
      // Authored by: GitHub Copilot
      func GetStructNamespaceDescription(targetStruct interface{}, fieldNamespace string) (namespace, fieldName string) {
      <...>
      }
      ```

> [!IMPORTANT]
> Code reviews must evaluate source code in all laguages cited below.

## Go standards

- all go code is located in the `src/main/go` directory
- when declaring a variable, give preference to `var` over `:=` as it is more explicit and more similar to other
  languages
    - exception, multiple varible declarations with reusage, e.g.:
        ```go
        err := doSomething()
        <...>
        result, err := doSomethingElse()
        ```
- most of the project's generic, reusable code can be found in the following listed packages. New code should be, in
  general, attentive to those packages to be DRY.
    - [infra](src/main/go/infra): represents the DDD infrastructure layer, and includes a lot of stack and utilitary
      code;
    - [inttestinfra](src/main/go/inttest/infra): represents the DDD infrastructure layer specific for integration tests,
      and includes a lot of stack and utilitary code;
    - [langext](src/main/go/langext): includes implementations that extend the Go language and are not available in the
      standard implementations at the time of writing.
- all go integration tests and related components are located in the `src/main/go/inttest` directory

## Javascript, Typescript, and HTML standards

- all javascript, typescript, and HTML code is located in the `src/main/web-static` directory

## SQL standards

### Migrations

- all SQL migration files are located in the `src/main/flyway/migrations` directory
- all SQL migration commands generated by AI must have a comment with the authoring information