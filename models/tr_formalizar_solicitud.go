package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrFormalizarSolicitud struct {
	RespuestasSolicitudesPrevias *[]RespuestaSolicitud
	RespuestasSolicitudesFormalizadas *[]RespuestaSolicitud
}

//funcion para la transaccion de formalizar solicitudes
func AddTransaccionFormalizarSolicitud(m *TrFormalizarSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")

	// Update respuesta solicitud
	for _, v := range *m.RespuestasSolicitudesPrevias {
		if _, err = o.Update(&v); err != nil {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		}
	}

	// Insert respuesta solicitud
	for _, v := range *m.RespuestasSolicitudesFormalizadas {
		if _, err = o.Insert(&v); err != nil {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_SOLICITUDES_1")
		}
	}
	
	err = o.Commit()
	return
}
