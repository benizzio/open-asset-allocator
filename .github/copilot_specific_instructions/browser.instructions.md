## Javascript, Typescript, and HTML standards

- all javascript, typescript, and HTML code is located in the `src/main/web-static` directory

## Code style standards beyond linters

- EXAMPLE REF: CODE TOO BIG
    ```typescript
        // bad example
        function handleHierarchicalIdLevelChange(targetElement: HTMLInputElement) {
            const ancestorTable = targetElement.closest("form");
            if(!ancestorTable) {
                return;
            }
            const targetElementName = targetElement.getAttribute("name");
            const fieldsToUpdate =
                ancestorTable.querySelectorAll<HTMLInputElement>(`[data-bind-to-name$='${ targetElementName }']`);
            fieldsToUpdate.forEach((field) => {
                field.value = targetElement.value;
            });
            const spansToUpdate =
                ancestorTable.querySelectorAll<HTMLSpanElement>(`[data-label-for-name='${ targetElementName }']`);
            spansToUpdate.forEach((span) => {
                span.textContent = targetElement.value;
            });
        };
  
        // good example
        function handleHierarchicalIdLevelChange(targetElement: HTMLInputElement) {

            const ancestorTable = targetElement.closest("form");
        
            if(!ancestorTable) {
                return;
            }
        
            const targetElementName = targetElement.getAttribute("name");
        
            const fieldsToUpdate =
                ancestorTable.querySelectorAll<HTMLInputElement>(`[data-bind-to-name$='${ targetElementName }']`);
            fieldsToUpdate.forEach((field) => {
                field.value = targetElement.value;
            });
        
            const spansToUpdate =
                ancestorTable.querySelectorAll<HTMLSpanElement>(`[data-label-for-name='${ targetElementName }']`);
            spansToUpdate.forEach((span) => {
                span.textContent = targetElement.value;
            });
        };
    ```
- EXAMPLE REF: CODE DOCS
  ```typescript
      // private function:
      /**
       * Checks if a value is serializable to JSON.
       * 
       * @author GitHub Copilot
       */
      function isSerializable(value: unknown): boolean {}
  
      // public function:
      /**
       * Converts a JavaScript value to a JSON string.
       *
       * @param object - The value to serialize.
       * @returns A JSON string representation of the value.
       * 
       * @example
       * <pre>{{stringify data}}</pre>
       * <span>{{stringify user}}</span>
       * 
       * @author GitHub Copilot
       */
      function stringifyHelper(object: unknown): string {}
  ```