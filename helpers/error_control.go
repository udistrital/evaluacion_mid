package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func ErrorControl(c beego.Controller, controller string) {
	if err := recover(); err != nil {
		logs.Error(err)
		localError := err.(map[string]interface{})
		c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + controller + "/" + (localError["funcion"]).(string))
		//fmt.Println(c.Data["mesaage"])
		c.Data["data"] = (localError["err"])
		if status, ok := localError["status"]; ok {
			c.Abort(status.(string))
		} else {
			c.Abort("404")
		}
	}
}