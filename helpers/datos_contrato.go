package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// InfoContrato ...
func InfoContrato(NumeroContrato string, vigencia string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia, "0", "")
	if resultContrato != nil {
		estadoContrato, err2 := ObtenerEstadoContrato(NumeroContrato, vigencia)
		if estadoContrato != nil {
			resultActividades, err3 := ObtenerActividadContrato(NumeroContrato, vigencia)
			if resultActividades != nil {
				infoProveedor, errProv := models.InfoProveedorID(fmt.Sprintf("%v", resultContrato[0]["Contratista"]))
				if infoProveedor != nil {
					lugarEjecucion := resultContrato[0]["LugarEjecucion"].(map[string]interface{})
					infoDependencia, errDependencia := GetGependencia(fmt.Sprintf("%v", lugarEjecucion["Dependencia"]))
					if infoDependencia != nil {
						documentoSupervisor := fmt.Sprintf("%d", (resultContrato[0]["Supervisor"].(map[string]interface{})["Documento"]).(float64))
						dependencuaSupervisor := fmt.Sprintf("%v", resultContrato[0]["Supervisor"].(map[string]interface{})["DependenciaSupervisor"])
						infoSupervisor, errSup := GetSupervisorContrato(documentoSupervisor, dependencuaSupervisor)
						if infoSupervisor != nil {
							infoOrganizada := models.OrganizarInfoContratoArgo(infoProveedor, resultContrato, estadoContrato, resultActividades, infoDependencia, infoSupervisor)
							return infoOrganizada, nil

						}
						return nil, errSup
					} else if lugarEjecucion["Dependencia"] == "" {
						documentoSupervisor := fmt.Sprintf("%d", (resultContrato[0]["Supervisor"].(map[string]interface{})["Documento"]).(float64))
						dependencuaSupervisor := fmt.Sprintf("%v", resultContrato[0]["Supervisor"].(map[string]interface{})["DependenciaSupervisor"])
						infoSupervisor, errSup2 := GetSupervisorContrato(documentoSupervisor, dependencuaSupervisor)
						if infoSupervisor != nil {
							infoOrganizadaSinDep := models.OrganizarInfoContratoSinDep(infoProveedor, resultContrato, estadoContrato, resultActividades, infoSupervisor)
							return infoOrganizadaSinDep, nil
						}
						return nil, errSup2
					}
					return nil, errDependencia
					// return infoProveedor, nil
				}
				return nil, errProv
				// return resultContrato, nil
			}
			return nil, err3
		}
		return nil, err2
	}
	return nil, err1
	// return nil, nil
}

// GetGependencia ...
func GetGependencia(CodDependencia string) (Dependencia []map[string]interface{}, outputError map[string]interface{}) {
	var dependencia []map[string]interface{}
	//error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"dependencia_SIC?query=ESFCODIGODEP:"+CodDependencia+",EstadoRegistro:true&sortby=Id&order=desc&limit=1", &dependencia)
	if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"dependencia_SIC?query=ESFCODIGODEP:"+CodDependencia+",EstadoRegistro:true&sortby=Id&order=desc&limit=1", &dependencia); (err == nil) && (response == 200) {
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetGependencia1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	if len(dependencia) < 1 {
		//fmt.Println(error)
		//errorProv := models.CrearError("no se pudo traer la info de la dependencia")
		//return nil, errorProv
		outputError = map[string]interface{}{"funcion": "/GetGependencia2", "err": "No se pudo traer la info de la dependencia", "status": "204"}
		return nil, outputError
	} else {
		return dependencia, nil
	}
}

// GetSupervisorContrato ...
func GetSupervisorContrato(numeroDocSupervisor string, dependenciaSupervisor string) (supervisorResult []map[string]interface{}, outputError map[string]interface{}) {
	var supervisor []map[string]interface{}
	//error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"supervisor_contrato/?query=Documento:"+numeroDocSupervisor+"&DependenciaSupervisor:"+dependenciaSupervisor+"&sortby=FechaInicio&order=desc&limit=1", &supervisor)
	if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"supervisor_contrato/?query=Documento:"+numeroDocSupervisor+"&DependenciaSupervisor:"+dependenciaSupervisor+"&sortby=FechaInicio&order=desc&limit=1", &supervisor); (err == nil) && (response == 200) {
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetSupervisorContrato1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	if len(supervisor) < 1 {
		//fmt.Println(error)
		//errorProv := models.CrearError("no se pudo traer la info del supervisor con documento:" + numeroDocSupervisor + " de la dependencia: " + dependenciaSupervisor)
		//return nil, errorProv
		texto_error := "No se pudo traer la info del supervisor con documento:" + numeroDocSupervisor + " de la dependencia: " + dependenciaSupervisor
		outputError = map[string]interface{}{"funcion": "/GetSupervisorContrato2", "err": texto_error, "status": "204"}
		return nil, outputError
	} else {
		return supervisor, nil
	}
}

// ObtenerActividadContrato
func ObtenerActividadContrato(NumContrato string, vigencia string) (contrato map[string]interface{}, outputError map[string]interface{}) {
	var ActividadesContrato map[string]interface{}
	_, err := getJsonWSO2Test(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"informacion_contrato/"+NumContrato+"/"+vigencia, &ActividadesContrato)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/ObtenerActividadContrato", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return ActividadesContrato, nil
	}
}

// ObtenerEstadoContrato
func ObtenerEstadoContrato(NumContrato string, vigencia string) (res map[string]interface{}, outputError map[string]interface{}) {
	var EstadoContrato map[string]interface{}
	_, err := getJsonWSO2Test(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contrato_estado/"+NumContrato+"/"+vigencia, &EstadoContrato)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/ObtenerEstadoContrato", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return EstadoContrato, nil
	}
}
