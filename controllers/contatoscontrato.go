package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	logs.Info(NumContrato)
	SupervisorIdent := c.GetString("SupID")
	logs.Info(Vigencia)
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
	fmt.Println("error  contrato", err1)
	if resultContrato != nil {
		fmt.Println("entro a no nil")
		// fmt.Println(resultContrato)
		InfoOrg := models.OrganizarInfoContratosMultipleProv(resultContrato)
		resultDependencia := models.ObtenerDependencias(supervidorIdent)
		InfoFiltrada, err2 := models.FiltroDependencia(InfoOrg, resultDependencia)
		// fmt.Println("INFO DE FILTRO", InfoFiltrada)
		// fmt.Println("ERROR DE FILTRO", err2)
		if InfoFiltrada != nil {
			return InfoFiltrada, nil

		} else {
			return nil, err2
		}
		// return resultContrato, nil
	} else {
		fmt.Println("entro a si nil contrato")
		return nil, err1
	}
	// return nil, nil
}

// ObtenerContratosContrato ...
func ObtenerContratosContrato(NumContrato string, vigencia string) (contrato []map[string]interface{}, outputError interface{}) {
	var ContratosProveedor []map[string]interface{}
	var error error
	if vigencia == "0" {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato, &ContratosProveedor)
	} else {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",VigenciaContrato:"+vigencia, &ContratosProveedor)
	}
	fmt.Println(len(ContratosProveedor))
	if len(ContratosProveedor) < 1 {
		fmt.Println(error)
		fmt.Println("entro al error")
		errorContrato := models.CrearError("no se encontraron contratos")
		return nil, errorContrato
	} else {
		fmt.Println("ok")
		return ContratosProveedor, nil
	}
}
