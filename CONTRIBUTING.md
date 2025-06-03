# Contibuting guidelines for colaborators (including context for LLM code assistants and agents)

## Code of conduct (humans only)

Any collaboration to the project is welcomed if created on the intent to improve or advance the project. The author's
main goal is to solve a real world problem and help others solve theirs. Solving real problems means that one has to
understand them to a minimum reasonable depth, and the best practices for that are:

- to have that problem myself and take the steps to solve it myself. If you want to "eat your own dog food", you are in
  the correct place;
- to inform myself on the problem domain and chosen philosophy, principles and rules chosen to solve it. For that,
  make sure to read provided base documentantion and its references and make as many questions as needed or for more
  references. I am here to help;
- to listen to experience in that problem domain.

**Important PS**: This project contains explorations of technologies the author had not extensive experience on the time
of creation and wanted to learn about, but it is not a technology exploration project. It a technology or stack choice
demonstrates empirically to be unfit, it may be replaced, but that is under the author's discretion. The main objective
must always remain to solve a real problem.

## Code syntax preferences

### Language agnostic standards

- use [clean code principles](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29)
- follow [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
- be [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)
    - when a block of code gets too big (not a hard rule, but around more than 3 instructions with multiple lines), the
      human eye perceives it better if it is divided into smaller contextual blocks, which facilitates understanding, so
      use them. Examples:
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

### Go

- when declaring a variable, give preference to `var` over `:=` as it is more explicit and more similar to other
  languages
    - exception, multiple varible declarations with reusage, e.g.:
        ```go
        err := doSomething()
        <...>
        result, err := doSomethingElse()
        ```
