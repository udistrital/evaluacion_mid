package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// DatosContratoController permite traer los datos necesarios para el contrato, dichos datos son consultados de diferentes apis
type DatosContratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *DatosContratoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description obtiene los didatos de contrato general,informacion del proveedor y dependencias del supervidor
// @Param	NumContrato	query	string	true		"Numero del contrato"
// @Param	VigenciaContrato	query	string	true		"año de vigencia del contrato"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *DatosContratoController) GetAll() {
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("VigenciaContrato")
	resultContratos, err1 := InfoContrato(NumContrato, Vigencia)
	if resultContratos != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = resultContratos
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, err1)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}

// InfoContrato ...
func InfoContrato(NumeroContrato string, vigencia string) (contrato []map[string]interface{}, outputError interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia)
	if resultContrato != nil {
		infoProveedor, errProv := models.InfoProveedorID(fmt.Sprintf("%v", resultContrato[0]["Contratista"]))
		if infoProveedor != nil {
			lugarEjecucion := resultContrato[0]["LugarEjecucion"].(map[string]interface{})
			infoDependencia, errDependencia := GetGependencia(fmt.Sprintf("%v", lugarEjecucion["Dependencia"]))
			if infoDependencia != nil {
				documentoSupervisor := fmt.Sprintf("%v", resultContrato[0]["Supervisor"].(map[string]interface{})["Documento"])
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
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"dependencia_SIC?query=ESFCODIGODEP:"+CodDependencia+",EstadoRegistro:true&sortby=Id&order=desc&limit=1", &dependencia)
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
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"supervisor_contrato?query=Documento:"+numeroDocSupervisor+"DependenciaSupervisor:"+dependenciaSupervisor+"&sortby=FechaInicio&order=desc&limit=1", &supervisor)
	if len(supervisor) < 1 {
		fmt.Println(error)
		errorProv := models.CrearError("no se pudo traer la info del supervisor con documento:" + numeroDocSupervisor + " de la dependencia: " + dependenciaSupervisor)
		return nil, errorProv
	} else {
		return supervisor, nil
	}
}
