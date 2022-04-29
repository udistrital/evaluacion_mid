package models

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// PostClasificacion ...
func PostClasificacion(clasificaciones interface{}, Plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError map[string]interface{}) {
	clasificacionesMap, errMap := GetElementoMaptoStringToMapArray(clasificaciones)
	ArrayClasificacionesDB := make([]map[string]interface{}, 0)
	if clasificacionesMap != nil {
		for i := 0; i < len(clasificacionesMap); i++ {
			getClasificacion, err1 := GetClasificacionParametrica(clasificacionesMap[i])
			if err1 == nil {
				if getClasificacion != nil {
					ArrayClasificacionesDB = append(ArrayClasificacionesDB, getClasificacion[0])
				} else {
					postClasificacion, errClasif := PostClasificacionParametrica(clasificacionesMap[i])
					if errClasif != nil {
						logs.Error("hubo error en ingresar la clasificacion:", clasificacionesMap[i])
						logs.Error("el error presentado es: ", errClasif)
						return nil, errClasif
					} else {
						ArrayClasificacionesDB = append(ArrayClasificacionesDB, postClasificacion)
					}
				}
			} else {
				return nil, err1
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
func PostClasificacionParametrica(clasificacionEnviar map[string]interface{}) (clasificacionesResult map[string]interface{}, outputError map[string]interface{}) {
	var clasificacionIngresada map[string]interface{}
	if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/clasificacion", "POST", &clasificacionIngresada, clasificacionEnviar); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/PostClasificacionParametrica", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return clasificacionIngresada, nil
	}
}

// GetClasificacionParametrica ... saber si ya existe para no crearla de nuevo
func GetClasificacionParametrica(clasificacion map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError map[string]interface{}) {
	var clasificacionGet map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", clasificacion["Nombre"]) + ",LimiteInferior:" + fmt.Sprintf("%v", clasificacion["LimiteInferior"]) + ",LimiteSuperior:" + fmt.Sprintf("%v", clasificacion["LimiteSuperior"]) + ",Activo:true&limit=1"
	//error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/clasificacion?query="+query, &clasificacionGet)
	if response, err1 := getJsonTest(beego.AppConfig.String("evaluacion_crud_url")+"v1/clasificacion?query="+query, &clasificacionGet); (err1 == nil) && (response == 200) {
		aux := reflect.ValueOf(clasificacionGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(clasificacionGet["Data"].([]interface{}))
				if err2 := json.Unmarshal(temp, &clasificacionesResult); err2 == nil {
					return clasificacionesResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetClasificacionParametrica4", "err": err2.Error(), "status": "204"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/GetClasificacionParametrica3", "err": "La longitud de los datos obtenidos es 0", "status": "502"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/GetClasificacionParametrica2", "err": "El valor no es valido", "status": "502"}
			return nil, outputError
		}

	} else {
		logs.Error(err1)
		outputError = map[string]interface{}{"funcion": "/GetClasificacionParametrica1", "err": err1.Error(), "status": "502"}
		return nil, outputError
	}
	/*if error != nil {
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
	}*/

}

// PostClasificacionPlantilla ... a tabla de rompimiento
func PostClasificacionPlantilla(clasificaciones []map[string]interface{}, Plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError map[string]interface{}) {
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
		if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/clasificacion_plantilla", "POST", &clasificacionPlantillaIngresada, datoContruirdo); err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/PostClasificacionPlantilla", "err": err.Error(), "status": "502"}
			return nil, outputError
		} else {
			ArrayClasificacionesPlantillaDB = append(ArrayClasificacionesPlantillaDB, clasificacionPlantillaIngresada)
		}
	}

	return ArrayClasificacionesPlantillaDB, nil
}

// GetClasicacionesPlntilla ...
func GetClasicacionesPlntilla(plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError map[string]interface{}) {
	ArrayClasificacionesPlantillaDB := make([]map[string]interface{}, 0)
	query := "?query=IdPlantilla:" + fmt.Sprintf("%v", plantilla["Id"])
	clasificacionesPlantilla, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("clasificacion_plantilla", query)
	if clasificacionesPlantilla != nil {
		for i := 0; i < len(clasificacionesPlantilla); i++ {
			ArrayClasificacionesPlantillaDB = append(ArrayClasificacionesPlantillaDB, clasificacionesPlantilla[i]["IdClasificacion"].(map[string]interface{}))
		}
		return ArrayClasificacionesPlantillaDB, nil
	} else {
		return nil, errTablaCrudEvaluacion
	}
}
