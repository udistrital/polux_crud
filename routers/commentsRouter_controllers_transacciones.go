package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrSolicitudController"] = append(beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrTrabajoGradoController"] = append(beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrTrabajoGradoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrFormatoController"] = append(beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrFormatoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrFormatoController"] = append(beego.GlobalControllerRouter["Polux_API_Crud/controllers:TrFormatoController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

		beego.GlobalControllerRouter["Polux_API_Crud/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["Polux_API_Crud/controllers:SolicitudMateriasController"],
			beego.ControllerComments{
				Method: "FormalizacionesPendientes",
				Router: `/FormalizacionesPendientes/:estado`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

}
