package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratosProveedor ...
func ListaContratosProveedor(IdentProv, supervisor, tipo string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	resultProv, outputError := InfoProveedor(IdentProv)
	if resultProv == nil || outputError != nil {
		return
	}

	IDProveedor := models.GetElementoMaptoString(resultProv, "Id")
	resultContrato, outputError := ObtenerContratosProveedor(IDProveedor, supervisor, tipo)
	if resultContrato == nil || outputError != nil {
		return
	}

	contratos = models.OrganizarInfoContratos(resultProv, resultContrato)
	return

}

// InfoProveedor ...
func InfoProveedor(IdentProv string) (proveedor []map[string]interface{}, outputError map[string]interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	//error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=NumDocumento:"+IdentProv+"&limit=0", &infoProveedor)
	if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=NumDocumento:"+IdentProv+"&limit=0", &infoProveedor); (err == nil) && (response == 200) {
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InfoProveedor1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	if len(infoProveedor) < 1 {
		//fmt.Println(error)
		//errorProv := models.CrearError("no se pudo traer la info del proveedor")
		//return nil, errorProv
		outputError = map[string]interface{}{"funcion": "/InfoProveedor2", "err": "No se pudo traer la info del proveedor", "status": "204"}
		return nil, outputError
	} else {
		return infoProveedor, nil
	}
}

// ObtenerContratosProveedor ...
func ObtenerContratosProveedor(IDProv, supervisor, tipo string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratosProveedor []map[string]interface{}
	//error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"contrato_general?query=Contratista:"+IDProv, &ContratosProveedor)
	urlCRUD := beego.AppConfig.String("administrativa_amazon_api_url") + beego.AppConfig.String("administrativa_amazon_api_version") + "contrato_general?query="
	query := CrearQueryContratoGeneral(IDProv, "0", "0", supervisor, tipo)

	response, err := getJsonTest(urlCRUD+query, &ContratosProveedor)
	if err != nil || response != 200 {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosProveedor1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}

	if len(ContratosProveedor) < 1 {
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosProveedor2", "err": "No se encontraron contratos", "status": "204"}
		return nil, outputError
	} else {
		return ContratosProveedor, nil
	}
}
