// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/Polux_API_Crud/controllers"

	"github.com/astaxie/beego"
	/*Incluyendo líbreria de auditoría
	"github.com/udistrital/auditoria"*/)

func init() {
	//Iniciando middleware
	//auditoria.InitMiddleware()

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_solicitud",
			beego.NSInclude(
				&controllers.EstadoSolicitudController{},
			),
		),

		beego.NSNamespace("/documento_solicitud",
			beego.NSInclude(
				&controllers.DocumentoSolicitudController{},
			),
		),

		beego.NSNamespace("/respuesta_solicitud",
			beego.NSInclude(
				&controllers.RespuestaSolicitudController{},
			),
		),

		beego.NSNamespace("/usuario_solicitud",
			beego.NSInclude(
				&controllers.UsuarioSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_solicitud",
			beego.NSInclude(
				&controllers.TipoSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_detalle",
			beego.NSInclude(
				&controllers.TipoDetalleController{},
			),
		),

		beego.NSNamespace("/detalle",
			beego.NSInclude(
				&controllers.DetalleController{},
			),
		),

		beego.NSNamespace("/detalle_tipo_solicitud",
			beego.NSInclude(
				&controllers.DetalleTipoSolicitudController{},
			),
		),

		beego.NSNamespace("/detalle_solicitud",
			beego.NSInclude(
				&controllers.DetalleSolicitudController{},
			),
		),

		beego.NSNamespace("/documento_entidad",
			beego.NSInclude(
				&controllers.DocumentoEntidadController{},
			),
		),

		beego.NSNamespace("/pregunta",
			beego.NSInclude(
				&controllers.PreguntaController{},
			),
		),

		beego.NSNamespace("/respuesta",
			beego.NSInclude(
				&controllers.RespuestaController{},
			),
		),

		beego.NSNamespace("/respuesta_formato",
			beego.NSInclude(
				&controllers.RespuestaFormatoController{},
			),
		),

		beego.NSNamespace("/respuesta_evaluacion",
			beego.NSInclude(
				&controllers.RespuestaEvaluacionController{},
			),
		),

		beego.NSNamespace("/comentario",
			beego.NSInclude(
				&controllers.ComentarioController{},
			),
		),

		beego.NSNamespace("/correccion",
			beego.NSInclude(
				&controllers.CorreccionController{},
			),
		),

		beego.NSNamespace("/carrera_elegible",
			beego.NSInclude(
				&controllers.CarreraElegibleController{},
			),
		),

		beego.NSNamespace("/espacios_academicos_elegibles",
			beego.NSInclude(
				&controllers.EspaciosAcademicosElegiblesController{},
			),
		),

		beego.NSNamespace("/estructura_investigacion_trabajo_grado",
			beego.NSInclude(
				&controllers.EstructuraInvestigacionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/areas_docente",
			beego.NSInclude(
				&controllers.AreasDocenteController{},
			),
		),

		beego.NSNamespace("/area_conocimiento",
			beego.NSInclude(
				&controllers.AreaConocimientoController{},
			),
		),

		beego.NSNamespace("/areas_trabajo_grado",
			beego.NSInclude(
				&controllers.AreasTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/documento_escrito",
			beego.NSInclude(
				&controllers.DocumentoEscritoController{},
			),
		),

		beego.NSNamespace("/documento_trabajo_grado",
			beego.NSInclude(
				&controllers.DocumentoTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/modalidad",
			beego.NSInclude(
				&controllers.ModalidadController{},
			),
		),

		beego.NSNamespace("/formato",
			beego.NSInclude(
				&controllers.FormatoController{},
			),
		),

		beego.NSNamespace("/rol_trabajo_grado",
			beego.NSInclude(
				&controllers.RolTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estado_espacio_academico_inscrito",
			beego.NSInclude(
				&controllers.EstadoEspacioAcademicoInscritoController{},
			),
		),

		beego.NSNamespace("/espacio_academico_inscrito",
			beego.NSInclude(
				&controllers.EspacioAcademicoInscritoController{},
			),
		),

		beego.NSNamespace("/vinculacion_trabajo_grado",
			beego.NSInclude(
				&controllers.VinculacionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/formato_evaluacion_carrera",
			beego.NSInclude(
				&controllers.FormatoEvaluacionCarreraController{},
			),
		),

		beego.NSNamespace("/socializacion",
			beego.NSInclude(
				&controllers.SocializacionController{},
			),
		),

		beego.NSNamespace("/evaluacion_trabajo_grado",
			beego.NSInclude(
				&controllers.EvaluacionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estado_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/distincion_trabajo_grado",
			beego.NSInclude(
				&controllers.DistincionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/trabajo_grado",
			beego.NSInclude(
				&controllers.TrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estado_revision_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoRevisionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/revision_trabajo_grado",
			beego.NSInclude(
				&controllers.RevisionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/tipo_pregunta",
			beego.NSInclude(
				&controllers.TipoPreguntaController{},
			),
		),

		beego.NSNamespace("/pregunta_formato",
			beego.NSInclude(
				&controllers.PreguntaFormatoController{},
			),
		),

		beego.NSNamespace("/estado_estudiante_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoEstudianteTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estudiante_trabajo_grado",
			beego.NSInclude(
				&controllers.EstudianteTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estado_asignatura_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoAsignaturaTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/asignatura_trabajo_grado",
			beego.NSInclude(
				&controllers.AsignaturaTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/modalidad_tipo_solicitud",
			beego.NSInclude(
				&controllers.ModalidadTipoSolicitudController{},
			),
		),

		beego.NSNamespace("/solicitud_trabajo_grado",
			beego.NSInclude(
				&controllers.SolicitudTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/tr_solicitud",
			beego.NSInclude(
				&controllers.TrSolicitudController{},
			),
		),

		beego.NSNamespace("/tr_trabajo_grado",
			beego.NSInclude(
				&controllers.TrTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/tr_respuesta_solicitud",
			beego.NSInclude(
				&controllers.TrRespuestaSolicitudController{},
			),
		),

		beego.NSNamespace("/tr_registrar_materias_posgrado",
			beego.NSInclude(
				&controllers.TrRegistrarMateriasPosgrado{},
			),
		),

		beego.NSNamespace("/tr_registrar_respuestas_solicitudes",
			beego.NSInclude(
				&controllers.TrRegistrarRespuestasSolicitudesController{},
			),
		),

		beego.NSNamespace("/tr_registrar_nota",
			beego.NSInclude(
				&controllers.TrRegistrarNota{},
			),
		),

		beego.NSNamespace("/tr_revisar_anteproyecto",
			beego.NSInclude(
				&controllers.TrRevisarAnteproyecto{},
			),
		),

		beego.NSNamespace("/tr_actualizar_documento_tg",
			beego.NSInclude(
				&controllers.TrActualizarDocumentoTg{},
			),
		),
	)
	beego.AddNamespace(ns)
}
