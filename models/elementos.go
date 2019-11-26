package models

import (
	"fmt"
	"reflect"
)

func GetElemento(objeto interface{}, item string) interface{} {
	var subobjeto interface{}
	subobjeto = objeto.(map[string]interface{})[item]
	return subobjeto
}

func GetElementoMaptoString(objeto interface{}, item string) string {
	value := reflect.ValueOf(objeto)
	var resuesta string
	if value.Len() > 0 {
		aux := value.Index(0).Interface().(map[string]interface{})
		resuesta = fmt.Sprintf("%v", aux[item])
	}
	if value.Len() == 0 {
		resuesta = fmt.Sprintf("Objeto de longitud cero")
	}

	return resuesta
}
