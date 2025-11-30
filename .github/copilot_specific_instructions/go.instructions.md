# Go standards

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

## Code style standards beyond linters

- EXAMPLE REF: CODE TOO BIG
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
- EXAMPLE REF: CODE DOCS
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