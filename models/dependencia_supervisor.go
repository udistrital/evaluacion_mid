package models

import (
	"github.com/astaxie/beego"
)

// ObtenerDependenciasSic ...
func ObtenerDependenciasSic(supervidorIdent string) (dependencias map[string]interface{}, outputError map[string]interface{}) {
	var DependenciasSuervisor map[string]interface{}
	_, err := getJsonWSO2Test(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"dependencias_sic/"+supervidorIdent, &DependenciasSuervisor)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/ObtenerDependenciasSic", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return DependenciasSuervisor, nil

	}
}

// ObtenerDependenciasSup ...
func ObtenerDependenciasSup(supervidorIdent string) (dependencias map[string]interface{}, outputError map[string]interface{}) {
	var DependenciasSuervisor map[string]interface{}
	_, err := getJsonWSO2Test(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"dependencias_supervisor/"+supervidorIdent, &DependenciasSuervisor)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/ObtenerDependenciasSup", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return DependenciasSuervisor, nil

	}
}
