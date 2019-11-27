package models

import (
	"fmt"

	"github.com/astaxie/beego"
)

// ObtenerDependencias ...
func ObtenerDependencias(supervidorIdent string) (dependencias map[string]interface{}) {
	var DependenciasSuervisor map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/"+supervidorIdent, &DependenciasSuervisor)
	fmt.Println("error", error)
	return DependenciasSuervisor
}
