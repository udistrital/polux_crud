package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrVinculadoRegistrarNota struct {
	TrabajoGrado            *TrabajoGrado
	DocumentoEscrito        *DocumentoEscrito
	VinculacionTrabajoGrado *VinculacionTrabajoGrado
	EvaluacionTrabajoGrado  *EvaluacionTrabajoGrado
}

// AddTransaccionVinculadoRegistrarNota función para registrar la nota de un trabajo de grado
func AddTransaccionVinculadoRegistrarNota(m *TrVinculadoRegistrarNota) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	alerta = append(alerta, "Success")
	fmt.Println("Transaccion para registrar la nota")
	//Se inserta la evaluación
	if _, err = o.Insert(m.EvaluacionTrabajoGrado); err != nil {
		fmt.Println(err)
		err = o.Rollback()
		alerta[0] = "Error"
		alerta = append(alerta, "REGISTRAR_NOTA.ERROR.REGISTRANDO_NOTA")
	}
	//Se inserta el documento escrito si no es nulo
	if m.DocumentoEscrito != nil {
		if idDocumentoEscrito, err := o.Insert(m.DocumentoEscrito); err == nil {
			m.DocumentoEscrito.Id = int(idDocumentoEscrito)
			//Se crea el documento trabajo de grado
			documentoTrabajoGrado := DocumentoTrabajoGrado{
				Id:               0,
				DocumentoEscrito: m.DocumentoEscrito,
				TrabajoGrado:     m.TrabajoGrado,
			}
			//SE inserta el documento trabajo de grado
			if _, err = o.Insert(&documentoTrabajoGrado); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "REGISTRAR_NOTA.ERROR.REGISTRANDO_ACTA_SUSTENTACION")
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "REGISTRAR_NOTA.ERROR.REGISTRANDO_ACTA_SUSTENTACION")
		}
	}

	var actualizarNotasTg bool
	var promedio float64
	var notaDirector float64
	/*
		if(m.VinculacionTrabajoGrado.RolTrabajoGrado.Id == 1){
			//Si es director interno se registra directalemte
			actualizarNotasTg = true
			promedio = m.EvaluacionTrabajoGrado.Nota
		} else if (m.VinculacionTrabajoGrado.RolTrabajoGrado.Id == 3){
	*/
	//Se busca la cantidad de personas vinculadas (Director interno y evaluadores) y se registra la nota para cada vinculado
	var vinculacionesTG []VinculacionTrabajoGrado
	if _, err := o.QueryTable(new(VinculacionTrabajoGrado)).RelatedSel().Filter("rol_trabajo_grado__in", 1, 3).Filter("Activo", true).Filter("trabajo_grado", m.TrabajoGrado.Id).All(&vinculacionesTG); err == nil {
		if len(vinculacionesTG) == 1 {
			//Si la cantidad de evaluadores registrados es 1 entonces se actualiza la nota
			actualizarNotasTg = true
			promedio = m.EvaluacionTrabajoGrado.Nota
			notaDirector = m.EvaluacionTrabajoGrado.Nota
		} else {
			//Si la cantidad de evaluadores registrados es mayor a 1 entonces se busca cuantos han registrado nota
			var notasRegistradas []EvaluacionTrabajoGrado
			if _, err := o.QueryTable(new(EvaluacionTrabajoGrado)).RelatedSel().Filter("vinculacion_trabajo_grado__in", vinculacionesTG).All(&notasRegistradas); err == nil {
				fmt.Println("Notas registradas", len(notasRegistradas))
				fmt.Println("vinculacionnes", len(vinculacionesTG))
				//Si hay algunar egistrada
				//Si no son tno se actualiza la nota
				if len(notasRegistradas) != len(vinculacionesTG) {
					actualizarNotasTg = false
				}
				//Si son todos se saca el promedio y se registra la nota
				if len(notasRegistradas) == len(vinculacionesTG) {
					fmt.Println("Actualiza notas")
					//Se busca el promedio
					var promedioTem float64
					promedioTem = 0
					for _, v := range notasRegistradas {
						promedioTem += v.Nota
						if v.VinculacionTrabajoGrado.RolTrabajoGrado.CodigoAbreviacion == "DIRECTOR" {
							notaDirector = v.Nota
						}
					}
					promedio = promedioTem / float64(len(notasRegistradas))
					fmt.Println("promedio total docentes", promedio)
					actualizarNotasTg = true
				}
			} else {
				//Error
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "REGISTRAR_NOTA.ERROR.CONSULTANDO_VINCULACIONES")
			}
		}
	} else {
		//Error
		fmt.Println(err)
		err = o.Rollback()
		alerta[0] = "Error"
		alerta = append(alerta, "REGISTRAR_NOTA.ERROR.CONSULTANDO_VINCULACIONES")
	}
	//}

	//Se actualizan las notas de TG teniendo en cuenta el número de evaluadores registrados y el tipo de vinculación
	if actualizarNotasTg {
		//Se buscan las notas del tg
		fmt.Println("Promedio a registrar: ", promedio)
		var asignaturaTrabajoGrado []AsignaturaTrabajoGrado
		if _, err := o.QueryTable(new(AsignaturaTrabajoGrado)).RelatedSel().Filter("trabajo_grado", m.TrabajoGrado.Id).All(&asignaturaTrabajoGrado); err == nil {
			for _, v := range asignaturaTrabajoGrado {
				//Se actualiza promedio
				if v.CodigoAsignatura == 1 {
					v.Calificacion = notaDirector
				} else {
					v.Calificacion = promedio
				}
				//Se actualiza estado a cursado
				v.EstadoAsignaturaTrabajoGrado.Id = 2
				if num, err := o.Update(&v, "Calificacion", "EstadoAsignaturaTrabajoGrado"); err == nil {
					fmt.Println("actualiza calificación TrabajoGrado")
					fmt.Println("Number of records updated in database:", num)
					//finaliza update
				} else {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "REGISTRAR_NOTA.ERROR.REGISTRANDO_NOTA")
				}
			}
			//Se actualiza el estado del trabajo de grado
			// a 19 Notificado a coordinación con calificación
			m.TrabajoGrado.EstadoTrabajoGrado.Id = 19
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "REGISTRAR_NOTA.ERROR.REGISTRANDO_NOTA")
		}
	} else {
		//Se actualiza el estado del trabajo de grado a 18 Sustentado
		m.TrabajoGrado.EstadoTrabajoGrado.Id = 18
		fmt.Println("No se actualizarNotasTg")
	}
	//Se actualiza el estado del tg
	if num, err := o.Update(m.TrabajoGrado, "EstadoTrabajoGrado"); err == nil {
		fmt.Println("Number of records updated in database:", num)
		//finaliza update
	} else {
		fmt.Println(err)
		err = o.Rollback()
		alerta[0] = "Error"
		alerta = append(alerta, "REGISTRAR_NOTA.ERROR.REGISTRANDO_NOTA")
	}
	err = o.Commit()
	//err = o.Rollback()
	return
}
