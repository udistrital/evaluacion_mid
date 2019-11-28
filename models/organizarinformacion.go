package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/nuxeo_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// OrganizarInfoContratos ...
func OrganizarInfoContratos(infoProveedor []map[string]interface{}, infoContratos []map[string]interface{}) (contratos []map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	// logs.Emergency(InfoOrganizada)
	// logs.Warning(len(infoContratos))
	// logs.Warning(infoContratos[0])
	for i := 0; i < len(infoContratos); i++ {
		// fmt.Println(infoContratos[i]["ContratoSuscrito"])
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"IdProveedor":      infoContratos[i]["Contratista"],
			"NombreProveedor":  infoProveedor[0]["NomProveedor"],
			"ContratoSuscrito": models.GetElementoMaptoString(infoContratos[i]["ContratoSuscrito"], "NumeroContratoSuscrito"),
			"Vigencia":         infoContratos[i]["VigenciaContrato"],
			// "Cotizacion":            infoContratos[i],
			"DependenciaSupervisor": models.GetElemento(infoContratos[i]["Supervisor"], "DependenciaSupervisor"),
		})
	}
	return InfoOrganizada
}

// OrganizarInfoContratosMultipleProv ...
func OrganizarInfoContratosMultipleProv(infoContratos []map[string]interface{}) (contratos []map[string]interface{}) {
	InfoOrganizada := []map[string]interface{}{}
	NomProveedor := []map[string]interface{}{}
	// logs.Emergency(InfoOrganizada)
	// logs.Warning(len(infoContratos))
	// logs.Warning(infoContratos[0])
	for i := 0; i < len(infoContratos); i++ {
		// fmt.Println(infoContratos[i]["ContratoSuscrito"])
		IDProv := fmt.Sprintf("%v", infoContratos[i]["Contratista"])
		resultProv, err := InfoProveedorID(IDProv)
		if resultProv != nil {
			NomProveedor = resultProv

		} else {
			fmt.Println("entro a si nil contrato", err)
			// return nil, err1
		}
		InfoOrganizada = append(InfoOrganizada, map[string]interface{}{
			"IdProveedor":           infoContratos[i]["Contratista"],
			"NombreProveedor":       NomProveedor[0]["NomProveedor"],
			"ContratoSuscrito":      models.GetElementoMaptoString(infoContratos[i]["ContratoSuscrito"], "NumeroContratoSuscrito"),
			"Vigencia":              infoContratos[i]["VigenciaContrato"],
			"DependenciaSupervisor": models.GetElemento(infoContratos[i]["Supervisor"], "DependenciaSupervisor"),
		})
	}
	return InfoOrganizada
}

// InfoProveedorID ...
func InfoProveedorID(IDProv string) (proveedor []map[string]interface{}, outputError interface{}) {
	// registroNovedadPost := make(map[string]interface{})
	var infoProveedor []map[string]interface{}
	error := request.GetJson(beego.AppConfig.String("administrativa_amazon_api_url")+beego.AppConfig.String("administrativa_amazon_api_version")+"informacion_proveedor?query=Id:"+IDProv+"&limit=0", &infoProveedor)
	fmt.Println(len(infoProveedor))
	if len(infoProveedor) < 1 {
		fmt.Println(error)
		fmt.Println("entro al error")
		errorProv := CrearError("no se pudo traer la info del proveedor")
		return nil, errorProv
	} else {
		fmt.Println("ok")
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
				// fmt.Println(Dep)
				if Dep == infoContratos[i]["DependenciaSupervisor"] {
					fmt.Println("son iguales y es", Dep, infoContratos[i]["DependenciaSupervisor"])
					InfoFiltrada = append(InfoFiltrada, infoContratos[i])
				}
			}
		}
		if len(InfoFiltrada) > 0 {
			return InfoFiltrada, nil
		} else {
			errorContratos := CrearError("Segun las dependencias de las que es supervisor no tiene cntratos disponibles")
			return nil, errorContratos
		}
	} else {
		return nil, errElmento
	}

}
