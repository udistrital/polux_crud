package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetallePasantiaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrActualizarDocumentoTg"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrActualizarDocumentoTg"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrPublicarAsignaturasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrPublicarAsignaturasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarActaSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarActaSeguimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarMateriasPosgrado"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarMateriasPosgrado"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarNota"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarNota"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarRespuestasSolicitudesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarRespuestasSolicitudesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarRevisionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRegistrarRevisionTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRevisarAnteproyecto"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrRevisarAnteproyecto"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrVinculadoRegistrarNotaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrVinculadoRegistrarNotaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
