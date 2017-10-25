package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRespuestaSolicitud struct {
	TrabajoGrado            *TrabajoGrado
	EstudianteTrabajoGrado  *[]EstudianteTrabajoGrado
	DocumentoEscrito        *DocumentoEscrito
	DocumentoTrabajoGrado   *DocumentoTrabajoGrado
	AreasTrabajoGrado       *[]AreasTrabajoGrado
	VinculacionTrabajoGrado *[]VinculacionTrabajoGrado
}

//funcion para la transaccion de solicitudes
func AddTransaccionRespuestaSolicitud(m *TrRespuestaSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
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

		/*
			m.Respuesta.SolicitudTrabajoGrado.Id = int(id)
			if _, err = o.Insert(m.Respuesta); err != nil {
				fmt.Println(err)
				err = o.Rollback()

				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_SOLICITUDES_3")
			}
			for _, v := range *m.DetallesSolicitud {
				v.SolicitudTrabajoGrado.Id = int(id)
				if _, err = o.Insert(&v); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_SOLICITUDES_1")
				}
			}

			for _, j := range *m.UsuariosSolicitud {
				fmt.Println(err)
				j.SolicitudTrabajoGrado.Id = int(id)
				if _, err = o.Insert(&j); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_SOLICITUDES_2")
				}
			}

		*/
		err = o.Commit()
	} else {
		fmt.Println(err)
		err = o.Rollback()
	}
	return
}
