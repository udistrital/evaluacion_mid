package helpers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// ListaContratosContrato ...
func ListaContratosContrato(NumeroContrato string, vigencia string, supervidorIdent string) (contratos []map[string]interface{}, outputError interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia)
	if resultContrato != nil {
		InfoOrg := models.OrganizarInfoContratosMultipleProv(resultContrato)
		resultDependencia, errDep := models.ObtenerDependencias(supervidorIdent)
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

		// return resultContrato, nil
	} else {
		return nil, err1
	}
	// return nil, nil
}

// ObtenerContratosContrato ...
func ObtenerContratosContrato(NumContrato string, vigencia string) (contrato []map[string]interface{}, outputError interface{}) {
	var ContratosProveedor []map[string]interface{}
	var error interface{}
	if vigencia == "0" {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato, &ContratosProveedor)
	} else {
		error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",VigenciaContrato:"+vigencia, &ContratosProveedor)
	}
	if len(ContratosProveedor) < 1 {
		error = models.CrearError("no se encontraron contratos")
		return nil, error
	} else {
		return ContratosProveedor, nil
	}
}
