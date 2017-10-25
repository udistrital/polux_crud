package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRespuestaSolicitud struct {
	RespuestaSolicitud *RespuestaSolicitud
	DocumentoEscrito   *DocumentoEscrito
}

//funcion para la transaccion de solicitudes
func AddTransaccionRespuestaSolicitud(m *TrRespuestaSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	if id, err := o.Insert(m.RespuestaSolicitud); err == nil {
		fmt.Println(id)

		if id_documento, err := o.Insert(m.DocumentoEscrito); err == nil {
			fmt.Println(id_documento)
			/*
				m.DocumentoTrabajoGrado.TrabajoGrado.Id = int(id)
				m.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(id_documento)

				if id_documento, err := o.Insert(m.DocumentoTrabajoGrado); err == nil {
					fmt.Println(id_documento)
				}*/
		}

		//err = o.Commit()
	} else {
		fmt.Println(err)
		err = o.Rollback()
	}
	return
}
