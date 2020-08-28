package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
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
		alertErr.Code = "404"
		alertas = append(alertas, err1)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()
}

// ListaContratosProveedor ...
func ListaContratosProveedor(IdentProv string, Idsuper string) (contratos []map[string]interface{}, outputError interface{}) {
	resultProv, err1 := InfoProveedor(IdentProv)
	if resultProv != nil {
		IDProveedor := models.GetElementoMaptoString(resultProv, "Id")
		resultContrato, err2 := ObtenerContratosProveedor(IDProveedor)
		if resultContrato != nil {
			InfoOrg := models.OrganizarInfoContratos(resultProv, resultContrato)
			if Idsuper == "0" {
				return InfoOrg, nil
			} else {
				resultDependencia, errDep := models.ObtenerDependencias(Idsuper)
				if errDep != nil {
					return nil, errDep
				} else {
					InfoFiltrada, err2 := models.FiltroDependencia(InfoOrg, resultDependencia)
					if InfoFiltrada != nil {
						return InfoFiltrada, nil
					} else {
						return nil, err2
					}
				}
			}
		} else {
			return nil, err2
		}
	} else {
		return nil, err1
	}

}

// InfoProveedor ...
func InfoProveedor(IdentProv string) (proveedor []map[string]interface{}, outputError interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=NumDocumento:"+IdentProv+"&limit=0", &infoProveedor)
	if len(infoProveedor) < 1 {
		fmt.Println(error)
		errorProv := models.CrearError("no se pudo traer la info del proveedor")
		return nil, errorProv
	} else {
		return infoProveedor, nil
	}
}

// ObtenerContratosProveedor ...
func ObtenerContratosProveedor(IDProv string) (contrato []map[string]interface{}, outputError interface{}) {
	var ContratosProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=Contratista:"+IDProv, &ContratosProveedor)
	if len(ContratosProveedor) < 1 {
		fmt.Println(error)
		errorContrato := models.CrearError("no se encontraron contratos")
		return nil, errorContrato
	} else {
		return ContratosProveedor, nil
	}
}
