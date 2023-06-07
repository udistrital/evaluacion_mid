package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
	_ "github.com/udistrital/utils_oas/request"
)

// ContatoscontratoController ... Filtro para tener lista de contratos segun su vigencia y los proveedores de estos
type ContratoscontratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContratoscontratoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get Contatoscontrato
// @Param	NumContrato	query	string	true	"Numero del contrato"
// @Param	Vigencia	query	string	true	"Vigencia del contrato,, para evitar el filtro se debe de mandar un 0 (cero)"
// @Param	Supervisor	query	string	true	"Supervisor del contrato. Para evitar el filtro se debe enviar un 0"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContratoscontratoController) GetAll() {

	defer helpers.ErrorControl(c.Controller, "ContratoscontratoController")

	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("Vigencia")
	Supervisor := c.GetString("Supervisor", "0")

	_, err1 := strconv.Atoi(NumContrato)
	_, err2 := strconv.Atoi(Vigencia)

	if err1 != nil || err2 != nil {
		panic(map[string]interface{}{"funcion": "GetAll", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	resultContratos, err := helpers.ListaContratosContrato(NumContrato, Vigencia, Supervisor)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err)
	}
	c.ServeJSON()
}
