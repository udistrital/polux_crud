package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarNota struct {
	EspaciosAcademicosCalificados		 *[]EspacioAcademicoInscrito
}

// AddTransaccionRegistrarNota Función para la transaccion de notas obtenidas en espacios académicos de posgrado
func AddTransaccionRegistrarNota(m *TrRegistrarNota) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")

	// Update de las respuestas antiguas
	for _, v := range *m.EspaciosAcademicosCalificados {
		if num, err := o.Update(&v, "Nota"); err == nil {
			fmt.Println("Number of records updated in database:", num)
			fmt.Println(&v)
		} else {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
			err = o.Rollback()
		}
	}

	err = o.Commit()
	return
}
