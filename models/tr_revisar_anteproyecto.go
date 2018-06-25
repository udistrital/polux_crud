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
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		err = o.Rollback()
	}
	return
}
