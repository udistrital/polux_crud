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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ContactoEntidadController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoRevisionController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoContactoController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoIdentificacionController"],
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

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrSolicitudController"],
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
