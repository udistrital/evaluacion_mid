package models

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
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
			// Ingreso de las secciones
			seccionesResult, errSecciones := PostSecciones(plantillaArray[0]["Secciones"], plantillaBase)
			if seccionesResult != nil {
				plantillaBase["SeccionesIngresadas"] = seccionesResult
				plantillaRetorno, errRetorno := FinalizarPlantilla(plantillaBase)
				if plantillaRetorno != nil {
					return plantillaRetorno, nil
				} else {
					return nil, errRetorno
				}
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
	if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/plantilla", "POST", &plantillaPost, plantilla); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/PostPlantilla", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		plantillaResultante := plantillaPost["Data"].(map[string]interface{})
		return plantillaResultante, nil
	}
}

// FinalizarPlantilla ... proceso en el cual todas las   plantillas anteriores pasan a estar inactivas y la creada actual quedara activa, es el paso final
func FinalizarPlantilla(plantillaCreada map[string]interface{}) (plantillaResult map[string]interface{}, outputError map[string]interface{}) {
	plantillasObtenidas, errPlantActivas := GetPlantillasActivas()
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
		if errPlantActivas != nil {
			return nil, errPlantActivas
		}
		plantillaActivada, errActivar := ActivarPlantilla(plantillaCreada)
		if plantillaActivada != nil {
			return plantillaActivada, nil
		} else {
			return nil, errActivar
		}
	}
}

// GetPlantillasActivas ...
func GetPlantillasActivas() (plantillaResult []map[string]interface{}, outputError map[string]interface{}) {
	var plantillasGet map[string]interface{}
	query := "Activo:true"
	if response, err := getJsonTest(beego.AppConfig.String("evaluacion_crud_url")+"v1/plantilla?query="+query, &plantillasGet); (err == nil) && (response == 200) {
		aux := reflect.ValueOf(plantillasGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(plantillasGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &plantillaResult); err == nil {
					return plantillaResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetPlantillasActivas4", "err": err.Error(), "status": "502"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/GetPlantillasActivas3", "err": "Cantidad de elementos vacia", "status": "502"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/GetPlantillasActivas2", "err": "Los valores no son validos", "status": "502"}
			return nil, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetPlantillasActivas1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	/*if error != nil {
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
	}*/

}

// DesactivarPlantillas ...
func DesactivarPlantillas(plantillasActivas []map[string]interface{}) (plantillasResult []map[string]interface{}, outputError map[string]interface{}) {
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
		if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/plantilla/"+fmt.Sprintf("%v", plantillasActivas[i]["Id"]), "PUT", &platillaActualizada, datoContruirdo); err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/DesactivarPlantillas", "err": err.Error(), "status": "502"}
			return nil, outputError
		} else {
			platillaActualizadaData := platillaActualizada["Data"].(map[string]interface{})
			arrayPlantillasIngresadas = append(arrayPlantillasIngresadas, platillaActualizadaData)
		}
	}
	return arrayPlantillasIngresadas, nil
}

// ActivarPlantilla ...
func ActivarPlantilla(plantillasParaActivar map[string]interface{}) (plantillasResult map[string]interface{}, outputError map[string]interface{}) {
	var platillaActualizada map[string]interface{}
	plantillaGet, errGetPlantilla := GetPlantilla(plantillasParaActivar)
	if plantillaGet != nil {
		datoContruirdo := make(map[string]interface{})
		datoContruirdo = map[string]interface{}{
			"Activo":        true,
			"Descripcion":   plantillaGet[0]["Descripcion"],
			"FechaCreacion": plantillaGet[0]["FechaCreacion"],
			"Id":            plantillaGet[0]["Id"],
			"Usuario":       plantillaGet[0]["Usuario"],
		}
		if err := sendJson(beego.AppConfig.String("evaluacion_crud_url")+"v1/plantilla/"+fmt.Sprintf("%v", plantillaGet[0]["Id"]), "PUT", &platillaActualizada, datoContruirdo); err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ActivarPlantilla", "err": err.Error(), "status": "502"}
			return nil, outputError
		} else {
			platillaActualizadaData := platillaActualizada["Data"].(map[string]interface{})
			return platillaActualizadaData, nil
		}
	} else {
		return nil, errGetPlantilla
	}
	return nil, nil
}

// GetPlantilla ...
func GetPlantilla(plantilla map[string]interface{}) (plantillaResult []map[string]interface{}, outputError map[string]interface{}) {
	var plantillaGet map[string]interface{}
	query := "Id:" + fmt.Sprintf("%v", plantilla["Id"])
	if response, err := getJsonTest(beego.AppConfig.String("evaluacion_crud_url")+"v1/plantilla?query="+query, &plantillaGet); (err == nil) && (response == 200) {
		aux := reflect.ValueOf(plantillaGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(plantillaGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &plantillaResult); err == nil {
					return plantillaResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetPlantilla4", "err": err.Error(), "status": "502"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/GetPlantilla3", "err": "Cantidad de elementos vacia", "status": "502"}
				return nil, outputError
			}
		} else {
			outputError = map[string]interface{}{"funcion": "/GetPlantilla2", "err": "Los valores no son validos", "status": "502"}
			return nil, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetPlantilla1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}

}

// ObtenerPlantillas ...
func ObtenerPlantillas() (plantillaResult map[string]interface{}, outputError map[string]interface{}) {
	var plantillaConstruida map[string]interface{}
	query := "?query=Activo:true"
	plantillaActiva, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("plantilla", query)
	if plantillaActiva != nil {
		plantillaConstruida = plantillaActiva[0]
		fmt.Println("tenemos plantilla")
		// return plantillaConstruida, nil
		clasificaciones, errClasificaciones := GetClasicacionesPlntilla(plantillaConstruida)
		if clasificaciones != nil {
			plantillaConstruida["Clasificaciones"] = clasificaciones
			secciones, errSecciones := GetSecciones(plantillaConstruida)
			if secciones != nil {
				plantillaConstruida["Secciones"] = secciones
				return plantillaConstruida, nil
			}
			return nil, errSecciones
		}
		return nil, errClasificaciones
	}
	fmt.Println("no tenemos plantilla")
	return nil, errTablaCrudEvaluacion
}

// ObternerPlantillaPorID ...
func ObtenerPlantillaPorID(IDPlantilla string) (plantillaResult map[string]interface{}, outputError map[string]interface{}) {
	var plantillaConstruida map[string]interface{}
	query := "?query=Id:" + IDPlantilla
	plantillaBusqueda, errTablaCrudEvaluacion := GetTablaCrudEvaluacion("plantilla", query)
	if plantillaBusqueda != nil {
		if len(plantillaBusqueda[0]) == 0 {
			texto_error := "La plantilla con id " + IDPlantilla + " no existe"
			fmt.Println(texto_error)
			outputError = map[string]interface{}{"funcion": "/ObtenerPlantillaPorID", "err": texto_error, "status": "204"}
			return nil, outputError
		}
		plantillaConstruida = plantillaBusqueda[0]
		fmt.Println("tenemos plantilla")
		clasificaciones, errClasificaciones := GetClasicacionesPlntilla(plantillaConstruida)
		if clasificaciones != nil {
			plantillaConstruida["Clasificaciones"] = clasificaciones
			secciones, errSecciones := GetSecciones(plantillaConstruida)
			if secciones != nil {
				plantillaConstruida["Secciones"] = secciones
				return plantillaConstruida, nil
			}
			return nil, errSecciones
		}
		return nil, errClasificaciones
	} else {
		return nil, errTablaCrudEvaluacion
	}
}
