package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratoMixto ...
func ListaContratoMixto(IdentificacionProveedor, NumeroContrato, vigencia, supervisor, tipo string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	ProveedorInfo, outputError := InfoProveedor(IdentificacionProveedor)
	if ProveedorInfo == nil || outputError != nil {
		return
	}

	IDProveedor := models.GetElementoMaptoString(ProveedorInfo, "Id")
	resultContrato, outputError := ObtenerContratoProveedor(IDProveedor, NumeroContrato, vigencia, supervisor, tipo)
	if resultContrato == nil || outputError != nil {
		return
	}

	contratos = models.OrganizarInfoContratos(ProveedorInfo, resultContrato)
	return

}

// ObtenerContratoProveedor ...
func ObtenerContratoProveedor(ProveedorID, NumContrato, vigencia, supervisor, tipo string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratoProveedor []map[string]interface{}
	var urlCRUD = beego.AppConfig.String("administrativa_amazon_api_url") + beego.AppConfig.String("administrativa_amazon_api_version") + "contrato_general?query="
	query := CrearQueryContratoGeneral(ProveedorID, NumContrato, vigencia, supervisor, tipo)

	response, err := getJsonTest(urlCRUD+query, &ContratoProveedor)
	if err != nil || response != 200 {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}

	if len(ContratoProveedor) < 1 {
		outputError = map[string]interface{}{"funcion": "/ObtenerContratoProveedor2", "err": "No se encontraron contratos", "status": "204"}
		return nil, outputError
	}

	return ContratoProveedor, nil
}
