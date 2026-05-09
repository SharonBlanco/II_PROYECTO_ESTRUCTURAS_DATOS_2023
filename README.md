# Sistema de Transporte - Grafos y Arboles

## Descripcion general

Este proyecto es un sistema de planificacion de rutas de transporte para la zona norte y central de Costa Rica. Utiliza un grafo para representar un mapa de localidades conectadas por carreteras, y un arbol para clasificar y generar estadisticas sobre las personas que consultan rutas.

El objetivo principal es encontrar la ruta mas corta entre dos lugares que ademas ofrezca la mayor cantidad de actividades turisticas segun las preferencias del usuario.

---

## Que hace el sistema

El programa tiene tres grandes funciones:

### Planificacion de rutas

El usuario selecciona un punto de partida, un destino y las categorias de actividades que le interesan (montana, acuatico, entretenimiento, cultural o naturaleza). El sistema calcula la ruta mas corta que pase por la mayor cantidad de actividades de esas categorias. Al finalizar la consulta, se registran los datos del usuario en un archivo JSON para futuras estadisticas.

### Gestion del grafo (mapa)

El sistema permite modificar el mapa en tiempo real a traves de un menu interactivo con las siguientes opciones:

- Agregar, modificar o eliminar lugares (vertices).
- Agregar, modificar o eliminar rutas entre lugares (arcos) con sus distancias en kilometros.
- Agregar, modificar o eliminar actividades turisticas asociadas a cada lugar.
- Consultar las actividades disponibles en cada lugar.
- Consultar la lista completa de actividades del sistema.

### Estadisticas generales

A partir de los datos almacenados en el archivo JSON, el sistema construye un arbol de clasificacion y genera estadisticas como:

- Porcentaje de hombres y mujeres en la poblacion registrada.
- Porcentaje de personas por lugar de residencia, separado por genero.
- Distribucion por rangos de edad (18-30, 31-64, mayor de 64).
- Porcentaje de personas que realizan una actividad especifica.

---

## Estructura del codigo

El programa esta escrito en un unico archivo Go y se organiza en las siguientes secciones:

### Estructuras de datos

| Estructura | Descripcion |
|---|---|
| Persona | Genero, edad, lugar de residencia y actividad preferida. |
| Actividad | Nombre y categoria (Montana, Acuatico, Entretenimiento, Cultural, Naturaleza). |
| Vertice | Un lugar en el mapa con nombre, lista de actividades y conexiones a otros lugares. |
| Arco | Una ruta entre dos lugares con su distancia en kilometros. |
| VerticeArbol | Un nodo del arbol de clasificacion con nombre, cantidad y conexiones. |
| ArcoArbol | Una conexion entre nodos del arbol. |

### Funciones principales

| Funcion | Que hace |
|---|---|
| cargarDatos | Crea los 10 vertices y 20+ arcos iniciales del grafo. |
| rutaCortaConCategorias | Algoritmo recursivo que encuentra la ruta mas corta con mas actividades de interes. |
| cargarArbol | Construye el arbol de clasificacion a partir de los datos del archivo JSON. |
| EstadisticasGenerales | Calcula y muestra porcentajes de genero, edad, residencia y actividades. |
| ModificarVertice | Menu principal interactivo con 13 opciones de gestion. |
| agregarPersona | Registra una nueva persona en el archivo JSON. |

---

## Mapa precargado

El grafo viene con 10 localidades conectadas entre si:

Santa Clara, San Ramon, Zarcero, Santa Rosa, Palmares, Naranjo, San Jose, Alajuela, Ciudad Quesada y Calle Blancos.

Cada localidad tiene actividades turisticas asignadas. Por ejemplo, Zarcero tiene Natacion y Senderismo; Ciudad Quesada tiene Termales, Ciclismo y Natacion; San Jose tiene Parque de Diversiones y Visitar Museos.

Las distancias entre localidades estan en kilometros (por ejemplo, Zarcero a San Ramon son 12 km, Alajuela a San Jose son 33 km).

---

## Actividades disponibles

El sistema incluye 14 actividades organizadas en 5 categorias:

| Categoria | Actividades |
|---|---|
| Montana | Senderismo, Volcanes, Ciclismo, Natacion, Canopy |
| Acuatico | Pesca, Termales |
| Entretenimiento | Cine, Parque de Diversiones |
| Cultural | Visitar Monumentos, Visitar Museos |
| Naturaleza | Paseos a Caballo, Mariposario, Miradero |

---

## Datos de personas

El archivo personas.json contiene 30 registros de personas con su genero, edad, lugar de residencia y actividad preferida. Estos datos se utilizan para construir el arbol de clasificacion y generar las estadisticas. Cada vez que un usuario consulta una ruta, sus datos se agregan automaticamente a este archivo.

---

## Tecnologias utilizadas

- Lenguaje: Go (Golang)
- Almacenamiento de datos: archivo JSON (personas.json)
- Interfaz: consola de texto con menu interactivo
- Estructuras de datos: grafo implementado con slices, arbol implementado con punteros

---

## Instrucciones de ejecucion

1. Tener Go instalado en el sistema.
2. Colocar los archivos Proyecto2ED.go, go.mod y personas.json en la misma carpeta.
3. Abrir una terminal en esa carpeta.
4. Ejecutar: `go run Proyecto2ED.go`
5. Seguir las opciones del menu interactivo (numeros del 0 al 13).

---

## Estructura de archivos

```
Proyecto2ED.go              Codigo fuente del programa (Go)
go.mod                      Archivo de modulo de Go
personas.json               Datos de 30 personas precargadas
DE_Proyecto2ED.docx         Documentacion del proyecto (analisis, solucion,
                            resultados, conclusiones y bitacora)
```

---

## Contexto academico

Proyecto desarrollado para el curso de Estructura de Datos (IC-2001), Tecnologico de Costa Rica, Campus San Carlos, segundo semestre 2023.
