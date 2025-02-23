package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReporteSolicitud struct {
	Id                      int    `orm:"column(id)"`
	TrabajoGrado            int    `orm:"column(trabajo_grado)"`
	Titulo                  string `orm:"column(titulo)"`
	Modalidad               string `orm:"column(modalidad)"`
	EstadoTrabajoGrado      string `orm:"column(estado_trabajo_grado)"`
	IdEstudiante            string `orm:"column(id_estudiante)"`
	NombreEstudiante        string
	IdCoestudiante          string `orm:"column(id_coestudiante)"`
	NombreCoestudiante      string
	ProgramaAcademico       string
	NombreCoordinador       string
	DocenteDirector         int `orm:"column(docente_director)"`
	NombreDocenteDirector   string
	DocenteCodirector       int `orm:"column(docente_codirector)"`
	NombreDocenteCodirector string
	Evaluador               int `orm:"column(evaluador)"`
	NombreEvaluador         string
	FechaSolicitud          time.Time `orm:"column(fecha_solicitud);type(timestamp without time zone)"`
	FechaRevision           time.Time `orm:"column(fecha_revision);type(timestamp without time zone);null"`
	Solicitud               string    `orm:"column(tipo_solicitud)"`
	Observacion             string    `orm:"column(justificacion)"`
	Respuesta               string    `orm:"column(estado_solicitud)"`
}

func init() {
	orm.RegisterModel(new(ReporteSolicitud))
}

func GetReporteSolicitud(Filtro *FiltrosReporte) (map[string]interface{}, error) {
	o := orm.NewOrm()

	var results []orm.Params

	var query = `
		WITH estudiantes AS (
			SELECT
				trabajo_grado,
				MAX(CASE WHEN row_number = 1 THEN estudiante ELSE NULL END) AS id_estudiante,
				MAX(CASE WHEN row_number = 2 THEN estudiante ELSE NULL END) AS id_coestudiante
			FROM (
				SELECT
					trabajo_grado,
					estudiante,
					ROW_NUMBER() OVER (PARTITION BY trabajo_grado ORDER BY estudiante) AS row_number
				FROM
					academica.estudiante_trabajo_grado
			) sub
			GROUP BY
				trabajo_grado
		),
		usuarios AS (
			SELECT
				trabajo_grado,
				MAX(CASE WHEN codigo_abreviacion = 'DIRECTOR_PLX' THEN usuario ELSE NULL END) AS docente_director,
				MAX(CASE WHEN codigo_abreviacion = 'CODIRECTOR_PLX' THEN usuario ELSE NULL END) AS docente_codirector,
				MAX(CASE WHEN codigo_abreviacion = 'EVALUADOR_PLX' THEN usuario ELSE NULL END) AS evaluador
			FROM
				academica.vinculacion_trabajo_grado
			GROUP BY
				trabajo_grado
		)

		SELECT 
			stg.id,
			stg.trabajo_grado,
			tg.titulo,
			mts.modalidad,
			tg.estado_trabajo_grado,
			e.id_estudiante,
			e.id_coestudiante,
			u.docente_director,
			u.docente_codirector,
			u.evaluador,
			stg.fecha AS fecha_solicitud,
			rs.fecha AS fecha_revision,
			mts.tipo_solicitud,
			rs.justificacion,
			rs.estado_solicitud
		FROM 
			academica.solicitud_trabajo_grado stg
		JOIN 
			academica.modalidad_tipo_solicitud mts
			ON stg.modalidad_tipo_solicitud = mts.id
		JOIN 
			academica.trabajo_grado tg
			ON stg.trabajo_grado = tg.id
		LEFT JOIN 
			academica.respuesta_solicitud rs
			ON stg.id = rs.solicitud_trabajo_grado
		LEFT JOIN 
			estudiantes e
			ON stg.trabajo_grado = e.trabajo_grado
		LEFT JOIN 
			usuarios u
			ON stg.trabajo_grado = u.trabajo_grado
		`

	var f_i = Filtro.FechaInicio.Format("2006-01-02")
	var f_f = Filtro.FechaFin.Format("2006-01-02")

	//se aplica el filtro de fechas
	query += `WHERE stg.fecha BETWEEN '`+ f_i + `' AND '`+ f_f+`'`

	//se aplica el filtro de estado
	if Filtro.Estado == "ACTIVOS" { //Si el estado seleccionado es activo, se traen todas las solicitudes con estado diferenta a aprobado o rechazado por coordinacion
		query += ` AND NOT rs.estado_solicitud = `+ strconv.Itoa(Filtro.IdEstFinalizado) +` AND NOT rs.estado_solicitud = ` + strconv.Itoa(Filtro.IdEstCancelado)
	} else if Filtro.Estado == "CULMINADOS" {///Si el estado seleccionado es culminado, se traen todas las solicitudes con estado aprobado o rechazado por coordinacion
		query += ` AND (rs.estado_solicitud = ` + strconv.Itoa(Filtro.IdEstFinalizado) + ` OR rs.estado_solicitud = ` + strconv.Itoa(Filtro.IdEstCancelado) + `)`
	}

	query += ` ORDER BY stg.id DESC`
	
	_, err := o.Raw(query).Values(&results)

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return map[string]interface{}{
			"Success": false,
			"Status":  "404",
			"Message": "No se encontraron Resultados",
			"Data":    []interface{}{},
		}, nil
	}

	var reporteSolicitud []ReporteSolicitud

	for _, result := range results {
		v := ReporteSolicitud{}

		if val, ok := result["id"].(string); ok {
			v.Id, _ = strconv.Atoi(val)
		}

		if val, ok := result["trabajo_grado"].(string); ok {
			v.TrabajoGrado, _ = strconv.Atoi(val)
		}

		if val, ok := result["titulo"].(string); ok && val != "" {
			v.Titulo = val
		}

		if val, ok := result["modalidad"].(string); ok && val != "" {
			v.Modalidad = val
		}

		if val, ok := result["estado_trabajo_grado"].(string); ok && val != "" {
			v.EstadoTrabajoGrado = val
		}

		if val, ok := result["id_estudiante"].(string); ok && val != "" {
			v.IdEstudiante = val
		}

		if val, ok := result["id_coestudiante"].(string); ok && val != "" {
			v.IdCoestudiante = val
		}

		v.ProgramaAcademico = ""

		v.NombreCoordinador = ""

		if val, ok := result["docente_director"].(string); ok {
			v.DocenteDirector, _ = strconv.Atoi(val)
		}

		if val, ok := result["docente_codirector"].(string); ok && val != "" {
			v.DocenteCodirector, _ = strconv.Atoi(val)
		}

		if val, ok := result["evaluador"].(string); ok && val != "" {
			v.Evaluador, _ = strconv.Atoi(val)
		}

		if val, ok := result["fecha_solicitud"].(string); ok {
			fechaSolicitud, _ := time.Parse(time.RFC3339, val)
			v.FechaSolicitud = fechaSolicitud
		}

		if val, ok := result["fecha_revision"].(string); ok && val != "" {
			fechaRevision, _ := time.Parse(time.RFC3339, val)
			v.FechaRevision = fechaRevision
		}

		if val, ok := result["tipo_solicitud"].(string); ok && val != "" {
			v.Solicitud = val
		}

		if val, ok := result["justificacion"].(string); ok && val != "" {
			v.Observacion = val
		}

		if val, ok := result["estado_solicitud"].(string); ok && val != "" {
			v.Respuesta = val
		}

		reporteSolicitud = append(reporteSolicitud, v)
	}

	return map[string]interface{}{
		"Success": true,
		"Status":  "201",
		"Message": "Reporte Solicitud generado correctamente",
		"Data":    reporteSolicitud,
	}, nil
}
