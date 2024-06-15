package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/helpers"
)

type InformacionCertificacionDveController struct {
	beego.Controller
}

// URLMapping maps the InformacionCertificacionDveController methods to POST requests.
func (c *InformacionCertificacionDveController) URLMapping() {
	c.Mapping("PostInformacionCertificacionDve", c.PostInformacionCertificacionDve)
}

// PostInformacionCertificacionDve handles POST requests to get certification information.
// @Title PostInformacionCertificacionDve
// @Description get certification information by various parameters
// @Param   numero_documento      query  []string  true  "List of document numbers"
// @Param   periodo_inicial       query  []string  true  "List of initial periods"
// @Param   periodo_final         query  []string  true  "List of final periods"
// @Param   vinculaciones         query  []string  true  "List of linkages"
// @Success 200 {object} models.InformacionCertificacionDve
// @Failure 400 Bad request
// @router / [post]
func (c *InformacionCertificacionDveController) PostInformacionCertificacionDve() {

	type BodyParams struct {
		NumeroDocumento string `json:"numero_documento"`
		PeriodoInicial  string `json:"periodo_inicial"`
		PeriodoFinal    string `json:"periodo_final"`
		Vinculaciones   string `json:"vinculaciones"`
	}
	var v BodyParams

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	numeroDocumento := helpers.StringToSlice(v.NumeroDocumento)
	periodoInicial := helpers.StringToSlice(v.PeriodoInicial)
	periodoFinal := helpers.StringToSlice(v.PeriodoFinal)
	vinculaciones := helpers.StringToSlice(v.Vinculaciones)

	certificacion, err := helpers.InformacionCertificacionDve(numeroDocumento, periodoInicial, periodoFinal, vinculaciones)

	if err != nil {
		c.Ctx.Output.SetStatus(204)
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 204, "Message": "No hay datos que coincidan con los filtros", "Data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta exitosa", "Data": certificacion}
	}

	c.ServeJSON()
}
