package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaInscritaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturaTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:AsignaturasElegiblesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:DocumentoTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstadoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstructuraInvestigacionTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EstudianteTgController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PaqueteRespuestasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:SolicitudMateriasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TipoVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"] = append(beego.GlobalControllerRouter["apiPolux/controllers:VinculacionExternaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
