package models

import (
	"fmt"
	"strconv"
	"strings"
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

func GetProyectoVinculacionTrabajoGrado(proyectos string) ([]EstudianteVinculacionTrabajoGrado, error) {
	o := orm.NewOrm()

	// Convertir string a slice de enteros
	proyectosInt, err := parseProyectos(proyectos)
	if err != nil {
		return nil, fmt.Errorf("Error al convertir proyectos a enteros: %v", err)
	}

	// Construir placeholders (generar dinámicamente una lista de ? según la cantidad de proyectos)
	placeholders := buildPlaceholders(len(proyectosInt))

	query := fmt.Sprintf(`
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
		LEFT JOIN 
			academica.espacio_academico_inscrito eai ON tg.id = eai.trabajo_grado
		WHERE 
			vt.activo = true 
			AND vt.codigo_abreviacion = 'COR_POSGRADO_PLX' 
			AND eai.proyecto_curricular_tg IN (%s)
		GROUP BY 
			vt.id, vt.usuario, vt.activo, vt.fecha_inicio, vt.fecha_fin, vt.rol_trabajo_grado, vt.trabajo_grado,
			tg.titulo, tg.modalidad, tg.estado_trabajo_grado, tg.periodo_academico, tg.objetivo
		ORDER BY 
			vt.id DESC`, placeholders)

	var results []orm.Params

	// Convertir el slice de enteros en un slice de interface{} para pasarlo como argumentos a la consulta.
	args := make([]interface{}, len(proyectosInt))
	for i, p := range proyectosInt {
		args[i] = p
	}

	// Ejecutar la consulta SQL
	_, err = o.Raw(query, args...).Values(&results)
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

// Función helper para convertir string de proyectos en slice de enteros
func parseProyectos(proyectos string) ([]int, error) {
	proyectosStr := strings.Split(proyectos, ",")
	proyectosInt := make([]int, len(proyectosStr))

	for i, p := range proyectosStr {
		id, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		proyectosInt[i] = id
	}

	return proyectosInt, nil
}

// Función helper para construir placeholders para la consulta SQL por proyectos
func buildPlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := range placeholders {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ",")
}
