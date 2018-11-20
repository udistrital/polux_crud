package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarMateriasPosgrado struct {
	RespuestaPrevia        				*RespuestaSolicitud
	RespuestaActualizada   				*RespuestaSolicitud
	TrabajoGrado           				*TrabajoGrado
	EstudianteTrabajoGrado 				*EstudianteTrabajoGrado
	EspaciosAcademicos		 				*[]EspacioAcademicoInscrito
	AsignaturasDeTrabajoDeGrado 	*[]AsignaturaTrabajoGrado
}

// Función para la transaccion de solicitudes de materias de posgrado
func AddTransaccionRegistrarMateriasPosgrado(m *TrRegistrarMateriasPosgrado) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")
	
	// Update de la respuesta previa
	if num, err := o.Update(m.RespuestaPrevia, "Activo"); err == nil {
		fmt.Println("Number of requests updated in database:", num)
		// Insert de la nueva respuesta
		if idRespuestaNueva, err := o.Insert(m.RespuestaActualizada); err == nil {
			fmt.Println("Answered request inserted:", idRespuestaNueva)
			// Insert del trabajo grado
			if idTrabajoGrado, err := o.Insert(m.TrabajoGrado); err == nil {
				fmt.Println("Degree assignment inserted", idTrabajoGrado)
				// Insert del estudiante trabajo grado
				m.EstudianteTrabajoGrado.TrabajoGrado.Id = int(idTrabajoGrado)
				if idEstudianteTrabajoGrado, err := o.Insert(m.EstudianteTrabajoGrado); err == nil {
					fmt.Println("Degree assignment student inserted:", idEstudianteTrabajoGrado)
					// Insert de espacios academicos inscritos
					for _, espacioAcademico := range *m.EspaciosAcademicos {
						espacioAcademico.TrabajoGrado.Id = int(idTrabajoGrado)
						if idEspacioAcademicoInscrito, err := o.Insert(&espacioAcademico); err == nil {
							fmt.Println("Academic space inserted:", idEspacioAcademicoInscrito)
						} else {
							fmt.Println(err)
							alerta[0] = "Error"
							alerta = append(alerta, "ERROR_SOLICITUDES_1")
							err = o.Rollback()
						}
					}
					// Insert de asignaturas trabajo grado
					for _, asignaturaTrabajoGrado := range *m.AsignaturasDeTrabajoDeGrado {
						asignaturaTrabajoGrado.TrabajoGrado.Id = int(idTrabajoGrado)
						if idAsignaturaTrabajoGrado, err := o.Insert(&asignaturaTrabajoGrado); err == nil {
							fmt.Println("Degree assignment subject inserted:", idAsignaturaTrabajoGrado)
						} else {
							fmt.Println(err)
							alerta[0] = "Error"
							alerta = append(alerta, "ERROR_SOLICITUDES_1")
							err = o.Rollback()
						}
					}
					// Actualización de la solicitud de trabajo de grado
					// la solicitud inicial queda relacionada al trabajo de grado
					m.RespuestaPrevia.SolicitudTrabajoGrado.TrabajoGrado.Id = int(idTrabajoGrado)
					if _, err := o.Update(m.RespuestaPrevia.SolicitudTrabajoGrado, "TrabajoGrado"); err == nil {
						fmt.Println("Solicitud de trabajo de grado actualizada")
					} else {
						fmt.Println(err)
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_SOLICITUDES_3")
						err = o.Rollback()
					}
					err = o.Commit()
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
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
			err = o.Rollback()
		}
	} else {
		fmt.Println(err)
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		err = o.Rollback();
	}
	return
}
