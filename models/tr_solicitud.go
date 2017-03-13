package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrSolicitud struct {
	//Tg *TrabajoGrado
	Solicitud *SolicitudMaterias
	//EstudianteTg        *EstudianteTg
	MateriasSolicitadas *[]AsignaturaInscrita
}

//funcion para la transaccion de solicitudes
func AddTransaccionSolicitud(m *TrSolicitud) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	if id, err := o.Insert(m.Solicitud); err == nil {
		for _, v := range *m.MateriasSolicitadas {
			v.IdSolicitudMaterias.Id = int(id)
			if _, err = o.Insert(&v); err != nil {
				fmt.Println("entro aqui?")
				err = o.Rollback()
			}
		}
		err = o.Commit()
	} else {
		err = o.Rollback()
	}
	return
}
