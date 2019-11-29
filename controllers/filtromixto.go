package controllers

import (
	"github.com/astaxie/beego"
)

// FiltromixtoController operations for Filtromixto
type FiltromixtoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FiltromixtoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get Filtromixto
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Filtromixto
// @Failure 403
// @router / [get]
func (c *FiltromixtoController) GetAll() {

}
