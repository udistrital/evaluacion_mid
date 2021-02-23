package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
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

	defer helpers.ErrorControl(c.Controller, "DatosContratoController")

	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("VigenciaContrato")
	resultContratos, err1 := helpers.InfoContrato(NumContrato, Vigencia)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err1)
	}
	c.ServeJSON()

}
