// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/evaluacion_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/filtroContrato",
			beego.NSInclude(
				&controllers.ContatoscontratoController{},
			),
		),
		beego.NSNamespace("/filtroProveedor",
			beego.NSInclude(
				&controllers.ContratosProveedorController{},
			),
		),
		beego.NSNamespace("/filtroMixto",
			beego.NSInclude(
				&controllers.FiltromixtoController{},
			),
		),
		beego.NSNamespace("/plantilla",
			beego.NSInclude(
				&controllers.PlantillaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
