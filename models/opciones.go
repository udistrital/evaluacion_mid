package models

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/udistrital/utils_oas/request"
)

// PostOpcionesItem ...
func PostOpcionesItem(opcionesItemMapeo []map[string]interface{}, itemDB map[string]interface{}) (ItemsResult []map[string]interface{}, outputError map[string]interface{}) {
	arrayOpcionesItemsIngresados := make([]map[string]interface{}, 0)
	for i := 0; i < len(opcionesItemMapeo); i++ {
		opcionParametrica, errOpcParametrica := GetOpcionesParametrica(opcionesItemMapeo[i]["IdOpciones"].(map[string]interface{}))
		if opcionParametrica != nil {
			// se puede ingresar la de rompimiento
			opcionItemIngreso, erroOpIt := IngresoOpcionesItem(opcionParametrica[0], itemDB)
			if opcionItemIngreso == nil && erroOpIt != nil {
				return nil, erroOpIt
			}
			arrayOpcionesItemsIngresados = append(arrayOpcionesItemsIngresados, opcionItemIngreso)
		} else {
			if errOpcParametrica != nil {
				return nil, errOpcParametrica
			}
			postOpcionParametrica, errOpt := PostOpcionesParametrica(opcionesItemMapeo[i]["IdOpciones"].(map[string]interface{}))
			if postOpcionParametrica != nil {
				opcionItemIngreso, erroOpIt := IngresoOpcionesItem(postOpcionParametrica, itemDB)
				if opcionItemIngreso == nil && erroOpIt != nil {
					return nil, erroOpIt
				}
				arrayOpcionesItemsIngresados = append(arrayOpcionesItemsIngresados, opcionItemIngreso)

			} else {
				logs.Error("hubo error en ingresar la opcion:", opcionesItemMapeo[i]["IdOpciones"].(map[string]interface{}))
				logs.Error("el error presentado es: ", errOpt)
			}
		}
	}
	return arrayOpcionesItemsIngresados, nil
}

// GetOpcionesParametrica ...
func GetOpcionesParametrica(opciones map[string]interface{}) (opcionesResult []map[string]interface{}, outputError map[string]interface{}) {
	var opcionesGet map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", opciones["Nombre"]) + ",Valor:" + fmt.Sprintf("%v", opciones["Valor"]) + ",Activo:true&limit=1"
	if response, err := getJsonTest(beego.AppConfig.String("evaluacion_crud_url")+"v1/opciones?query="+query, &opcionesGet); (err == nil) && (response == 200) {
		aux := reflect.ValueOf(opcionesGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(opcionesGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &opcionesResult); err == nil {
					return opcionesResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetOpcionesParametrica4", "err": err.Error(), "status": "502"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/GetOpcionesParametrica3", "err": "Cantidad de elementos vacia", "status": "502"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/GetOpcionesParametrica2", "err": "Cantidad de elementos vacia", "status": "502"}
			return nil, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetOpcionesParametrica1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	/*if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(opcionesGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return opcionesGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}*/

}

// PostOpcionesParametrica ...
func PostOpcionesParametrica(opcionEnviar map[string]interface{}) (opcionResult map[string]interface{}, outputError map[string]interface{}) {
	var opcionIngresada map[string]interface{}
	datoContruirdo := make(map[string]interface{})
	datoContruirdo = map[string]interface{}{
		"Activo": true,
		"Valor":  opcionEnviar["Valor"],
		"Nombre": opcionEnviar["Nombre"],
	}
	if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/opciones", "POST", &opcionIngresada, datoContruirdo); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/PostOpcionesParametrica", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		opcionIngresadaData := opcionIngresada["Data"].(map[string]interface{})
		return opcionIngresadaData, nil
	}
}

// IngresoOpcionesItem ...
func IngresoOpcionesItem(opcionDB map[string]interface{}, itemDB map[string]interface{}) (ItemsResult map[string]interface{}, outputError map[string]interface{}) {
	var opcionItemIngresada map[string]interface{}
	datoContruirdo := make(map[string]interface{})
	datoContruirdo = map[string]interface{}{
		"Activo": true,
		"IdItem": map[string]interface{}{
			"Id": itemDB["Id"],
		},
		"IdOpciones": map[string]interface{}{
			"Id": opcionDB["Id"],
		},
	}
	if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/opcion_item", "POST", &opcionItemIngresada, datoContruirdo); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/IngresoOpcionesItem", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		opcionItemIngresadaData := opcionItemIngresada["Data"].(map[string]interface{})
		return opcionItemIngresadaData, nil
	}
}

// GetOpciones ...
func GetOpciones(item map[string]interface{}) (condicionesResult []map[string]interface{}) {
	arrayVacio := make([]map[string]interface{}, 0)
	query := "?query=IdItem:" + fmt.Sprintf("%v", item["Id"]) + "&limit=0&fields=IdOpciones&sortby=Id&order=asc"
	opciones, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("opcion_item", query)
	if opciones != nil {
		return opciones
	} else {
		logs.Error(errTablaCrudEvaluacion)
	}
	return arrayVacio
}
