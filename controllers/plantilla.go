package controllers

import (
	"github.com/astaxie/beego"
)

// PlantillaController operations for Plantilla
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
// @Description create Plantilla
// @Param	body		body 	models.Plantilla	true		"body for Plantilla content"
// @Success 201 {object} models.Plantilla
// @Failure 403 body is empty
// @Failure 400 Bad Request
// @router / [post]
func (c *PlantillaController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Plantilla by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Plantilla
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PlantillaController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Plantilla
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Plantilla
// @Failure 403
// @router / [get]
func (c *PlantillaController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Plantilla
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Plantilla	true		"body for Plantilla content"
// @Success 200 {object} models.Plantilla
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
