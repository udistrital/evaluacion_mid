package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// PostCondiciones ...
func PostCondiciones(condicionesMap []map[string]interface{}, arraySecciones []map[string]interface{}) (condicionesResult map[string]interface{}, outputError map[string]interface{}) {
	if len(arraySecciones) > 0 {
		for i := 0; i < len(condicionesMap); i++ {
			// se verifica si la seccion penultima es la de la condicion
			if fmt.Sprintf("%v", condicionesMap[i]["Nombre_seccion_condicion"]) == fmt.Sprintf("%v", arraySecciones[len(arraySecciones)-2]["Nombre"]) {
				opcionDB, errOpcParametrica := GetOpcionesParametrica(condicionesMap[i])
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
											} else if errCondicion != nil{
												return nil, errCondicion
											}
										}
									}
								} else if errMapOpciones != nil{
									return nil, errMapOpciones
								}
							}
						}
					} else if errMapItems != nil {
						return nil , errMapItems
					}

				}
				if errOpcParametrica != nil{
					return nil, errOpcParametrica
				}
			}
		}
	}

	return nil, nil
}

// PostCondicionDB ...
func PostCondicionDB(seccionHijaActual map[string]interface{}, seccionCondicion map[string]interface{}, opcionItem map[string]interface{}) (condicionResult map[string]interface{}, outputError map[string]interface{}) {
	var condicionIngresada map[string]interface{}
	datoContruirdo := make(map[string]interface{})
	datoContruirdo = map[string]interface{}{
		"Activo": true,
		"IdSeccion": map[string]interface{}{
			"Id": seccionHijaActual["Id"],
		},
		"OpcionItemId":         opcionItem["Id"].(float64),
		"SeccionDependenciaId": seccionCondicion["Id"].(float64),
	}
	if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/condicion", "POST", &condicionIngresada, datoContruirdo); err != nil{
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/PostCondicionDB", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return condicionIngresada, nil
	}
}

// GetCondiciones ...
func GetCondiciones(seccion map[string]interface{}) (condicionesResult []map[string]interface{}, outputError map[string]interface{}) {
	arrayCondiciones := make([]map[string]interface{}, 0)
	query := "?query=IdSeccion:" + fmt.Sprintf("%v", seccion["Id"])
	condicionesSeccion, errTablaCrudEvaluacion1 := GetTablaCrudEvaluacion("condicion", query)
	if condicionesSeccion != nil {
		aux := reflect.ValueOf(condicionesSeccion[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				querySeccion := "?query=Id:" + fmt.Sprintf("%v", condicionesSeccion[0]["SeccionDependenciaId"]) + "&fields=Nombre"

				seccionDeCondicion, errTablaCrudEvaluacion2 := GetTablaCrudEvaluacion("seccion", querySeccion)
				if seccionDeCondicion == nil {
					return nil, errTablaCrudEvaluacion2
				}
				queryOpcion := "?query=Id:" + fmt.Sprintf("%v", condicionesSeccion[0]["OpcionItemId"]) + "&fields=Nombre,Valor"
				opcionCondicion, errTablaCrudEvaluacion3 := GetTablaCrudEvaluacion("opciones", queryOpcion)
				if opcionCondicion == nil {
					return nil, errTablaCrudEvaluacion3
				}
				datoContruirdo := make(map[string]interface{})
				datoContruirdo = map[string]interface{}{
					"Nombre_seccion_condicion": seccionDeCondicion[0]["Nombre"],
					"Nombre":                   opcionCondicion[0]["Nombre"],
					"Valor":                    opcionCondicion[0]["Valor"],
				}
				arrayCondiciones = append(arrayCondiciones, datoContruirdo)
				return arrayCondiciones, nil
			}
		}

	} else if errTablaCrudEvaluacion1 != nil{
		return nil, errTablaCrudEvaluacion1
	}
	return arrayCondiciones, nil
}
