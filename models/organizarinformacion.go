package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
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
func OrganizarInfoContratosMultipleProv(infoContratos []map[string]interface{}) (contratos []map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	NomProveedor := []map[string]interface{}{}
	for i := 0; i < len(infoContratos); i++ {
		IDProv := fmt.Sprintf("%v", infoContratos[i]["Contratista"])
		resultProv, err := InfoProveedorID(IDProv)
		if resultProv != nil {
			NomProveedor = resultProv

		} else {
			fmt.Println("entro a si nil contrato", err)
		}
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"IdProveedor":           infoContratos[i]["Contratista"],
			"NombreProveedor":       NomProveedor[0]["NomProveedor"],
			"ContratoSuscrito":      GetElementoMaptoString(infoContratos[i]["ContratoSuscrito"], "NumeroContratoSuscrito"),
			"Vigencia":              infoContratos[i]["VigenciaContrato"],
			"DependenciaSupervisor": GetElemento(infoContratos[i]["Supervisor"], "DependenciaSupervisor"),
		})
	}
	return InfoOrganizada
}

// InfoProveedorID ...
func InfoProveedorID(IDProv string) (proveedor []map[string]interface{}, outputError interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=Id:"+IDProv+"&limit=0", &infoProveedor)
	if len(infoProveedor) < 1 {
		fmt.Println(error)
		errorProv := CrearError("no se pudo traer la info del proveedor")
		return nil, errorProv
	} else {
		return infoProveedor, nil
	}
}

// FiltroDependencia ...
func FiltroDependencia(infoContratos []map[string]interface{}, dependencias map[string]interface{}) (listaFiltrada []map[string]interface{}, outputError interface{}) {
	DependenciasSic := make([]map[string]interface{}, 0)
	InfoFiltrada := make([]map[string]interface{}, 0)
	DependenciasSic = append(DependenciasSic, dependencias)
	Dependencia := GetElemento(DependenciasSic[0]["DependenciasSic"], "Dependencia")
	ArrayDependencia, errElmento := GetElementoMaptoStringToArray(Dependencia, "ESFCODIGODEP")
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
			errorContratos := CrearError("Segun las dependencias de las que es supervisor no tiene contratos disponibles")
			return nil, errorContratos
		}
	} else {
		return nil, errElmento
	}

}

// OrganizarInfoContratoArgo ...
func OrganizarInfoContratoArgo(infoProveedor []map[string]interface{}, infoContrato []map[string]interface{}, infoDependencia []map[string]interface{}) (infoOrganizada []map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	for i := 0; i < len(infoContrato); i++ {
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"contrato_general":      infoContrato[0],
			"informacion_proveedor": infoProveedor[0],
			"dependencia_SIC":       infoDependencia[0],
		})
	}
	return InfoOrganizada
}
