package helpers

import (
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/astaxie/beego"
)

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
		error = getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID, &ContratoProveedor)
	} else {
		error = getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor)
	}
	if len(ContratoProveedor) < 1 {
		error = models.CrearError("no se encontraron contratos")
		return nil, error
	} else {
		return ContratoProveedor, nil
	}
}
