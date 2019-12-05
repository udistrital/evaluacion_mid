package models

import (
	"fmt"
	"reflect"
)

// GetElemento ...
func GetElemento(objeto interface{}, item string) interface{} {
	var subobjeto interface{}
	subobjeto = objeto.(map[string]interface{})[item]
	return subobjeto
}

// GetElementoMaptoString ...
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

// GetElementoMaptoStringToArray ...
func GetElementoMaptoStringToArray(objeto interface{}, item string) (ElementosArray []string, outputError interface{}) {
	value := reflect.ValueOf(objeto)
	ArrayRespuesta := make([]string, 0)
	var resuesta string
	if value.IsValid() {
		if value.Len() > 0 {
			for i := 0; i < value.Len(); i++ {
				aux := value.Index(i).Interface().(map[string]interface{})
				resuesta = fmt.Sprintf("%v", aux[item])
				ArrayRespuesta = append(ArrayRespuesta, resuesta)
			}
		}
		if value.Len() == 0 {
			errorElementos := CrearError("Longitud cero de elmentos")
			return nil, errorElementos
		}
		return ArrayRespuesta, nil

	} else {
		errorElementos := CrearError("No se obtivueron dependencias para el documento de identificacion")
		return nil, errorElementos
	}

}

// GetElementoMaptoStringToMapArray ...
func GetElementoMaptoStringToMapArray(objeto interface{}) (ElementosMapArray []map[string]interface{}, outputError interface{}) {
	value := reflect.ValueOf(objeto)
	// ArrayRespuesta := make([]string, 0)
	// var resuesta string
	// var elementoIndividual map[string]interface{}
	ArrayElementos := make([]map[string]interface{}, 0)
	if value.IsValid() {
		if value.Len() > 0 {
			for i := 0; i < value.Len(); i++ {
				aux := value.Index(i).Interface().(map[string]interface{})
				ArrayElementos = append(ArrayElementos, aux)
				// resuesta = fmt.Sprintf("%v", aux[item])
				// ArrayRespuesta = append(ArrayRespuesta, resuesta)
			}
		}
		if value.Len() == 0 {
			errorElementos := CrearError("Longitud cero de elmentos")
			return nil, errorElementos
		}
		return ArrayElementos, nil

	} else {
		errorElementos := CrearError("json con datos incompletos")
		return nil, errorElementos
	}

}
