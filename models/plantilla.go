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

// PostClasificacion ...
func PostClasificacion(clasificaciones interface{}, Plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError interface{}) {
	clasificacionesMap, errMap := GetElementoMaptoStringToMapArray(clasificaciones)
	ArrayClasificacionesDB := make([]map[string]interface{}, 0)
	if clasificacionesMap != nil {
		for i := 0; i < len(clasificacionesMap); i++ {
			getClasificacion := GetClasificacionParametrica(clasificacionesMap[i])
			if getClasificacion != nil {
				ArrayClasificacionesDB = append(ArrayClasificacionesDB, getClasificacion[0])

			} else {
				postClasificacion, errClasif := PostClasificacionParametrica(clasificacionesMap[i])
				if errClasif != nil {
					logs.Error("hubo error en ingresar la clasificacion:", clasificacionesMap[i])
					logs.Error("el error presentado es: ", errClasif)
				} else {
					ArrayClasificacionesDB = append(ArrayClasificacionesDB, postClasificacion)
				}
			}
		}
		clasificacionesPlantilla, errClsPln := PostClasificacionPlantilla(ArrayClasificacionesDB, Plantilla)
		if errClsPln != nil {
			return nil, errClsPln
		} else {
			return clasificacionesPlantilla, nil
		}
	} else {
		fmt.Println("valio verga", errMap)
		return nil, errMap
	}
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
func PostClasificacionPlantilla(clasificaciones []map[string]interface{}, Plantilla map[string]interface{}) (clasificacionesResult []map[string]interface{}, outputError interface{}) {
	var clasificacionPlantillaIngresada map[string]interface{}
	ArrayClasificacionesPlantillaDB := make([]map[string]interface{}, 0)

	for i := 0; i < len(clasificaciones); i++ {
		datoContruirdo := make(map[string]interface{})

		datoContruirdo = map[string]interface{}{
			"Activo": true,
			"IdClasificacion": map[string]interface{}{
				"Id": clasificaciones[i]["Id"],
			},
			"IdPlantilla": map[string]interface{}{
				"Id": Plantilla["Id"],
			},
		}
		error := request.SendJson(beego.AppConfig.String("evaluacion_crud_url")+"clasificacion_plantilla", "POST", &clasificacionPlantillaIngresada, datoContruirdo)
		if error != nil {
			logs.Error("Ocurrio un error al ingresar el dato: ", clasificaciones[i], " el error es:", error)
			return nil, error
		} else {
			ArrayClasificacionesPlantillaDB = append(ArrayClasificacionesPlantillaDB, clasificacionPlantillaIngresada)
			// return clasificacionPlantillaIngresada, nil
		}
	}

	return ArrayClasificacionesPlantillaDB, nil
}
