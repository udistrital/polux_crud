package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrPublicarAsignaturas struct {
	CarreraElegible             *CarreraElegible
	EspaciosAcademicosElegibles []EspaciosAcademicosElegibles
}

// AddTransaccionPublicarAsignaturas funcion para registrar las materias de posgrado
func AddTransaccionPublicarAsignaturas(m *TrPublicarAsignaturas) (alerta []string, err error) {

	o := orm.NewOrm()

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

	err = o.Begin()
	if err != nil {
		panic(err)
	}

	alerta = append(alerta, "Success")

	var carreraElegible CarreraElegible
	err = o.QueryTable(new(CarreraElegible)).RelatedSel().
		Filter("CodigoCarrera", m.CarreraElegible.CodigoCarrera).
		Filter("Periodo", m.CarreraElegible.Periodo).
		Filter("Anio", m.CarreraElegible.Anio).
		Filter("CodigoPensum", m.CarreraElegible.CodigoPensum).
		One(&carreraElegible)

	if err != nil {
		if err.Error() != orm.ErrNoRows.Error() {
			alerta = append(alerta, "ERROR.CARGAR_CARRERA_ELEGIBLE")
			panic(err)
		}

		idCarreraElegible, err := o.Insert(m.CarreraElegible)
		if err != nil {
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.INSERTAR_CARRERA_ELEGIBLE")
			panic(err)
		}
		m.CarreraElegible.Id = int(idCarreraElegible)
	} else {
		m.CarreraElegible.Id = carreraElegible.Id
	}

	// Registrar y desactivar asignaturas
	espaciosNuevos := []int{}
	for _, v := range m.EspaciosAcademicosElegibles {
		espaciosNuevos = append(espaciosNuevos, v.CodigoAsignatura)
		v.CarreraElegible = &CarreraElegible{Id: m.CarreraElegible.Id}

		var asignatura EspaciosAcademicosElegibles
		err = o.QueryTable(new(EspaciosAcademicosElegibles)).RelatedSel().
			Filter("CodigoAsignatura", v.CodigoAsignatura).
			Filter("CarreraElegible", v.CarreraElegible.Id).
			One(&asignatura)

		if err == nil { // Asignatura registrada, se actualiza el estado
			if asignatura.Activo {
				continue
			}

			v.Id = asignatura.Id
			v.Activo = true
			_, err = o.Update(&v, "Activo")
			if err != nil {
				alerta[0] = "Error"
				alerta = append(alerta, "ACTUALIZAR_ESPACIO_ACADEMICO_ELEGIBLE")
				panic(err)
			}
		} else if err.Error() == orm.ErrNoRows.Error() { // La asignatura no esta registrada, se registra
			_, err = o.Insert(&v)
			if err != nil {
				alerta[0] = "Error"
				alerta = append(alerta, "INSTERTAR_ESPACIO_ACADEMICO_ELEGIBLE")
				panic(err)
			}
		} else {
			alerta[0] = "Error"
			alerta = append(alerta, "CARGAR_ESPACIO_ACADEMICO_ELEGIBLE")
			panic(err)
		}
	}

	var espaciosDesactivar []EspaciosAcademicosElegibles
	qs := o.QueryTable(new(EspaciosAcademicosElegibles)).RelatedSel()
	if len(espaciosNuevos) > 0 {
		qs = qs.Exclude("CodigoAsignatura__in", espaciosNuevos)
	}

	_, err = qs.Filter("CarreraElegible", m.CarreraElegible.Id).
		Filter("Activo", true).
		All(&espaciosDesactivar)
	if err != nil {
		alerta[0] = "Error"
		alerta = append(alerta, "ACTUALIZAR_ESPACIO_ACADEMICO_ELEGIBLE")
		panic(err)
	}

	for _, e := range espaciosDesactivar {
		e.Activo = false
		_, err = o.Update(&e, "Activo")
		if err != nil {
			alerta[0] = "Error"
			alerta = append(alerta, "ACTUALIZAR_ESPACIO_ACADEMICO_ELEGIBLE")
			panic(err)
		}
	}

	return
}
