package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratoMixto ...
func ListaContratoMixto(IdentificacionProveedor string, NumeroContrato string, vigencia string, supervisorIdent string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	ProveedorInfo, errorProv := InfoProveedor(IdentificacionProveedor)
	if ProveedorInfo != nil {
		IDProveedor := models.GetElementoMaptoString(ProveedorInfo, "Id")
		resultContrato, errContrato := ObtenerContratoProveedor(IDProveedor, NumeroContrato, vigencia)
		if resultContrato != nil {
			InfoOrg := models.OrganizarInfoContratos(ProveedorInfo, resultContrato)
			if supervisorIdent == "0" {
				return InfoOrg, nil
			} else {
				resultDependenciaSic, errDep := models.ObtenerDependenciasSic(supervisorIdent)
				if errDep != nil {
					return nil, errDep
				} else if models.GetElemento(resultDependenciaSic["DependenciasSic"], "Dependencia") == nil {
					resultDependenciaSup, errDep2 := models.ObtenerDependenciasSup(supervisorIdent)
					if errDep2 != nil {
						return nil, errDep2
					} else {
						InfoFiltrada, errFiltro := models.FiltroDependenciaSup(InfoOrg, resultDependenciaSup)
						if errFiltro != nil {
							return nil, errFiltro
						} else {
							return InfoFiltrada, nil
						}
					}
				} else {
					InfoFiltrada, errFiltro2 := models.FiltroDependenciaSic(InfoOrg, resultDependenciaSic)
					if InfoFiltrada != nil {
						return InfoFiltrada, nil
					} else {
						return nil, errFiltro2
					}
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
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor1", "err": err.Error(), "status": "502"}
			return nil, outputError
		}
	} else if NumContrato == "0" {
		if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor); (err == nil) && (response == 200) {
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor2", "err": err.Error(), "status": "502"}
		}
	} else {
		//error = getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor)
		if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",Contratista:"+ProveedorID+",VigenciaContrato:"+vigencia, &ContratoProveedor); (err == nil) && (response == 200) {
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor3", "err": err.Error(), "status": "502"}
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
