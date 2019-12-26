package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// PostOpcionesItem ...
func PostOpcionesItem(opcionesItemMapeo []map[string]interface{}, itemDB map[string]interface{}) (ItemsResult []map[string]interface{}, outputError interface{}) {
	arrayOpcionesItemsIngresados := make([]map[string]interface{}, 0)
	for i := 0; i < len(opcionesItemMapeo); i++ {
		opcionParametrica := GetOpcionesParametrica(opcionesItemMapeo[i]["Id_opciones"].(map[string]interface{}))
		if opcionParametrica != nil {
			// se puede ingresar la de rompimiento
			opcionItemIngreso, erroOpIt := IngresoOpcionesItem(opcionParametrica[0], itemDB)
			if opcionItemIngreso == nil && erroOpIt != nil {
				return nil, erroOpIt
			}
			arrayOpcionesItemsIngresados = append(arrayOpcionesItemsIngresados, opcionItemIngreso)
		} else {
			postOpcionParametrica, errOpt := PostOpcionesParametrica(opcionesItemMapeo[i]["Id_opciones"].(map[string]interface{}))
			if postOpcionParametrica != nil {
				opcionItemIngreso, erroOpIt := IngresoOpcionesItem(postOpcionParametrica, itemDB)
				if opcionItemIngreso == nil && erroOpIt != nil {
					return nil, erroOpIt
				}
				arrayOpcionesItemsIngresados = append(arrayOpcionesItemsIngresados, opcionItemIngreso)

			} else {
				logs.Error("hubo error en ingresar la opcion:", opcionesItemMapeo[i]["Id_opciones"].(map[string]interface{}))
				logs.Error("el error presentado es: ", errOpt)
			}
		}
	}
	return arrayOpcionesItemsIngresados, nil
}

// GetOpcionesParametrica ...
func GetOpcionesParametrica(opciones map[string]interface{}) (opcionesResult []map[string]interface{}) {
	var opcionesGet []map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", opciones["Nombre"]) + ",Valor:" + fmt.Sprintf("%v", opciones["Valor"]) + ",Activo:true&limit=1"
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/opciones?query="+query, &opcionesGet)
	if error != nil {
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
	}

}

// PostOpcionesParametrica ...
func PostOpcionesParametrica(opcionEnviar map[string]interface{}) (opcionResult map[string]interface{}, outputError interface{}) {
	var opcionIngresada map[string]interface{}
	datoContruirdo := make(map[string]interface{})
	datoContruirdo = map[string]interface{}{
		"Activo": true,
		"Valor":  opcionEnviar["Valor"],
		"Nombre": opcionEnviar["Nombre"],
	}
	error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/opciones", "POST", &opcionIngresada, datoContruirdo)
	if error != nil {
		return nil, error
	} else {
		return opcionIngresada, nil
	}
}

// IngresoOpcionesItem ...
func IngresoOpcionesItem(opcionDB map[string]interface{}, itemDB map[string]interface{}) (ItemsResult map[string]interface{}, outputError interface{}) {
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
	error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/opcion_item", "POST", &opcionItemIngresada, datoContruirdo)
	if error != nil {
		return nil, error
	} else {
		return opcionItemIngresada, nil
	}
}

// GetOpciones ...
func GetOpciones(item map[string]interface{}) (condicionesResult []map[string]interface{}) {
	arrayVacio := make([]map[string]interface{}, 0)
	query := "?query=IdItem:" + fmt.Sprintf("%v", item["Id"]) + "&limit=0&fields=IdOpciones&sortby=Id&order=asc"
	opciones := GetTablaCrudEvaluacion("opcion_item", query)
	if opciones != nil {
		return opciones
	}
	return arrayVacio
}
