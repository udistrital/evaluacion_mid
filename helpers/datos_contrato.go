package helpers

import (
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/astaxie/beego"
	"fmt"
)

// InfoContrato ...
func InfoContrato(NumeroContrato string, vigencia string) (contrato []map[string]interface{}, outputError interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia)
	if resultContrato != nil {
		infoProveedor, errProv := models.InfoProveedorID(fmt.Sprintf("%v", resultContrato[0]["Contratista"]))
		if infoProveedor != nil {
			lugarEjecucion := resultContrato[0]["LugarEjecucion"].(map[string]interface{})
			infoDependencia, errDependencia := GetGependencia(fmt.Sprintf("%v", lugarEjecucion["Dependencia"]))
			if infoDependencia != nil {
				documentoSupervisor := fmt.Sprintf("%d", (resultContrato[0]["Supervisor"].(map[string]interface{})["Documento"]).(float64))
				dependencuaSupervisor := fmt.Sprintf("%v", resultContrato[0]["Supervisor"].(map[string]interface{})["DependenciaSupervisor"])
				infoSupervisor, errSup := GetSupervisorContrato(documentoSupervisor, dependencuaSupervisor)
				if infoSupervisor != nil {
					infoOrganizada := models.OrganizarInfoContratoArgo(infoProveedor, resultContrato, infoDependencia, infoSupervisor)
					return infoOrganizada, nil

				}
				return nil, errSup

			}
			return nil, errDependencia
			// return infoProveedor, nil
		}
		return nil, errProv
		// return resultContrato, nil
	}
	return nil, err1
	// return nil, nil
}

// GetGependencia ...
func GetGependencia(CodDependencia string) (Dependencia []map[string]interface{}, outputError interface{}) {
	var dependencia []map[string]interface{}
	error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"dependencia_SIC?query=ESFCODIGODEP:"+CodDependencia+",EstadoRegistro:true&sortby=Id&order=desc&limit=1", &dependencia)
	if len(dependencia) < 1 {
		fmt.Println(error)
		errorProv := models.CrearError("no se pudo traer la info de la dependencia")
		return nil, errorProv
	} else {
		return dependencia, nil
	}
}

// GetSupervisorContrato ...
func GetSupervisorContrato(numeroDocSupervisor string, dependenciaSupervisor string) (supervisorResult []map[string]interface{}, outputError interface{}) {
	var supervisor []map[string]interface{}
	error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"supervisor_contrato/?query=Documento:"+numeroDocSupervisor+"&DependenciaSupervisor:"+dependenciaSupervisor+"&sortby=FechaInicio&order=desc&limit=1", &supervisor)
	if len(supervisor) < 1 {
		fmt.Println(error)
		errorProv := models.CrearError("no se pudo traer la info del supervisor con documento:" + numeroDocSupervisor + " de la dependencia: " + dependenciaSupervisor)
		return nil, errorProv
	} else {
		return supervisor, nil
	}
}
