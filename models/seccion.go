package models

import (
	"fmt"

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
		fmt.Println("valio verga", errMap)
		return nil, errMap
	}
}

// IngresarSeccionesPadre ..4.
func IngresarSeccionesPadre(secciones []map[string]interface{}, Plantilla map[string]interface{}) (seccionesResult []map[string]interface{}, outputError interface{}) {
	arraySeccionesIngresadas := make([]map[string]interface{}, 0)
	var seccionIngresada map[string]interface{}
	for i := 0; i < len(secciones); i++ {
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
			arraySeccionesIngresadas = append(arraySeccionesIngresadas, seccionIngresada)
			seccionesHijasResult, errSecHija := IngresoSeccionHija(secciones[i], seccionIngresada, Plantilla)
			if errSecHija != nil {
				logs.Error("no hijas")
			} else {
				logs.Info("hijas:", seccionesHijasResult)
			}

			// return clasificacionPlantillaIngresada, nil
		}
	}
	return arraySeccionesIngresadas, nil
}

// IngresoSeccionHija ...
func IngresoSeccionHija(seccion map[string]interface{}, seccionPadre map[string]interface{}, Plantilla map[string]interface{}) (seccionesHijasResult []map[string]interface{}, outputError interface{}) {
	seccionMap, errMap := GetElementoMaptoStringToMapArray(seccion["Seccion_hija_id"])
	var seccionHijaIngresada map[string]interface{}
	arraySeccionesHijasIngresadas := make([]map[string]interface{}, 0)

	if seccionMap != nil {
		for i := 0; i < len(seccionMap); i++ {
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
			} else {
				arraySeccionesHijasIngresadas = append(arraySeccionesHijasIngresadas, seccionHijaIngresada)
				// AQUI SE INGRESARA LOS ITEMS
			}
		}
		return arraySeccionesHijasIngresadas, nil

	} else {
		fmt.Println("valio verga", errMap)
		return nil, errMap
	}
}
