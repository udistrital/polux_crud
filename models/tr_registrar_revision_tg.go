package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarRevisionTg struct {
	Correcciones							*[]Correccion
	RevisionTrabajoGrado			*RevisionTrabajoGrado
}

// Función para la transaccion de revisiones de anteproyectos
func AddTransaccionRegistrarRevisionTg(m *TrRegistrarRevisionTg) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	
	// Update de la revisión del trabajo de grado
	if num, err := o.Update(m.RevisionTrabajoGrado, "EstadoRevisionTrabajoGrado", "FechaRevision"); err == nil {
		fmt.Println("Number of reviews updated:", num)
		for _, v := range *m.Correcciones {
			if num, err := o.Insert(&v); err == nil {
				fmt.Println("Number of records inserted in database:", num)
				fmt.Println(&v)
			} else {
				fmt.Println(err)
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR.REGISTRANDO_CORRECCIONES")
				err = o.Rollback()
			}
		}
		err = o.Commit()
	} else {
		fmt.Println(err)
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR.REGISTRANDO_REVISION")
		err = o.Rollback()
	}
	return
}
