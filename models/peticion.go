package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/udistrital/utils_oas/request"
)

// GetJSONJBPM ...
func GetJSONJBPM(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

// GetTablaCrudEvaluacion ...
func GetTablaCrudEvaluacion(tabla string, query string) (objetoResult []map[string]interface{}, outputError map[string]interface{}) {
	var objetiGet map[string]interface{}
	var url string
	if query != "" {
		url = beego.AppConfig.String("evaluacion_crud_url") + "v1/" + tabla + query
	} else {
		url = beego.AppConfig.String("evaluacion_crud_url") + "v1/" + tabla
	}
	if response, err := getJsonTest(url, &objetiGet); (response == 200) && (err == nil) {
		aux := reflect.ValueOf(objetiGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(objetiGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &objetoResult); err == nil {
					return objetoResult, nil
				} else {
					outputError = map[string]interface{}{"funcion": "/GetTablaCrudEvaluacion", "err": err.Error(), "status": "502"}
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
		fmt.Println("error en get tabla", tabla, "con la peticion: ", url)
		return nil
	} else {
		aux := reflect.ValueOf(objetiGet["Data"])
		if aux.IsValid() {
			if aux.Len() > 0 {
				temp, _ := json.Marshal(objetiGet["Data"].([]interface{}))
				if err := json.Unmarshal(temp, &objetoResult); err == nil {
					return objetoResult
				} else {
					return nil
				}
			} else {
				return nil
			}
		} else {
			return nil
		}
	}*/
}
