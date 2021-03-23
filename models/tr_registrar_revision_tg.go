package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarRevisionTg struct {
	Comentarios						*[]Comentario
	RevisionTrabajoGrado			*RevisionTrabajoGrado
}

// Función para la transaccion de revisiones de anteproyectos
func AddTransaccionRegistrarRevisionTg(m *TrRegistrarRevisionTg) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")
	
	// Update de la revisión del trabajo de grado
	if num, err := o.Update(m.RevisionTrabajoGrado, "EstadoRevisionTrabajoGrado", "FechaRevision"); err == nil {
		fmt.Println("Number of reviews updated:", num)
		// Insert de las correcciones y comentarios
		for _, v := range *m.Comentarios {
			if idCorreccion, err := o.Insert(v.Correccion); err == nil {
				v.Correccion.Id = int(idCorreccion)
				fmt.Println("Correction inserted", idCorreccion)
				if idComentario, err := o.Insert(&v); err == nil {
					fmt.Println("Comentario inserted", idComentario)
				} else {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR.INSERTANDO_REVISIONES")
			}
			} else {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR.INSERTANDO_REVISIONES")
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
