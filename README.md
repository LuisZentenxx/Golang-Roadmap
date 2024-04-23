# Golang-Roadmap

¡Bienvenido al roadmap de Go! Este repositorio está diseñado para ayudarte a aprender Go desde cero hasta un nivel intermedio, cubriendo lo básico y algunas características más avanzadas del lenguaje.

## Contenidos

- [Golang-Roadmap](#golang-roadmap)
  - [Contenidos](#contenidos)
  - [Introducción a Go](#introduccion-a-go)
  - [Tipos de Datos](#tipos-de-datos)
  - [Variables y Declaración](#variables-y-declaración)
  - [Condicionales](#condicionales)
  - [Loops](#loops)
  - [Funciones](#funciones)
  - [Colecciones](#colecciones)
  - [Package & Modulos](#package--modulos)
  - [Concurrencia](#concurrencia)
  - [Uso](#uso)
  - [Contribuciones](#contribuciones)
  - [Recursos Adicionales](#recursos-adicionales)

## Introduccion a Go

En esta sección, cubriremos los conceptos básicos de Go y cómo configurar tu entorno de desarrollo.

## Tipos de Datos

En Go, hay varios tipos de datos que se utilizan para almacenar diferentes tipos de valores. Algunos de los tipos de datos más comunes incluyen:

- **int:** Usado para almacenar enteros, como 1, -5, 100, etc.
- **cadena:** Se utiliza para almacenar cadenas de texto, como "Hola Mundo", "Gopher", etc.
- **byte:** Es un alias del tipo `uint8` y se utiliza para almacenar datos binarios o caracteres ASCII.
- **uint:** Representa enteros sin signo, como 0, 1, 2, etc.
- **complex:** Se utiliza para representar números complejos, como 3 + 2i.
- **bool:** Se utiliza para representar valores booleanos, es decir, verdadero (`true`) o falso (`false`).
- **float:** Se utiliza para representar números de coma flotante, como 3,14, -0,5, etc.

Estos son sólo algunos de los tipos de datos disponibles en Go. Cada tipo de dato tiene un tamaño y rango específico, y se utilizan para diferentes propósitos en programación.

## Variables y Declaración

En esta sección, aprenderemos sobre las variables y cómo declararlas en Go.

En Go, puedes declarar una variable usando la palabra clave `var`, seguida del nombre de la variable y su tipo.

## Condicionales

## Loops

## Funciones

## Colecciones

## Package & Modulos

En Go, los paquetes son una forma de organizar y reutilizar código. Un paquete consiste en uno o más archivos fuente Go que juntos proporcionan un conjunto de funciones, tipos y variables relacionadas.

#### Paquetes

- **Paquetes de Librería Estándar**: Go viene con una rica librería estándar de paquetes para tareas comunes como entrada/salida, redes y concurrencia.
- **Paquetes personalizados**: Los desarrolladores pueden crear sus **propios paquetes** para encapsular funcionalidad y promover la reutilización de código dentro de sus proyectos.

#### Módulos

Los módulos en Go proporcionan una forma de gestionar las dependencias y el versionado de tus proyectos. Un módulo es una colección de paquetes Go relacionados que se versionan juntos como una sola unidad.

- Archivo **go.mod**: Cada módulo Go está definido por un archivo `go.mod`, que especifica el nombre del módulo, la versión y las dependencias.
- **Versiones**: Los módulos permiten a los desarrolladores especificar las versiones de las dependencias requeridas por su proyecto, asegurando construcciones consistentes en diferentes entornos.
- **Gestión de dependencias**: Los módulos Go simplifican la gestión de dependencias descargando y gestionando automáticamente las dependencias basadas en el archivo `go.mod`.

### Ejemplo:

Consideremos un proyecto que requiere el uso de una librería cliente HTTP de terceros. Especificando la librería como una dependencia en el fichero `go.mod`, los módulos Go se asegurarán de que la versión correcta de la librería es descargada y usada en el proyecto.

```go
// go.mod
module example.com/myproject

go 1.15

require (
  github.com/ThirdParty/library v1.2.3
)
```

## Concurrencia

La concurrencia es una característica fundamental de Go que permite la ejecución simultánea de múltiples cálculos. Go proporciona un soporte robusto para la concurrencia a través de goroutines y canales.

#### Goroutines

- **Goroutines**: Las goroutines son hilos de ejecución ligeros gestionados por el tiempo de ejecución de Go. Permiten la ejecución concurrente de funciones sin la sobrecarga de los hilos tradicionales del sistema operativo.
- **Concurrencia**: Las Goroutines permiten la ejecución concurrente de múltiples funciones, haciendo posible la realización de tareas de forma asíncrona o en paralelo.
- **Sintaxis sencilla**: Las Goroutines se crean usando la palabra clave `go` seguida de una llamada a una función, haciendo que la concurrencia sea fácil de implementar en los programas Go.

```go
func main() {
  // Comienza una nueva Goroutines
  go doSomething()

  // Continúa ejecutando otro código simultáneamente 
}

func doSomething() {
  // Realizar alguna tarea de forma concurrente
}
```

#### Canales

**Channels**: Los canales son un potente medio de comunicación y sincronización entre goroutines en Go.
**Comunicación tipada**: Los canales facilitan el intercambio de datos tipados entre goroutines, permitiendo una comunicación segura sin condiciones de carrera.
**Sincronización**: Los canales se pueden utilizar para sincronizar la ejecución de goroutines, asegurando que ciertas operaciones se realizan en un orden específico.
**Operaciones de bloqueo**: El envío o recepción de datos en un canal se bloqueará hasta que otra goroutine esté lista para recibir o enviar, proporcionando un medio simple y eficiente de sincronización.

> NOTA: Los canales son una potente característica de Go que permite una comunicación segura y eficiente entre goroutines. Mediante el uso de canales, los desarrolladores pueden escribir programas concurrentes que son fáciles de entender y mantener.

## Uso

1. Clona este repositorio: `git clone https://github.com/LuisZentenxx/Golang-Roadmap.git`.
2. Explora los archivos y directorios para acceder a la información de cada tema.
3. Sigue el roadmap paso a paso para aprender Go de forma efectiva.

## Contribuciones

¡Las contribuciones son bienvenidas! Si encuentras errores, quieres añadir más información o mejorar el contenido, ¡no dudes en hacerlo! Abre un PR y estaremos encantados de revisarlo.

## Recursos Adicionales

- [Documentación oficial de Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/): Ejemplos prácticos de Go.
- [A Tour of Go](https://tour.golang.org/): Una introducción interactiva al lenguaje Go.
- [Awesome Go](https://github.com/avelino/awesome-go): Una lista de recursos y herramientas Go.

---

Disfruta aprendiendo Go y ¡feliz programación!



