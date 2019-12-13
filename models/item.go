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
			// logs.Info("si se ingresaron los items:", ItemsIngresados)
			return ItemsIngresados, nil
		} else {
			logs.Error("error en items ingresados:", errItems)
			return nil, errItems
		}
	} else {
		fmt.Println("valio verga", errMap)
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
					// fmt.Println(datoContruirdo)
					error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"item", "POST", &itemIngresado, datoContruirdo)
					if error != nil {
						logs.Error("Ocurrio un error al ingresar el dato: ", itemIngresado, " el error es:", error)
						return nil, error
					} else {
						opcionesItemsMap, errMapOpciones := GetElementoMaptoStringToMapArray(items[i]["Opcion_item"])
						if opcionesItemsMap != nil {
							// logs.Info("si hay opciones a ingresar")
							opcionesIngresadas, errOp := PostOpcionesItem(opcionesItemsMap, itemIngresado)
							if opcionesIngresadas == nil && errOp != nil {
								return nil, errOp
							}
							itemIngresado["OpcionesIngresadas"] = opcionesIngresadas
						} else {
							logs.Error("no hay opciones a ingresar", errMapOpciones)
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
			fmt.Println("valio verga", errTipoMap)
			return nil, errTipoMap
		}

	}
	return arrayitemsIngresados, nil
}

// GetTipoItemParametrica ...
func GetTipoItemParametrica(tipoItem map[string]interface{}) (tipoItemResult []map[string]interface{}) {
	var tipoItemGet []map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", tipoItem["Nombre"]) + ",CodigoAbreviacion:" + fmt.Sprintf("%v", tipoItem["CodigoAbreviacion"]) + ",Activo:true&limit=1"
	// fmt.Println("query", query)
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"tipo_item?query="+query, &tipoItemGet)
	if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(tipoItemGet[0])
		// fmt.Println("aux: ", aux.Len())
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
	// fmt.Println("query", query)
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"estilo_pipe?query="+query, &estiloPipeGet)
	if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(estiloPipeGet[0])
		// fmt.Println("aux: ", aux.Len())
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
