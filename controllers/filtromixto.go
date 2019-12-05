package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
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
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	IdentificacionProveedor := c.GetString("IdentProv")
	NumContrato := c.GetString("NumContrato")
	Vigencia := c.GetString("Vigencia")
	SupervisorIdent := c.GetString("SupID")
	resultContratos, err1 := ListaContratoMixto(IdentificacionProveedor, NumContrato, Vigencia, SupervisorIdent)
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

// ListaContratoMixto ...
func ListaContratoMixto(IdentificacionProveedor string, NumeroContrato string, vigencia string, supervidorIdent string) (contratos []map[string]interface{}, outputError interface{}) {
	ProveedorInfo, errorProv := InfoProveedor(IdentificacionProveedor)
	if ProveedorInfo != nil {
		IDProveedor := models.GetElementoMaptoString(ProveedorInfo, "Id")
		resultContrato, errContrato := ObtenerContratoProveedor(IDProveedor, NumeroContrato, vigencia)
		if resultContrato != nil {
			InfoOrg := models.OrganizarInfoContratos(ProveedorInfo, resultContrato)
			resultDependencia, errDep := models.ObtenerDependencias(supervidorIdent)
			if errDep != nil {
				return nil, errDep
			} else {
				InfoFiltrada, errFiltro := models.FiltroDependencia(InfoOrg, resultDependencia)
				if InfoFiltrada != nil {
					return InfoFiltrada, nil
				} else {
					return nil, errFiltro
				}
			}

		} else {
			return nil, errContrato

		}

	} else {
		return nil, errorProv
	}
}

// ObtenerContratoProveedor ...
func ObtenerContratoProveedor(ProveedorID string, NumContrato string, vigencia string) (contrato []map[string]interface{}, outputError interface{}) {
	var ContratoProveedor []map[string]interface{}
	var error interface{}
	if vigencia == "0" {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID, &ContratoProveedor)
	} else {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor)
	}
	if len(ContratoProveedor) < 1 {
		error = models.CrearError("no se encontraron contratos")
		return nil, error
	} else {
		return ContratoProveedor, nil
	}
}
