package models

import (
	"fmt"
	"strings"
	"github.com/astaxie/beego/orm"
)

type TrRevision struct {
	TrabajoGrado    	   *TrabajoGrado
	Vinculaciones          *[]VinculacionTrabajoGrado //Cambio de director o evaluador
	DocumentoEscrito	   *DocumentoEscrito
	DocumentoTrabajoGrado  *DocumentoTrabajoGrado 	  
}

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
	DetallesPasantia       *DetallePasantia  //SOlicitud inicial de pasantia
	TrRevision			   *TrRevision	 //Solicitud de revisión

}

// AddTransaccionRespuestaSolicitud funcion para dar respuesta a las solicitudes
func AddTransaccionRespuestaSolicitud(m *TrRespuestaSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	err = o.Begin()
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
				// Si la solicitud es de pasantia se agrega el detalel de la pasantia
				if m.DetallesPasantia != nil {
					m.DetallesPasantia.TrabajoGrado.Id = int(id_TrabajoGrado)
					if _, err = o.Insert(m.DetallesPasantia); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_15")
					}
				}
			} else {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_3")
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
		var idVinculadoAntiguo int
		var idVinculadoNuevo int64
		for _, v := range *m.Vinculaciones {
			//Si esta activo es nuevo y se inserta sino se actualiza la fecha de fin y el activo
			if v.Activo == true {
				// Se busca si el docente ya estuvo vinculado y se actualiza
				var vinculado VinculacionTrabajoGrado
				if err = o.QueryTable(new(VinculacionTrabajoGrado)).RelatedSel().Filter("TrabajoGrado",v.TrabajoGrado).Filter("Usuario",v.Usuario).Filter("RolTrabajoGrado",v.RolTrabajoGrado).One(&vinculado); err == nil {
					//SI si se encuentra 
					idVinculadoNuevo = int64(vinculado.Id)
					vinculado.Activo = v.Activo
					vinculado.FechaFin = v.FechaFin
					vinculado.FechaInicio = v.FechaInicio
					fmt.Println("Se actualiza vinculado", vinculado)
					if _, err = o.Update(&vinculado); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
					} 
				} else if (err == orm.ErrNoRows) {
					// Si no se encuentra 
					fmt.Println("Se inserta vinculado", v)
					if idVinculadoNuevo, err = o.Insert(&v); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
					} 
				} else {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
				}
			} else {
				idVinculadoAntiguo = v.Id;
				if _, err = o.Update(&v, "Activo", "FechaFin"); err != nil {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_6")
				}
			}
		}

		//Se busca si el vinculado antiguo tiene una revision pendiente
		var revisionTrabajoGrado RevisionTrabajoGrado
		if err = o.QueryTable(new(RevisionTrabajoGrado)).RelatedSel().Filter("VinculacionTrabajoGrado",idVinculadoAntiguo).Filter("estado_revision_trabajo_grado",1).One(&revisionTrabajoGrado); err == nil {
			//Se actualiza el id de la vinculación
			revisionTrabajoGrado.VinculacionTrabajoGrado.Id = int(idVinculadoNuevo)
			if _, err = o.Update(&revisionTrabajoGrado, "VinculacionTrabajoGrado"); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
			}
		} else if (err == orm.ErrNoRows) {
			// Si no se encuentra 
			fmt.Println("El vinculado no tiene revisiones pendientes")
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
		}

		// Si  el cambio es de director externo, se recibe la data del detalle de la pasantia y 
		// se actualiza
		if m.DetallesPasantia != nil {
			//Se busca el detalle de la pasantia asociado al tg
			var detallePasantia DetallePasantia
			if err = o.QueryTable(new(DetallePasantia)).RelatedSel().Filter("TrabajoGrado",m.DetallesPasantia.TrabajoGrado.Id).One(&detallePasantia); err == nil {
				detallePasantia.Observaciones = strings.Split(detallePasantia.Observaciones," y dirigida por ")[0];
				detallePasantia.Observaciones += m.DetallesPasantia.Observaciones
				if num, err = o.Update(&detallePasantia,"Observaciones"); err == nil {
					fmt.Println("Number of records updated in database:", num)
				} else {
					fmt.Println(err)
					err = o.Rollback()
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_16")
				}
			} else {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_16")
			}
		}

		//Si el estado del trabajo de grado es  no viable 7 para primer evaluador se cambia el estado del trabajo de grado a 
		// En evaluación por segundo revisor 9
		/*
		trabajo := (*m.Vinculaciones)[0].TrabajoGrado;
		if (trabajo.EstadoTrabajoGrado.Id == 7) {
			trabajo.EstadoTrabajoGrado.Id = 9
			// Se actualiza el trabajo de grado
			if num, err = o.Update(trabajo, "EstadoTrabajoGrado"); err == nil {
				fmt.Println("Number of rows updated with trabajo de grado", num)
			} else {
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_17")
				fmt.Println(err)
				err = o.Rollback()
			}
		}
		*/

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

	// Solicitud de revisión del trabajo de grado
	if m.TrRevision != nil {
		// Se actualiza el trabajo de grado
		if num, err = o.Update(m.TrRevision.TrabajoGrado, "EstadoTrabajoGrado"); err == nil {
			fmt.Println("Number of rows updated with trabajo de grado", num)
		} else {
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_17")
			fmt.Println(err)
			err = o.Rollback()
		}
		//Se inserta el documento final de la revisión y se relaciona con el trabajo grado
		if id_documento, err := o.Insert(m.TrRevision.DocumentoEscrito); err == nil {
			fmt.Println("Id documento final",id_documento)
			m.TrRevision.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(id_documento)
			//Se inserta el nuevo documento escrito
			if num, err = o.Insert(m.TrRevision.DocumentoTrabajoGrado); err == nil {
				fmt.Println("Number of rows updated with documento trabajo de grado", num)
			} else {
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_18")
				fmt.Println(err)
				err = o.Rollback()
			}
			/*var documentoTrabajoGrado DocumentoTrabajoGrado
			if err = o.QueryTable(new(DocumentoTrabajoGrado)).RelatedSel().Filter("documento_escrito__tipo_documento_escrito",4).Filter("trabajo_grado",m.TrRevision.TrabajoGrado.Id).One(&documentoTrabajoGrado); err == nil {
				m.TrRevision.DocumentoTrabajoGrado.Id = documentoTrabajoGrado.Id 
				if num, err = o.Update(m.TrRevision.DocumentoTrabajoGrado, "DocumentoEscrito"); err == nil {
					fmt.Println("Number of rows updated with documento trabajo de grado", num)
				} else {
					alerta[0] = "Error"
					alerta = append(alerta, "ERROR_RTA_SOLICITUD_18")
					fmt.Println(err)
					err = o.Rollback()
				}
			} else {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_18")
			}*/
		} else {
			fmt.Println(err)
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_18")
			err = o.Rollback()
		}
		//Se actualizan las vinculaciones
		for _, v := range *m.TrRevision.Vinculaciones {
			//Si esta activo es nuevo y se inserta sino se actualiza la fecha de fin y el activo
			if v.Activo == true {
				// Se busca si el docente ya estuvo vinculado y se actualiza
				var vinculado VinculacionTrabajoGrado
				if err = o.QueryTable(new(VinculacionTrabajoGrado)).RelatedSel().Filter("TrabajoGrado",v.TrabajoGrado).Filter("Usuario",v.Usuario).Filter("RolTrabajoGrado",v.RolTrabajoGrado).One(&vinculado); err == nil {
					//SI si se encuentra 
					vinculado.Activo = v.Activo
					vinculado.FechaFin = v.FechaFin
					vinculado.FechaInicio = v.FechaInicio
					fmt.Println("Se actualiza vinculado", vinculado)
					if _, err = o.Update(&vinculado); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
					} 
				} else if (err == orm.ErrNoRows) {
					// Si no se encuentra 
					fmt.Println("Se inserta vinculado", v)
					if _, err = o.Insert(&v); err != nil {
						fmt.Println(err)
						err = o.Rollback()
						alerta[0] = "Error"
						alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
					} 
				} else {
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
	err = o.Commit()
	//err = o.Rollback()
	return
}
