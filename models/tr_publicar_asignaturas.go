package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrPublicarAsignaturas struct {
	CarreraElegible              *CarreraElegible
	EspaciosAcademicosElegibles  *[]EspaciosAcademicosElegibles
}

// AddTransaccionPublicarAsignaturas funcion para registrar las materias de posgrado
func AddTransaccionPublicarAsignaturas(m *TrPublicarAsignaturas) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	// Buscar si la carrera elegible ta esta registrada
	var carreraElegible CarreraElegible
	//se busca la carrera elegible
	if err := o.QueryTable(new(CarreraElegible)).RelatedSel().Filter("CodigoCarrera",m.CarreraElegible.CodigoCarrera).Filter("Periodo",m.CarreraElegible.Periodo).Filter("Anio",m.CarreraElegible.Anio).Filter("CodigoPensum",m.CarreraElegible.CodigoPensum).One(&carreraElegible); err == nil{
		//La carrera  esta registrada
		m.CarreraElegible.Id = carreraElegible.Id
		fmt.Println("CarreraElegible", m.CarreraElegible.Id)
	} else if err == orm.ErrNoRows{
		//La carrera no esta registrada
		if idCarreraElegible, err := o.Insert(m.CarreraElegible); err == nil {
			fmt.Println("CarreraElegible", idCarreraElegible);
			m.CarreraElegible.Id = int(idCarreraElegible)
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "ERROR.INSERTAR_CARRERA_ELEGIBLE")
		}
	} else {
		fmt.Println(err)
		err = o.Rollback()
		alerta[0] = "Error"
		alerta = append(alerta, "ERROR.CARGAR_CARRERA_ELEGIBLE")
	}
	for _, v := range *m.EspaciosAcademicosElegibles {
		v.CarreraElegible.Id = m.CarreraElegible.Id
		//buscar asignatura elegible
		var asignatura EspaciosAcademicosElegibles
		if err := o.QueryTable(new(EspaciosAcademicosElegibles)).RelatedSel().Filter("CodigoAsignatura",v.CodigoAsignatura).Filter("CarreraElegible",v.CarreraElegible.Id).One(&asignatura); err == nil{
			//Asignatura registrada, se actualiza el estado
			fmt.Println("update", v);
			v.Id = asignatura.Id;
			v.Activo = !asignatura.Activo;
			if _, err = o.Update(&v,"Activo"); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "ACTUALIZAR_ESPACIO_ACADEMICO_ELEGIBLE")
			}			
		} else if err == orm.ErrNoRows{
			//La asignatura no esta registrada, se actualiza
			fmt.Println("instert",v);
			if _, err = o.Insert(&v); err != nil {
				fmt.Println(err)
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "INSTERTAR_ESPACIO_ACADEMICO_ELEGIBLE")
			}
		} else {
			fmt.Println(err)
			err = o.Rollback()
			alerta[0] = "Error"
			alerta = append(alerta, "CARGAR_ESPACIO_ACADEMICO_ELEGIBLE")
		}
	}
	err = o.Commit()
	return
}
