package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRevisarAnteproyecto struct {
	TrabajoGrado           				*TrabajoGrado
	RevisionTrabajoGrado					*RevisionTrabajoGrado
	Correccion										*Correccion
}

// Función para la transaccion de revisiones de anteproyectos
func AddTransaccionRevisarAnteproyecto(m *TrRevisarAnteproyecto) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	
	// Update del trabajo de grado
	if num, err := o.Update(m.TrabajoGrado, "EstadoTrabajoGrado"); err == nil {
		fmt.Println("Number of degree work records updated:", num)
		// Insert de la revisión del anteproyecto
		if idRevisionTrabajoGrado, err := o.Insert(m.RevisionTrabajoGrado); err == nil {
			fmt.Println("Degree work review inserted:", idRevisionTrabajoGrado)
			// Insert de la corrección
			m.Correccion.RevisionTrabajoGrado.Id = int(idRevisionTrabajoGrado)
			if idCorreccion, err := o.Insert(m.Correccion); err == nil {
				fmt.Println("Correction inserted", idCorreccion)
			} else {
				fmt.Println(err)
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
				err = o.Rollback()
			}
		} else {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
			err = o.Rollback()
		}
	} else {
		fmt.Println(err)
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		err = o.Rollback()
	}

	// Si el anteproyecto es no viable, se cancela el trabajo de grado
	if m.TrabajoGrado.EstadoTrabajoGrado.Id == 2 {
		var estudiantesTrabajoGrado []EstudianteTrabajoGrado
		//Se buscan los estudiantes relacionados con el trabajo de grado y se cambia el estado
		if _, err := o.QueryTable(new(EstudianteTrabajoGrado)).RelatedSel().Filter("trabajo_grado",m.TrabajoGrado.Id).Filter("estado_estudiante_trabajo_grado", 1).All(&estudiantesTrabajoGrado); err == nil {
			for _, v := range estudiantesTrabajoGrado {
				v.EstadoEstudianteTrabajoGrado.Id = 3
				if _, err = o.Update(&v, "EstadoEstudianteTrabajoGrado"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
				}
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
		}
		//Se buscan las asignaturas del trabajo grado y se cambia el estado
		var asignaturasTrabajoGrado []AsignaturaTrabajoGrado
		if _, err := o.QueryTable(new(AsignaturaTrabajoGrado)).RelatedSel().Filter("trabajo_grado",m.TrabajoGrado.Id).Filter("estado_asignatura_trabajo_grado", 1).All(&asignaturasTrabajoGrado); err == nil {
			for _, v := range asignaturasTrabajoGrado {
				v.EstadoAsignaturaTrabajoGrado.Id = 3
				if _, err = o.Update(&v, "EstadoAsignaturaTrabajoGrado"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
				}
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
		}	
		//Se buscan las vinculaciones del trabajo de grado y se cambia el estado
		var vinculacionesTrabajoGrado []VinculacionTrabajoGrado
		if _, err := o.QueryTable(new(VinculacionTrabajoGrado)).RelatedSel().Filter("trabajo_grado",m.TrabajoGrado.Id).Filter("Activo", true).All(&vinculacionesTrabajoGrado); err == nil {
			for _, v := range vinculacionesTrabajoGrado {
				v.Activo = false
				if _, err = o.Update(&v, "Activo"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
				}
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
		}
		//Se buscan las áreas del trabajo de grado y se cambia el estado
		var areasTrabajoGrado []AreasTrabajoGrado
		if _, err := o.QueryTable(new(AreasTrabajoGrado)).RelatedSel().Filter("trabajo_grado",m.TrabajoGrado.Id).Filter("Activo", true).All(&areasTrabajoGrado); err == nil {
			for _, v := range areasTrabajoGrado {
				v.Activo = false
				if _, err = o.Update(&v, "Activo"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
				}
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.CANCELANDO_TRABAJO_GRADO")
		}
	}

	err = o.Commit()
	return
}
