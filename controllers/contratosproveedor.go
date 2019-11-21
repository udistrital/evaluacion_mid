package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// ContratosProveedorController operations for ContratosProveedor
type ContratosProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContratosProveedorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Create
// @Description create ContratosProveedor
// @Param	body		body 	{}	true		"body for ContratosProveedor content"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [post]
func (c *ContratosProveedorController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get ContratosProveedor by id
// @Param	id		path 	string	false		"The key for staticblock"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:id [get]
func (c *ContratosProveedorController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get ContratosProveedor
// @Param	ProvID	query	string	true		"ID del Proveedor"
// @Param	SupID	query	string	true		"ID del supervisor"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContratosProveedorController) GetAll() {
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	logs.Info("viva el get")
	ProveedorID := c.GetString("ProvID")
	logs.Info(ProveedorID)
	SupervisorID := c.GetString("SupID")
	logs.Info(SupervisorID)
	resultContratos, err1 := ListaContratos(ProveedorID, SupervisorID)
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

func ListaContratos(Idprov string, Idsuper string) (novedad []map[string]interface{}, outputError interface{}) {
	// fmt.Println(Idprov, Idsuper)
	// fmt.Println(beego.AppConfig.String("administrativa_amazon_api_url"), beego.AppConfig.String("administrativa_amazon_api_version"))
	resultProv, err1 := InfoProveedor(Idprov)
	fmt.Println("error  lista", err1)
	// fmt.Println("resultado lista", resultProv)
	if resultProv != nil {
		fmt.Println("entro a no nil")
		return resultProv, nil
	} else {
		fmt.Println("entro a si nil")
		// err1 =interface{"error": "no" }
		return nil, err1
	}
}

func InfoProveedor(Idprov string) (novedad []map[string]interface{}, outputError interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=NumDocumento:"+Idprov, &infoProveedor)
	fmt.Println(error)
	fmt.Println(len(infoProveedor))
	if len(infoProveedor) < 1 {
		fmt.Println("entro al error")
		return nil, error
	} else {
		fmt.Println("ok")
		return infoProveedor, nil
	}
}
