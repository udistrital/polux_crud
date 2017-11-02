package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRespuestaSolicitud struct {
	RespuestaAnterior  *RespuestaSolicitud
	RespuestaNueva     *RespuestaSolicitud
	DocumentoSolicitud *DocumentoSolicitud
	TipoSolicitud      *TipoSolicitud
	Vinculaciones      *[]VinculacionTrabajoGrado
}

//funcion para la transaccion de solicitudes
func AddTransaccionRespuestaSolicitud(m *TrRespuestaSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")

	var num int64
	//update del estado de la ultima solicitud
	if num, err = o.Update(m.RespuestaAnterior); err == nil {
		fmt.Println("Number of records updated in database:", num)
		fmt.Println(m.RespuestaNueva)
		//insert de la nueva rta
		if id_rta, err := o.Insert(m.RespuestaNueva); err == nil {
			fmt.Println(id_rta)

			//insert documento asociado a la nueva rta
			m.DocumentoSolicitud.SolicitudTrabajoGrado.Id = int(id_rta)
			if id_documento, err := o.Insert(m.DocumentoSolicitud); err == nil {
				fmt.Println(id_documento)

				//seguir la transacción según el caso: de acuerdo al tipo de solicitud
				fmt.Println(m.TipoSolicitud.Nombre)
				fmt.Println(m.TipoSolicitud.Id)
				if m.TipoSolicitud.Id == 4 { //solicitud de cambio de director interno
					for _, v := range *m.Vinculaciones {
						fmt.Println(v)
						if v.Activo == true {
							if _, err = o.Insert(&v); err != nil {
								fmt.Println(err)
								err = o.Rollback()
								alerta[0] = "Error"
								alerta = append(alerta, "ERROR_SOLICITUDES_1")
							}
						} else {
							if _, err = o.Update(&v); err != nil {
								fmt.Println(err)
								err = o.Rollback()
								alerta[0] = "Error"
								alerta = append(alerta, "ERROR_SOLICITUDES_1")
							}
						}

					}
				}

				err = o.Commit()
			} else {
				fmt.Println(err)
				err = o.Rollback()
			}

		} else {
			fmt.Println(err)
			err = o.Rollback()
		}

	} else {
		fmt.Println(err)
		err = o.Rollback()
	}
	return
}
