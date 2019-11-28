package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// ContatoscontratoController ... Filtro para tener lista de contratos segun su vigencia y los proveedores de estos
type ContatoscontratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContatoscontratoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Create
// @Description create Contatoscontrato
// @Param	body		body 	models.Contatoscontrato	true		"body for Contatoscontrato content"
// @Success 201 {object} models.Contatoscontrato
// @Failure 403 body is empty
// @router / [post]
func (c *ContatoscontratoController) Post() {

}

// GetAll ...
// @Title GetAll
// @Description get Contatoscontrato
// @Param	NumContrato	query	string	true		"Numero del contrato"
// @Param	Vigencia	query	string	false		"Vigencia del contrato"
// @Param	SupID	query	string	false		"ID del supervisor"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContatoscontratoController) GetAll() {
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("Vigencia")
	SupervisorIdent := c.GetString("SupID")
	resultContratos, err1 := ListaContratosContrato(NumContrato, Vigencia, SupervisorIdent)
	if resultContratos != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = resultContratos
	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err1)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()
}

// ListaContratosContrato ...
func ListaContratosContrato(NumeroContrato string, vigencia string, supervidorIdent string) (contratos []map[string]interface{}, outputError interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia)
	if resultContrato != nil {
		InfoOrg := models.OrganizarInfoContratosMultipleProv(resultContrato)
		resultDependencia, errDep := models.ObtenerDependencias(supervidorIdent)
		if errDep != nil {
			return nil, errDep
		} else {
			InfoFiltrada, err2 := models.FiltroDependencia(InfoOrg, resultDependencia)

			if InfoFiltrada != nil {
				return InfoFiltrada, nil

			} else {
				return nil, err2
			}
		}

		// return resultContrato, nil
	} else {
		return nil, err1
	}
	// return nil, nil
}

// ObtenerContratosContrato ...
func ObtenerContratosContrato(NumContrato string, vigencia string) (contrato []map[string]interface{}, outputError interface{}) {
	var ContratosProveedor []map[string]interface{}
	var error interface{}
	if vigencia == "0" {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato, &ContratosProveedor)
	} else {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",VigenciaContrato:"+vigencia, &ContratosProveedor)
	}
	if len(ContratosProveedor) < 1 {
		error = models.CrearError("no se encontraron contratos")
		return nil, error
	} else {
		return ContratosProveedor, nil
	}
}
