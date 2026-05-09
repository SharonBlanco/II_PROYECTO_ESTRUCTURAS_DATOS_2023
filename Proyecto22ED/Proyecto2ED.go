// Paquete main proporciona una breve descripción del propósito de este programa en Go.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

// La estructura Persona representa una persona con género, edad, residencia y detalles de actividad.
type Persona struct {
	Genero     string    // Género de la persona.
	Edad       uint8     // Edad de la persona.
	Residencia string    // Residencia de la persona.
	Actividad  Actividad // Detalles de la actividad de la persona.
}

// La estructura Actividad representa una actividad con nombre y categoría.
type Actividad struct {
	Nombre    string // Nombre de la actividad.
	Categoria string // Categoría de la actividad.
}

// La estructura Arco representa una arista en el grafo con vértice de origen, vértice de destino y distancia.
type Arco struct {
	Origen    *Vertice // Vértice de origen de la arista.
	Destino   *Vertice // Vértice de destino de la arista.
	Distancia int      // Distancia entre los vértices de origen y destino.
}

// La estructura Vertice representa un vértice en el grafo con nombre, lista de actividades, aristas conectadas y estado de visita.
type Vertice struct {
	Nombre      string       // Nombre del vértice.
	Actividades []*Actividad // Lista de actividades asociadas al vértice.
	ListaArcos  []*Arco      // Lista de aristas conectadas al vértice.
	visitado    bool         // Estado de visita del vértice.
}

// La estructura VerticeArbol representa un vértice en el árbol con nombre, cantidad, siguiente vértice y lista de aristas conectadas.
type VerticeArbol struct {
	nombre        string        // Nombre del vértice en el árbol.
	cantidad      int           // Cantidad de ocurrencias del vértice.
	sigV          *VerticeArbol // Siguiente vértice en el árbol.
	subListaArcos *ArcoArbol    // Lista de aristas conectadas al vértice en el árbol.
}

// La estructura ArcoArbol representa una arista en el árbol con vértice de origen, vértice de destino y siguiente arista en la lista.
type ArcoArbol struct {
	origen  *VerticeArbol // Vértice de origen de la arista en el árbol.
	destino *VerticeArbol // Vértice de destino de la arista en el árbol.
	sigA    *ArcoArbol    // Siguiente arista en la lista.
}

// ListaActividades es una variable global que representa una lista de actividades.
var ListaActividades []Actividad

// grafo es una variable global que representa un grafo con una lista de vértices.
var grafo []Vertice

// arbol es una variable global que representa un árbol con un vértice raíz.
var arbol *VerticeArbol

// ImprimirMenu imprime un menú basado en la lista de vértices proporcionada en el grafo.
// Recibe un slice de Vertice y muestra el índice y nombre de cada vértice en el menú.
func ImprimirMenu(grafo []Vertice) {
	fmt.Println()
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Println("Menu")
	for i := 0; i < len(grafo); i++ {
		fmt.Println(i, ".", grafo[i].Nombre)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------------------")
}

// traductorMenu traduce un número a un nombre específico de lugar basado en su posición en el menú.
// Recibe un byte que representa la posición en el menú y devuelve el nombre correspondiente del lugar.
func traductorMenu(numero byte) string {
	switch numero {
	case 0:
		return "Santa Clara"
	case 1:
		return "San Ramon"
	case 2:
		return "Zarcero"
	case 3:
		return "Santa Rosa"
	case 4:
		return "Palmares"
	case 5:
		return "Naranjo"
	case 6:
		return "San Jose"
	case 7:
		return "Alajuela"
	case 8:
		return "Ciudad Quesada"
	case 9:
		return "Calle Blancos"
	default:
		return "Número no válido"
	}
}

// traductorCategorias traduce un número a una categoría específica basada en su valor.
// Recibe un byte que representa la categoría y devuelve el nombre correspondiente de la categoría.
func traductorCategorias(numero byte) string {
	switch numero {
	case 1:
		return "Montaña"
	case 2:
		return "Acuatico"
	case 3:
		return "Entretenimiento"
	case 4:
		return "Cultural"
	case 5:
		return "Naturaleza"
	default:
		return "Número no válido"
	}
}

// traductorActividades traduce un número a una actividad específica basada en su valor.
// Recibe un byte que representa la actividad y devuelve el nombre correspondiente de la actividad.
func traductorActividades(numero byte) string {
	switch numero {
	case 1:
		return "Senderismo"
	case 2:
		return "Pesca"
	case 3:
		return "Termales"
	case 4:
		return "Volcanes"
	case 5:
		return "Paseos a caballo"
	case 6:
		return "Mariposario"
	case 7:
		return "Ciclismo"
	case 8:
		return "Miradero"
	case 9:
		return "Natacion"
	case 10:
		return "Cine"
	case 11:
		return "Parque de Diversiones"
	case 12:
		return "Visitar monumentos"
	case 13:
		return "Visitar museos"
	case 14:
		return "Canopy"
	default:
		return "Número no válido"
	}
}

// cargarDatos es una función que carga datos predefinidos en el grafo proporcionado.
// Recibe un puntero a un slice de Vertice y agrega vértices y aristas según las conexiones especificadas.
func cargarDatos(grafo *[]Vertice) {
	CrearVertice("Santa Clara", grafo)
	CrearVertice("San Ramon", grafo)
	CrearVertice("Zarcero", grafo)
	CrearVertice("Santa Rosa", grafo)
	CrearVertice("Palmares", grafo)
	CrearVertice("Naranjo", grafo)
	CrearVertice("San Jose", grafo)
	CrearVertice("Alajuela", grafo)
	CrearVertice("Ciudad Quesada", grafo)
	CrearVertice("Calle Blancos", grafo)

	insertarArco("Zarcero", 12, "San Ramon", grafo)
	insertarArco("San Ramon", 12, "Zarcero", grafo)
	insertarArco("Zarcero", 23, "Santa Rosa", grafo)
	insertarArco("Santa Rosa", 23, "Zarcero", grafo)

	insertarArco("Naranjo", 32, "San Ramon", grafo)
	insertarArco("San Ramon", 32, "Naranjo", grafo)
	insertarArco("Palmares", 25, "San Ramon", grafo)
	insertarArco("San Ramon", 25, "Palmares", grafo)

	insertarArco("Palmares", 31, "Naranjo", grafo)
	insertarArco("Naranjo", 31, "Palmares", grafo)
	insertarArco("Palmares", 40, "Alajuela", grafo)
	insertarArco("Alajuela", 40, "Palmares", grafo)
	insertarArco("Palmares", 34, "Ciudad Quesada", grafo)
	insertarArco("Ciudad Quesada", 34, "Palmares", grafo)
	insertarArco("Palmares", 18, "Ciudad Quesada", grafo)
	insertarArco("Ciudad Quesada", 18, "Palmares", grafo)
	insertarArco("Palmares", 68, "Santa Rosa", grafo)
	insertarArco("Santa Rosa", 68, "Palmares", grafo)

	insertarArco("Naranjo", 57, "San Jose", grafo)
	insertarArco("San Jose", 57, "Naranjo", grafo)

	insertarArco("Alajuela", 33, "San Jose", grafo)
	insertarArco("San Jose", 33, "Alajuela", grafo)

	insertarArco("Alajuela", 8, "Ciudad Quesada", grafo)
	insertarArco("Ciudad Quesada", 8, "Alajuela", grafo)

	insertarArco("Santa Clara", 13, "Ciudad Quesada", grafo)
	insertarArco("Ciudad Quesada", 13, "Santa Clara", grafo)
	insertarArco("Santa Clara", 22, "San Jose", grafo)
	insertarArco("San Jose", 22, "Santa Clara", grafo)
	insertarArco("Calle Blancos", 11, "San Jose", grafo)
	insertarArco("San Jose", 11, "Calle Blancos", grafo)
	insertarArco("Calle Blancos", 8, "Santa Clara", grafo)
	insertarArco("Santa Clara", 8, "Calle Blancos", grafo)

	insertarArco("Santa Rosa", 38, "Ciudad Quesada", grafo)
	insertarArco("Ciudad Quesada", 38, "Santa Rosa", grafo)
}

// CrearVertice crea un nuevo vértice con el nombre especificado y lo agrega al grafo.
// Recibe el nombre del vértice como un string y un puntero a un slice de Vertice (grafo).
func CrearVertice(nombre string, grafo *[]Vertice) {
	vertice := Vertice{Nombre: nombre} // Crea un nuevo vértice con el nombre especificado.
	*grafo = append(*grafo, vertice)   // Agrega el vértice al grafo usando el puntero al slice de Vertice.
}

// ModificarVertice permite realizar diversas operaciones en un grafo y una lista de actividades.
// El usuario puede elegir entre varias opciones para modificar un vértice y las actividades asociadas.
// Parámetros:
//   - nombre: Nombre del vértice que se desea modificar.
//   - grafo: Puntero a un slice de estructuras Vertice que representa el grafo.
//   - listaActividades: Puntero a un slice de estructuras Actividad que contiene la lista de actividades.
func ModificarVertice(grafo *[]Vertice, listaActividades *[]Actividad) {
	// Inicialización del lector de entrada estándar.
	userInput := bufio.NewReader(os.Stdin)
	// Buscar el vértice en el grafo.
	//vertice := buscarVertice(nombre, grafo)

	// Mostrar las opciones disponibles al usuario.
	/*fmt.Println("1. Agregar lugar (vértice) \n2. Modificar nombre del lugar (vértice) \n3. 2. Modificar nombre del lugar (vértice) " +
	"\n4. Agregar actividades al lugar (vértice) \n5. Modificar actividades del lugar (vértice) \n6. Borrar actividades del lugar (vértice) " +
	"\n7. Agregar ruta (arco) \n8. Modificar ruta (arco) \n9. Borrar ruta (arco) \n10. Consultar ruta \n11. Ver estadísticas generales" +
	"\n12. Ver lista de actividades \n13. Ver actividades por cada vértice")*/

	/*// Verificar si el vértice existe en el grafo.
	if vertice == nil {
		fmt.Println("El vértice no existe")
		return
	}*/

loop:
	// Bucle infinito para mantener el programa en ejecución.
	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("Menu")
		fmt.Println("1. Agregar lugar (vértice)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("2. Modificar nombre del lugar (vértice)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("3. Eliminar nombre del lugar (vértice)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("4. Agregar actividades al lugar (vértice)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("5. Modificar actividades")
		fmt.Print("\033[H\033[2J")
		fmt.Println("6. Borrar actividades del lugar (vértice)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("7. Agregar ruta (arco)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("8. Modificar ruta (arco)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("9. Borrar ruta (arco)")
		fmt.Print("\033[H\033[2J")
		fmt.Println("10. Consultar ruta")
		fmt.Print("\033[H\033[2J")
		fmt.Println("11. Ver estadísticas generales")
		fmt.Print("\033[H\033[2J")
		fmt.Println("12. Ver lista de actividades")
		fmt.Print("\033[H\033[2J")
		fmt.Println("13. Ver actividades por cada vértice")
		fmt.Print("\033[H\033[2J")
		fmt.Println("0. Para salir del programa")
		fmt.Print("\033[H\033[2J")

		var feature byte
		// Leer la opción del usuario desde la entrada estándar.
		for {
			_, err := fmt.Scan(&feature)
			if err != nil {
				fmt.Println("Número inválido.")
				continue
			}
			break
		}
		// Switch para manejar las diferentes características que el usuario puede seleccionar.
		switch feature {

		case 1:
			// Leer la entrada del usuario para el nuevo vértice.
			userInput.ReadString('\n')
			var nuevoLugar string
			fmt.Print("Ingresa el nuevo vértice:")
			nuevoLugar, _ = userInput.ReadString('\n')
			nuevoLugar = strings.TrimRight(nuevoLugar, "\r\n")
			// Crear un nuevo vértice en el grafo.
			CrearVertice(nuevoLugar, grafo)
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		case 2:
			// Leer la entrada del usuario para el nuevo nombre del vértice.
			userInput.ReadString('\n')
			ImprimirMenu(*grafo)
			var numVerticeAcambiar byte
			fmt.Print("Ingrese el numero del lugar(vertice) que desea modificar: ")
			for {

				_, err := fmt.Scan(&numVerticeAcambiar)
				if err != nil || numVerticeAcambiar > 9 || numVerticeAcambiar < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			vertice := buscarVertice(traductorMenu(numVerticeAcambiar), grafo)
			var nuevoNombre string
			fmt.Print("Ingresa el nuevo nombre del vértice: ")
			userInput.ReadString('\n')
			nuevoNombre, _ = userInput.ReadString('\n')
			nuevoNombre = strings.TrimRight(nuevoNombre, "\r\n")
			// Actualizar el nombre del vértice existente en el grafo.
			vertice.Nombre = nuevoNombre
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		case 3:
			// Leer la entrada del usuario para el vértice a eliminar.
			ImprimirMenu(*grafo)
			userInput.ReadString('\n')
			var verticeAEliminar byte
			fmt.Print("Ingrese el numero del lugar (vertice) a eliminar: ")
			for {
				_, err := fmt.Scan(&verticeAEliminar)
				if err != nil || verticeAEliminar > 9 || verticeAEliminar < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Verificar si el vértice a eliminar existe y eliminarlo del grafo si es el caso.
			if buscarVertice(traductorMenu(verticeAEliminar), grafo) != nil {
				BorrarVertice(traductorMenu(verticeAEliminar), grafo)
				fmt.Println()
				fmt.Println("Los cambios han sido guardados")
				fmt.Println()
			} else {
				fmt.Print("El vertice a eliminar no existe")
			}
			continue

		case 4:
			// Leer la entrada del usuario para el nombre de la nueva actividad.
			ImprimirMenu(*grafo)
			var numVerticeAcambiar byte
			fmt.Print("Ingrese el numero del lugar(vertice) que desea modificar: ")
			for {
				_, err := fmt.Scan(&numVerticeAcambiar)
				if err != nil || numVerticeAcambiar > 9 || numVerticeAcambiar < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			vertice := buscarVertice(traductorMenu(numVerticeAcambiar), grafo)
			fmt.Println()
			fmt.Println()
			fmt.Println("-----------------------------------------------------------------------------------------")
			fmt.Println("1. Senderismo \n2. Pesca \n3. Termales " +
				"\n4. Volcanes \n5. Paseos a caballo \n6. Mariposario \n7. Ciclismo \n8. Miradero \n9. Natacion" +
				"\n10. Cine \n11. Parque de Diversiones \n12. Vistar monumentos \n13. Visistar museos \n14. Canopy")
			userInput.ReadString('\n')
			var actividad byte
			fmt.Print("Ingrese el numero de la nueva actividad: ")
			for {
				_, err := fmt.Scan(&actividad)
				if err != nil || actividad > 14 || actividad < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Verificar si la actividad existe y agregarla al vértice si es el caso.
			if buscarActividad(listaActividades, traductorActividades(actividad)) == nil {
				fmt.Println("La actividad no existe")
			}

			agregarActividadVertice(vertice, traductorActividades(actividad), listaActividades)
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		case 5:
			// Leer la entrada del usuario para el nombre de la actividad a modificar.
			userInput.ReadString('\n')
			fmt.Println()
			fmt.Println()
			fmt.Println("-----------------------------------------------------------------------------------------")
			fmt.Println("1. Senderismo \n2. Pesca \n3. Termales " +
				"\n4. Volcanes \n5. Paseos a caballo \n6. Mariposario \n7. Ciclismo \n8. Miradero \n9. Natacion" +
				"\n10. Cine \n11. Parque de Diversiones \n12. Vistar monumentos \n13. Visistar museos \n14. Canopy")

			var actividadAModificar byte
			fmt.Print("Ingrese el numero de la actividad a modificar: ")
			for {
				_, err := fmt.Scan(&actividadAModificar)
				if err != nil || actividadAModificar > 14 || actividadAModificar < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Verificar si la actividad existe y modificarla si es el caso.
			if buscarActividad(listaActividades, traductorActividades(actividadAModificar)) == nil {
				fmt.Println("La actividad no existe")
				return
			} else {

				ModificarActividades(traductorActividades(actividadAModificar), listaActividades)
				fmt.Println()
				fmt.Println("Los cambios han sido guardados")
				fmt.Println()

			}
			continue

		case 6:
			// Leer la entrada del usuario para el nombre de la actividad a eliminar.
			ImprimirMenu(*grafo)
			var numVerticeAcambiar byte
			fmt.Print("Ingrese el numero del lugar(vertice) que desea modificar: ")
			for {
				_, err := fmt.Scan(&numVerticeAcambiar)
				if err != nil || numVerticeAcambiar > 9 || numVerticeAcambiar < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			vertice := buscarVertice(traductorMenu(numVerticeAcambiar), grafo)
			userInput.ReadString('\n')
			var actividad byte
			fmt.Println()
			fmt.Println()
			fmt.Println("-----------------------------------------------------------------------------------------")
			fmt.Println("1. Senderismo \n2. Pesca \n3. Termales " +
				"\n4. Volcanes \n5. Paseos a caballo \n6. Mariposario \n7. Ciclismo \n8. Miradero \n9. Natacion" +
				"\n10. Cine \n11. Parque de Diversiones \n12. Vistar monumentos \n13. Visistar museos \n14. Canopy")
			fmt.Print("Ingrese el numero de la actividad que desea eliminar: ")
			for {
				_, err := fmt.Scan(&actividad)
				if err != nil || actividad > 14 || actividad < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Verificar si la actividad existe y eliminarla del vértice si es el caso.
			if buscarActividad(listaActividades, traductorActividades(actividad)) == nil {
				fmt.Println("La actividad no existe")
				return
			}

			eliminarActividadVertice(vertice, buscarActividad(listaActividades, traductorActividades(actividad)).Nombre)
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		case 7:
			// Mostrar el menú del grafo para que el usuario seleccione el origen y destino del nuevo arco.
			ImprimirMenu(*grafo)

			var origen byte
			fmt.Print("Ingrese el numero del origen del nuevo arco: ")
			for {
				_, err := fmt.Scan(&origen)
				if err != nil || origen > 9 || origen < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			var destino byte
			fmt.Print("Ingrese el numero del destino del nuevo arco: ")
			for {
				_, err := fmt.Scan(&destino)
				if err != nil || destino > 9 || destino < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			var distancia int
			fmt.Print("Ingrese la distancia entre el vertice y el destino: ")
			for {
				_, err := fmt.Scan(&distancia)
				if err != nil {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Insertar un nuevo arco en el grafo con el origen, destino y distancia proporcionados.
			insertarArco(traductorMenu(origen), distancia, traductorMenu(destino), grafo)
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		//Modificar ruta (Arco)
		case 8:
			// Mostrar el menú del grafo para que el usuario seleccione el origen y destino del arco a modificar.
			ImprimirMenu(*grafo)
			fmt.Print("Ingrese el numero del origen del arco que desea modificar: ")
			var origenArco byte
			for {
				_, err := fmt.Scan(&origenArco)
				if err != nil || origenArco > 9 || origenArco < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			var destinoArco byte
			fmt.Print("Ingrese el destino del arco que desea modificar: ")
			for {
				_, err := fmt.Scan(&destinoArco)
				if err != nil || destinoArco > 9 || destinoArco < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Modificar el arco en el grafo con el origen y destino proporcionados.
			ModificarArco(traductorMenu(origenArco), traductorMenu(destinoArco), grafo)
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		case 9:
			// Mostrar el menú del grafo para que el usuario seleccione el origen y destino del arco a eliminar.
			ImprimirMenu(*grafo)
			fmt.Print("Ingrese el numero del origen del arco que desea eliminar: ")
			var origenArco byte
			for {
				_, err := fmt.Scan(&origenArco)
				if err != nil || origenArco > 9 || origenArco < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			var destinoArco byte
			fmt.Print("Ingrese el destino del arco que desea eliminar: ")
			for {
				_, err := fmt.Scan(&destinoArco)
				if err != nil || destinoArco > 9 || destinoArco < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Eliminar el arco del grafo con el origen y destino proporcionados.
			BorrarArco(traductorMenu(origenArco), traductorMenu(destinoArco), grafo)
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			continue

		//Consultar ruta
		case 10:
			// Mostrar el menú del grafo para que el usuario seleccione el origen y destino del arco para la consulta de ruta.
			ImprimirMenu(*grafo)
			fmt.Print("Ingrese el numero del origen del arco del que desea partir: ")
			var origenArco byte
			for {
				_, err := fmt.Scan(&origenArco)
				if err != nil || origenArco > 9 || origenArco < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			var destinoArco byte
			fmt.Print("Ingrese el destino del arco al que desea ir: ")
			for {
				_, err := fmt.Scan(&destinoArco)
				if err != nil || destinoArco > 9 || destinoArco < 0 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			var categorias []string

			// Mostrar las opciones de categorías para que el usuario seleccione.
			fmt.Println()
			fmt.Println()
			fmt.Println("-----------------------------------------------------------------------------------------")
			fmt.Print("Ingrese las categorías que desea buscar: ")
			fmt.Println("\n1. Montaña \n2. Acuático \n3. Entretenimiento " +
				"\n4. Cultural \n5. Naturaleza ")

			for {
				var cat byte
				_, err := fmt.Scan(&cat)
				if err != nil || cat > 5 || cat < 1 {
					fmt.Print("Número inválido. Por favor, ingrese un número del 1 al 5.")
					userInput.ReadString('\n')
					continue
				}

				// Traducir el número de categoría a su nombre correspondiente y agregarlo a la lista de categorías.
				categorias = append(categorias, traductorCategorias(cat))
				fmt.Println("¿Desea agregar más categorías? (S/N): ")

				var respuesta string
				fmt.Scan(&respuesta)
				if strings.ToUpper(respuesta) != "S" {
					break
				}

				// Mostrar las categorías seleccionadas hasta el momento.
				fmt.Println("Categorías seleccionadas:")
				for _, categoria := range categorias {
					fmt.Println(categoria)
				}
			}

			// Variables para almacenar los resultados de la consulta de ruta.
			var ruta string
			var dis int
			var cantidadActividades int

			// Llamar a la función para encontrar la ruta más corta con las categorías seleccionadas.
			laruta, distancia, cantact := rutaCortaConCategorias(buscarVertice(traductorMenu(origenArco), grafo), traductorMenu(destinoArco), ruta, dis, cantidadActividades, grafo, categorias)
			fmt.Printf("La ruta más corta con la mayor cantidad de actividades de su interes es %s con una distancia de %v km y una cantidad de actividades de %v", laruta, distancia, cantact)
			// Desmarcar los nodos del grafo para futuras consultas de ruta.
			desmarcar(*grafo)
			agregarUnanuevaPersona("personas.json", *listaActividades)
			continue

			//Estadisticas generales
		case 11:

			var personasLoaded []Persona
			file, err := os.Open("personas.json")
			if err != nil {
				fmt.Println("Error al abrir el archivo:", err)
				return
			}
			defer file.Close()

			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&personasLoaded); err != nil {
				fmt.Println("Error al leer el archivo JSON:", err)
				return
			}
			for _, personaData := range personasLoaded {
				cargarArbol(&arbol, personaData)
			}
			// Mostrar el menú del grafo para que el usuario seleccione un lugar y una actividad para calcular estadísticas generales.
			ImprimirMenu(*grafo)
			var lugar byte
			fmt.Print("Ingrese un lugar para poder calcular estadísticas generales: ")
			for {
				_, err := fmt.Scan(&lugar)
				if err != nil || lugar < 0 || lugar > 9 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}
			var actividad byte

			fmt.Println()
			fmt.Println()
			fmt.Println("-----------------------------------------------------------------------------------------")
			fmt.Println("1. Senderismo \n2. Pesca \n3. Termales " +
				"\n4. Volcanes \n5. Paseos a caballo \n6. Mariposario \n7. Ciclismo \n8. Miradero \n9. Natacion" +
				"\n10. Cine \n11. Parque de Diversiones \n12. Vistar monumentos \n13. Visistar museos \n14. Canopy")
			fmt.Print("Ingrese una actividad para poder calcular estadísticas generales: ")
			for {
				_, err := fmt.Scan(&actividad)
				if err != nil || actividad <= 0 || actividad > 14 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Calcular estadísticas generales para el lugar y la actividad proporcionados.
			EstadisticasGenerales(&arbol, traductorMenu(lugar), traductorActividades(actividad))
			arbol = nil
			continue

		//Ver lista de actividades
		case 12:
			// Mostrar la lista de actividades disponibles.
			MostrarActividades(*listaActividades)
			continue

		//Ver actividades por cada vertice
		case 13:
			// Leer la entrada del usuario para obtener el nombre del vértice y mostrar las actividades asociadas.
			userInput.ReadString('\n')
			ImprimirMenu(*grafo)
			fmt.Print("Ingresa el nombre del vértice del que deseas ver las actividades: ")
			var lugar byte
			for {
				_, err := fmt.Scan(&lugar)
				if err != nil || lugar < 0 || lugar > 9 {
					fmt.Println("Número inválido.")
					userInput.ReadString('\n')
					continue
				}
				break
			}

			// Verificar si el vértice existe y mostrar las actividades asociadas si es el caso.
			if buscarVertice(traductorMenu(lugar), grafo) != nil {
				fmt.Println()
				fmt.Println("Las actividades del vertice son: ")
				fmt.Println()
				MostrarActividadesPorVertice(buscarVertice(traductorMenu(lugar), grafo))
			} else {
				fmt.Println("El vertice no se encontró!!!")
			}
			continue

		case 0:
			fmt.Println("El programa ha finalizado")
			break loop

		default:
			// Mostrar un mensaje de error para opciones inválidas.
			fmt.Println()
			fmt.Println("Esta opción no es valida, ingrese un numero valido")
			fmt.Println()
			continue

		}
	}
}

// buscarVertice busca un vértice por su nombre en un grafo dado.
// Recibe el nombre del vértice a buscar (origen) y un puntero al slice de estructuras Vertice (grafo).
// Si encuentra el vértice, devuelve un puntero al vértice encontrado; de lo contrario, devuelve nil.
func buscarVertice(origen string, grafo *[]Vertice) *Vertice {
	// Verifica si el grafo es nulo.
	if grafo == nil {
		return nil
	}
	// Convierte el puntero del grafo a un slice de estructuras Vertice.
	elgrafo := *grafo
	// Itera a través del grafo para buscar el vértice por su nombre.
	for i := 0; i < len(elgrafo); i++ {
		// Compara el nombre del vértice actual con el nombre de origen.
		if elgrafo[i].Nombre == origen {
			// Si encuentra el vértice, devuelve un puntero al vértice encontrado.
			return &elgrafo[i]
		}
	}
	// Si no encuentra el vértice, devuelve nil.
	return nil
}

// BorrarVertice elimina un vértice del grafo según su nombre.
// Recibe el nombre del vértice a borrar (nombre) y un puntero al slice de estructuras Vertice (grafo).
func BorrarVertice(nombre string, grafo *[]Vertice) {
	// Crea una nueva lista para almacenar los vértices que no deben ser eliminados.
	nuevaLista := []Vertice{}
	// Convierte el puntero del grafo a un slice de estructuras Vertice.
	copiaGrafo := *grafo
	// Itera a través del grafo y agrega los vértices que no coinciden con el nombre proporcionado a la nueva lista.
	for i := 0; i < len(copiaGrafo); i++ {
		if copiaGrafo[i].Nombre != nombre {
			nuevaLista = append(nuevaLista, copiaGrafo[i])
		}
	}
	// Asigna la nueva lista al puntero del grafo, eliminando así el vértice del grafo original.
	*grafo = nuevaLista
}

// desmarcar restablece las marcas "visitado" en todos los vértices del grafo.
// Recibe un slice de estructuras Vertice (grafo).
func desmarcar(grafo []Vertice) {
	// Itera a través de los vértices y establece el campo "visitado" en false para cada uno.
	for i := range grafo {
		grafo[i].visitado = false
	}
}

//==================================Actividades Vertice===============================

// agregarActividadVertice agrega una actividad al slice de actividades de un vértice específico.
// Recibe un puntero al vértice (vertice), el nombre de la actividad a agregar (nombreActividad),
// y un puntero al slice de estructuras Actividad (listaActividades).
func agregarActividadVertice(vertice *Vertice, nombreActividad string, listaActividades *[]Actividad) {
	// Verifica si la lista de actividades o el vértice son nulos.
	if listaActividades == nil || vertice == nil {
		return
	}
	// Obtiene el slice de actividades desde el puntero proporcionado.
	lista := *listaActividades
	// Itera a través de la lista de actividades para encontrar la actividad con el nombre proporcionado.
	for i := 0; i < len(lista); i++ {
		// Comprueba si la actividad actual no es nula y tiene el mismo nombre que el proporcionado.
		if &lista[i] != nil && lista[i].Nombre == nombreActividad {
			// Agrega un puntero a la actividad encontrada al slice de actividades del vértice.
			vertice.Actividades = append(vertice.Actividades, &lista[i])
			return
		}
	}
}

// eliminarActividadVertice elimina una actividad del slice de actividades de un vértice específico.
// Recibe un puntero al vértice (vertice) y el nombre de la actividad a eliminar (nombreActividad).
func eliminarActividadVertice(vertice *Vertice, nombreActividad string) {
	// Crea un nuevo slice para almacenar las actividades que no deben ser eliminadas.
	newLista := []*Actividad{}
	// Itera a través de las actividades del vértice y agrega las actividades que no coinciden con el nombre proporcionado al nuevo slice.
	for i := 0; i < len(vertice.Actividades); i++ {
		if vertice.Actividades[i].Nombre != nombreActividad {
			newLista = append(newLista, vertice.Actividades[i])
		}
	}
	// Asigna el nuevo slice de actividades al campo "Actividades" del vértice, eliminando así la actividad del vértice.
	vertice.Actividades = newLista
}

// MostrarActividadesPorVertice muestra las actividades asociadas a un vértice específico.
// Recibe un puntero al vértice (vertice).
func MostrarActividadesPorVertice(vertice *Vertice) {
	fmt.Println()
	fmt.Println(vertice.Nombre)
	for _, actividad := range vertice.Actividades {
		fmt.Println(actividad.Nombre, actividad.Categoria)
	}
}

//==================================Actividades====================================

// RegistrarActividad registra una nueva actividad en la lista de actividades.
// Recibe un puntero al slice de estructuras Actividad (listaActividades),
// el nombre de la actividad a registrar (nombreActividad), y la categoría de la actividad (categoria).
func RegistrarActividad(listaActividades *[]Actividad, nombreActividad string, categoria string) {
	// Crea una nueva estructura Actividad con el nombre y la categoría proporcionados.
	actividad := Actividad{Nombre: nombreActividad, Categoria: categoria}
	// Agrega la nueva actividad al slice de actividades usando append.
	*listaActividades = append(*listaActividades, actividad)
}

// buscarActividad busca una actividad por su nombre en una lista de actividades.
// Recibe un puntero al slice de estructuras Actividad (listaActividades) y el nombre de la actividad a buscar (nombre).
// Si encuentra la actividad, devuelve un puntero a la estructura Actividad; de lo contrario, devuelve nil.
func buscarActividad(listaActividades *[]Actividad, nombre string) *Actividad {
	// Itera a través de la lista de actividades y compara los nombres.
	for i := range *listaActividades {
		if (*listaActividades)[i].Nombre == nombre {
			// Si encuentra la actividad, devuelve un puntero a la estructura Actividad.
			return &(*listaActividades)[i]
		}
	}
	// Si no encuentra la actividad, devuelve nil.
	return nil
}

// ModificarActividades permite modificar el nombre o la categoría de una actividad en una lista de actividades.
// Recibe el nombre de la actividad a modificar (nombreActividad) y un puntero al slice de estructuras Actividad (listaActividades).
func ModificarActividades(nombreActividad string, listaActividades *[]Actividad) {
	// Crear un lector para la entrada del usuario.
	userInput := bufio.NewReader(os.Stdin)
	// Buscar la actividad en la lista de actividades.
	actividad := buscarActividad(listaActividades, nombreActividad)
	// Verificar si la actividad existe.
	if actividad == nil {
		fmt.Println("La actividad no existe")
		return
	}
	// Mostrar opciones al usuario.
	fmt.Println("1. Para modificar el nombre \n2. Para modificar la categoria")
	var feature byte
	for {
		// Leer la opción del usuario y manejar errores de entrada.
		_, err := fmt.Scan(&feature)
		if err != nil {
			fmt.Println("Número inválido.")
			userInput.ReadString('\n')
			continue
		}
		break
	}

	// Limpiar el búfer del lector.
	userInput.ReadString('\n')

	if actividad == nil {
		fmt.Println("La actividad no existe")
		return
	}
	for {
		// Realizar la modificación correspondiente.
		switch feature {
		case 1:

			var nombre string
			fmt.Print("Ingresa el nuevo nombre: ")
			nombre, _ = userInput.ReadString('\n')
			nombre = strings.TrimRight(nombre, "\r\n")
			actividad.Nombre = nombre
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			return
		case 2:

			var categoria string
			fmt.Print("Ingresa la nueva categoría: ")
			categoria, _ = userInput.ReadString('\n')
			categoria = strings.TrimRight(categoria, "\r\n")
			actividad.Categoria = categoria
			fmt.Println()
			fmt.Println("Los cambios han sido guardados")
			fmt.Println()
			return
		default:
			fmt.Println()
			fmt.Println("This feature doesn't exist")
			fmt.Println()
			return
		}
	}
}

// BorrarActividades elimina una actividad específica de la lista de actividades.
// Recibe un puntero al slice de estructuras Actividad (listaActividades) y el nombre de la actividad a eliminar (nombre).
func BorrarActividades(listaActividades *[]Actividad, nombre string) {
	// Crea un nuevo slice para almacenar las actividades que no deben ser eliminadas.
	nuevaLista := []Actividad{}
	// Itera a través de las actividades y agrega las actividades que no coinciden con el nombre proporcionado al nuevo slice.
	for i := 0; i < len(*listaActividades); i++ {
		lista := *listaActividades
		if lista[i].Nombre != nombre {
			nuevaLista = append(nuevaLista, lista[i])
		}
		// Actualiza el puntero listaActividades para que apunte al nuevo slice sin la actividad eliminada.
		listaActividades = &nuevaLista
	}
}

// MostrarActividades muestra una lista de actividades en la salida estándar.
// Esta función toma una lista de actividades ([]Actividad) como argumento y muestra
// el nombre y la categoría de cada actividad en la salida estándar.
func MostrarActividades(listaActividades []Actividad) {
	for i := 0; i < len(listaActividades); i++ {
		// Imprimir el nombre y la categoría de la actividad.
		fmt.Printf("%s, %s\n", listaActividades[i].Nombre, listaActividades[i].Categoria)
	}
}

//===================================Arco======================================

// insertarArco inserta un arco en un grafo representado como un slice de vértices.
// Esta función toma el nombre del vértice de origen, la distancia entre los vértices y el nombre del vértice de destino.
// También toma un puntero a un slice de vértices que representa el grafo.
// La función busca los vértices de origen y destino en el grafo, luego crea un arco con la distancia especificada
// y lo agrega a la lista de arcos del vértice de origen.
func insertarArco(origen string, dis int, des string, grafo *[]Vertice) {
	// Buscar el vértice de origen en el grafo.
	vOrigen := buscarVertice(origen, grafo)

	// Buscar el vértice de destino en el grafo.
	vDestino := buscarVertice(des, grafo)

	// Verificar si el vértice de origen no se encuentra en el grafo.
	if vOrigen == nil {
		fmt.Println("No se encuentra el vértice de origen.")
		return
	}

	// Verificar si el vértice de destino no se encuentra en el grafo.
	if vDestino == nil {
		fmt.Println("No se encuentra el vértice de destino.")
		return
	}

	// Crear un nuevo arco con la distancia especificada, el vértice de origen y el vértice de destino.
	arco := Arco{Distancia: dis, Origen: vOrigen, Destino: vDestino}

	// Agregar el arco a la lista de arcos del vértice de origen.
	vOrigen.ListaArcos = append(vOrigen.ListaArcos, &arco)
}

// ModificarArco permite modificar la distancia de un arco en un grafo representado como un slice de vértices.
// Esta función toma el nombre del vértice de origen, el nombre del vértice de destino y un puntero al grafo.
// Primero, busca el vértice de origen en el grafo y muestra un menú de opciones para modificar el arco.
// El usuario puede elegir modificar la distancia (feature 1) o cancelar la operación. Después de cada modificación,
// se informa al usuario si los cambios han sido guardados.
func ModificarArco(origen, destino string, grafo *[]Vertice) {
	// Buscar el vértice de origen en el grafo.
	vertice := buscarVertice(origen, grafo)

	// Verificar si el vértice de origen no se encuentra en el grafo.
	if vertice == nil {
		fmt.Println("El vértice no existe.")
		return
	}

	fmt.Println("1. Para modificar la distancia")
	var feature byte
	for {
		_, err := fmt.Scan(&feature)
		if err != nil {
			fmt.Println("Número inválido.")
			continue
		}
		break
	}

	arco := buscarArco(origen, destino, grafo)

	for {
		switch feature {
		case 1:
			// Leer la nueva distancia del arco.
			var distancia int
			fmt.Print("Ingresa la nueva distancia del arco: ")
			_, _ = fmt.Scan(&distancia)
			arco.Distancia = distancia
			fmt.Println()
			fmt.Println("Los cambios han sido guardados.")
			fmt.Println()
			return
		default:
			fmt.Println()
			fmt.Println("Esta función no existe.")
			fmt.Println()
			return
		}
	}
}

// buscarArco busca un arco en un grafo representado como un slice de vértices.
// Esta función toma el nombre del vértice de origen, el nombre del vértice de destino y un puntero al grafo.
// Itera a través de los arcos del vértice de origen para encontrar el arco que conecta el vértice de origen con el vértice de destino.
// Si se encuentra el arco, se devuelve una referencia a ese arco. Si no se encuentra, se devuelve nil.
func buscarArco(origen, destino string, grafo *[]Vertice) *Arco {
	// Buscar el vértice de origen en el grafo.
	vertice := buscarVertice(origen, grafo)

	// Iterar a través de los arcos del vértice de origen.
	for i := 0; i < len(vertice.ListaArcos); i++ {
		// Verificar si el arco conecta el vértice de origen con el vértice de destino.
		if vertice.ListaArcos[i].Origen.Nombre == origen && vertice.ListaArcos[i].Destino.Nombre == destino {
			// Devolver una referencia al arco encontrado.
			return vertice.ListaArcos[i]
		}
	}

	// Si no se encuentra el arco, devolver nil.
	return nil
}

// BorrarArco elimina un arco que conecta un vértice de origen con un vértice de destino en un grafo representado como un slice de vértices.
// Esta función toma el nombre del vértice de origen, el nombre del vértice de destino y un puntero al grafo.
// Primero, busca el vértice de origen en el grafo. Si el vértice de origen no se encuentra, la función se detiene.
// Luego, verifica si el vértice de origen tiene una lista de arcos. Si no tiene una lista de arcos, la función se detiene.
// Si el vértice de origen tiene una lista de arcos, itera a través de los arcos y crea una nueva lista de arcos que excluye el arco que conecta el vértice de origen con el vértice de destino.
// Finalmente, actualiza la lista de arcos del vértice de origen con la nueva lista de arcos, eliminando así el arco.
func BorrarArco(origen string, destino string, grafo *[]Vertice) {
	// Buscar el vértice de origen en el grafo.
	vertice := buscarVertice(origen, grafo)

	// Verificar si el vértice de origen no se encuentra en el grafo.
	if vertice == nil {
		return
	}

	// Verificar si el vértice de origen tiene una lista de arcos.
	if vertice.ListaArcos == nil {
		return
	}

	// Crear una nueva lista de arcos que excluye el arco que conecta el vértice de origen con el vértice de destino.
	listaArcos := vertice.ListaArcos
	nuevaLista := []*Arco{}
	for i := 0; i < len(listaArcos); i++ {
		if listaArcos[i].Origen.Nombre != origen || listaArcos[i].Destino.Nombre != destino {
			nuevaLista = append(nuevaLista, listaArcos[i])
		}
	}

	// Actualizar la lista de arcos del vértice de origen con la nueva lista de arcos, eliminando así el arco.
	vertice.ListaArcos = nuevaLista
}

//============================Funciones para grafo====================================

// rutaCortaConCategorias encuentra la ruta más corta desde un vértice de origen hasta un vértice de destino en un grafo, considerando restricciones de categorías de actividades.
// Esta función toma un vértice de origen, el nombre del vértice de destino, una cadena de ruta parcial, la distancia acumulada, la cantidad de actividades acumuladas, un puntero al grafo, y un slice de categorías de actividades.
// La función busca la ruta más corta desde el vértice de origen hasta el vértice de destino, teniendo en cuenta las categorías de actividades permitidas.
// Devuelve la ruta más corta encontrada, su distancia y la cantidad de actividades que cumplan con las restricciones de categorías.
// Si no se encuentra una ruta válida, se devuelve una ruta vacía, distancia máxima y 0 actividades.
func rutaCortaConCategorias(origen *Vertice, destino string, ruta string, dis int, cantidadActividades int, grafo *[]Vertice, categorias []string) (string, int, int) {
	if origen == nil || origen.visitado {
		return "", math.MaxInt64, 0 // Devuelve una ruta inválida con distancia máxima y 0 actividades
	}
	if origen.Nombre == destino {
		return ruta + "-->" + destino, dis, cantidadActividades
	}

	origen.visitado = true
	var rutaMejor string
	var distanciaMejor = math.MaxInt64
	var actividadesMejor = 0

	for _, arco := range origen.ListaArcos {
		ruta, distancia, actividades := rutaCortaConCategorias(buscarVertice(arco.Destino.Nombre, grafo), destino, ruta+"-->"+origen.Nombre, dis+arco.Distancia, cantidadActividades+contarActividadesPorCategoria(origen, categorias), grafo, categorias)

		if actividades >= actividadesMejor {
			if distancia < distanciaMejor || (distancia == distanciaMejor && actividades > actividadesMejor) {
				rutaMejor = ruta
				distanciaMejor = distancia
				actividadesMejor = actividades
			}
		}
	}

	origen.visitado = false

	return rutaMejor, distanciaMejor, actividadesMejor
}

// contarActividadesPorCategoria cuenta la cantidad de actividades en un vértice que pertenecen a categorías especificadas.
// Esta función toma un puntero a un vértice (*Vertice) que representa el vértice en el grafo y un slice de categorías de actividades.
// Itera a través de las actividades del vértice y verifica si cada actividad pertenece a alguna de las categorías especificadas en el slice.
// Devuelve el conteo de actividades que cumplen con las categorías especificadas.
func contarActividadesPorCategoria(verticePuntero *Vertice, categorias []string) int {
	count := 0
	vertice := verticePuntero
	for i := 0; i < len(vertice.Actividades); i++ {
		for j := 0; j < len(categorias); j++ {
			if vertice.Actividades[i].Categoria == categorias[j] {
				count++
				break
			}
		}
	}
	return count
}

//==============================Vertice Arbol=========================================

// insertarVerticeArbol inserta un nuevo vértice en una estructura de árbol (enlazada) representada por un puntero doble a un vértice de árbol (**VerticeArbol).
// Esta función toma el nombre del nuevo vértice y un puntero doble al árbol. Crea un nuevo vértice con el nombre especificado y lo enlaza como el siguiente vértice del árbol.
// Luego, actualiza el puntero doble al árbol para apuntar al nuevo vértice recién insertado.
// Devuelve un puntero al nuevo vértice que se ha insertado en el árbol.
func insertarVerticeArbol(nom string, arbol **VerticeArbol) *VerticeArbol {
	// Crear un nuevo vértice con el nombre especificado y enlazarlo como el siguiente vértice del árbol.
	nuevoVertice := &VerticeArbol{nombre: nom, sigV: *arbol}

	// Actualizar el puntero doble al árbol para apuntar al nuevo vértice recién insertado.
	*arbol = nuevoVertice

	// Devolver un puntero al nuevo vértice insertado en el árbol.
	return nuevoVertice
}

// buscarVerticeArbol busca un vértice en una estructura de árbol (enlazada) por su nombre.
// Esta función toma el nombre del vértice que se desea buscar y un puntero doble al árbol.
// Itera a través del árbol enlazado y compara el nombre del vértice con el nombre proporcionado.
// Si se encuentra un vértice con el nombre coincidente, se devuelve un puntero a ese vértice. Si no se encuentra, se devuelve nil.
func buscarVerticeArbol(origen string, arbol **VerticeArbol) *VerticeArbol {
	tempV := *arbol
	for tempV != nil {
		if tempV.nombre == origen {
			return tempV
		}
		tempV = tempV.sigV
	}
	return nil
}

//==============================Arco Arbol=========================================

// insertarArcoArbol inserta un arco entre dos vértices en una estructura de árbol (enlazada) por sus nombres y una distancia específica.
// Esta función toma un puntero doble al vértice de origen (*VerticeArbol), un puntero doble al vértice de destino (*VerticeArbol) y un puntero doble al árbol (estructura de árbol enlazada).
// Verifica si los vértices de origen y destino existen en el árbol. Si el vértice de origen no se encuentra, muestra un mensaje de error y termina la operación.
// Si el vértice de destino no se encuentra, muestra un mensaje de error y termina la operación.
// Luego, crea un nuevo arco con el vértice de destino y lo enlaza en la lista de arcos del vértice de origen.
func insertarArcoArbol(vOrigen **VerticeArbol, des **VerticeArbol, arbol **VerticeArbol) {
	if *vOrigen == nil {
		fmt.Println("No se encuentra el origen.")
		return
	}
	if *des == nil {
		fmt.Println("No se encuentra el destino.")
		return
	}

	// Crear un nuevo arco con el vértice de destino y enlazarlo en la lista de arcos del vértice de origen.
	nuevoArco := &ArcoArbol{destino: *des, sigA: nil}
	origen := *vOrigen
	nuevoArco.sigA = origen.subListaArcos
	origen.subListaArcos = nuevoArco
}

// buscarArcoArbol busca un arco en la lista de arcos de un vértice de árbol (estructura de árbol enlazada) por el nombre del destino.
// Esta función toma un puntero doble al vértice de árbol (*VerticeArbol) y el nombre del destino que se desea buscar.
// Primero, verifica si el puntero doble al vértice de árbol es nulo. Si es nulo, muestra un mensaje de error y devuelve nil.
// Luego, itera a través de la lista de arcos del vértice de árbol para encontrar un arco cuyo destino coincida con el nombre especificado.
// Si se encuentra un arco con el nombre coincidente, se devuelve un puntero a ese arco. Si no se encuentra, se devuelve nil.
func buscarArcoArbol(vertice **VerticeArbol, destino string) *ArcoArbol {
	if *vertice == nil {
		fmt.Println("Error en buscar arco")
		return nil
	}
	elvertice := *vertice
	listaArcos := elvertice.subListaArcos
	for listaArcos != nil {
		if listaArcos.destino.nombre == destino {
			return listaArcos
		}
		listaArcos = listaArcos.sigA
	}
	return nil
}

//====================================Arbol=======================================

// anadirInformacion añade información a un nodo de un árbol de vértices (estructura de árbol enlazada).
// Esta función toma un puntero doble al nodo del árbol (*VerticeArbol).
// Verifica si el puntero doble al nodo del árbol es nulo. Si es nulo, muestra un mensaje de error y termina la operación.
// Luego, incrementa la cantidad almacenada en el nodo del árbol en 1 unidad.
func anadirInformacion(nodoArbol **VerticeArbol) {
	if nodoArbol == nil {
		fmt.Println("Error en añadir Información")
		return
	}
	nodoDelArbol := *nodoArbol
	nodoDelArbol.cantidad = nodoDelArbol.cantidad + 1
}

// cargarArbol carga información de una persona en una estructura de árbol (enlazada) según sus atributos.
// Esta función toma un puntero doble al árbol de vértices (*VerticeArbol) y los datos de una persona (Persona).
// Primero, verifica si el árbol de vértices está vacío. Si está vacío, crea un vértice raíz.
// Luego, verifica si el género de la persona ya existe como un vértice en el árbol. Si no existe, crea un nuevo vértice para el género y establece un arco entre el vértice raíz y el vértice del género.
// Además, incrementa la información del vértice del género.
// Luego, verifica si el lugar de residencia de la persona ya existe como un vértice en el árbol. Si no existe, crea un nuevo vértice para el lugar de residencia y establece un arco entre el vértice del género y el vértice del lugar de residencia. Incrementa la información del vértice del lugar de residencia.
// Según la edad de la persona, crea vértices y arcos adicionales para representar la edad en el árbol, incrementando la información correspondiente.
// Finalmente, crea un vértice para la actividad de la persona y establece un arco desde el vértice de la edad al vértice de la actividad, incrementando la información del vértice de la actividad si ya existe.
func cargarArbol(arbol **VerticeArbol, personaData Persona) {
	if *arbol == nil {
		insertarVerticeArbol("Raiz", arbol)
	}
	if buscarVerticeArbol(personaData.Genero, arbol) == nil {
		insertarVerticeArbol(personaData.Genero, arbol)
		vertice1 := buscarVerticeArbol("Raiz", arbol)
		vertice2 := buscarVerticeArbol(personaData.Genero, arbol)

		insertarArcoArbol(&vertice1, &vertice2, arbol)
		genero := buscarVerticeArbol(personaData.Genero, arbol)
		anadirInformacion(&genero)
	} else if buscarVerticeArbol(personaData.Genero, arbol) != nil {
		genero := buscarVerticeArbol(personaData.Genero, arbol)
		anadirInformacion(&genero)
	}
	elvertice := buscarVerticeArbol(personaData.Genero, arbol)
	if buscarArcoArbol(&elvertice, personaData.Residencia) == nil {
		residencia := insertarVerticeArbol(personaData.Residencia, arbol)
		vertice3 := buscarVerticeArbol(personaData.Genero, arbol)
		insertarArcoArbol(&vertice3, &residencia, arbol)
		anadirInformacion(&residencia)
	} else if buscarArcoArbol(&elvertice, personaData.Residencia) != nil {
		anadirInformacion(&buscarArcoArbol(&elvertice, personaData.Residencia).destino)
	}
	if personaData.Edad > 18 && personaData.Edad <= 30 {
		if buscarArcoArbol(&elvertice, personaData.Residencia) == nil {
			fmt.Println("error")
			return
		}
	}
	if buscarArcoArbol(&elvertice, personaData.Residencia) == nil {
		fmt.Println("Error")
		return
	}
	residencia1 := buscarArcoArbol(&elvertice, personaData.Residencia).destino
	if buscarArcoArbol(&residencia1, "18/30") == nil {
		edad := insertarVerticeArbol("18/30", arbol)
		insertarArcoArbol(&residencia1, &edad, arbol)
		anadirInformacion(&edad)

	} else if buscarArcoArbol(&residencia1, "18/30") != nil {
		anadirInformacion(&buscarArcoArbol(&residencia1, "18/30").destino)
	}

	if personaData.Edad >= 31 && personaData.Edad <= 64 {
		residencia2 := buscarArcoArbol(&elvertice, personaData.Residencia).destino
		if buscarArcoArbol(&residencia2, "31/64") == nil {
			edad := insertarVerticeArbol("31/64", arbol)
			insertarArcoArbol(&residencia2, &edad, arbol)
			anadirInformacion(&edad)
		} else if buscarArcoArbol(&residencia2, "31/64") != nil {
			anadirInformacion(&buscarArcoArbol(&residencia2, "31/64").destino)
		}
	}
	if personaData.Edad > 64 {
		residencia3 := buscarArcoArbol(&elvertice, personaData.Residencia).destino
		if buscarArcoArbol(&residencia3, "64") == nil {
			edad := insertarVerticeArbol("64", arbol)
			insertarArcoArbol(&residencia3, &edad, arbol)

		} else if buscarArcoArbol(&residencia3, "64") != nil {
			anadirInformacion(&buscarArcoArbol(&residencia3, "64").destino)
		}
	}
	//for i := 0; i < len(personaData.Actividad); i++ {
	residencia4 := buscarArcoArbol(&elvertice, personaData.Residencia).destino
	var variable string
	if personaData.Edad >= 18 || personaData.Edad <= 30 {
		variable = "18/30"
	} else if personaData.Edad >= 31 || personaData.Edad <= 64 {
		variable = "31/64"
	} else {
		variable = "64"
	}
	edad := buscarArcoArbol(&residencia4, variable).destino
	if buscarArcoArbol(&edad, personaData.Actividad.Nombre) == nil {
		actividad := insertarVerticeArbol(personaData.Actividad.Nombre, arbol)
		insertarArcoArbol(&edad, &actividad, arbol)
		anadirInformacion(&actividad)
	} else if buscarArcoArbol(&edad, personaData.Actividad.Nombre) != nil {
		anadirInformacion(&buscarArcoArbol(&edad, personaData.Actividad.Nombre).destino)
	}
}

// ============================Funciones para arbol=============================
// EstadisticasGenerales calcula estadísticas generales de una población almacenada en una estructura de árbol (enlazada).
// Esta función toma un puntero doble al árbol de vértices (*VerticeArbol), un lugar de residencia (lugar) y una actividad específica (actividad).
// Calcula el porcentaje de mujeres y hombres en la población, el porcentaje de personas en un lugar de residencia específico,
// el porcentaje de personas en diferentes grupos de edad y el porcentaje de personas que realizan una actividad específica.
func EstadisticasGenerales(arbol **VerticeArbol, lugar string, actividad string) {
	mujer := buscarVerticeArbol("Femenino", arbol)
	hombre := buscarVerticeArbol("Masculino", arbol)

	porcentajemujer := 0
	cantMujer := 0
	cantHombre := 0
	if mujer != nil {
		porcentajemujer = (mujer.cantidad * 100) / (mujer.cantidad + hombre.cantidad)
		cantMujer = mujer.cantidad
	}
	porcentajehombre := 0
	if hombre != nil {
		porcentajehombre = (hombre.cantidad * 100) / (mujer.cantidad + hombre.cantidad)
		cantHombre = hombre.cantidad
	}
	totalPob := cantMujer + cantHombre
	fmt.Println("El porcentaje de mujeres es ", porcentajemujer, "% y el porcentaje de hombres es ", porcentajehombre, "%")
	porcentajeLugarHombre := 0
	porcentajeLugarMujer := 0
	if buscarArcoArbol(&hombre, lugar) != nil {
		porcentajeLugarHombre = (buscarArcoArbol(&hombre, lugar).destino.cantidad * 100) / cantHombre
	}
	if buscarArcoArbol(&mujer, lugar) != nil {
		porcentajeLugarMujer = (buscarArcoArbol(&mujer, lugar).destino.cantidad * 100) / cantMujer
	}
	fmt.Println("El porcentaje de mujeres en", lugar, "es", porcentajeLugarMujer, "% y el porcentaje de hombres es ", porcentajeLugarHombre, "%")
	elarbol := *arbol
	cont1830 := 0
	cont3164 := 0
	cont64 := 0
	for elarbol != nil {
		if elarbol.nombre == "18/30" {
			cont1830 = cont1830 + 1
		} else if elarbol.nombre == "31/64" {
			cont3164 = cont3164 + 1
		} else if elarbol.nombre == "64" {
			cont64 = cont64 + 1
		}
		elarbol = elarbol.sigV
	}
	if totalPob == 0 {
		fmt.Println("No hay población registrada")
		return
	}
	fmt.Println("El porcentaje de personas entre 18 y 30 es ", (cont1830*100)/totalPob, "%")
	fmt.Println("El porcentaje de personas entre 31 y 64 es ", (cont3164*100)/totalPob, "%")
	fmt.Println("El porcentaje de personas mayor a 64 es", (cont64*100)/totalPob, "%")
	elarbol = *arbol
	contactividad := 0
	for elarbol != nil {
		if elarbol.nombre == actividad {
			contactividad = contactividad + 1
		}
		elarbol = elarbol.sigV
	}
	fmt.Println("El porcentaje de población que realiza la actividad de", actividad, "es del", (contactividad*100)/totalPob, "%")
}

// ===============================Funciones de archivos para agregar personas==========================================
// guardarPersonas guarda una lista de personas en un archivo JSON con el nombre especificado.
// Esta función toma el nombre del archivo (filename) y una lista de personas (personas) que se desea guardar en el archivo.
// Primero, crea o abre el archivo con el nombre especificado y verifica si hay algún error al hacerlo.
// Luego, crea un codificador JSON y escribe la lista de personas en el archivo.
// Si ocurre algún error durante la escritura, la función devuelve un error.
// Si la escritura es exitosa, la función muestra un mensaje indicando que los datos se han escrito en el archivo y devuelve nil.
func guardarPersonas(filename string, personas []Persona) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(personas); err != nil {
		return err
	}
	fmt.Println("Datos escritos en el archivo personas.json")
	return nil
}

// agregarPersona agrega una nueva persona a un archivo JSON existente con el nombre especificado.
// Esta función toma el nombre del archivo (filename) y una nueva persona (nuevaPersona) que se desea agregar al archivo.
// Primero, carga la lista de personas existente desde el archivo especificado utilizando la función cargarPersonas.
// Luego, agrega la nueva persona a la lista de personas cargadas.
// Finalmente, guarda la lista actualizada en el archivo utilizando la función guardarPersonas.
// Si se producen errores durante la carga, adición o guardado, la función devuelve un error correspondiente.
// Si la operación es exitosa, muestra un mensaje indicando que la nueva persona ha sido agregada al archivo y devuelve nil.
func agregarPersona(filename string, nuevaPersona Persona) error {
	// Cargar la lista de personas desde el archivo
	personasCargadas := cargarPersonas(filename)

	// Agregar la nueva persona a la lista
	personasCargadas = append(personasCargadas, nuevaPersona)

	// Guardar la lista actualizada en el archivo
	if err := guardarPersonas(filename, personasCargadas); err != nil {
		return err
	}
	fmt.Println("Nueva persona agregada al archivo personas.json")
	return nil
}

// cargarPersonas carga una lista de personas desde un archivo JSON con el nombre especificado.
// Esta función toma el nombre del archivo (filename) desde el cual se desea cargar la lista de personas.
// Primero, intenta abrir el archivo y verifica si hay errores al hacerlo.
// Luego, crea un decodificador JSON y lee la lista de personas desde el archivo.
// Si ocurre algún error durante la lectura, la función muestra un mensaje de error y devuelve una lista de personas vacía.
// Si la lectura es exitosa, la función devuelve la lista de personas cargada desde el archivo.
func cargarPersonas(filename string) []Persona {
	var personasLoaded []Persona
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return personasLoaded
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&personasLoaded); err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
	}

	return personasLoaded
}

// agregarUnanuevaPersona permite al usuario ingresar información de una nueva persona y agregarla a una lista de personas en un archivo.
// Esta función toma el nombre del archivo (nombreArchivo) en el que se agregará la nueva persona y una lista de actividades (listaActividades) como referencia.
// Primero, solicita al usuario ingresar su género, cantón de residencia, edad y el nombre de su actividad.
// Luego, crea una instancia de la estructura Persona con la información ingresada y la actividad correspondiente buscada en la lista de actividades.
// Finalmente, agrega la nueva persona al archivo utilizando la función agregarPersona.
func agregarUnanuevaPersona(nombreArchivo string, listaActividades []Actividad) {
	userInput := bufio.NewReader(os.Stdin)
	fmt.Print("\nIngrese su genero \n1. Genero masculino \n2. Genero femenino\n")
	var cat byte
	for {
		_, err := fmt.Scan(&cat)
		if err != nil || cat > 2 || cat < 1 {
			fmt.Print("Número inválido. Por favor, ingrese el 1 o el 2.")
			userInput.ReadString('\n')
			continue
		}
		break
	}
	var genero string
	if cat == 1 {
		genero = "Masculino"
	} else if cat == 2 {
		genero = "Femenino"
	}
	userInput.ReadString('\n')
	var nuevoLugar string
	fmt.Println("Ingrese su canton de residencia")
	nuevoLugar, _ = userInput.ReadString('\n')
	nuevoLugar = strings.TrimRight(nuevoLugar, "\r\n")
	fmt.Println("Ingrese su edad: ")
	var edad byte
	for {
		_, err := fmt.Scan(&edad)
		if err != nil || edad < 18 || edad > 125 {
			fmt.Print("Número inválido.")
			continue
		}
		break
	}
	userInput.ReadString('\n')
	var actividad string
	fmt.Println("Ingrese el nombre de su actividad")
	actividad, _ = userInput.ReadString('\n')
	actividad = strings.TrimRight(actividad, "\r\n")
	actividadStruct := buscarActividad(&listaActividades, actividad)
	laactividadStruct := *actividadStruct
	persona := Persona{genero, edad, nuevoLugar, laactividadStruct}
	agregarPersona(nombreArchivo, persona)
}

//===============================Main==========================================

func main() {
	//Añadir actividades a la lista de actividades
	RegistrarActividad(&ListaActividades, "Senderismo", "Montaña")
	RegistrarActividad(&ListaActividades, "Pesca", "Acuatico")
	RegistrarActividad(&ListaActividades, "Termales", "Acuatico")
	RegistrarActividad(&ListaActividades, "Volcanes", "Montaña")
	RegistrarActividad(&ListaActividades, "Paseos a caballo", "Naturaleza")
	RegistrarActividad(&ListaActividades, "Mariposario", "Naturaleza")
	RegistrarActividad(&ListaActividades, "Ciclismo", "Montaña")
	RegistrarActividad(&ListaActividades, "Miradero", "Naturaleza")
	RegistrarActividad(&ListaActividades, "Natacion", "Montaña")
	RegistrarActividad(&ListaActividades, "Cine", "Entretenimiento")
	RegistrarActividad(&ListaActividades, "Parque de Diversiones", "Entretenimeinto")
	RegistrarActividad(&ListaActividades, "Visitar monumentos", "Cultural")
	RegistrarActividad(&ListaActividades, "Visitar museos", "Cultural")
	RegistrarActividad(&ListaActividades, "Canopy", "Montaña")

	//Cargar los vértices y arcos al grafo
	cargarDatos(&grafo)

	//Agregar las actividades correspondientes a los vertices
	agregarActividadVertice(buscarVertice("Zarcero", &grafo), "Natacion", &ListaActividades)
	agregarActividadVertice(buscarVertice("Zarcero", &grafo), "Senderismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("Santa Rosa", &grafo), "Miradero", &ListaActividades)
	agregarActividadVertice(buscarVertice("Santa Rosa", &grafo), "Senderismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("Ciudad Quesada", &grafo), "Termales", &ListaActividades)
	agregarActividadVertice(buscarVertice("Ciudad Quesada", &grafo), "Ciclismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("Ciudad Quesada", &grafo), "Natacion", &ListaActividades)
	agregarActividadVertice(buscarVertice("Palmares", &grafo), "Natacion", &ListaActividades)
	agregarActividadVertice(buscarVertice("Palmares", &grafo), "Ciclismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("Palmares", &grafo), "Mariposario", &ListaActividades)
	agregarActividadVertice(buscarVertice("San Ramon", &grafo), "Ciclismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("San Ramon", &grafo), "Cine", &ListaActividades)
	agregarActividadVertice(buscarVertice("Alajuela", &grafo), "Ciclismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("Alajuela", &grafo), "Visitar monumentos", &ListaActividades)
	agregarActividadVertice(buscarVertice("San Jose", &grafo), "Parque de diversiones", &ListaActividades)
	agregarActividadVertice(buscarVertice("San Jose", &grafo), "Visitar museos", &ListaActividades)
	agregarActividadVertice(buscarVertice("Santa Clara", &grafo), "Termales", &ListaActividades)
	agregarActividadVertice(buscarVertice("Santa Clara", &grafo), "Senderismo", &ListaActividades)
	agregarActividadVertice(buscarVertice("Calle Blancos", &grafo), "Canopy", &ListaActividades)
	agregarActividadVertice(buscarVertice("Naranjo", &grafo), "Paseos a caballo", &ListaActividades)

	//Crear los datos de las personas en una lista
	personas := []Persona{
		{
			Genero:     "Masculino",
			Edad:       25,
			Residencia: "Zarcero",
			Actividad:  *buscarActividad(&ListaActividades, "Senderismo"),
		},
		{
			Genero:     "Femenino",
			Edad:       30,
			Residencia: "Santa Rosa",
			Actividad:  *buscarActividad(&ListaActividades, "Pesca"),
		},
		{
			Genero:     "Femenino",
			Edad:       20,
			Residencia: "San Ramon",
			Actividad:  *buscarActividad(&ListaActividades, "Termales"),
		},
		{
			Genero:     "Masculino",
			Edad:       46,
			Residencia: "Santa Clara",
			Actividad:  *buscarActividad(&ListaActividades, "Volcanes"),
		},
		{
			Genero:     "Femenino",
			Edad:       23,
			Residencia: "Alajuela",
			Actividad:  *buscarActividad(&ListaActividades, "Paseos a caballo"),
		},
		{
			Genero:     "Femenino",
			Edad:       26,
			Residencia: "Alajuela",
			Actividad:  *buscarActividad(&ListaActividades, "Mariposario"),
		},
		{
			Genero:     "Masculino",
			Edad:       35,
			Residencia: "San Jose",
			Actividad:  *buscarActividad(&ListaActividades, "Ciclismo"),
		},
		{
			Genero:     "Masculino",
			Edad:       40,
			Residencia: "Naranjo",
			Actividad:  *buscarActividad(&ListaActividades, "Miradero"),
		},
		{
			Genero:     "Femenino",
			Edad:       28,
			Residencia: "Palmares",
			Actividad:  *buscarActividad(&ListaActividades, "Natacion"),
		},
		{
			Genero:     "Masculino",
			Edad:       20,
			Residencia: "Zarcero",
			Actividad:  *buscarActividad(&ListaActividades, "Cine"),
		},
		{
			Genero:     "Masculino",
			Edad:       26,
			Residencia: "Ciudad Quesada",
			Actividad:  *buscarActividad(&ListaActividades, "Parque de Diversiones"),
		},
		{
			Genero:     "Femenino",
			Edad:       19,
			Residencia: "Santa Rosa",
			Actividad:  *buscarActividad(&ListaActividades, "Visitar monumentos"),
		},
		{
			Genero:     "Femenino",
			Edad:       29,
			Residencia: "Naranjo",
			Actividad:  *buscarActividad(&ListaActividades, "Visitar museos"),
		},
		{
			Genero:     "Masculino",
			Edad:       27,
			Residencia: "San Jose",
			Actividad:  *buscarActividad(&ListaActividades, "Canopy"),
		},

		{
			Genero:     "Masculino",
			Edad:       24,
			Residencia: "Alajuela",
			Actividad:  *buscarActividad(&ListaActividades, "Pesca"),
		},
		{
			Genero:     "Femenino",
			Edad:       32,
			Residencia: "San Ramon",
			Actividad:  *buscarActividad(&ListaActividades, "Natacion"),
		},
		{
			Genero:     "Femenino",
			Edad:       39,
			Residencia: "Palmares",
			Actividad:  *buscarActividad(&ListaActividades, "Volcanes"),
		},
		{
			Genero:     "Masculino",
			Edad:       44,
			Residencia: "Ciudad Quesada",
			Actividad:  *buscarActividad(&ListaActividades, "Ciclismo"),
		},
		{
			Genero:     "Femenino",
			Edad:       48,
			Residencia: "Naranjo",
			Actividad:  *buscarActividad(&ListaActividades, "Visitar monumentos"),
		},
		{
			Genero:     "Masculino",
			Edad:       35,
			Residencia: "Alajuela",
			Actividad:  *buscarActividad(&ListaActividades, "Canopy"),
		},
		{
			Genero:     "Masculino",
			Edad:       46,
			Residencia: "Santa Rosa",
			Actividad:  *buscarActividad(&ListaActividades, "Parque de Diversiones"),
		},
		{
			Genero:     "Femenino",
			Edad:       45,
			Residencia: "Santa Clara",
			Actividad:  *buscarActividad(&ListaActividades, "Volcanes"),
		},
		{
			Genero:     "Femenino",
			Edad:       52,
			Residencia: "San Jose",
			Actividad:  *buscarActividad(&ListaActividades, "Miradero"),
		},
		{
			Genero:     "Femenino",
			Edad:       31,
			Residencia: "Alajuela",
			Actividad:  *buscarActividad(&ListaActividades, "Visitar museos"),
		},
		{
			Genero:     "Masculino",
			Edad:       27,
			Residencia: "San Ramon",
			Actividad:  *buscarActividad(&ListaActividades, "Paseos a caballo"),
		},
		{
			Genero:     "Femenino",
			Edad:       39,
			Residencia: "Alajuela",
			Actividad:  *buscarActividad(&ListaActividades, "Ciclismo"),
		},
		{
			Genero:     "Femenino",
			Edad:       48,
			Residencia: "Ciudad Quesada",
			Actividad:  *buscarActividad(&ListaActividades, "Visitar monumentos"),
		},
		{
			Genero:     "Masculino",
			Edad:       45,
			Residencia: "Santa Rosa",
			Actividad:  *buscarActividad(&ListaActividades, "Mariposario"),
		},
		{
			Genero:     "Femenino",
			Edad:       65,
			Residencia: "San Ramon",
			Actividad:  *buscarActividad(&ListaActividades, "Ciclismo"),
		},
		{
			Genero:     "Masculino",
			Edad:       23,
			Residencia: "Zarcero",
			Actividad:  *buscarActividad(&ListaActividades, "Cine"),
		},
	}
	//Añadir esos datos al archivo jason
	file, err := os.Create("personas.json")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(personas); err != nil {
		fmt.Println("Error al escribir en el archivo JSON:", err)
		return
	}
	//fmt.Println("Datos escritos en el archivo personas.json")

	// Cargar la lista de personas desde el archivo JSON
	var personasLoaded []Persona
	file, err = os.Open("personas.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&personasLoaded); err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return
	}

	// Ahora personasLoaded contiene la lista de personas recuperada desde el archivo JSON
	//or _, personaData := range personasLoaded {
	//fmt.Println("Persona:", personaData)
	//copia := personaData
	//cargarArbol(&arbol, copia)
	ModificarVertice(&grafo, &ListaActividades)
	//fmt.Println(rutaCortaConCategorias(buscarVertice("Zarcero", &grafo), "Naranjo", "", 0, 0, &grafo, []string{}))
}
