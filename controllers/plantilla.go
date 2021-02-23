package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
	"github.com/udistrital/evaluacion_mid/helpers"
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

	defer helpers.ErrorControl(c.Controller, "PlantillaController")

	var plantillaRecivida map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &plantillaRecivida); err == nil {

		plantillaRespuerta, errPlantilla := models.IngresoPlantilla(plantillaRecivida)

		if plantillaRespuerta != nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": models.CrearSuccess("La plantilla se ingresó con exito")}
		} else {
			panic(errPlantilla)
		}

	} else {
		panic(err)
	}
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

	defer helpers.ErrorControl(c.Controller, "PlantillaController")

	idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.Atoi(idStr)
	// fmt.Println(id)

	plantilla, errPlantilla := models.ObtenerPlantillaPorID(idStr)
	if plantilla != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": plantilla}
	} else {
		panic(errPlantilla)
	}
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

	defer helpers.ErrorControl(c.Controller, "PlantillaController")

	plantilla, errPlantilla := models.ObtenerPlantillas()
	if plantilla != nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "successful", "Data": plantilla}
	} else {
		panic(errPlantilla)
	}
	c.ServeJSON()

}
