package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrSolicitud struct {
	Solicitud         *SolicitudTrabajoGrado
	Respuesta         *RespuestaSolicitud
	DetallesSolicitud *[]DetalleSolicitud
	UsuariosSolicitud *[]UsuarioSolicitud
}

// AddTransaccionSolicitud funcion para registrar solicitudes (detalles y usuarios)
func AddTransaccionSolicitud(m *TrSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	if id, err := o.Insert(m.Solicitud); err == nil {
		fmt.Println(id)
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
		err = o.Commit()
	} else {
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_CARGAR_SOLICITUDES")
		fmt.Println(err)
		err = o.Rollback()
	}
	return
}
