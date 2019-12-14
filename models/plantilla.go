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
			// logs.Info("se ingresaron las clasificaciones con exito: ")
			// aqui ira el ingreso de la secciones
			seccionesResult, errSecciones := PostSecciones(plantillaArray[0]["Secciones"], plantillaBase)
			// fmt.Println(seccionesResult)
			if seccionesResult != nil {
				plantillaBase["SeccionesIngresadas"] = seccionesResult
				platillaRetono, errRetorno := FinalizarPlantilla(plantillaBase)
				if platillaRetono != nil {
					logs.Info("finalizacion completa")
					return platillaRetono, nil

				} else {
					logs.Error("error en finalizacion")
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
		fmt.Println("si habian plantillas true")
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
		fmt.Println("NO habian plantillas true")
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
				logs.Warning(plantillasGet)
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
	fmt.Println("desactivar plantillas")
	arrayPlantillasIngresadas := make([]map[string]interface{}, 0)

	for i := 0; i < len(plantillasActivas); i++ {
		var platillaActualizada map[string]interface{}

		datoContruirdo := make(map[string]interface{})
		logs.Emergency(fmt.Sprintf("%v", plantillasActivas[i]["Id"]))
		// logs.Emergency(time_bogota.TiempoCorreccionFormato(fmt.Sprintf("%v", plantillasActivas[i]["FechaCreacion"])))
		datoContruirdo = map[string]interface{}{
			"Activo":        false,
			"Descripcion":   plantillasActivas[i]["Descripcion"],
			"FechaCreacion": plantillasActivas[i]["FechaCreacion"],
			"Id":            plantillasActivas[i]["Id"],
			"Usuario":       plantillasActivas[i]["Usuario"],
		}
		// fmt.Println(datoContruirdo)
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla/"+fmt.Sprintf("%v", plantillasActivas[i]["Id"]), "PUT", &platillaActualizada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al desactivar una plantilla: ", platillaActualizada, " el error es:", error)
			return nil, error
		}
		logs.Info("platilla desactivada: ", platillaActualizada)
		arrayPlantillasIngresadas = append(arrayPlantillasIngresadas, platillaActualizada)
	}
	return arrayPlantillasIngresadas, nil
}

// ActivarPlantilla ...
func ActivarPlantilla(plantillasParaActivar map[string]interface{}) (plantillasResult map[string]interface{}, outputError interface{}) {
	fmt.Println("activar plantilla")
	var platillaActualizada map[string]interface{}
	plantillaGet := GetPlantilla(plantillasParaActivar)
	if plantillaGet != nil {
		datoContruirdo := make(map[string]interface{})
		logs.Emergency(plantillaGet[0]["FechaCreacion"])
		datoContruirdo = map[string]interface{}{
			"Activo":        true,
			"Descripcion":   plantillaGet[0]["Descripcion"],
			"FechaCreacion": plantillaGet[0]["FechaCreacion"],
			"Id":            plantillaGet[0]["Id"],
			"Usuario":       plantillaGet[0]["Usuario"],
		}
		// fmt.Println(datoContruirdo)
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"plantilla/"+fmt.Sprintf("%v", plantillaGet[0]["Id"]), "PUT", &platillaActualizada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al activar la plantilla: ", platillaActualizada, " el error es:", error)
			return nil, error
		}
		logs.Info("plantilla activada", platillaActualizada)
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
				logs.Warning(plantillaGet)
				return plantillaGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

}
