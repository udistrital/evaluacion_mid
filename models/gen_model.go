package models

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

// CrearError ...
func CrearError(contenido string) interface{} {
	var ErrorCreado interface{}

	ContendoBody, errBody := json.Marshal(map[string]string{
		"Error": contenido,
	})
	if errBody != nil {
		logs.Error("fallo el objeto a enviar: ", errBody)
	}
	json.Unmarshal([]byte(ContendoBody), &ErrorCreado)
	return ErrorCreado
}
