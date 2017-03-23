package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrFormatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_Crud/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
