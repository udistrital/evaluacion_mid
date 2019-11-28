package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
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

// GetAll ...
// @Title GetAll
// @Description get ContratosProveedor
// @Param	ProvID	query	string	true		"ID del Proveedor"
// @Param	SupID	query	string	false		"ID del supervisor"
// @Success 200 {}
// @Failure 404 not found resource
// @router / [get]
func (c *ContratosProveedorController) GetAll() {
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	ProveedorIdent := c.GetString("ProvID")
	SupervisorIdent := c.GetString("SupID")
	resultContratos, err1 := ListaContratosProveedor(ProveedorIdent, SupervisorIdent)
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

// ListaContratosProveedor ...
func ListaContratosProveedor(IdentProv string, Idsuper string) (contratos []map[string]interface{}, outputError interface{}) {
	// fmt.Println(IdentProv, Idsuper)
	// fmt.Println(beego.AppConfig.String("administrativa_amazon_api_url"), beego.AppConfig.String("administrativa_amazon_api_version"))
	resultProv, err1 := InfoProveedor(IdentProv)
	fmt.Println("error  lista", err1)
	// fmt.Println(resultProv)
	// fmt.Println(models.GetElementoMaptoString(resultProv, "Id"))
	if resultProv != nil {
		fmt.Println("entro a no nil")
		IDProveedor := models.GetElementoMaptoString(resultProv, "Id")
		fmt.Println(IDProveedor)
		resultContrato, err2 := ObtenerContratosProveedor(IDProveedor)
		fmt.Println("error  contrato", err2)
		// fmt.Println(resultProv)
		// fmt.Println(models.GetElementoMaptoString(resultProv, "Id"))
		if resultContrato != nil {
			fmt.Println("entro a no nil")
			// fmt.Println(resultContrato)
			pruebaOrg := models.OrganizarInfoContratos(resultProv, resultContrato)
			return pruebaOrg, nil
		} else {
			fmt.Println("entro a si nil contrato")
			return nil, err2
		}
		// return resultProv, nil
	} else {
		fmt.Println("entro a si nil")
		return nil, err1
	}

}

// InfoProveedor ...
func InfoProveedor(IdentProv string) (proveedor []map[string]interface{}, outputError interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=NumDocumento:"+IdentProv+"&limit=0", &infoProveedor)
	fmt.Println(len(infoProveedor))
	if len(infoProveedor) < 1 {
		fmt.Println(error)
		fmt.Println("entro al error")
		errorProv := models.CrearError("no se pudo traer la info del proveedor")
		return nil, errorProv
	} else {
		fmt.Println("ok")
		return infoProveedor, nil
	}
}

// ObtenerContratosProveedor ...
func ObtenerContratosProveedor(IDProv string) (contrato []map[string]interface{}, outputError interface{}) {
	var ContratosProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=Contratista:"+IDProv, &ContratosProveedor)
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
