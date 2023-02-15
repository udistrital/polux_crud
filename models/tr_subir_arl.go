package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrSubirArl struct {
	DocumentoEscrito      *DocumentoEscrito
	DocumentoTrabajoGrado *DocumentoTrabajoGrado
}

func AddTransaccionArl(m *TrSubirArl) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")
	if id, err := o.Insert(m.DocumentoEscrito); err == nil {

		m.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(id)
		if _, err := o.Insert(m.DocumentoTrabajoGrado); err == nil {
			err = o.Commit()
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_SUBIR_DOCUMENTO_TR_GRADO")
		}

	} else {
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_SUBIR_ARL")
		fmt.Println(err)
		err = o.Rollback()
	}

	return
}
