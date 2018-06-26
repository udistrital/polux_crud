package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrTrabajoGrado struct {
	TrabajoGrado            *TrabajoGrado
	EstudianteTrabajoGrado  *[]EstudianteTrabajoGrado
	DocumentoEscrito        *DocumentoEscrito
	DocumentoTrabajoGrado   *DocumentoTrabajoGrado
	AreasTrabajoGrado       *[]AreasTrabajoGrado
	VinculacionTrabajoGrado *[]VinculacionTrabajoGrado
	AsignaturasTrabajoGrado *[]AsignaturaTrabajoGrado
}

// AddTransaccionTrabajoGrado funcion para crear un trabajo de grado desde una solicitud
func AddTransaccionTrabajoGrado(m *TrTrabajoGrado) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")
	if id, err := o.Insert(m.TrabajoGrado); err == nil {
		fmt.Println(id)

		for _, v := range *m.EstudianteTrabajoGrado {
			v.TrabajoGrado.Id = int(id)
			if _, err = o.Insert(&v); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_SOLICITUDES_1")
			}
		}

		for _, v := range *m.AreasTrabajoGrado {
			v.TrabajoGrado.Id = int(id)
			if _, err = o.Insert(&v); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_SOLICITUDES_2")
			}
		}

		for _, v := range *m.VinculacionTrabajoGrado {
			v.TrabajoGrado.Id = int(id)
			if _, err = o.Insert(&v); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_SOLICITUDES_3")
			}
		}

		if id_documento, err := o.Insert(m.DocumentoEscrito); err == nil {
			fmt.Println(id_documento)
			m.DocumentoTrabajoGrado.TrabajoGrado.Id = int(id)
			m.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(id_documento)

			if id_documento, err := o.Insert(m.DocumentoTrabajoGrado); err == nil {
				fmt.Println(id_documento)
			}
		}

		err = o.Commit()
	} else {
		fmt.Println(err)
		err = o.Rollback()
	}
	return
}
