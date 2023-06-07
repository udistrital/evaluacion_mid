package helpers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratosContrato ...
func ListaContratosContrato(NumeroContrato, vigencia, supervisor string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	resultContrato, err1 := ObtenerContratosContrato(NumeroContrato, vigencia, supervisor)
	if resultContrato != nil {
		InfoOrg, err2 := models.OrganizarInfoContratosMultipleProv(resultContrato)
		if err2 != nil {
			return nil, err2
		} else {
			return InfoOrg, nil
		}
		// return resultContrato, nil
	} else {
		return nil, err1
	}
	// return nil, nil
}

// ObtenerContratosContrato ...
func ObtenerContratosContrato(NumContrato string, vigencia, supervisor string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratosProveedor []map[string]interface{}
	var urlCRUD = beego.AppConfig.String("administrativa_amazon_api_url") + beego.AppConfig.String("administrativa_amazon_api_version") + "contrato_general?query="
	var query []string

	if NumContrato != "0" {
		query = append(query, "ContratoSuscrito.NumeroContratoSuscrito:"+NumContrato)
	}

	if vigencia != "0" {
		query = append(query, "VigenciaContrato:"+vigencia)
	}

	if supervisor != "0" {
		query = append(query, "Supervisor__Documento:"+supervisor)
	}

	query_ := strings.Join(query, ",")
	response, err := getJsonTest(urlCRUD+query_, &ContratosProveedor)
	if err != nil || response != 200 {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosContrato1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}

	if len(ContratosProveedor) < 1 {
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosContrato3", "err": "No se encontraron contratos", "status": "204"}
		return nil, outputError
	} else {
		return ContratosProveedor, nil
	}
}
