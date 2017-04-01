package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type TrPregunta struct {
	Pregunta   *PreguntaFormato
	Respuestas *[]RespuestaFormato
}

type TrFormato struct {
	Formato     *Formato
	TrPreguntas *[]TrPregunta
}

type TrUpdatePregunta struct {
	Pregunta   *PreguntaFormato
	TypeUpdate string
	Respuestas *[]TrUpdateRespuesta
}

type TrUpdateRespuesta struct {
	Respuesta  *RespuestaFormato
	TypeUpdate string
}

type TrUpdateFormato struct {
	Formato    *Formato
	TypeUpdate string
	Preguntas  *[]TrUpdatePregunta
}

func UpdateTrFormato(m *TrUpdateFormato) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	if m.TypeUpdate == "update" {
		if _, err = o.Update(m.Formato); err != nil {
			o.Rollback()
			return 0, err
		}
	}
	for _, preg := range *m.Preguntas {
		if preg.Pregunta.IdPregunta.Id == 0 {
			if _, err = o.Insert(preg.Pregunta.IdPregunta); err != nil {
				o.Rollback()
				return 0, err
			}
		}
		switch tu := preg.TypeUpdate; tu {
		case "update":
			if _, err = o.Update(preg.Pregunta); err != nil {
				o.Rollback()
				return 0, err
			}
		case "delete":
			if _, err = o.Delete(preg.Pregunta.Id); err != nil {
				o.Rollback()
				return 0, err
			}
		case "insert":
			if _, err = o.Insert(preg.Pregunta); err != nil {
				o.Rollback()
				return 0, err
			}
		}
		for _, res := range *preg.Respuestas {
			if res.Respuesta.IdRespuesta.Id == 0 {
				if _, err = o.Insert(res.Respuesta.IdRespuesta); err != nil {
					o.Rollback()
					return 0, err
				}
			}
			switch tus := res.TypeUpdate; tus {
			case "update":
				if _, err = o.Update(res.Respuesta); err != nil {
					o.Rollback()
					return 0, err
				}
			case "delete":
				if _, err = o.Delete(res.Respuesta.Id); err != nil {
					o.Rollback()
					return 0, err
				}
			case "insert":
				if _, err = o.Insert(res.Respuesta); err != nil {
					o.Rollback()
					return 0, err
				}
			}
		}

	}
	o.Commit()
	return 0, nil
}

//funcion para la transaccion de solicitudes
func AddTrFormato(m *TrFormato) (id int64, err error) {
	fmt.Println("formato:", m.Formato)
	o := orm.NewOrm()
	o.Begin()
	if id, err = o.Insert(m.Formato); err == nil {
		for _, v := range *m.TrPreguntas {
			v.Pregunta.IdFormato.Id = int(id)
			fmt.Println("pregunta:", v.Pregunta)
			//---
			if v.Pregunta.IdPregunta.Id == 0 {
				if _, err = o.Insert(v.Pregunta.IdPregunta); err != nil {
					err = o.Rollback()
					break
				}
				fmt.Println("preguntica:", v.Pregunta.IdPregunta)
			}
			//--
			if _, err = o.Insert(v.Pregunta); err != nil {
				fmt.Println("ocurrio algo!", err)
				err = o.Rollback()
				break
			} else {

				//--
				for _, vr := range *v.Respuestas {
					vr.IdPreguntaFormato = v.Pregunta

					//---
					if vr.IdRespuesta.Id == 0 {
						if _, err = o.Insert(vr.IdRespuesta); err != nil {
							err = o.Rollback()
							break
						}
						fmt.Println("respuesticca:", vr.IdRespuesta)
					}
					//--

					if _, err = o.Insert(&vr); err != nil {
						fmt.Println("ocurrio algo con respuestas!", err)
						err = o.Rollback()
						break
					}
				}

				//--
			}
		}
		err = o.Commit()
	} else {
		err = o.Rollback()
	}

	return id, err
}

//funcion para la transaccion de solicitudes
func GetTrFormato(id int) (v *TrFormato, err error) {
	fmt.Println("formato:saf", id)

	o := orm.NewOrm()
	var f = new(Formato)
	f.Id = id
	//v.Formato = &Formato{Id: id}
	if err = o.Read(f); err == nil {
		fmt.Println("formato:", *f)
		var trf = new(TrFormato)
		trf.Formato = f

		var preguntas []PreguntaFormato

		idformato := strconv.Itoa(id)
		//_, err = o.Raw("select * from polux.pregunta_formato where id_formato=" + idformato).QueryRows(&preguntas)

		o.QueryTable(new(PreguntaFormato)).Filter("id_formato", idformato).RelatedSel().All(&preguntas)

		var tam = len(preguntas)
		fmt.Println(tam)

		var tr_preguntas []TrPregunta

		fmt.Println("preguntas query: ", preguntas)

		for _, v := range preguntas {
			p := v
			//o.QueryTable(&p).Filter("Id", p.Id).RelatedSel().All(&p)
			var trpregunta TrPregunta
			var respuestas []RespuestaFormato
			trpregunta.Pregunta = &p
			o.QueryTable(new(RespuestaFormato)).Filter("id_pregunta_formato", p.Id).RelatedSel().All(&respuestas)
			trpregunta.Respuestas = &respuestas

			tr_preguntas = append(tr_preguntas, trpregunta)
		}

		//o.QueryTable(&preguntas).Filter("IdFormato", id).RelatedSel().All(&preguntas)

		trf.TrPreguntas = &tr_preguntas

		return trf, nil
	}
	return nil, err
}
