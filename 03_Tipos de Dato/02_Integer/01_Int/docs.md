# Enteros (Integers) en Go

- **Definición:**
Los enteros en Go son un tipo de dato utilizado para representar números enteros, es decir, números sin parte decimal. Estos tipos de datos permiten almacenar valores numéricos enteros, ya sea con o sin signo, y están diseñados para manejar una variedad de rangos de valores.
Para qué sirve:
- **Para qué sirve:** Los enteros son fundamentales en Go para realizar cálculos matemáticos, llevar a cabo operaciones aritméticas y comparaciones de valores numéricos dentro del programa. Permiten a los desarrolladores trabajar con números enteros de diferentes magnitudes y especificar cómo se almacenan y manipulan estos valores en las variables, lo que garantiza la precisión y consistencia en las operaciones numéricas.

- **Sintaxis:** 
    - **Enteros con signo:**
        - `int`: Depende de la arquitectura del sistema, normalmente 32 o 64 bits.
        - `int8`: Rango de -128 a 127
        - `int16`: Rango de -32768 a 32767
        - `int32`: Rango de -2147483648 a 2147483647
        - `int64`: Rango de -9223372036854775808 a 9223372036854775807
        
    - **Enteros sin signo:**
        - `uint8`: Rango de 0 a 255
        - `uint16`: Rango de 0 a 65535
        - `uint32`: Rango de 0 a 4294967295
        - `uint64`: Rango de 0 a 18446744073709551615
- **Uso:** 
  Los enteros en Go se utilizan para definir variables y constantes que almacenan valores numéricos enteros. Estos valores se utilizan en operaciones aritméticas, comparaciones y otras manipulaciones numéricas dentro del programa.

- **Convenciones:** 
  Al trabajar con enteros en Go, es recomendable seguir estas convenciones:
    - Utilizar el tipo de dato más específico que satisfaga las necesidades del rango de valores esperados.
    - Asignar nombres descriptivos a las variables que indiquen claramente que almacenan valores enteros.
    - Ser cuidadoso al realizar conversiones de tipos para evitar pérdida de datos o desbordamientos.

        
### **Conversión de tipos (Type Conversion)**
    La conversión de tipos se realiza utilizando directamente el nombre de la variable como función para convertir tipos. Si se convierte a un tipo que tiene un rango inferior al actual, puede producirse pérdida de datos.
        
### **Desbordamiento (Overflow)**
    Si se asigna un tipo y luego se utiliza un número mayor que el rango del tipo para asignarlo, se producirá un error de desbordamiento.


#### Autor `Luis Zenteno`