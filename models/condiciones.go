package models

import "fmt"

// funcion recibe el array para tener todas hasta el momento, buscara si la ultima es la de la condicion,
// tambien recibira la seccion que tiene la condicion
// las seeciones deben de tener a los items dentro de ellas ( hare lo posible), en caso dadu se buscaran los items
// se recibira el map de condiciones
// luego con los items y as opciones se buscara la tabla de rompimiento que cumpla la condicion, si no se encuentr manda error, si se encuentra se realiza post de condicion
// puedo reutilizar la busqueda de las opciones, el metod ya existe

// PostCondiciones ...
func PostCondiciones(condicionesMap []map[string]interface{}, arraySecciones []map[string]interface{}, seccionHijaActual map[string]interface{}) (condicionesResult []map[string]interface{}, outputError interface{}) {
	if len(arraySecciones) > 0 {
		// fmt.Println("nombre de anterior seccion", arraySecciones[len(arraySecciones)-2]["Nombre"])
		// fmt.Println("------------------------------------------------------")
		// fmt.Println("condiciones:", condicionesMap)
		// fmt.Println("------------------------------------------------------")

		// fmt.Println("nombre de ultima seccion", arraySecciones[len(arraySecciones)-1]["Nombre"])
		// fmt.Println("------------------------------------------------------")

		for i := 0; i < len(condicionesMap); i++ {
			// se verifica si la seccion penultima es la de la condicion
			if fmt.Sprintf("%v", condicionesMap[i]["Nombre_seccion_condicion"]) == fmt.Sprintf("%v", arraySecciones[len(arraySecciones)-2]["Nombre"]) {
				fmt.Println("coincide seccion y condicion")
				opcionDB := GetOpcionesParametrica(condicionesMap[i])
				if opcionDB != nil {

				}
			}
		}
	}

	return nil, nil
}
