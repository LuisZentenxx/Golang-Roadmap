# Variables en Go

- **Definición:** 
  Una variable es un espacio de almacenamiento nombrado que se utiliza para
  contener valores. Estos valores pueden ser de diferentes tipos, como enteros, cadenas de
  texto, booleanos, etc.

- **Sintaxis:**
    - `var nombreVariable tipoDato`: Esta es la forma clásica de declarar una variable en Go. Se especifica el nombre de la variable seguido de su tipo de dato. La variable se declara pero no se inicializa, por lo que su valor inicial será el cero del tipo de dato especificado.
    
    - `var nombreVariable tipoDato = valor`: Esta forma permite declarar una variable y asignarle un valor inicial en el mismo paso. Se especifica el nombre de la variable, su tipo de dato y el valor que se le asignará.
    
    - `nombreVariable := valor`: Esta es una forma abreviada y conveniente de declarar e inicializar una variable en Go. Utilizando el operador `:=` el tipo de dato se infiere automáticamente del valor proporcionado. Esta forma es útil cuando se conoce el valor inicial de la variable y se desea que el tipo de dato se determine automáticamente.
- **Uso:**
  Las variables en Go se utilizan para almacenar y manipular datos durante la ejecución del programa. Pueden ser utilizadas para almacenar resultados intermedios, valores de entrada del usuario, valores de configuración, entre otros usos.

- **Convenciones:**
  En Go, se siguen ciertas convenciones al nombrar variables:
    - Los nombres de variables deben ser significativos y descriptivos.
    - Se utiliza el estilo CamelCase, donde la primera letra de cada palabra, excepto la primera, se escribe en mayúscula sin separación adicional (ejemplo: `nombreVariable`).
    - Para nombres de variables globales, se prefieren las iniciales en mayúscula (por ejemplo, `NombreVariable`).
    - Evite utilizar nombres de una sola letra o abreviaturas a menos que el contexto sea claro y específico.
