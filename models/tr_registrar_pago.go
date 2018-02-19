package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarPago struct {
	RespuestaAnterior      *RespuestaSolicitud
	RespuestaNueva         *RespuestaSolicitud
}

//funcion para la transaccion de solicitudes
func AddTransaccionRegistrarPago(m *TrRegistrarPago) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	fmt.Println(m.RespuestaNueva.Activo)
	alerta = append(alerta, "Success")

	if num, err := o.Update(m.RespuestaAnterior, "Activo"); err == nil {
		fmt.Println("Number of records updated in database:", num)

		// Insert nueva respuesta
		if id, err := o.Insert(m.RespuestaNueva); err == nil {
			fmt.Println(id)
			err = o.Commit()
		} else {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
			err = o.Rollback()
		}

	} else {
		fmt.Println("Error")
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		err = o.Rollback();
	}
	return
}
