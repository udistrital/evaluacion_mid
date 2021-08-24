package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// OrganizarInfoContratos ...
func OrganizarInfoContratos(infoProveedor []map[string]interface{}, infoContratos []map[string]interface{}) (contratos []map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	for i := 0; i < len(infoContratos); i++ {
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"IdProveedor":      infoContratos[i]["Contratista"],
			"NombreProveedor":  infoProveedor[0]["NomProveedor"],
			"ContratoSuscrito": GetElementoMaptoString(infoContratos[i]["ContratoSuscrito"], "NumeroContratoSuscrito"),
			"Vigencia":         infoContratos[i]["VigenciaContrato"],
			// "Cotizacion":            infoContratos[i],
			"DependenciaSupervisor": GetElemento(infoContratos[i]["Supervisor"], "DependenciaSupervisor"),
		})
	}
	return InfoOrganizada
}

// OrganizarInfoContratosMultipleProv ...
func OrganizarInfoContratosMultipleProv(infoContratos []map[string]interface{}) (contratos []map[string]interface{}, outputError map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	NomProveedor := []map[string]interface{}{}
	for i := 0; i < len(infoContratos); i++ {
		IDProv := fmt.Sprintf("%v", infoContratos[i]["Contratista"])
		resultProv, err := InfoProveedorID(IDProv)
		if resultProv != nil {
			NomProveedor = resultProv

		} else {
			return nil, err
		}
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"IdProveedor":           infoContratos[i]["Contratista"],
			"NombreProveedor":       NomProveedor[0]["NomProveedor"],
			"ContratoSuscrito":      GetElementoMaptoString(infoContratos[i]["ContratoSuscrito"], "NumeroContratoSuscrito"),
			"Vigencia":              infoContratos[i]["VigenciaContrato"],
			"DependenciaSupervisor": GetElemento(infoContratos[i]["Supervisor"], "DependenciaSupervisor"),
		})
	}
	return InfoOrganizada, nil
}

// InfoProveedorID ...
func InfoProveedorID(IDProv string) (proveedor []map[string]interface{}, outputError map[string]interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	//error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=Id:"+IDProv+"&limit=0", &infoProveedor)
	if response, err := getJsonTest(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=Id:"+IDProv+"&limit=0", &infoProveedor); (err == nil) && (response == 200) {
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InfoProveedorID1", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	if len(infoProveedor) < 1 {
		outputError = map[string]interface{}{"funcion": "/InfoProveedorID2", "err": "No se pudo traer la info del proveedor", "status": "502"}
		return nil, outputError
	} else {
		return infoProveedor, nil
	}
}

// FiltroDependencia ...
func FiltroDependencia(infoContratos []map[string]interface{}, dependencias map[string]interface{}) (listaFiltrada []map[string]interface{}, outputError map[string]interface{}) {
	DependenciasSic := make([]map[string]interface{}, 0)
	InfoFiltrada := make([]map[string]interface{}, 0)
	DependenciasSic = append(DependenciasSic, dependencias)
	Dependencia := GetElemento(DependenciasSic[0]["DependenciasSic"], "Dependencia")
	ArrayDependencia, errElemento := GetElementoMaptoStringToArray(Dependencia, "ESFCODIGODEP")
	if ArrayDependencia != nil {
		for i := 0; i < len(infoContratos); i++ {
			for _, Dep := range ArrayDependencia {
				if Dep == infoContratos[i]["DependenciaSupervisor"] {
					InfoFiltrada = append(InfoFiltrada, infoContratos[i])
				}
			}
		}
		if len(InfoFiltrada) > 0 {
			return InfoFiltrada, nil
		} else {
			outputError = map[string]interface{}{"funcion": "/FiltroDependencia1", "err": "Segun las dependencias de las que es supervisor no tiene contratos disponibles", "status": "204"}
			return nil, outputError
			//	errorContratos := CrearError("Segun las dependencias de las que es supervisor no tiene contratos disponibles")
			//	return nil, errorContratos
		}
	} else {
		return nil, errElemento
	}

}

// OrganizarInfoContratoArgo ...
func OrganizarInfoContratoArgo(infoProveedor []map[string]interface{}, infoContrato []map[string]interface{}, infoActividades map[string]interface{}, infoDependencia []map[string]interface{}, infoSupervisor []map[string]interface{}) (infoOrganizada []map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	for i := 0; i < len(infoContrato); i++ {
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"contrato_general":      infoContrato[0],
			"informacion_proveedor": infoProveedor[0],
			"dependencia_SIC":       infoDependencia[0],
			"supervisor_contrato":   infoSupervisor[0],
			"actividades_contrato":  infoActividades,
		})
	}
	return InfoOrganizada
}
