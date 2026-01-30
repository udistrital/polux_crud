package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRegistrarActaSeguimiento struct {
	Acta *DocumentoTrabajoGrado
}

// AddTransaccionRespuestaSolicitud funcion para dar respuesta a las solicitudes
func RegistrarActaSeguimiento(m *TrRegistrarActaSeguimiento) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")
	//Se inserta el documento escrito
	fmt.Println(m.Acta)
	if idDocumentoEscrito, err := o.Insert(m.Acta.DocumentoEscrito); err == nil {
		m.Acta.DocumentoEscrito.Id = int(idDocumentoEscrito)
		//Se inserta el documento trabajo de grado
		if _, err := o.Insert(m.Acta); err == nil {
			err = o.Commit()
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "PASANTIA.ERROR.CARGANDO_ACTAS_SEGUIMIENTO")
		}
	} else {
		fmt.Println(err)
		err = o.Rollback()
		alerta[0] = "Error"
		alerta = append(alerta, "PASANTIA.ERROR.CARGANDO_ACTAS_SEGUIMIENTO")
	}

	return
}
