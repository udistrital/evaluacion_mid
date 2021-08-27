package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// PostSecciones ...
func PostSecciones(secciones interface{}, Plantilla map[string]interface{}) (seccionesResult []map[string]interface{}, outputError map[string]interface{}) {
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
func IngresarSeccionesPadre(secciones []map[string]interface{}, Plantilla map[string]interface{}) (seccionesResult []map[string]interface{}, outputError map[string]interface{}) {
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
		if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/seccion", "POST", &seccionIngresada, datoContruirdo); err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/IngresarSeccionesPadre", "err": err.Error(), "status": "502"}
			return nil, outputError
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
func IngresoSeccionHija(seccion map[string]interface{}, seccionPadre map[string]interface{}, Plantilla map[string]interface{}) (seccionesHijasResult []map[string]interface{}, outputError map[string]interface{}) {
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
			if err := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/seccion", "POST", &seccionHijaIngresada, datoContruirdo); err != nil {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/IngresoSeccionHija", "err": err.Error(), "status": "502"}
				return nil, outputError
			}

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

// GetSecciones ...
func GetSecciones(plantilla map[string]interface{}) (seccionesResult []map[string]interface{}, outputError map[string]interface{}) {
	ArraySeccionesPlantillaDB := make([]map[string]interface{}, 0)
	query := "?query=IdPlantilla:" + fmt.Sprintf("%v", plantilla["Id"]) + "&sortby=Id&order=asc&limit=0"
	seccionesPlantilla, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("seccion", query)
	if seccionesPlantilla != nil {
		for i := 0; i < len(seccionesPlantilla); i++ {
			if seccionesPlantilla[i]["SeccionPadreId"] == nil {
				queryHija := "?query=IdPlantilla:" + fmt.Sprintf("%v", plantilla["Id"]) + ",SeccionPadreId:" + fmt.Sprintf("%v", seccionesPlantilla[i]["Id"]) + "&sortby=Id&order=asc&limit=0"
				seccionesHijas, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("seccion", queryHija)
				if errTablaCrudEvaluacion == nil {
					for j := 0; j < len(seccionesHijas); j++ {
						condicion, errCondicion := GetCondiciones(seccionesHijas[j])
						if condicion != nil {
							seccionesHijas[j]["Condicion"] = condicion
						} else {
							return nil, errCondicion
						}
						itemSeccion, errItems := GetItems(seccionesHijas[j])
						if itemSeccion != nil {
							seccionesHijas[j]["Item"] = itemSeccion
						} else {
							return nil, errItems
						}
					}
					seccionesPlantilla[i]["Seccion_hija_id"] = seccionesHijas
					ArraySeccionesPlantillaDB = append(ArraySeccionesPlantillaDB, seccionesPlantilla[i])
				} else {
					return nil, errTablaCrudEvaluacion
				}
			}
		}
		// return seccionesPlantilla, nil
		return ArraySeccionesPlantillaDB, nil
	} else {
		return nil, errTablaCrudEvaluacion
	}
	/*error := CrearError("no se encontraron secciones para la plantilla")
	return nil, error*/
	// return nil, nil
}
