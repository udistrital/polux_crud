package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrActualizarDocumentoTg struct {
	DocumentoEscrito							*DocumentoEscrito
	DocumentoTrabajoGrado					*DocumentoTrabajoGrado
	TrabajoGrado           				*TrabajoGrado
}

// Función para la transaccion de revisiones de anteproyectos
func AddTransaccionActualizarDocumentoTg(m *TrActualizarDocumentoTg) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	
	// Insert del documento escrito
	if idDocumentoEscrito, err := o.Insert(m.DocumentoEscrito); err == nil {
		fmt.Println("Written document inserted:", idDocumentoEscrito)
		// Update del documento del trabajo de grado
		m.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(idDocumentoEscrito)
		if num, err := o.Update(m.DocumentoTrabajoGrado, "DocumentoEscrito"); err == nil {
			fmt.Println("Number of degree work documents updated:", num)
			// Update del trabajo de grado
			if num, err := o.Update(m.TrabajoGrado, "EstadoTrabajoGrado"); err == nil {
				fmt.Println("Number of degree work records updated:", num)
				err = o.Commit()
			} else {
				fmt.Println(err)
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
				err = o.Rollback()
			}
		} else {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
			err = o.Rollback()
		}
	} else {
		fmt.Println(err)
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
		err = o.Rollback()
	}
	return
}
