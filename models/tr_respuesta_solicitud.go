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
	TrTrabajoGrado         *TrTrabajoGrado            //Solictudes iniciales
	ModalidadTipoSolicitud *ModalidadTipoSolicitud    //Para saber el tipo de solicitud inicial
	TrabajoGrado           *TrabajoGrado              //Cambio Titulo
}

//funcion para la transaccion de solicitudes
func AddTransaccionRespuestaSolicitud(m *TrRespuestaSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	var num int64

	//solicitud rechazada
	fmt.Println(m.RespuestaNueva.EstadoSolicitud.Id)
	if m.RespuestaNueva.EstadoSolicitud.Id == 5 {

		//update del estado de la ultima solicitud
		if num, err = o.Update(m.RespuestaAnterior, "Activo"); err == nil {
			fmt.Println("Number of records updated in database:", num)
			fmt.Println(m.RespuestaNueva)
			//insert de la nueva rta
			if id_rta, err := o.Insert(m.RespuestaNueva); err == nil {
				fmt.Println(id_rta)
				//insert documento asociado a la nueva rta
				//m.DocumentoSolicitud.SolicitudTrabajoGrado.Id = int(id_rta)
				if id_documento, err := o.Insert(m.DocumentoSolicitud); err == nil {
					fmt.Println(id_documento)
					err = o.Commit()
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
	} else {

		if n, er := o.Update(m.RespuestaAnterior); er == nil {
			fmt.Println("Number of records updated in database:", n)
			fmt.Println(m.RespuestaNueva)
			//insert de la nueva rta
			if id_rta, err := o.Insert(m.RespuestaNueva); err == nil {
				fmt.Println(id_rta)

				//insert documento asociado a la nueva rta
				//m.DocumentoSolicitud.SolicitudTrabajoGrado.Id = int(id_rta)
				if id_documento, err := o.Insert(m.DocumentoSolicitud); err == nil {
					fmt.Println(id_documento)

					//seguir la transacción según el caso: de acuerdo al tipo de solicitud
					fmt.Println(m.TipoSolicitud.Nombre)
					fmt.Println(m.TipoSolicitud.Id)

					if (m.TipoSolicitud.Id == 4) || (m.TipoSolicitud.Id == 10) { //solicitud de cambio de director interno, o evaluador
						for _, v := range *m.Vinculaciones {
							fmt.Println(v)
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
						err = o.Commit()
					}

					if m.TipoSolicitud.Id == 8 { //Solicitud de cambio tÍtulo trabajo de grado

						if _, err = o.Update(m.TrabajoGrado); err != nil {
							fmt.Println(err)
							err = o.Rollback()
							alerta[0] = "Error"
							alerta = append(alerta, "ERROR_RTA_SOLICITUD_13")
						} else {
							err = o.Commit()
						}
					}

					if m.TipoSolicitud.Id == 6 || m.TipoSolicitud.Id == 7 { //Solicitud de socialización o prórroga
						err = o.Commit()
					}

					if m.TipoSolicitud.Id == 3 { //solicitud de cancelación de la modalidad
						if num, err = o.Update(m.EstudianteTrabajoGrado); err == nil {
							fmt.Println("Number of records updated in database:", num)
							//consultar # estudiantes del trabajo de grado
							cnt, _ := o.QueryTable("estudiante_trabajo_grado").Filter("trabajo_grado", m.EstudianteTrabajoGrado.TrabajoGrado.Id).Count()
							fmt.Println("Count:", cnt)
							//estado(cancelado) = 2
							cnt2, _ := o.QueryTable("estudiante_trabajo_grado").Filter("trabajo_grado", m.EstudianteTrabajoGrado.TrabajoGrado.Id).Filter("estado_estudiante_trabajo_grado", 2).Count()
							fmt.Println("Count:", cnt2)

							if cnt == cnt2 {
								//se inactivan las vinculaciones
								for _, v := range *m.Vinculaciones {
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
									err = o.Commit()
								} else {
									fmt.Println(err)
									err = o.Rollback()
									alerta[0] = "Error"
									alerta = append(alerta, "ERROR_RTA_SOLICITUD_8")
								}
							} else {
								err = o.Commit()
							}
						} else {
							fmt.Println(err)
							err = o.Rollback()
							alerta[0] = "Error"
							alerta = append(alerta, "ERROR_RTA_SOLICITUD_7")
						}
					}

					//solicitudes iniciales
					if m.TipoSolicitud.Id == 2 {
						//Monografia, Proyecto de emprendimento, Creación e Interpretación, Producción académica
						if m.ModalidadTipoSolicitud.Id == 20 || m.ModalidadTipoSolicitud.Id == 46 || m.ModalidadTipoSolicitud.Id == 38 || m.ModalidadTipoSolicitud.Id == 55 {

							if id, err := o.Insert(m.TrTrabajoGrado.TrabajoGrado); err == nil {
								fmt.Println(id)

								for _, v := range *m.TrTrabajoGrado.EstudianteTrabajoGrado {
									v.TrabajoGrado.Id = int(id)
									if _, err = o.Insert(&v); err != nil {
										fmt.Println(err)
										err = o.Rollback()
										alerta[0] = "Error"
										alerta = append(alerta, "ERROR_RTA_SOLICITUD_9")
									}
								}

								for _, v := range *m.TrTrabajoGrado.AreasTrabajoGrado {
									v.TrabajoGrado.Id = int(id)
									if _, err = o.Insert(&v); err != nil {
										fmt.Println(err)
										err = o.Rollback()
										alerta[0] = "Error"
										alerta = append(alerta, "ERROR_RTA_SOLICITUD_10")
									}
								}

								for _, v := range *m.TrTrabajoGrado.VinculacionTrabajoGrado {
									v.TrabajoGrado.Id = int(id)
									if _, err = o.Insert(&v); err != nil {
										fmt.Println(err)
										err = o.Rollback()
										alerta[0] = "Error"
										alerta = append(alerta, "ERROR_RTA_SOLICITUD_5")
									}
								}

								if id_documento, err := o.Insert(m.TrTrabajoGrado.DocumentoEscrito); err == nil {
									fmt.Println(id_documento)
									m.TrTrabajoGrado.DocumentoTrabajoGrado.TrabajoGrado.Id = int(id)
									m.TrTrabajoGrado.DocumentoTrabajoGrado.DocumentoEscrito.Id = int(id_documento)

									if id_documento, err := o.Insert(m.TrTrabajoGrado.DocumentoTrabajoGrado); err == nil {
										fmt.Println(id_documento)
										err = o.Commit()
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
								alerta = append(alerta, "ERROR_RTA_SOLICITUD_4")
							}

						} else if m.ModalidadTipoSolicitud.Id == 13 || m.ModalidadTipoSolicitud.Id == 16 { //espacios académicos de posgrado o profundización
							err = o.Commit()
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
				alerta = append(alerta, "ERROR_RTA_SOLICITUD_2")
			}

		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR_RTA_SOLICITUD_1")
		}
	}

	return
}
