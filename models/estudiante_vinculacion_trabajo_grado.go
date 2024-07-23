package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type EstudianteVinculacionTrabajoGrado struct {
	Id              int           `orm:"column(id);pk;auto"`
	Usuario         int           `orm:"column(usuario)"`
	Activo          bool          `orm:"column(activo)"`
	FechaInicio     time.Time     `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin        time.Time     `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	RolTrabajoGrado int           `orm:"column(rol_trabajo_grado)"`
	TrabajoGrado    *TrabajoGrado `orm:"column(trabajo_grado);rel(fk)"`
	Estudiante      *string       `orm:"column(estudiante)"`
}

func init() {
	orm.RegisterModel(new(EstudianteVinculacionTrabajoGrado))
}

func GetEstudianteVinculacionTrabajoGrado(documento string) ([]EstudianteVinculacionTrabajoGrado, error) {
	o := orm.NewOrm()

	var results []orm.Params

	_, err := o.Raw(`
		SELECT 
			vt.id, vt.usuario, vt.activo, vt.fecha_inicio, vt.fecha_fin, vt.rol_trabajo_grado, vt.trabajo_grado,
			tg.titulo, tg.modalidad, tg.estado_trabajo_grado, tg.periodo_academico, tg.objetivo,
			STRING_AGG(et.estudiante, ', ') AS estudiante
		FROM 
			academica.vinculacion_trabajo_grado vt
		LEFT JOIN 
			academica.estudiante_trabajo_grado et ON vt.trabajo_grado = et.trabajo_grado
		LEFT JOIN 
			academica.trabajo_grado tg ON vt.trabajo_grado = tg.id
		WHERE 
			vt.activo = true AND vt.usuario = ?
		GROUP BY 
			vt.id, vt.usuario, vt.activo, vt.fecha_inicio, vt.fecha_fin, vt.rol_trabajo_grado, vt.trabajo_grado,
		    tg.titulo, tg.modalidad, tg.estado_trabajo_grado, tg.periodo_academico, tg.objetivo
		ORDER BY vt.id DESC`, documento).Values(&results)

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("No se encontraron resultados")
	}

	var vinculaciones []EstudianteVinculacionTrabajoGrado

	for _, result := range results {
		v := EstudianteVinculacionTrabajoGrado{
			TrabajoGrado: &TrabajoGrado{},
		}

		if val, ok := result["id"].(string); ok {
			v.Id, _ = strconv.Atoi(val)
		}

		if val, ok := result["usuario"].(string); ok {
			v.Usuario, _ = strconv.Atoi(val)
		}

		if val, ok := result["activo"].(string); ok {
			v.Activo, _ = strconv.ParseBool(val)
		}

		if val, ok := result["fecha_inicio"].(string); ok {
			fechaInicio, _ := time.Parse(time.RFC3339, val)
			v.FechaInicio = fechaInicio
		}

		if val, ok := result["fecha_fin"].(string); ok && val != "" {
			fechaFin, _ := time.Parse(time.RFC3339, val)
			v.FechaFin = fechaFin
		}

		if val, ok := result["rol_trabajo_grado"].(string); ok {
			v.RolTrabajoGrado, _ = strconv.Atoi(val)
		}

		if val, ok := result["trabajo_grado"].(string); ok {
			v.TrabajoGrado.Id, _ = strconv.Atoi(val)
		}

		if val, ok := result["titulo"].(string); ok && val != "" {
			v.TrabajoGrado.Titulo = val
		}

		if val, ok := result["modalidad"].(string); ok {
			v.TrabajoGrado.Modalidad, _ = strconv.Atoi(val)
		}

		if val, ok := result["estado_trabajo_grado"].(string); ok {
			v.TrabajoGrado.EstadoTrabajoGrado, _ = strconv.Atoi(val)
		}

		if val, ok := result["periodo_academico"].(string); ok && val != "" {
			v.TrabajoGrado.PeriodoAcademico = val
		}

		if val, ok := result["objetivo"].(string); ok && val != "" {
			v.TrabajoGrado.Objetivo = val
		}

		if val, ok := result["estudiante"].(string); ok && val != "" {
			v.Estudiante = &val
		}

		vinculaciones = append(vinculaciones, v)
	}

	return vinculaciones, nil
}
