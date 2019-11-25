package controllers

import (
	"github.com/astaxie/beego"
)

// ContatoscontratoController operations for Contatoscontrato
type ContatoscontratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContatoscontratoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Create
// @Description create Contatoscontrato
// @Param	body		body 	models.Contatoscontrato	true		"body for Contatoscontrato content"
// @Success 201 {object} models.Contatoscontrato
// @Failure 403 body is empty
// @router / [post]
func (c *ContatoscontratoController) Post() {

}

// GetAll ...
// @Title GetAll
// @Description get Contatoscontrato
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Contatoscontrato
// @Failure 403
// @router / [get]
func (c *ContatoscontratoController) GetAll() {

}
