package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarNota struct {
	EspaciosAcademicosCalificados *[]EspacioAcademicoInscrito
	AsignaturasDeTrabajoDeGrado   *[]AsignaturaTrabajoGrado
	TrabajoDeGradoTerminado       *TrabajoGrado
	EvaluacionTrabajoGrado        *EvaluacionTrabajoGrado
}

// AddTransaccionRegistrarNota Función para la transaccion de notas obtenidas en espacios académicos de posgrado
func AddTransaccionRegistrarNota(m *TrRegistrarNota) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")

	// Update del trabajo de grado

	if num, err := o.Update(m.TrabajoDeGradoTerminado); err == nil {
		fmt.Println("Number of degree assigments updated in database:", num)
		// Update de los espacios académicos inscritos
		for _, espacioAcademicoCalificado := range *m.EspaciosAcademicosCalificados {
			if num, err := o.Update(&espacioAcademicoCalificado); err == nil {
				fmt.Println("Number of academic spaces updated in database:", num)
				fmt.Println(&espacioAcademicoCalificado)
			} else {
				fmt.Println(err)
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
				err = o.Rollback()
			}
		}
		// Update de las asignaturas de trabajo de grado
		for _, asignaturaTrabajoGrado := range *m.AsignaturasDeTrabajoDeGrado {
			if num, err := o.Update(&asignaturaTrabajoGrado); err == nil {
				fmt.Println("Number of degree assigment subjects updated in database:", num)
				fmt.Println(&asignaturaTrabajoGrado)
			} else {
				fmt.Println(err)
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_3")
				err = o.Rollback()
			}
		}
		// Insert de evaluación trabajo grado
		fmt.Println(m.EvaluacionTrabajoGrado)
		if _, err = o.Insert(m.EvaluacionTrabajoGrado); err != nil {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_4")
		}

		err = o.Commit()
	} else {
		fmt.Println(err)
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		err = o.Rollback()
	}
	return
}
