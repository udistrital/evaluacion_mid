package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/formatdata"
)

func ObtenerDependencias(supervidorIdent string) (dependencias map[string]interface{}, outputError interface{}) {
	var DependenciasSuervisor map[string]interface{}
	error := GetJsonJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/"+supervidorIdent, &DependenciasSuervisor)
	fmt.Println("error", error)
	fmt.Println("super")
	fmt.Println(formatdata.JsonPrint(DependenciasSuervisor))
	logs.Info("super format")
	fmt.Println(formatdata.JsonPrint(DependenciasSuervisor))
	formatdata.JsonPrint(DependenciasSuervisor)
	return DependenciasSuervisor, nil
}
