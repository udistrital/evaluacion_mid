package models

import (
	"github.com/udistrital/nuxeo_mid/models"
)

func OrganizarInfoContratos(infoProveedor []map[string]interface{}, infoContratos []map[string]interface{}) (novedad []map[string]interface{}) {
	registrojbpm := []map[string]interface{}{}
	// logs.Emergency(registrojbpm)
	// logs.Warning(len(infoContratos))
	// logs.Warning(infoContratos[0])
	for i := 0; i < len(infoContratos); i++ {
		// fmt.Println(infoContratos[i]["ContratoSuscrito"])
		registrojbpm = append(registrojbpm, map[string]interface{}{
			"IdProveedor":      infoContratos[i]["Contratista"],
			"NombreProveedor":  infoProveedor[0]["NomProveedor"],
			"ContratoSuscrito": models.GetElementoMaptoString(infoContratos[i]["ContratoSuscrito"], "NumeroContratoSuscrito"),
			"vigencia":         infoContratos[i]["VigenciaContrato"],
			"supervisor":       infoContratos[i]["Supervisor"],
		})
	}
	return registrojbpm
}
