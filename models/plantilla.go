package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// IngresoPlantilla ...
func IngresoPlantilla(plantilla map[string]interface{}) (plantillaResult map[string]interface{}, outputError interface{}) {
	plantillaIngresada := make(map[string]interface{})
	plantillaArray := make([]map[string]interface{}, 0)
	plantillaArray = append(plantillaArray, plantilla)
	plantillaIngresada = map[string]interface{}{
		"Activo":      plantillaArray[0]["Activo"],
		"Descripcion": plantillaArray[0]["Descripcion"],
		"Usuario":     plantillaArray[0]["Usuario"],
	}
	plantillaBase, errPlantilla := PostPlantilla(plantillaIngresada)
	fmt.Println(plantillaBase["Id"])
	if errPlantilla != nil {
		return nil, errPlantilla
	} else {
		clasificacionesResult, errClasificaciones := PostClasificacion(plantillaArray[0]["Clasificacion"])
		fmt.Println(clasificacionesResult)
		fmt.Println(errClasificaciones)
		return plantillaBase, nil

	}
}

// PostPlantilla ...
func PostPlantilla(plantilla map[string]interface{}) (plantillaResult map[string]interface{}, outputError interface{}) {
	var plantillaPost map[string]interface{}
	error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla", "POST", &plantillaPost, plantilla)
	if error != nil {
		return nil, error
	} else {
		return plantillaPost, nil
	}
}

// PostClasificacion ...
func PostClasificacion(clasificaciones interface{}) (clasificacionesResult map[string]interface{}, outputError interface{}) {
	fmt.Println(clasificaciones)
	// fmt.Println(clasificaciones.(map[string]interface{}))
	clasificacionesMap, errMap := GetElementoMaptoStringToMapArray(clasificaciones)
	ArrayClasificacionesDB := make([]map[string]interface{}, 0)
	if clasificacionesMap != nil {
		// fmt.Println(clasificacionesMap)
		// fmt.Println(clasificacionesMap[0]["Nombre"])
		for i := 0; i < len(clasificacionesMap); i++ {
			getClasificacion := GetClasificacionParametrica(clasificacionesMap[i])
			if getClasificacion != nil {
				logs.Info("si existe clasificacion para", clasificacionesMap[i]["Nombre"])
				ArrayClasificacionesDB = append(ArrayClasificacionesDB, getClasificacion[0])

			} else {
				logs.Info("NO existe clasificacion para", clasificacionesMap[i]["Nombre"])
				postClasificacion, errClasif := PostClasificacionParametrica(clasificacionesMap[i])
				if errClasif != nil {
					logs.Error("hubo error en ingresar la clasificacion:", clasificacionesMap[i])
					logs.Error("el error presentado es: ", errClasif)
				} else {
					ArrayClasificacionesDB = append(ArrayClasificacionesDB, postClasificacion)
				}
			}
		}
		logs.Info("clasificaciones: ", ArrayClasificacionesDB)
	} else {
		fmt.Println("valio verga", errMap)
	}
	// for _, clasificacion := range clasificaciones.(map[string]interface{})) {
	// 	fmt.Println(clasificacion)
	// }
	return nil, nil
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
	// fmt.Println(clasificacion["Nombre"])
	// var infoClasificacion []map[string]interface{}
	query := "Nombre:" + fmt.Sprintf("%v", clasificacion["Nombre"]) + ",LimiteInferior:" + fmt.Sprintf("%v", clasificacion["LimiteInferior"]) + ",LimiteSuperior:" + fmt.Sprintf("%v", clasificacion["LimiteSuperior"]) + ",Activo:true&limit=1"
	// fmt.Println("query", query)
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"clasificacion?query="+query, &clasificacionGet)
	if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(clasificacionGet[0])
		// fmt.Println("aux: ", aux.Len())
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
func PostClasificacionPlantilla() {

}
