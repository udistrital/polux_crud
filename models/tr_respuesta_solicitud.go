package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrRespuestaSolicitud struct {
	RespuestaAnterior      *RespuestaSolicitud
	RespuestaNueva         *RespuestaSolicitud
	DocumentoSolicitud     *DocumentoSolicitud
	TipoSolicitud          *TipoSolicitud
	Vinculaciones          *[]VinculacionTrabajoGrado //Cambio de director o evaluador..//Cancelación TG
	EstudianteTrabajoGrado *EstudianteTrabajoGrado    //Cancelación trabajo grado
	VinculacionesCancelacion *[]VinculacionTrabajoGrado //Vinculaciones para cancelacion de trabajo de grado
	TrTrabajoGrado         *TrTrabajoGrado            //Solictudes iniciales
	ModalidadTipoSolicitud *ModalidadTipoSolicitud    //Para saber el tipo de solicitud inicial
	TrabajoGrado           *TrabajoGrado              //Cambio Titulo
	SolicitudTrabajoGrado  *SolicitudTrabajoGrado     //solicitud inicial
	EspaciosAcademicos     *[]EspacioAcademicoInscrito //Solicitud de cambio de asignaturas
}

// AddTransaccionRespuestaSolicitud funcion para dar respuesta a las solicitudes
func AddTransaccionRespuestaSolicitud(m *TrRespuestaSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	var num int64

	//actualizar estados de solicitud
	if num, err = o.Update(m.RespuestaAnterior, "Activo"); err == nil {
		fmt.Println("Number of records updated in database:", num)
		//insert de la nueva rta
		if id_respuesta, err := o.Insert(m.RespuestaNueva); err == nil {
			fmt.Println(id_respuesta)
			//insert documento asociado a la nueva rta
			if id_documento, err := o.Insert(m.DocumentoSolicitud); err == nil {
				fmt.Println(id_documento)
			} else {
				fmt.Println(err)
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_3")
				err = o.Rollback()
			}
		} else {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
			err = o.Rollback()
		}
	} else {
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		fmt.Println(err)
		err = o.Rollback()
	}

	//solicitud inicial, se crea trabajo de grado
	if m.TrTrabajoGrado != nil {
		if id_TrabajoGrado, err := o.Insert(m.TrTrabajoGrado.TrabajoGrado); err == nil {
			fmt.Println(id_TrabajoGrado)
			//la solicitud inicial queda relacionada al trabajo de grado
			m.SolicitudTrabajoGrado.TrabajoGrado.Id = int(id_TrabajoGrado)
			if id_sols, err := o.Update(m.SolicitudTrabajoGrado, "TrabajoGrado"); err == nil {
				fmt.Println(id_sols)
				// Se agregan asignaturas de trabajo de grado
				for _, v := range *m.TrTrabajoGrado.AsignaturasTrabajoGrado {
					v.TrabajoGrado.Id = int(id_TrabajoGrado)
					if _, err = o.Insert(&v); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_14")
					}
				}
				// Se agregan estudiantes
				for _, v := range *m.TrTrabajoGrado.EstudianteTrabajoGrado {
					v.TrabajoGrado.Id = int(id_TrabajoGrado)
					if _, err = o.Insert(&v); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_9")
					}
				}
				// Se agregan areas de conocimiento
				for _, v := range *m.TrTrabajoGrado.AreasTrabajoGrado {
					v.TrabajoGrado.Id = int(id_TrabajoGrado)
					if _, err = o.Insert(&v); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_10")
					}
				}
				// Se agregan vinculaciones
				for _, v := range *m.TrTrabajoGrado.VinculacionTrabajoGrado {
					v.TrabajoGrado.Id = int(id_TrabajoGrado)
					if _, err = o.Insert(&v); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
					}
				}
				// Se agrega documento Estrito
				if id_documento, err := o.Insert(m.TrTrabajoGrado.DocumentoEscrito); err == nil {
					fmt.Println(id_documento)
					m.TrTrabajoGrado.DocumentoTrabajoGrado.TrabajoGrado.Id = int(id_TrabajoGrado)
					m.TrTrabajoGrado.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(id_documento)
					if id_documento, err := o.Insert(m.TrTrabajoGrado.DocumentoTrabajoGrado); err == nil {
						fmt.Println(id_documento)
						//Se terminan de insertar
					} else {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_12")
					}
				} else {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_11")
				}
			} else {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_333")
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_4")
		}
	}

	//Solicitud de cambio de docente evaluador o docente director
	if m.Vinculaciones != nil {
		fmt.Println(m.Vinculaciones);
		for _, v := range *m.Vinculaciones {
			//Si esta activo es nuevo y se inserta sino se actualiza la fecha de fin y el activo
			if v.Activo == true {
				if _, err = o.Insert(&v); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
				}
			} else {
				if _, err = o.Update(&v, "Activo", "FechaFin"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_6")
				}
			}

		}
	}

	//Solicitud de cambio de nombre del trabajo de grado
	if m.TrabajoGrado != nil {
		//Actualiza el nombre del tg
		if _, err = o.Update(m.TrabajoGrado,"Titulo"); err != nil {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_13")
		}
	}

	//Solicitud de cancelacion de modalidad
	if m.EstudianteTrabajoGrado != nil {
		var estudianteTrabajoGrado EstudianteTrabajoGrado
		//se busca al estudiante con el trabajo de grado
		if err := o.QueryTable(new(EstudianteTrabajoGrado)).Filter("trabajo_grado",m.EstudianteTrabajoGrado.TrabajoGrado.Id).Filter("estudiante", m.EstudianteTrabajoGrado.Estudiante).Filter("estado_estudiante_trabajo_grado", 1).One(&estudianteTrabajoGrado); err == nil{
			m.EstudianteTrabajoGrado.Id = estudianteTrabajoGrado.Id
			if num, err = o.Update(m.EstudianteTrabajoGrado); err == nil {
				fmt.Println("Number of records updated in database:", num)
				//consultar # estudiantes del trabajo de grado que estan activos
				cnt, _ := o.QueryTable("estudiante_trabajo_grado").Filter("trabajo_grado", m.EstudianteTrabajoGrado.TrabajoGrado.Id).Filter("estado_estudiante_trabajo_grado", 1).Count()
				fmt.Println("Count estudiantes activos tg:", cnt)
				// si no hay estudiantes vinculados se inactivan las vinculaciones y se cancela el tg
				if cnt == 0 {
					//se inactivan las vinculaciones
					for _, v := range *m.VinculacionesCancelacion {
						v.Activo = false
						if _, err = o.Update(&v, "Activo", "FechaFin"); err != nil {
							fmt.Println(err)
							err = o.Rollback()
							alerta[0] = "Error"
							alerta = append(alerta, "ERROR_RTA_SOLICITUD_14")
						}
					}
					//se cancela el trabajo de grado
					tg := m.EstudianteTrabajoGrado.TrabajoGrado
					tg.EstadoTrabajoGrado.Id = 2
					if num, err = o.Update(tg); err == nil {
						fmt.Println("Number of records updated in database:", num)
						//finalizan los inserts
					} else {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_8")
					}
				} 
			} else {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_7")
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_7")
		}

		// actualziar Asignaturas trabajo de grado a cancelado
		var asignaturasTrabajoGrado []AsignaturaTrabajoGrado
		//se busca asingaturas trabajo grado
		if _, err := o.QueryTable(new(AsignaturaTrabajoGrado)).RelatedSel().Filter("trabajo_grado",m.EstudianteTrabajoGrado.TrabajoGrado.Id).All(&asignaturasTrabajoGrado); err == nil{
			fmt.Println("asignaturasTrabajoGrado",asignaturasTrabajoGrado);
			for _, v := range asignaturasTrabajoGrado {
				//Id de la asignatura 3 o cancelado
				v.EstadoAsignaturaTrabajoGrado.Id = 3
				if _, err = o.Update(&v, "EstadoAsignaturaTrabajoGrado"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_8")
				}
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_7")
		}

		//Actualizar espacios_academicos_inscritos
		var espaciosAcademicosInscritos []EspacioAcademicoInscrito
		//se buscan espacios academicos inscritos activos
		if _, err := o.QueryTable(new(EspacioAcademicoInscrito)).RelatedSel().Filter("trabajo_grado",m.EstudianteTrabajoGrado.TrabajoGrado.Id).Filter("estado_espacio_academico_inscrito",1).All(&espaciosAcademicosInscritos); err == nil{
			fmt.Println("espaciosAcademicosInscritos",espaciosAcademicosInscritos);
			for _, v := range espaciosAcademicosInscritos {
				//Id del espacio 2  cancelado
				v.EstadoEspacioAcademicoInscrito.Id = 2
				if _, err = o.Update(&v, "EstadoEspacioAcademicoInscrito"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_8")
				}
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_7")
		}
	}

	//Solicitud de cambio de materias
	if m.EspaciosAcademicos != nil {
		var espacioActual EspacioAcademicoInscrito
		var espacioNuevo EspacioAcademicoInscrito
		for _, v := range *m.EspaciosAcademicos {
			fmt.Println(v);
			if v.EstadoEspacioAcademicoInscrito.Id == 1 {
				espacioNuevo = v
			} else if v.EstadoEspacioAcademicoInscrito.Id == 2 {
				//espacio que se cancela 
				if err = o.QueryTable(new(EspacioAcademicoInscrito)).RelatedSel().Filter("TrabajoGrado",v.TrabajoGrado).Filter("EspaciosAcademicosElegibles__CodigoAsignatura",v.EspaciosAcademicosElegibles.CodigoAsignatura).One(&espacioActual); err == nil {
					espacioActual.EstadoEspacioAcademicoInscrito = v.EstadoEspacioAcademicoInscrito
					if num, err = o.Update(&espacioActual,"EstadoEspacioAcademicoInscrito"); err == nil {
						fmt.Println("Number of records updated in database:", num)
						//finaliza update
					} else {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_8")
					}
				} else {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_13")
				}
			}
		}
		// Buscar id de asignatura del espacio nuevo
		var espacioElegibleNuevo EspaciosAcademicosElegibles
		if err = o.QueryTable(new(EspaciosAcademicosElegibles)).RelatedSel().Filter("CarreraElegible",espacioActual.EspaciosAcademicosElegibles.CarreraElegible).Filter("CodigoAsignatura",espacioNuevo.EspaciosAcademicosElegibles.CodigoAsignatura).One(&espacioElegibleNuevo); err == nil {
			espacioNuevo.EspaciosAcademicosElegibles = &espacioElegibleNuevo
			//Se inserta el nuevo espacio
			if idEspacioNuevo, err := o.Insert(&espacioNuevo); err == nil {
				fmt.Println(idEspacioNuevo)
			}else{
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_13")
		}
	}
	err = o.Commit()

	return
}
