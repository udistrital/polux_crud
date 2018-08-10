package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrActualizarDocumentoTg struct {
	DocumentoEscrito						*DocumentoEscrito
	DocumentoTrabajoGrado				*DocumentoTrabajoGrado
	TrabajoGrado           			*TrabajoGrado
}

// Función para la transaccion de revisiones de anteproyectos
func AddTransaccionActualizarDocumentoTg(m *TrActualizarDocumentoTg) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	
	if (m.TrabajoGrado.EstadoTrabajoGrado.Id == 4 || m.TrabajoGrado.EstadoTrabajoGrado.Id == 15) {
		fmt.Println("Degree work state (4 or 15):", m.TrabajoGrado.EstadoTrabajoGrado.Id)
		// Update del documento del trabajo de grado
		if num, err := o.Update(m.DocumentoEscrito); err == nil {
			fmt.Println("Number of degree work documents updated:", num)
			// Update del trabajo de grado
			if num, err := o.Update(m.TrabajoGrado); err == nil {
				fmt.Println("Number of degree work records updated:", num)
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
	} else if (m.TrabajoGrado.EstadoTrabajoGrado.Id == 13 || m.TrabajoGrado.EstadoTrabajoGrado.Id == 22) {
		fmt.Println("Degree work state (13):", m.TrabajoGrado.EstadoTrabajoGrado.Id)
		// Insert del documento escrito
		m.DocumentoEscrito.Id = 0
		if idDocumentoEscrito, err := o.Insert(m.DocumentoEscrito); err == nil {
			fmt.Println("Written document inserted:", idDocumentoEscrito)
			m.DocumentoTrabajoGrado.Id = 0
			m.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(idDocumentoEscrito)
			m.DocumentoTrabajoGrado.TrabajoGrado.Id = int(m.TrabajoGrado.Id)
			if idDocumentoTg, err := o.Insert(m.DocumentoTrabajoGrado); err == nil {
				fmt.Println("Degree work document created:", idDocumentoTg)// Update del trabajo de grado
				if num, err := o.Update(m.TrabajoGrado); err == nil {
					fmt.Println("Number of degree work records updated:", num)
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
	}

	err = o.Commit()
	return
}
