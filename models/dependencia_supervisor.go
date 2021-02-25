package models

import (
	"github.com/astaxie/beego"
)

// ObtenerDependencias ...
func ObtenerDependencias(supervidorIdent string) (dependencias map[string]interface{}, outputError map[string]interface{}) {
	var DependenciasSuervisor map[string]interface{}
	_, err := getJsonWSO2Test(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/"+supervidorIdent, &DependenciasSuervisor)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/ObtenerDependencias", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return DependenciasSuervisor, nil

	}
}
