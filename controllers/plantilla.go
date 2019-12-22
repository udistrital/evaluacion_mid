package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
)

// PlantillaController maneja el ingreso y optencion de plantillas para las evaluaciones
type PlantillaController struct {
	beego.Controller
}

// URLMapping ...
func (c *PlantillaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description se crea una nueva plantilla, para esto existe un json de guia para ello o se debera de crear un modulo para failitar el proceso
// @Param	body		body 	{}	true		"body for Plantilla content"
// @Success 201 {}
// @Failure 403 body is empty
// @Failure 400 Bad Request
// @router / [post]
func (c *PlantillaController) Post() {
	var plantillaRecivida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &plantillaRecivida); err == nil {

		plantillaRespuerta, errPlantilla := models.IngresoPlantilla(plantillaRecivida)

		if plantillaRespuerta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			// alertErr.Body = plantillaRespuerta
			alertErr.Body = models.CrearSuccess("la plantilla se ingreso con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errPlantilla)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}

// GetOne ...
// @Title GetOne
// @Description Obtiene la estructura de la platilla , segun el ID de la plnatilla enviado, similar al get all
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PlantillaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.Atoi(idStr)
	// fmt.Println(id)
	var alertErr models.Alert

	alertas := append([]interface{}{"Response:"})
	plantilla, errPlantilla := models.ObternerPlantillaPorID(idStr)
	if plantilla != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = plantilla
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errPlantilla)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description Obtiene la ultima plantilla activa en base de datos, la cual es un json con todas las propiedades necesarias para la interpretacion en el cliente
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {}
// @Failure 403
// @router / [get]
func (c *PlantillaController) GetAll() {
	var alertErr models.Alert

	alertas := append([]interface{}{"Response:"})
	plantilla, errPlantilla := models.ObtenerPlantillas()
	if plantilla != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = plantilla
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errPlantilla)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Plantilla
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	{}	true		"body for Plantilla content"
// @Success 200 {}
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PlantillaController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Plantilla
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PlantillaController) Delete() {

}
