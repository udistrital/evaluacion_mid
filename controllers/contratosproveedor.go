package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
)

// ContratosProveedorController ... Filtro para tener lista de contratos de un proveedor
type ContratosProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContratosProveedorController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get ContratosProveedor
// @Param	ProvID			query	string	true	"ID del Proveedor"
// @Param	Supervisor		query	string	false	"Supervisor del contrato. Para evitar el filtro se debe enviar un 0"
// @Param	TipoContrato	query	string	false	"Tipo de contrato. Soporta prefijo in y notin para indicar múltiples valores separados por |"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContratosProveedorController) GetAll() {

	defer helpers.ErrorControl(c.Controller, "ContratosProveedorController")

	ProveedorIdent := c.GetString("ProvID")
	Supervisor := c.GetString("Supervisor", "0")
	TipoContrato := c.GetString("TipoContrato")

	_, err1 := strconv.Atoi(ProveedorIdent)

	if err1 != nil {
		panic(map[string]interface{}{"funcion": "GetAll", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	resultContratos, err3 := helpers.ListaContratosProveedor(ProveedorIdent, Supervisor, TipoContrato)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err3)
	}
	c.ServeJSON()
}
