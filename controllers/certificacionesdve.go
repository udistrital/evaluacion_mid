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
	c.Mapping("GetVinculacionesDve", c.GetVinculacionesDve)
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
		NumeroDocumento string   `json:"numero_documento"`
		PeriodoInicial  string   `json:"periodo_inicial"`
		PeriodoFinal    string   `json:"periodo_final"`
		Vinculaciones   []string `json:"vinculaciones"`
		IncluirSalario  bool     `json:"incluir_salario"`
	}
	var v BodyParams

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	numeroDocumento := helpers.StringToSlice(v.NumeroDocumento)
	periodoInicial := helpers.StringToSlice(v.PeriodoInicial)
	periodoFinal := helpers.StringToSlice(v.PeriodoFinal)
	incluirSalario := v.IncluirSalario

	certificacion, err := helpers.InformacionCertificacionDve(numeroDocumento, periodoInicial, periodoFinal, v.Vinculaciones, incluirSalario)

	if err != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": err, "Data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta exitosa", "Data": certificacion}
	}

	c.ServeJSON()
}

// VinculacionesDve retrieves the teacher's linkage information.
// @Title VinculacionesDve
// @Description Fetches the linkage information of a teacher based on their document number.
// @Param   documento_docente     query  string  true  "Document number of the teacher"
// @Success 200 {object} []models.VinculacionesDocente "List of linkages for the teacher"
// @Failure 404 Not Found "The requested information could not be found"
// @Failure 502 Bad Gateway "Error processing the request due to an unexpected failure"
// @router /vinculaciones_docente/:documento_docente [get]
func (c *InformacionCertificacionDveController) GetVinculacionesDve() {

	defer helpers.ErrorControl(c.Controller, "InformacionCertificacionDveController")

	docDocente := c.Ctx.Input.Param(":documento_docente")
	if docDocente == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "Bad Request: docDocente is empty"}
		c.ServeJSON()
		return
	}

	vinculacionesDocente, outputError := helpers.VinculacionesDocenteDve(docDocente)
	if outputError != nil {
		c.Ctx.Output.SetStatus(outputError["Status"].(int))
		c.Data["json"] = outputError
	} else {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Vinculaciones obtenidas con éxito", "Data": vinculacionesDocente}
	}
	c.ServeJSON()
}
