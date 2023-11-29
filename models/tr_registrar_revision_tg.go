package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrRegistrarRevisionTg struct {
	Comentarios          []Comentario
	RevisionTrabajoGrado RevisionTrabajoGrado
}

// Función para la transaccion de revisiones de anteproyectos
func AddTransaccionRegistrarRevisionTg(m *TrRegistrarRevisionTg) (alerta []string, err error) {

	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	alerta = append(alerta, "Success")
	if m.RevisionTrabajoGrado.Id == 0 {

		// Se debe crear una nueva revisión
		estado, err := GetAllEstadoRevisionTrabajoGrado(map[string]string{"CodigoAbreviacion": "FINALIZADA"}, nil, nil, nil, 0, 1)
		if err != nil || len(estado) != 1 {
			return nil, err
		}

		estadoTr, err := GetAllEstadoTrabajoGrado(map[string]string{"CodigoAbreviacion": "EC"}, nil, nil, nil, 0, 1)
		if err != nil || len(estado) != 1 {
			return nil, err
		}

		// Se consulta el estado de la nueva revisión
		m.RevisionTrabajoGrado.EstadoRevisionTrabajoGrado = estado[0].(EstadoRevisionTrabajoGrado).Id
		m.RevisionTrabajoGrado.VinculacionTrabajoGrado.TrabajoGrado.EstadoTrabajoGrado = estadoTr[0].(EstadoTrabajoGrado).Id

		// Se consulta el número de revisiones realizadas a la fecha
		numRevisiones, err := o.QueryTable(new(RevisionTrabajoGrado)).Filter("DocumentoTrabajoGrado__TrabajoGrado__Id", fmt.Sprint(m.RevisionTrabajoGrado.DocumentoTrabajoGrado.TrabajoGrado.Id)).Count()
		if err != nil {
			return nil, err
		}

		m.RevisionTrabajoGrado.NumeroRevision = int(numRevisiones) + 1
		m.RevisionTrabajoGrado.FechaRecepcion = time.Now()
		m.RevisionTrabajoGrado.FechaRevision = time.Now()
		_, err = o.Insert(&m.RevisionTrabajoGrado)
		if err != nil {
			panic(err)
		}

		correccion := Correccion{
			RevisionTrabajoGrado: &m.RevisionTrabajoGrado,
			Documento:            false,
		}
		_, err = o.Insert(&correccion)
		if err != nil {
			panic(err)

		}

		for _, comment := range m.Comentarios {
			comment.Correccion = &correccion
			comment.Fecha = time.Now()
			_, err = o.Insert(&comment)
			if err != nil {
				alerta[0] = "Error"
				alerta = append(alerta, "ERROR.INSERTANDO_REVISIONES")
				panic(err)
			}
		}

		_, err = o.Update(m.RevisionTrabajoGrado.VinculacionTrabajoGrado.TrabajoGrado, "EstadoTrabajoGrado")
		if err != nil {
			panic(err)
		}

		return alerta, nil
	}

	// Update de la revisión del trabajo de grado
	_, err = o.Update(&m.RevisionTrabajoGrado, "EstadoRevisionTrabajoGrado", "FechaRevision")
	if err != nil {
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR.REGISTRANDO_REVISION")
		panic(err)
	}

	// Insert de las correcciones y comentarios
	for _, v := range m.Comentarios {
		idCorreccion, err := o.Insert(v.Correccion)
		if err != nil {
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.INSERTANDO_REVISIONES")
			panic(err)
		}

		v.Correccion = &Correccion{Id: int(idCorreccion)}
		_, err = o.Insert(&v)
		if err != nil {
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.INSERTANDO_REVISIONES")
			panic(err)
		}
	}

	return
}
