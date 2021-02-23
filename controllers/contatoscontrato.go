package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
	_ "github.com/udistrital/utils_oas/request"
)

// ContatoscontratoController ... Filtro para tener lista de contratos segun su vigencia y los proveedores de estos
type ContatoscontratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContatoscontratoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get Contatoscontrato
// @Param	NumContrato	query	string	true		"Numero del contrato"
// @Param	Vigencia	query	string	true		"Vigencia del contrato,, para evitar el filtro se debe de mandar un 0 (cero)"
// @Param	SupID	query	string	true		"Identificacion del supervisor"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContatoscontratoController) GetAll() {

	defer helpers.ErrorControl(c.Controller, "ContratoscontratoController")

	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("Vigencia")
	SupervisorIdent := c.GetString("SupID")
	resultContratos, err1 := helpers.ListaContratosContrato(NumContrato, Vigencia, SupervisorIdent)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err1)
	}
	c.ServeJSON()
}
