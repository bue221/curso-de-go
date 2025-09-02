# Curso Profesional de Go

Este repositorio contiene un curso completo de programación en Go (Golang) que abarca desde conceptos básicos hasta temas avanzados como concurrencia y programación orientada a objetos.

## 📚 Estructura del Curso

El curso está organizado en 13 módulos progresivos que van desde lo más básico hasta conceptos avanzados:

### 1. Hola Mundo (`1-hola_mundo/`)

- Introducción básica a Go
- Uso de paquetes externos (`rsc.io/quote`)
- Estructura básica de un programa en Go

### 2. Variables (`2-variables/`)

- Declaración y asignación de variables
- Múltiples asignaciones
- Tipos de datos básicos

### 3. Constantes (`3-constantes/`)

- Declaración de constantes
- Constantes agrupadas
- Uso de `iota` para enumeraciones
- Convenciones de nomenclatura

### 4. Tipos de Datos (`4-tipos_datos/`)

- Tipos numéricos (int, float, complex)
- Tipos booleanos
- Tipos de texto (string, rune)
- Límites de tipos numéricos

### 5. Ejercicios (`5-ejercicios/`)

- **Ejercicio práctico**: Calculadora de triángulos rectángulos
- Cálculo de hipotenusa usando teorema de Pitágoras
- Cálculo de área y perímetro
- Uso de funciones matemáticas

### 6. Control de Flujos (`6-control-de-flujos/`)

- Estructuras condicionales (`if`, `else`)
- Estructuras de control (`switch`)
- Uso de paquete `time` para fechas

### 7. Random (`7-random/`)

- Generación de números aleatorios
- Juego de adivinanza de números
- Bucles y control de flujo interactivo

### 8. Estructuras de Datos (`8-estructura-de-datos/`)

- **Arrays**: declaración y manipulación
- **Slices**: segmentación y operaciones
- **Maps**: diccionarios y asociaciones clave-valor
- Iteración con `range`

### 9. Gestor de Tareas (`9-gestor-de-tareas/`)

- **Aplicación completa**: Sistema de gestión de tareas
- Estructuras (`struct`) y métodos
- Interfaz de usuario por consola
- Operaciones CRUD (Crear, Leer, Actualizar, Eliminar)

### 10. Manejo de Errores (`10-manejo-errores/`)

- Manejo explícito de errores
- Uso de `defer` para limpieza de recursos
- Creación y manipulación de archivos
- Conceptos de `panic` y recuperación

### 11. Módulos y Paquetes (`11-greatings/`)

- Creación de paquetes personalizados
- Exportación de funciones
- Validación de parámetros
- Estructura de módulos Go

### 12. Programación Orientada a Objetos (`12-POO/`)

- Interfaces y polimorfismo
- Composición vs herencia
- Métodos y receptores
- Ejemplo con animales (perros y gatos)

### 13. Concurrencia (`13-concurrencia/`)

- Programación concurrente básica
- Verificación de APIs en paralelo
- Medición de tiempo de ejecución
- Uso de `goroutines` y `channels`

## 🚀 Cómo Usar Este Curso

### Prerrequisitos

- Go 1.19 o superior instalado
- Conocimientos básicos de programación (recomendado)

### Ejecución

Cada módulo es independiente y puede ejecutarse por separado:

```bash
# Navegar a cualquier módulo
cd 1-hola_mundo

# Ejecutar el programa
go run hola.go

# O compilar y ejecutar
go build hola.go
./hola
```

### Módulos con Dependencias

Algunos módulos requieren dependencias externas:

```bash
# Para el módulo 1 (hola_mundo)
cd 1-hola_mundo
go mod tidy
go run hola.go

# Para el módulo 11 (greatings)
cd 11-greatings
go mod tidy
go test
```

## 🎯 Objetivos del Curso

Al completar este curso, habrás aprendido:

- ✅ Sintaxis básica y estructura de Go
- ✅ Tipos de datos y variables
- ✅ Control de flujo y estructuras de datos
- ✅ Funciones y manejo de errores
- ✅ Programación orientada a objetos
- ✅ Concurrencia y paralelismo
- ✅ Gestión de paquetes y módulos
- ✅ Desarrollo de aplicaciones completas

## 📁 Estructura de Archivos

```
curso-profesional-go/
├── 1-hola_mundo/          # Introducción básica
├── 2-variables/           # Variables y tipos
├── 3-constantes/          # Constantes y enumeraciones
├── 4-tipos_datos/         # Tipos de datos en Go
├── 5-ejercicios/          # Ejercicios prácticos
├── 6-control-de-flujos/   # Condicionales y bucles
├── 7-random/              # Números aleatorios
├── 8-estructura-de-datos/ # Arrays, slices, maps
├── 9-gestor-de-tareas/    # Aplicación completa
├── 10-manejo-errores/     # Manejo de errores
├── 11-greatings/          # Módulos y paquetes
├── 12-POO/                # Programación orientada a objetos
└── 13-concurrencia/       # Concurrencia y paralelismo
```

## 🔧 Tecnologías Utilizadas

- **Go (Golang)**: Lenguaje de programación principal
- **Go Modules**: Gestión de dependencias
- **Standard Library**: Paquetes estándar de Go
- **Testing**: Framework de pruebas integrado

## 📝 Notas Importantes

- El código está diseñado para ser simple y sin abstracciones innecesarias
- Cada ejemplo incluye comentarios explicativos
- Los ejercicios prácticos refuerzan los conceptos teóricos
- El curso progresa de manera gradual en complejidad

## 🤝 Contribuciones

Este es un curso educativo. Si encuentras errores o tienes sugerencias de mejora, no dudes en contribuir.

## 📄 Licencia

Este proyecto es de uso educativo y está disponible para fines de aprendizaje.

---

**¡Disfruta aprendiendo Go! 🐹**
