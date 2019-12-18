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
	if errPlantilla != nil {
		return nil, errPlantilla
	} else {
		clasificacionesResult, errClasificaciones := PostClasificacion(plantillaArray[0]["Clasificacion"], plantillaBase)
		if clasificacionesResult != nil {
			// aqui ira el ingreso de la secciones
			seccionesResult, errSecciones := PostSecciones(plantillaArray[0]["Secciones"], plantillaBase)
			if seccionesResult != nil {
				plantillaBase["SeccionesIngresadas"] = seccionesResult
				platillaRetono, errRetorno := FinalizarPlantilla(plantillaBase)
				if platillaRetono != nil {
					return platillaRetono, nil
				} else {
					return nil, errRetorno
				}
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
	plantillasObtenidas := GetPlantillasActivas()
	if plantillasObtenidas != nil {
		plantillasDesactivadas, errDesactivar := DesactivarPlantillas(plantillasObtenidas)
		if plantillasDesactivadas != nil {
			plantillaActivada, errActivar := ActivarPlantilla(plantillaCreada)
			if plantillaActivada != nil {
				return plantillaActivada, nil
			} else {
				return nil, errActivar
			}
		} else {
			return nil, errDesactivar
		}
	} else {
		plantillaActivada, errActivar := ActivarPlantilla(plantillaCreada)
		if plantillaActivada != nil {
			return plantillaActivada, nil
		} else {
			return nil, errActivar
		}
	}
}

// GetPlantillasActivas ...
func GetPlantillasActivas() (plantillaResult []map[string]interface{}) {
	var plantillasGet []map[string]interface{}
	query := "Activo:true"
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla?query="+query, &plantillasGet)
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
func DesactivarPlantillas(plantillasActivas []map[string]interface{}) (plantillasResult []map[string]interface{}, outputError interface{}) {
	arrayPlantillasIngresadas := make([]map[string]interface{}, 0)

	for i := 0; i < len(plantillasActivas); i++ {
		var platillaActualizada map[string]interface{}

		datoContruirdo := make(map[string]interface{})
		datoContruirdo = map[string]interface{}{
			"Activo":        false,
			"Descripcion":   plantillasActivas[i]["Descripcion"],
			"FechaCreacion": plantillasActivas[i]["FechaCreacion"],
			"Id":            plantillasActivas[i]["Id"],
			"Usuario":       plantillasActivas[i]["Usuario"],
		}
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla/"+fmt.Sprintf("%v", plantillasActivas[i]["Id"]), "PUT", &platillaActualizada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al desactivar una plantilla: ", platillaActualizada, " el error es:", error)
			return nil, error
		}
		arrayPlantillasIngresadas = append(arrayPlantillasIngresadas, platillaActualizada)
	}
	return arrayPlantillasIngresadas, nil
}

// ActivarPlantilla ...
func ActivarPlantilla(plantillasParaActivar map[string]interface{}) (plantillasResult map[string]interface{}, outputError interface{}) {
	var platillaActualizada map[string]interface{}
	plantillaGet := GetPlantilla(plantillasParaActivar)
	if plantillaGet != nil {
		datoContruirdo := make(map[string]interface{})
		datoContruirdo = map[string]interface{}{
			"Activo":        true,
			"Descripcion":   plantillaGet[0]["Descripcion"],
			"FechaCreacion": plantillaGet[0]["FechaCreacion"],
			"Id":            plantillaGet[0]["Id"],
			"Usuario":       plantillaGet[0]["Usuario"],
		}
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla/"+fmt.Sprintf("%v", plantillaGet[0]["Id"]), "PUT", &platillaActualizada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al activar la plantilla: ", platillaActualizada, " el error es:", error)
			return nil, error
		}
		return platillaActualizada, nil
	}
	return nil, nil
}

// GetPlantilla ...
func GetPlantilla(plantilla map[string]interface{}) (plantillaResult []map[string]interface{}) {
	var plantillaGet []map[string]interface{}
	query := "Id:" + fmt.Sprintf("%v", plantilla["Id"])
	error := request.GetJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla?query="+query, &plantillaGet)
	if error != nil {
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(plantillaGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return plantillaGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

}

// ObtenerPlantillas ...
func ObtenerPlantillas() (plantillaResult map[string]interface{}, outputError interface{}) {
	var plantillaConstruida map[string]interface{}
	query := "?query=Activo:true"
	plantillaActiva := GetTablaCrudEvaluacion("plantilla", query)
	if plantillaActiva != nil {
		plantillaConstruida = plantillaActiva[0]
		fmt.Println("tenemos plantilla")
		clasificaciones, errClasificaciones := GetClasicacionesPlntilla(plantillaConstruida)
		if clasificaciones != nil {
			plantillaConstruida["Clasificaciones"] = clasificaciones
			return plantillaConstruida, nil
		}
		return nil, errClasificaciones
	}
	fmt.Println("no tenemos plantilla")
	error := CrearError("no se encontraron plantillas")
	return nil, error
}
