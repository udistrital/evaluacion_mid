package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// CesionesContratos Consulta las cesiones que se han realizado a una lista de contratos
func CesionesContratos(contratos []map[string]interface{}) (cesiones []map[string]interface{}, outputError map[string]interface{}) {

	basePath := beego.AppConfig.String("novedades_crud_url") + beego.AppConfig.String("novedades_crud_version")
	for _, contrato := range contratos {
		contratoSuscrito := contrato["ContratoSuscrito"]
		vigencia := contrato["Vigencia"]
		if contratoSuscrito == "" || vigencia == 0 {
			continue
		}

		var detalleCesion []map[string]interface{}
		query := "propiedad/?sortby=Id&order=desc&query=" + CrearQueryNovedadesCesion("0", fmt.Sprint(contratoSuscrito), fmt.Sprint(vigencia))

		//response, err := getJsonTest(basePath+query, &detalleCesion)
		response, err := request.GetJsonTest2(basePath+query, &detalleCesion)
		if err != nil || response != 200 {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/CesionesContratos2", "err": err.Error(), "status": "502"}
			return nil, outputError
		} else if len(detalleCesion) == 0 || detalleCesion[0]["Propiedad"] == nil {
			continue
		}

		nombreProv, outputError := models.InfoProveedorID(fmt.Sprint(detalleCesion[0]["Propiedad"]))
		if outputError != nil {
			return nil, outputError
		}

		cesion := map[string]interface{}{
			"ContratoSuscrito": contratoSuscrito,
			"IdProveedor":      detalleCesion[0]["Propiedad"],
			"NombreProveedor":  nombreProv[0]["NomProveedor"],
			"Vigencia":         vigencia,
		}
		cesiones = append(cesiones, cesion)

	}

	return
}
