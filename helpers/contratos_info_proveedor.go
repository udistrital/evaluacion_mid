package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
)

// ListaContratosProveedor ...
func ListaContratosProveedor(IdentProv string, Idsuper string) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	resultProv, err1 := InfoProveedor(IdentProv)
	if resultProv != nil {
		IDProveedor := models.GetElementoMaptoString(resultProv, "Id")
		resultContrato, err2 := ObtenerContratosProveedor(IDProveedor)
		if resultContrato != nil {
			InfoOrg := models.OrganizarInfoContratos(resultProv, resultContrato)
			if Idsuper == "0" {
				return InfoOrg, nil
			} else {
				resultDependencia, errDep := models.ObtenerDependencias(Idsuper)
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
			}
		} else {
			return nil, err2
		}
	} else {
		return nil, err1
	}

}

// InfoProveedor ...
func InfoProveedor(IdentProv string) (proveedor []map[string]interface{}, outputError map[string]interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	//error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+"informacion_proveedor?query=NumDocumento:"+IdentProv+"&limit=0", &infoProveedor)
	if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+"informacion_proveedor?query=NumDocumento:"+IdentProv+"&limit=0", &infoProveedor); (err == nil) && (response == 200) {
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
func ObtenerContratosProveedor(IDProv string) (contrato []map[string]interface{}, outputError map[string]interface{}) {
	var ContratosProveedor []map[string]interface{}
	//error := getJson(beego.AppConfig.String("administrativa_amazon_api_url")+"contrato_general?query=Contratista:"+IDProv, &ContratosProveedor)
	if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+"contrato_general?query=Contratista:"+IDProv, &ContratosProveedor); (err == nil) && (response == 200) {
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosProveedor1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	if len(ContratosProveedor) < 1 {
		///fmt.Println(error)
		//errorContrato := models.CrearError("no se encontraron contratos")
		//return nil, errorContrato
		outputError = map[string]interface{}{"funcion": "/ObtenerContratosProveedor2", "err": "No se encontraron contratos", "status": "204"}
		return nil, outputError
	} else {
		return ContratosProveedor, nil
	}
}
