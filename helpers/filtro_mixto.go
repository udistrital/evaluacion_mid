package helpers

import (
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// ListaContratoMixto ...
func ListaContratoMixto(IdentificacionProveedor string, NumeroContrato string, vigencia string, supervidorIdent string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
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
func ObtenerContratoProveedor(ProveedorID string, NumContrato string, vigencia string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratoProveedor []map[string]interface{}
	if vigencia == "0" {
		//error = getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID, &ContratoProveedor)
		if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID, &ContratoProveedor); (err == nil) && (response == 200) {
		}else{
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor1", "err": err.Error(), "status": "502"}
			return nil, outputError
		}
	} else {
		//error = getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor)
		if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor); (err == nil) && (response == 200) {
		}else{
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor2", "err": err.Error(), "status": "502"}
			return nil, outputError
		}
	}
	if len(ContratoProveedor) < 1 {
		//error = models.CrearError("no se encontraron contratos")
		//return nil, error
		outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor2", "err": "No se encontraron contratos", "status": "204"}
		return nil, outputError
	} else {
		return ContratoProveedor, nil
	}
}
