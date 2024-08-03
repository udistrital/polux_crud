package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReporteGeneral struct {
	Id                 int
	TrabajoGrado       int       `orm:"column(trabajo_grado)"`
	Titulo             string    `orm:"column(titulo)"`
	Modalidad          string    `orm:"column(modalidad)"`
	EstadoTrabajoGrado string    `orm:"column(estado)"`
	IdEstudiante       string    `orm:"column(id_estudiante)"`
	IdCoestudiante     string    `orm:"column(id_coestudiante)"`
	AreaConocimiento   string    `orm:"column(area_conocimiento)"`
	DocenteDirector    int       `orm:"column(docente_director)"`
	DocenteCodirector  int       `orm:"column(docente_codirector)"`
	Evaluador          int       `orm:"column(evaluador)"`
	FechaInicio        time.Time `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin           time.Time `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	CalificacionUno    float32   `orm:"column(calificacion_1)"`
	CalificacionDos    float32   `orm:"column(calificacion_2)"`
}

func init() {
	orm.RegisterModel(new(ReporteGeneral))
}

func GetReporteGeneral() (map[string]interface{}, error) {
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
		),
		notas AS (
			SELECT
				trabajo_grado,
				MAX(CASE WHEN row_number = 1 THEN calificacion ELSE NULL END) AS nota2,
				MAX(CASE WHEN row_number = 2 THEN calificacion ELSE NULL END) AS nota1
			FROM (
				SELECT
					trabajo_grado,
					calificacion,
					ROW_NUMBER() OVER (PARTITION BY trabajo_grado ORDER BY calificacion) AS row_number
				FROM
					academica.asignatura_trabajo_grado
			) sub
			GROUP BY
				trabajo_grado
		)

		SELECT
			tg.id AS trabajo_grado,
			tg.titulo,
			tg.modalidad,
			tg.estado_trabajo_grado,
			est.id_estudiante,
			est.id_coestudiante,
			atg.area_conocimiento,
			usr.docente_director,
			usr.docente_codirector,
			usr.evaluador,
			vtg.fecha_inicio,
			vtg.fecha_fin,
			notas.nota1 AS calificacion_1,
			notas.nota2 AS calificacion_2
		FROM
			academica.trabajo_grado tg
		LEFT JOIN
			estudiantes est ON tg.id = est.trabajo_grado
		LEFT JOIN
			academica.areas_trabajo_grado atg ON tg.id = atg.trabajo_grado
		LEFT JOIN
			academica.vinculacion_trabajo_grado vtg ON tg.id = vtg.trabajo_grado
		LEFT JOIN
			usuarios usr ON tg.id = usr.trabajo_grado
		LEFT JOIN
			notas ON tg.id = notas.trabajo_grado
		GROUP BY
			tg.id, tg.titulo, tg.modalidad, tg.estado_trabajo_grado, est.id_estudiante, est.id_coestudiante, atg.area_conocimiento, usr.docente_director, usr.docente_codirector, usr.evaluador, vtg.fecha_inicio, vtg.fecha_fin, notas.nota1, notas.nota2
		ORDER BY
			tg.id DESC`).Values(&results)

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

	var reporteGeneral []ReporteGeneral
	contador := 0

	for _, result := range results {
		v := ReporteGeneral{
			Id: contador,
		}

		if val, ok := result["trabajo_grado"].(string); ok {
			v.TrabajoGrado, _ = strconv.Atoi(val)
		}

		if val, ok := result["titulo"].(string); ok && val != "" {
			v.Titulo = val
		}

		if val, ok := result["modalidad"].(string); ok {
			v.Modalidad = val
		}

		if val, ok := result["estado_trabajo_grado"].(string); ok {
			v.EstadoTrabajoGrado = val
		}

		if val, ok := result["id_estudiante"].(string); ok && val != "" {
			v.IdEstudiante = val
		}

		if val, ok := result["id_coestudiante"].(string); ok && val != "" {
			v.IdCoestudiante = val
		}

		if val, ok := result["area_conocimiento"].(string); ok {
			v.AreaConocimiento = val
		}

		if val, ok := result["docente_director"].(string); ok {
			v.DocenteDirector, _ = strconv.Atoi(val)
		}

		if val, ok := result["docente_codirector"].(string); ok && val != "" {
			v.DocenteCodirector, _ = strconv.Atoi(val)
		}

		if val, ok := result["evaluador"].(string); ok && val != "" {
			v.Evaluador, _ = strconv.Atoi(val)
		}

		if val, ok := result["fecha_inicio"].(string); ok {
			fechaInicio, _ := time.Parse(time.RFC3339, val)
			v.FechaInicio = fechaInicio
		}

		if val, ok := result["fecha_fin"].(string); ok && val != "" {
			fechaFin, _ := time.Parse(time.RFC3339, val)
			v.FechaFin = fechaFin
		}

		if val, ok := result["calificacion_1"].(string); ok && val != "" {
			floatVal, _ := strconv.ParseFloat(val, 32)
			v.CalificacionUno = float32(floatVal)
		}

		if val, ok := result["calificacion_2"].(string); ok && val != "" {
			floatVal, _ := strconv.ParseFloat(val, 32)
			v.CalificacionDos = float32(floatVal)
		}

		reporteGeneral = append(reporteGeneral, v)
		contador++
	}

	return map[string]interface{}{
		"Success": true,
		"Data":    reporteGeneral,
	}, nil
}
