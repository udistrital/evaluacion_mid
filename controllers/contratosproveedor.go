package controllers

import (
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
// @Param	ProvID	query	string	true		"ID del Proveedor"
// @Param	SupID	query	string	true		"Identificacion del supervisor, para evitar el filtro se debe de mandar un 0 (cero)"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContratosProveedorController) GetAll() {

	defer helpers.ErrorControl(c.Controller, "ContratosProveedorController")

	ProveedorIdent := c.GetString("ProvID")
	SupervisorIdent := c.GetString("SupID")
	resultContratos, err1 := helpers.ListaContratosProveedor(ProveedorIdent, SupervisorIdent)
	if resultContratos != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": resultContratos}
	} else {
		panic(err1)
	}
	c.ServeJSON()
}


