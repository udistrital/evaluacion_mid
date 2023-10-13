package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratosContrato ...
func ListaContratosContrato(NumeroContrato, vigencia, supervisor, tipo string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	resultContrato, outputError := ObtenerContratosContrato(NumeroContrato, vigencia, supervisor, tipo)
	if outputError != nil || resultContrato == nil {
		return
	}

	InfoOrg, outputError := models.OrganizarInfoContratosMultipleProv(resultContrato)
	if outputError != nil {
		return
	}

	cesiones, outputError := CesionesContratos(InfoOrg)
	if outputError != nil {
		return
	}

	contratos = append(InfoOrg, cesiones...)
	return contratos, nil
}

// ObtenerContratosContrato ...
func ObtenerContratosContrato(NumContrato, vigencia, supervisor, tipo string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratosProveedor []map[string]interface{}
	var urlCRUD = beego.AppConfig.String("administrativa_amazon_api_url") + beego.AppConfig.String("administrativa_amazon_api_version") + "contrato_general?query="
	query := CrearQueryContratoGeneral("0", NumContrato, vigencia, supervisor, tipo)

	response, err := getJsonTest(urlCRUD+query, &ContratosProveedor)
	if err != nil || response != 200 {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosContrato1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}

	return ContratosProveedor, nil
}
