package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:ContratosProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:ContratosProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:ContratoscontratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:ContratoscontratoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:DatosContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:DatosContratoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:FiltromixtoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:FiltromixtoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:InformacionCertificacionDveController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:InformacionCertificacionDveController"],
        beego.ControllerComments{
            Method: "PostInformacionCertificacionDve",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
