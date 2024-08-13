package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReporteSolicitud struct {
	Id                 int       `orm:"column(id)"`
	TrabajoGrado       int       `orm:"column(trabajo_grado)"`
	Titulo             string    `orm:"column(titulo)"`
	Modalidad          string    `orm:"column(modalidad)"`
	EstadoTrabajoGrado string    `orm:"column(estado_trabajo_grado)"`
	IdEstudiante       string    `orm:"column(id_estudiante)"`
	IdCoestudiante     string    `orm:"column(id_coestudiante)"`
	ProgramaAcademico  string    `orm:"column(programa_academico)"`
	Coordinador        int       `orm:"column(coordinador)"`
	DocenteDirector    int       `orm:"column(docente_director)"`
	DocenteCodirector  int       `orm:"column(docente_codirector)"`
	Evaluador          int       `orm:"column(evaluador)"`
	FechaSolicitud     time.Time `orm:"column(fecha_solicitud);type(timestamp without time zone)"`
	FechaRevision      time.Time `orm:"column(fecha_revision);type(timestamp without time zone);null"`
	Solicitud          string    `orm:"column(tipo_solicitud)"`
	Observacion        string    `orm:"column(justificacion)"`
	Respuesta          string    `orm:"column(estado_solicitud)"`
}

func init() {
	orm.RegisterModel(new(ReporteSolicitud))
}

func GetReporteSolicitud() (map[string]interface{}, error) {
	o := orm.NewOrm()

	var results []orm.Params

	_, err := o.Raw(`
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
				MAX(CASE WHEN rol_trabajo_grado = 4593 THEN usuario ELSE NULL END) AS docente_director,
				MAX(CASE WHEN rol_trabajo_grado = 4594 THEN usuario ELSE NULL END) AS docente_codirector,
				MAX(CASE WHEN rol_trabajo_grado = 4595 THEN usuario ELSE NULL END) AS evaluador
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
		ORDER BY
			stg.id DESC`).Values(&results)

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return map[string]interface{}{
			"Success": false,
			"Data":    []interface{}{},
			"Message": "No se encontraron Resultados",
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

		v.Coordinador = 0

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
		"Data":    reporteSolicitud,
	}, nil
}
