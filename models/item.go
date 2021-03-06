package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

// PostItems ...
func PostItems(seccionConDatos map[string]interface{}, seccionHijaDB map[string]interface{}) (ItemsResult []map[string]interface{}, outputError interface{}) {
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
func IngresoItems(items []map[string]interface{}, SeccionDB map[string]interface{}) (itemsResult []map[string]interface{}, outputError interface{}) {
	arrayitemsIngresados := make([]map[string]interface{}, 0)

	for i := 0; i < len(items); i++ {
		var itemIngresado map[string]interface{}
		tipoItemMap, errTipoMap := GetElementoMaptoStringToMapArray(items[i]["Id_tipo_item"])
		if tipoItemMap != nil {
			tipoItemDB := GetTipoItemParametrica(tipoItemMap[0])
			if tipoItemDB != nil {
				pipeDB := GetEstiloPipeParametrica(items[i]["Estilo_pipe_id"].(map[string]interface{}))
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
						"SeccionHijaId": nil,
					}
					error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/item", "POST", &itemIngresado, datoContruirdo)
					if error != nil {
						logs.Error("Ocurrio un error al ingresar el dato: ", itemIngresado, " el error es:", error)
						return nil, error
					} else {
						opcionesItemsMap, errMapOpciones := GetElementoMaptoStringToMapArray(items[i]["Opcion_item"])
						if opcionesItemsMap != nil && errMapOpciones == nil {
							opcionesIngresadas, errOp := PostOpcionesItem(opcionesItemsMap, itemIngresado)
							if opcionesIngresadas == nil && errOp != nil {
								return nil, errOp
							}
							itemIngresado["OpcionesIngresadas"] = opcionesIngresadas
						} else {
							itemIngresado["OpcionesIngresadas"] = nil
						}
						arrayitemsIngresados = append(arrayitemsIngresados, itemIngresado)

					}
				} else {
					errorPipeDB := CrearError("no se pudo obtener el Pipe de estilo para el item" + fmt.Sprintf("%v", items[i]["Nombre"]))
					return nil, errorPipeDB
				}
			} else {
				errorTipoItem := CrearError("no se pudo obtener el tipo de item para el item" + fmt.Sprintf("%v", items[i]["Nombre"]))
				return nil, errorTipoItem
			}

		} else {
			return nil, errTipoMap
		}

	}
	return arrayitemsIngresados, nil
}

// GetTipoItemParametrica ...
func GetTipoItemParametrica(tipoItem map[string]interface{}) (tipoItemResult []map[string]interface{}) {
	var tipoItemGet []map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", tipoItem["Nombre"]) + ",CodigoAbreviacion:" + fmt.Sprintf("%v", tipoItem["CodigoAbreviacion"]) + ",Activo:true&limit=1"
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/tipo_item?query="+query, &tipoItemGet)
	if error != nil {
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
	}

}

// GetEstiloPipeParametrica ...
func GetEstiloPipeParametrica(pipe map[string]interface{}) (tipoItemResult []map[string]interface{}) {
	var estiloPipeGet []map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", pipe["Nombre"]) + ",CodigoAbreviacion:" + fmt.Sprintf("%v", pipe["CodigoAbreviacion"]) + ",Activo:true&limit=1"
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/estilo_pipe?query="+query, &estiloPipeGet)
	if error != nil {
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
	}

}

// GetItems ...
func GetItems(seccion map[string]interface{}) (itemsResult []map[string]interface{}, outputError interface{}) {
	campos := "&fields=IdEstiloPipe,IdTipoItem,Nombre,Tamano,Valor,Id&sortby=Id&order=asc&limit=0"
	query := "?query=IdSeccion:" + fmt.Sprintf("%v", seccion["Id"]) + campos
	items := GetTablaCrudEvaluacion("item", query)
	if items != nil {
		for i := 0; i < len(items); i++ {
			opcionesItem := GetOpciones(items[i])
			items[i]["Opcion_item"] = opcionesItem
		}
		return items, nil
	}
	error := CrearError("no se encontraron los items de la seccion")
	return nil, error
}
