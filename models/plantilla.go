package models

import (
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
	if errPlantilla != nil {
		return nil, errPlantilla
	} else {
		clasificacionesResult, errClasificaciones := PostClasificacion(plantillaArray[0]["Clasificacion"], plantillaBase)
		if clasificacionesResult != nil {
			// logs.Info("se ingresaron las clasificaciones con exito: ")
			// aqui ira el ingreso de la secciones
			seccionesResult, errSecciones := PostSecciones(plantillaArray[0]["Secciones"], plantillaBase)
			// fmt.Println(seccionesResult)
			if seccionesResult != nil {
				plantillaBase["SeccionesIngresadas"] = seccionesResult

				return plantillaBase, nil
				// return seccionesResult, nil
			} else {
				return nil, errSecciones
			}
		} else {
			logs.Error("error al ingresar alguna clasificacion: ", errClasificaciones)
			return nil, errClasificaciones
		}
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

// FinalizarPlantilla ... proceso en el cual todas las   plantillas anteriores pasan a estar inactivas y la creada actual quedara activa, es el paso final
func FinalizarPlantilla(plantillaCreada map[string]interface{}) (plantillaResult map[string]interface{}, outputError interface{}) {
	return nil, nil
}

// GetPantillasActivas ...
func GetPantillasActivas() (plantillaResult []map[string]interface{}) {
	var plantillasGet []map[string]interface{}
	query := "Activo:true"
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"clasificacion?query="+query, &plantillasGet)
	if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(plantillasGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return plantillasGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

}

// DesactivarPlantillas ...
func DesactivarPlantillas() {

}

// ActivarPlantilla ...
func ActivarPlantilla() {

}
