package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrSeleccionAdmitidos struct {
	RespuestasNuevas         *[]RespuestaSolicitud
	RespuestasAntiguas       *[]RespuestaSolicitud
}

//funcion para la transaccion de solicitudes
func TransaccionSeleccionAdmitidos(m *TrSeleccionAdmitidos) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	for _, v := range *m.RespuestasAntiguas{
		if num, err := o.Update(&v, "Activo"); err == nil {
			fmt.Println("Number of records updated in database:", num)
			fmt.Println(&v)
		}else{
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
			err = o.Rollback()
		}
	}
	for _, v := range *m.RespuestasNuevas{
		if _, err = o.Insert(&v); err != nil {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
			err = o.Rollback()
		}
	}
	err = o.Commit()
	return
}
