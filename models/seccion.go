package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// PostSecciones ...
func PostSecciones(secciones interface{}, Plantilla map[string]interface{}) (seccionesResult []map[string]interface{}, outputError interface{}) {
	seccionesMap, errMap := GetElementoMaptoStringToMapArray(secciones)
	if seccionesMap != nil {
		seccionesPadre, errSecPadre := IngresarSeccionesPadre(seccionesMap, Plantilla)
		if errSecPadre != nil {
			return nil, errSecPadre
		} else {
			return seccionesPadre, nil
		}
	} else {
		return nil, errMap
	}
}

// IngresarSeccionesPadre ...
func IngresarSeccionesPadre(secciones []map[string]interface{}, Plantilla map[string]interface{}) (seccionesResult []map[string]interface{}, outputError interface{}) {
	arraySeccionesIngresadas := make([]map[string]interface{}, 0)
	for i := 0; i < len(secciones); i++ {
		var seccionIngresada map[string]interface{}
		datoContruirdo := make(map[string]interface{})

		datoContruirdo = map[string]interface{}{
			"Activo": true,
			"Nombre": secciones[i]["Nombre"],
			"IdPlantilla": map[string]interface{}{
				"Id": Plantilla["Id"],
			},
			"SeccionHijaId": nil,
		}
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"seccion", "POST", &seccionIngresada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al ingresar el dato: ", secciones[i], " el error es:", error)
			return nil, error
		} else {
			seccionesHijasResult, errSecHija := IngresoSeccionHija(secciones[i], seccionIngresada, Plantilla)
			seccionIngresada["seccionesHijasIngresadas"] = seccionesHijasResult
			arraySeccionesIngresadas = append(arraySeccionesIngresadas, seccionIngresada)
			if (seccionesHijasResult == nil) && (errSecHija != nil) {
				return nil, errSecHija
			}
		}
	}
	return arraySeccionesIngresadas, nil
}

// IngresoSeccionHija ...
func IngresoSeccionHija(seccion map[string]interface{}, seccionPadre map[string]interface{}, Plantilla map[string]interface{}) (seccionesHijasResult []map[string]interface{}, outputError interface{}) {
	seccionMap, errMap := GetElementoMaptoStringToMapArray(seccion["Seccion_hija_id"])
	arraySeccionesHijasIngresadas := make([]map[string]interface{}, 0)

	if seccionMap != nil {
		for i := 0; i < len(seccionMap); i++ {
			var seccionHijaIngresada map[string]interface{}
			datoContruirdo := make(map[string]interface{})
			datoContruirdo = map[string]interface{}{
				"Activo": true,
				"Nombre": seccionMap[i]["Nombre"],
				"IdPlantilla": map[string]interface{}{
					"Id": Plantilla["Id"],
				},
				"SeccionPadreId": map[string]interface{}{
					"Id": seccionPadre["Id"],
				},
			}
			error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"seccion", "POST", &seccionHijaIngresada, datoContruirdo)
			if error != nil {
				logs.Error("Ocurrio un error al ingresar el dato: ", seccionMap[i], " el error es:", error)
				return nil, error
			} // else {
			itemsResult, errItems := PostItems(seccionMap[i], seccionHijaIngresada)
			if (itemsResult == nil) && (errItems != nil) {
				return nil, errItems
			}
			seccionHijaIngresada["ItemsIngresados"] = itemsResult
			arraySeccionesHijasIngresadas = append(arraySeccionesHijasIngresadas, seccionHijaIngresada)

			condicionesMap, errMapcondiciones := GetElementoMaptoStringToMapArray(seccionMap[i]["Condicion"])
			if condicionesMap != nil && errMapcondiciones == nil {
				condicionesIngresadas, errCondiciones := PostCondiciones(condicionesMap, arraySeccionesHijasIngresadas)
				if condicionesIngresadas != nil {
					seccionHijaIngresada["CondicionesIngresadas"] = condicionesIngresadas
				} else {
					return nil, errCondiciones
				}
			}

		}
		return arraySeccionesHijasIngresadas, nil

	} else {
		return nil, errMap
	}
}
