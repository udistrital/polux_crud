package models

import (

	"github.com/astaxie/beego/orm"
)

type TrabajoGradoTr struct {
	TrabajoGrado *TrabajoGrado
	EstudianteTg *EstudianteTg
}

// Transacción para registrar un trabajo de grado
func AddTrabajoGradoTr(m *TrabajoGradoTr) (id int64, err error) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Begin()
	if id_tg, err := o.Insert(m.TrabajoGrado); err == nil {
		m.EstudianteTg.IdTrabajoGrado.Id=int(id_tg)

		if id, err = o.Insert(m.EstudianteTg); err == nil {
			err = o.Commit()
		}else{
			err = o.Rollback()
		}

	}else {
		err = o.Rollback()
	}

	return
}
