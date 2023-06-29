package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
)

// FiltromixtoController ...  Filtro para tener lista de contratos segun el numero de contrato su vigencia y la identificacion del proveedor
type FiltromixtoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FiltromixtoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get Filtromixto
// @Param	IdentProv		query	string	true	"Identificacion del proveedor"
// @Param	NumContrato		query	string	true	"Numero del contrato"
// @Param	Vigencia		query	string	true	"Vigencia del contrato,, para evitar el filtro se debe de mandar un 0 (cero)"
// @Param	Supervisor		query	string	false	"Supervisor del contrato. Para evitar el filtro se debe enviar un 0"
// @Param	TipoContrato	query	string	false	"Tipo de contrato. Soporta prefijo in y notin para indicar múltiples valores separados por |"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *FiltromixtoController) GetAll() {

	defer helpers.ErrorControl(c.Controller, "FiltromixtoController")

	IdentificacionProveedor := c.GetString("IdentProv")
	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("Vigencia")
	Supervisor := c.GetString("Supervisor", "0")
	TipoContrato := c.GetString("TipoContrato")

	_, err1 := strconv.Atoi(IdentificacionProveedor)
	_, err2 := strconv.Atoi(NumContrato)
	_, err3 := strconv.Atoi(Vigencia)

	if (err1 != nil) || (err2 != nil) || (err3 != nil) {
		panic(map[string]interface{}{"funcion": "GetAll", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	resultContratos, err5 := helpers.ListaContratoMixto(IdentificacionProveedor, NumContrato, Vigencia, Supervisor, TipoContrato)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err5)
	}
	c.ServeJSON()
}
