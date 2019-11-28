package models

import (
	"github.com/astaxie/beego"
)

// ObtenerDependencias ...
func ObtenerDependencias(supervidorIdent string) (dependencias map[string]interface{}, outputError interface{}) {
	var DependenciasSuervisor map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/"+supervidorIdent, &DependenciasSuervisor)
	if error != nil {
		return nil, error
	} else {
		return DependenciasSuervisor, nil

	}
}
