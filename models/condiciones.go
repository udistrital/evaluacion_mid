package models

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// PostCondiciones ...
func PostCondiciones(condicionesMap []map[string]interface{}, arraySecciones []map[string]interface{}) (condicionesResult map[string]interface{}, outputError interface{}) {
	if len(arraySecciones) > 0 {
		logs.Warning("tamaño: ", len(arraySecciones))
		for i := 0; i < len(condicionesMap); i++ {
			// se verifica si la seccion penultima es la de la condicion
			if fmt.Sprintf("%v", condicionesMap[i]["Nombre_seccion_condicion"]) == fmt.Sprintf("%v", arraySecciones[len(arraySecciones)-2]["Nombre"]) {
				fmt.Println("coincide seccion y condicion")
				opcionDB := GetOpcionesParametrica(condicionesMap[i])
				if opcionDB != nil {
					ItemIngresadosMap, errMapItems := GetElementoMaptoStringToMapArray(arraySecciones[len(arraySecciones)-2]["ItemsIngresados"])
					seccionComparacion := arraySecciones[len(arraySecciones)-2]
					seccionHijaActual := arraySecciones[len(arraySecciones)-1]
					if ItemIngresadosMap != nil && errMapItems == nil {
						for j := 0; j < len(ItemIngresadosMap); j++ {
							if ItemIngresadosMap[j]["OpcionesIngresadas"] != nil {
								OpcionesIngresadasMap, errMapOpciones := GetElementoMaptoStringToMapArray(ItemIngresadosMap[j]["OpcionesIngresadas"])
								if OpcionesIngresadasMap != nil && errMapOpciones == nil {
									for k := 0; k < len(OpcionesIngresadasMap); k++ {
										if OpcionesIngresadasMap[k]["IdOpciones"].(map[string]interface{})["Id"] == opcionDB[0]["Id"] {
											condicionIngresada, errCondicion := PostCondicionDB(seccionHijaActual, seccionComparacion, OpcionesIngresadasMap[k]["IdOpciones"].(map[string]interface{}))
											if condicionIngresada != nil && errCondicion == nil {
												return condicionIngresada, nil
											}
										}
									}
								}
							}
						}
					}

				}
			}
		}
	}

	return nil, nil
}

// PostCondicionDB ...
func PostCondicionDB(seccionHijaActual map[string]interface{}, seccionCondicion map[string]interface{}, opcionItem map[string]interface{}) (condicionResult map[string]interface{}, outputError interface{}) {
	var condicionIngresada map[string]interface{}
	datoContruirdo := make(map[string]interface{})
	// datotest := opcionItem["Id"].(float64)
	// fmt.Println(datotest)
	datoContruirdo = map[string]interface{}{
		"Activo": true,
		"IdSeccion": map[string]interface{}{
			"Id": seccionHijaActual["Id"],
		},
		"OpcionItemId":         opcionItem["Id"].(float64),
		"SeccionDependenciaId": seccionCondicion["Id"].(float64),
	}
	error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"condicion", "POST", &condicionIngresada, datoContruirdo)
	if error != nil {
		return nil, error
	} else {
		return condicionIngresada, nil
	}
}
