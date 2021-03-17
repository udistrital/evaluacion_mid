package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
	"strconv"
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
// @Param	IdentProv	query	string	true		"Identificacion del proveedor"
// @Param	NumContrato	query	string	true		"Numero del contrato"
// @Param	Vigencia	query	string	true		"Vigencia del contrato,, para evitar el filtro se debe de mandar un 0 (cero)"
// @Param	SupID	query	string	true		"Identificacion del supervisor"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *FiltromixtoController) GetAll() {

	defer helpers.ErrorControl(c.Controller, "FiltromixtoController")

	IdentificacionProveedor := c.GetString("IdentProv")
	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("Vigencia")
	SupervisorIdent := c.GetString("SupID")

	_, err1 := strconv.Atoi(IdentificacionProveedor)
	_, err2 := strconv.Atoi(NumContrato)
	_, err3 := strconv.Atoi(Vigencia)
	_, err4 := strconv.Atoi(SupervisorIdent)

	if (err1 != nil) || (err2 != nil) || (err3 != nil) || (err4 != nil) {
		panic(map[string]interface{}{"funcion": "GetAll", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	resultContratos, err5 := helpers.ListaContratoMixto(IdentificacionProveedor, NumContrato, Vigencia, SupervisorIdent)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err5)
	}
	c.ServeJSON()
}
