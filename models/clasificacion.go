package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// PostClasificacion ...
func PostClasificacion(clasificaciones interface{}, Plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError interface{}) {
	clasificacionesMap, errMap := GetElementoMaptoStringToMapArray(clasificaciones)
	ArrayClasificacionesDB := make([]map[string]interface{}, 0)
	if clasificacionesMap != nil {
		for i := 0; i < len(clasificacionesMap); i++ {
			getClasificacion := GetClasificacionParametrica(clasificacionesMap[i])
			if getClasificacion != nil {
				ArrayClasificacionesDB = append(ArrayClasificacionesDB, getClasificacion[0])

			} else {
				postClasificacion, errClasif := PostClasificacionParametrica(clasificacionesMap[i])
				if errClasif != nil {
					logs.Error("hubo error en ingresar la clasificacion:", clasificacionesMap[i])
					logs.Error("el error presentado es: ", errClasif)
				} else {
					ArrayClasificacionesDB = append(ArrayClasificacionesDB, postClasificacion)
				}
			}
		}
		clasificacionesPlantilla, errClsPln := PostClasificacionPlantilla(ArrayClasificacionesDB, Plantilla)
		if errClsPln != nil {
			return nil, errClsPln
		} else {
			return clasificacionesPlantilla, nil
		}
	} else {
		return nil, errMap
	}
}

// PostClasificacionParametrica ... ingresar en tabla
func PostClasificacionParametrica(clasificacionEnviar map[string]interface{}) (clasificacionesResult map[string]interface{}, outputError interface{}) {
	var clasificacionIngresada map[string]interface{}
	error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"clasificacion", "POST", &clasificacionIngresada, clasificacionEnviar)
	if error != nil {
		return nil, error
	} else {
		return clasificacionIngresada, nil
	}
}

// GetClasificacionParametrica ... saber si ya existe para no crearla de nuevo
func GetClasificacionParametrica(clasificacion map[string]interface{}) (clasificacionesResult []map[string]interface{}) {
	var clasificacionGet []map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", clasificacion["Nombre"]) + ",LimiteInferior:" + fmt.Sprintf("%v", clasificacion["LimiteInferior"]) + ",LimiteSuperior:" + fmt.Sprintf("%v", clasificacion["LimiteSuperior"]) + ",Activo:true&limit=1"
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"clasificacion?query="+query, &clasificacionGet)
	if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(clasificacionGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return clasificacionGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

}

// PostClasificacionPlantilla ... a tabla de rompimiento
func PostClasificacionPlantilla(clasificaciones []map[string]interface{}, Plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError interface{}) {
	ArrayClasificacionesPlantillaDB := make([]map[string]interface{}, 0)

	for i := 0; i < len(clasificaciones); i++ {
		var clasificacionPlantillaIngresada map[string]interface{}
		datoContruirdo := make(map[string]interface{})

		datoContruirdo = map[string]interface{}{
			"Activo": true,
			"IdClasificacion": map[string]interface{}{
				"Id": clasificaciones[i]["Id"],
			},
			"IdPlantilla": map[string]interface{}{
				"Id": Plantilla["Id"],
			},
		}
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"clasificacion_plantilla", "POST", &clasificacionPlantillaIngresada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al ingresar el dato: ", clasificaciones[i], " el error es:", error)
			return nil, error
		} else {
			ArrayClasificacionesPlantillaDB = append(ArrayClasificacionesPlantillaDB, clasificacionPlantillaIngresada)
			// return clasificacionPlantillaIngresada, nil
		}
	}

	return ArrayClasificacionesPlantillaDB, nil
}

// GetClasicacionesPlntilla ...
func GetClasicacionesPlntilla(plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError interface{}) {
	ArrayClasificacionesPlantillaDB := make([]map[string]interface{}, 0)
	query := "?query=IdPlantilla:" + fmt.Sprintf("%v", plantilla["Id"])
	clasificacionesPlantilla := GetTablaCrudEvaluacion("clasificacion_plantilla", query)
	if clasificacionesPlantilla != nil {
		for i := 0; i < len(clasificacionesPlantilla); i++ {
			ArrayClasificacionesPlantillaDB = append(ArrayClasificacionesPlantillaDB, clasificacionesPlantilla[i]["IdClasificacion"].(map[string]interface{}))
		}
		return ArrayClasificacionesPlantillaDB, nil
	}
	error := CrearError("no se encontraron clasificaciones para la plantilla")

	return nil, error
}
