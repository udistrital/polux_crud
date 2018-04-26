package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarMateriasPosgrado struct {
	RespuestaPrevia        *RespuestaSolicitud
	RespuestaActualizada   *RespuestaSolicitud
	TrabajoGrado           *TrabajoGrado
	EstudianteTrabajoGrado *EstudianteTrabajoGrado
	EspaciosAcademicos		 *[]EspacioAcademicoInscrito
}

// Función para la transaccion de solicitudes de materias de posgrado
func AddTransaccionRegistrarMateriasPosgrado(m *TrRegistrarMateriasPosgrado) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	
	if num, err := o.Update(m.RespuestaPrevia, "Activo"); err == nil {
		fmt.Println("Number of records updated in database:", num)
		// Insert nueva respuesta
		if idRespuestaNueva, err := o.Insert(m.RespuestaActualizada); err == nil {
			fmt.Println(idRespuestaNueva)
			// Insert trabajo grado
			if idTrabajoGrado, err := o.Insert(m.TrabajoGrado); err == nil {
				fmt.Println(idTrabajoGrado)
				// Insert estudiante trabajo grado
				m.EstudianteTrabajoGrado.TrabajoGrado.Id = int(idTrabajoGrado)
				if idEstudianteTrabajoGrado, err := o.Insert(m.EstudianteTrabajoGrado); err == nil {
					fmt.Println(idEstudianteTrabajoGrado)
					// Insert espacios academicos inscritos
					for _, v := range *m.EspaciosAcademicos {
						v.TrabajoGrado.Id = int(idTrabajoGrado)
						if _, err = o.Insert(&v); err != nil {
							fmt.Println(err)
							err = o.Rollback()
							alerta[0] = "Error"
							alerta = append(alerta, "ERROR_SOLICITUDES_1")
						}
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