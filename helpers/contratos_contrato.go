package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratosContrato ...
func ListaContratosContrato(NumeroContrato string, vigencia string, supervidorIdent string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia)
	if resultContrato != nil {
		InfoOrg, err2 := models.OrganizarInfoContratosMultipleProv(resultContrato)
		if err2 != nil {
			return nil, err2
		}
		resultDependencia, errDep := models.ObtenerDependencias(supervidorIdent)
		if errDep != nil {
			return nil, errDep
		} else {
			InfoFiltrada, err3 := models.FiltroDependencia(InfoOrg, resultDependencia)

			if InfoFiltrada != nil {
				return InfoFiltrada, nil

			} else {
				return nil, err3
			}
		}

		// return resultContrato, nil
	} else {
		return nil, err1
	}
	// return nil, nil
}

// ObtenerContratosContrato ...
func ObtenerContratosContrato(NumContrato string, vigencia string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratosProveedor []map[string]interface{}
	//var err interface{}
	if vigencia == "0" {
		if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato, &ContratosProveedor); (err == nil) && (response == 200) {
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratosContrato1", "err": err.Error(), "status": "502"}
			return nil, outputError
		}
		//error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato, &ContratosProveedor)
	} else {
		if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",VigenciaContrato:"+vigencia, &ContratosProveedor); (err == nil) && (response == 200) {
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ObtenerContratosContrato2", "err": err.Error(), "status": "502"}
			return nil, outputError
		}
		//error = request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+"contrato_general?query=ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato+",VigenciaContrato:"+vigencia, &ContratosProveedor)
	}
	if len(ContratosProveedor) < 1 {
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosContrato3", "err": "No se encontraron contratos", "status": "204"}
		return nil, outputError
		//err = models.CrearError("no se encontraron contratos")
	} else {
		return ContratosProveedor, nil
	}
}

// ObtenerActividadContrato
func ObtenerActividadContrato(NumContrato string, vigencia string) (contrato map[string]interface{}, outputError map[string]interface{}) {
	var ActividadesContrato map[string]interface{}
	_, err := getJsonWSO2Test(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"informacion_contrato/"+NumContrato+"/"+vigencia, &ActividadesContrato)
	if err != nil {
		outputError = map[string]interface{}{"funcion": "/ObtenerActividadContrato", "err": err.Error(), "status": "502"}
		return nil, outputError
	} else {
		return ActividadesContrato, nil
	}
}
