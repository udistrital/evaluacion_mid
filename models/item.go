package models

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/udistrital/utils_oas/request"
)

// PostItems ...
func PostItems(seccionConDatos map[string]interface{}, seccionHijaDB map[string]interface{}) (ItemsResult []map[string]interface{}, outputError map[string]interface{}) {
	itemsMap, errMap := GetElementoMaptoStringToMapArray(seccionConDatos["Item"])

	if itemsMap != nil {
		ItemsIngresados, errItems := IngresoItems(itemsMap, seccionHijaDB)
		if ItemsIngresados != nil {
			return ItemsIngresados, nil
		} else {
			logs.Error("error en items ingresados:", errItems)
			return nil, errItems
		}
	} else {
		return nil, errMap
	}
}

// IngresoItems ...
func IngresoItems(items []map[string]interface{}, SeccionDB map[string]interface{}) (itemsResult []map[string]interface{}, outputError map[string]interface{}) {
	arrayitemsIngresados := make([]map[string]interface{}, 0)

	for i := 0; i < len(items); i++ {
		var itemIngresado map[string]interface{}
		tipoItemDB, errorTipoItem := GetTipoItemParametrica(items[i]["IdTipoItem"].(map[string]interface{}))
		if tipoItemDB != nil {
			pipeDB, errorPipeDB := GetEstiloPipeParametrica(items[i]["IdEstiloPipe"].(map[string]interface{}))
			if pipeDB != nil {
				datoContruirdo := make(map[string]interface{})

				datoContruirdo = map[string]interface{}{
					"Activo": true,
					"Valor":  items[i]["Valor"],
					"Tamano": items[i]["Tamano"],
					"Nombre": items[i]["Nombre"],
					"IdTipoItem": map[string]interface{}{
						"Id": tipoItemDB[0]["Id"],
					},
					"IdEstiloPipe": map[string]interface{}{
						"Id": pipeDB[0]["Id"],
					},
					"IdSeccion": map[string]interface{}{
						"Id": SeccionDB["Id"],
					},
				}
				if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/item", "POST", &itemIngresado, datoContruirdo); err != nil {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/IngresoItems", "err": err.Error(), "status": "502"}
					return nil, outputError
				} else {
					itemIngresadoData := itemIngresado["Data"].(map[string]interface{})
					opcionesItemsMap, errMapOpciones := GetElementoMaptoStringToMapArray(items[i]["Opcion_item"])
					if opcionesItemsMap != nil && errMapOpciones == nil && len(opcionesItemsMap[0]) != 0 {
						opcionesIngresadas, errOp := PostOpcionesItem(opcionesItemsMap, itemIngresadoData)
						if opcionesIngresadas == nil && errOp != nil {
							return nil, errOp
						}
						itemIngresadoData["OpcionesIngresadas"] = opcionesIngresadas
					} else {
						itemIngresadoData["OpcionesIngresadas"] = nil
					}
					arrayitemsIngresados = append(arrayitemsIngresados, itemIngresadoData)
				}
			} else {
				//errorPipeDB := CrearError("no se pudo obtener el Pipe de estilo para el item" + fmt.Sprintf("%v", items[i]["Nombre"]))
				return nil, errorPipeDB
			}
		} else {
			//errorTipoItem := CrearError("no se pudo obtener el tipo de item para el item" + fmt.Sprintf("%v", items[i]["Nombre"]))
			return nil, errorTipoItem
		}
	}
	return arrayitemsIngresados, nil
}

// GetTipoItemParametrica ...
func GetTipoItemParametrica(tipoItem map[string]interface{}) (tipoItemResult []map[string]interface{}, outputError map[string]interface{}) {
	var tipoItemGet map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", tipoItem["Nombre"]) + ",CodigoAbreviacion:" + fmt.Sprintf("%v", tipoItem["CodigoAbreviacion"]) + ",Activo:true&limit=1"
	if response, err := getJsonTest(beego.AppConfig.String("evaluacion_crud_url")+"v1/tipo_item?query="+query, &tipoItemGet); (err == nil) && (response == 200) {
		aux := reflect.ValueOf(tipoItemGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(tipoItemGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &tipoItemResult); err == nil {
					return tipoItemResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetTipoItemParametrica4", "err": err.Error(), "status": "502"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/GetTipoItemParametrica3", "err": "Cantidad de elementos vacia", "status": "502"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/GetTipoItemParametrica2", "err": "Los valores no son validos", "status": "502"}
			return nil, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetTipoItemParametrica1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	/*if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(tipoItemGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return tipoItemGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}*/

}

// GetEstiloPipeParametrica ...
func GetEstiloPipeParametrica(pipe map[string]interface{}) (tipoItemResult []map[string]interface{}, outputError map[string]interface{}) {
	var estiloPipeGet map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", pipe["Nombre"]) + ",CodigoAbreviacion:" + fmt.Sprintf("%v", pipe["CodigoAbreviacion"]) + ",Activo:true&limit=1"
	if response, err := getJsonTest(beego.AppConfig.String("evaluacion_crud_url")+"v1/estilo_pipe?query="+query, &estiloPipeGet); (err == nil) && (response == 200) {
		aux := reflect.ValueOf(estiloPipeGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(estiloPipeGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &tipoItemResult); err == nil {
					return tipoItemResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetEstiloPipeParametrica4", "err": err.Error(), "status": "502"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/GetEstiloPipeParametrica3", "err": "Cantidad de elementos vacia", "status": "502"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/GetEstiloPipeParametrica2", "err": "Los valores no son validos", "status": "502"}
			return nil, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetEstiloPipeParametrica1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	/*if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(estiloPipeGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return estiloPipeGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}*/

}

// GetItems ...
func GetItems(seccion map[string]interface{}) (itemsResult []map[string]interface{}, outputError map[string]interface{}) {
	campos := "&fields=IdEstiloPipe,IdTipoItem,Nombre,Tamano,Valor,Id&sortby=Id&order=asc&limit=0"
	query := "?query=IdSeccion:" + fmt.Sprintf("%v", seccion["Id"]) + campos
	items, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("item", query)
	if items != nil {
		for i := 0; i < len(items); i++ {
			opcionesItem := GetOpciones(items[i])
			items[i]["Opcion_item"] = opcionesItem
		}
		return items, nil
	}
	return nil, errTablaCrudEvaluacion
}
