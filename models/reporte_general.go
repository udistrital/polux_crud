package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReporteGeneral struct {
	Id                      int
	TrabajoGrado            int    `orm:"column(trabajo_grado)"`
	Titulo                  string `orm:"column(titulo)"`
	Modalidad               string `orm:"column(modalidad)"`
	EstadoTrabajoGrado      string `orm:"column(estado)"`
	IdEstudiante            string `orm:"column(id_estudiante)"`
	NombreEstudiante        string
	IdCoestudiante          string `orm:"column(id_coestudiante)"`
	NombreCoestudiante      string
	ProgramaAcademico       string
	AreaConocimiento        string `orm:"column(area_conocimiento)"`
	DocenteDirector         int    `orm:"column(docente_director)"`
	NombreDocenteDirector   string
	DocenteCodirector       int `orm:"column(docente_codirector)"`
	NombreDocenteCodirector string
	Evaluador               int `orm:"column(evaluador)"`
	NombreEvaluador         string
	FechaInicio             time.Time `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin                time.Time `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	CalificacionUno         float32   `orm:"column(calificacion_1)"`
	CalificacionDos         float32   `orm:"column(calificacion_2)"`
}

type FiltrosReporte struct {
	ProyectoCurricular	string
	Estado				string
	FechaInicio			time.Time
	FechaFin			time.Time
	IdEstFinalizado		int
	IdEstCancelado		int
}

func init() {
	orm.RegisterModel(new(ReporteGeneral))
}

func GetReporteGeneral(Filtro *FiltrosReporte) (map[string]interface{}, error) {
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
			(SELECT MAX(fecha_creacion)
			 FROM academica.asignatura_trabajo_grado atg
			 WHERE atg.trabajo_grado = tg.id
			 AND atg.codigo_asignatura = 2) 
			AS fecha_inicio,
			(SELECT MAX(fecha_modificacion)
			 FROM academica.asignatura_trabajo_grado atg
			 WHERE atg.trabajo_grado = tg.id
			 AND atg.codigo_asignatura = 2
			 AND atg.calificacion > 0) 
			AS fecha_fin,
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
		
		`
	var f_i = Filtro.FechaInicio.Format("2006-01-02")
	var f_f = Filtro.FechaFin.Format("2006-01-02")

	//se aplica el filtro de fechas
	query += `WHERE fecha_inicio BETWEEN '`+ f_i + `' AND '`+ f_f+`'`

	//se aplica el filtro de estado
	if Filtro.Estado == "ACTIVOS" { //Si el estado seleccionado es activo, se traen todos los tg con estado diferenta a cancelado y calificado
		query += ` AND NOT estado_trabajo_grado = `+ strconv.Itoa(Filtro.IdEstFinalizado) +` AND NOT estado_trabajo_grado = ` + strconv.Itoa(Filtro.IdEstCancelado)
	} else if Filtro.Estado == "CULMINADOS" {//Si el estado seleccionado es culminado, se traen todos los tg con estado calificado
		query += ` AND estado_trabajo_grado = `+ strconv.Itoa(Filtro.IdEstFinalizado)
	}
	
	query += ` GROUP BY tg.id, tg.titulo, tg.modalidad, tg.estado_trabajo_grado, est.id_estudiante, est.id_coestudiante, atg.area_conocimiento, usr.docente_director, usr.docente_codirector, usr.evaluador, notas.nota1, notas.nota2
			ORDER BY tg.id DESC`

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
		"Status":  "201",
		"Message": "Reporte General generado correctamente",
		"Data":    reporteGeneral,
	}, nil
}
