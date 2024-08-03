package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasDocenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AreasTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:AsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CarreraElegibleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:CorreccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetallePasantiaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DetalleTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DistincionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEntidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoEscritoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:DocumentoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EspaciosAcademicosElegiblesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoAsignaturaTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEspacioAcademicoInscritoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoEstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoRevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstadoTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstructuraInvestigacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EstudianteTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:Estudiante_Vinculacion_Trabajo_GradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:Estudiante_Vinculacion_Trabajo_GradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:EvaluacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:FormatoEvaluacionCarreraController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ModalidadTipoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:PreguntaFormatoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ReporteGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:ReporteGeneralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaEvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaFormatoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "GetByUser",
            Router: "/user/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RevisionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:RolTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SocializacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:SolicitudTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoDetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoPreguntaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TipoSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrActualizarDocumentoTg"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrActualizarDocumentoTg"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrPublicarAsignaturasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrPublicarAsignaturasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarActaSeguimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarActaSeguimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarMateriasPosgrado"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarMateriasPosgrado"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarNota"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarNota"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarRespuestasSolicitudesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarRespuestasSolicitudesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarRevisionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRegistrarRevisionTgController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRespuestaSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRevisarTg"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrRevisarTg"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrSubirArlController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrSubirArlController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrVinculadoRegistrarNotaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrVinculadoRegistrarNotaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:TrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:UsuarioSolicitudController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/polux_crud/controllers:VinculacionTrabajoGradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
