package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
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
func GetTablaCrudEvaluacion(tabla string, query string) (objetoResult []map[string]interface{}) {
	var objetiGet []map[string]interface{}
	var url string
	if query != "" {
		url = beego.AppConfig.String("evaluacion_crud_url") +"v1/" + tabla + query
	} else {
		url = beego.AppConfig.String("evaluacion_crud_url") + "v1/" + tabla
	}
	error := request.GetJson(url, &objetiGet)
	if error != nil {
		fmt.Println("error en get tabla", tabla, "con la peticion: ", url)
		logs.Error(error)
		return nil
	} else {
		aux := reflect.ValueOf(objetiGet[0])
		if aux.IsValid() {
			if aux.Len() > 0 {
				return objetiGet
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
}
