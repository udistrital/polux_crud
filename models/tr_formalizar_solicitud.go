package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrFormalizarSolicitud struct {
	SolicitudesActualizadas *[]RespuestaSolicitud
}

//funcion para la transaccion de formalizar solicitudes
func AddTransaccionFormalizarSolicitud(m *TrFormalizarSolicitud) (alerta []string, err error) {
	fmt.Println("Here is inside")
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	fmt.Println("Here is after orm")

	// Update respuesta solicitud
	for _, v := range *m.SolicitudesActualizadas {
		fmt.Println("Here is for")
		if _, err = o.Update(&v); err != nil {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		}
	}
	return
}
