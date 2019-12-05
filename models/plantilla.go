package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	fmt.Println(plantillaBase)
	if errPlantilla != nil {
		return nil, errPlantilla
	} else {
		return plantillaBase, nil
	}
}

// PostPlantilla ...
func PostPlantilla(plantilla map[string]interface{}) (plantillaResult map[string]interface{}, outputError interface{}) {
	var plantillaPost map[string]interface{}
	error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla", "POST", &plantillaPost, plantilla)
	logs.Error(error)
	logs.Info(plantillaPost)
	if error != nil {
		return nil, error
	} else {
		return plantillaPost, nil
	}
}
