package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RespuestaSolicitud struct {
	Id                    int                    `orm:"column(id);pk;auto"`
	Fecha                 time.Time              `orm:"column(fecha);type(timestamp without time zone)"`
	Justificacion         string                 `orm:"column(justificacion);null"`
	EnteResponsable       int                    `orm:"column(ente_responsable);null"`
	Usuario               int                    `orm:"column(usuario);null"`
	EstadoSolicitud       int                    `orm:"column(estado_solicitud)"`
	SolicitudTrabajoGrado *SolicitudTrabajoGrado `orm:"column(solicitud_trabajo_grado);rel(fk)"`
	Activo                bool                   `orm:"column(activo)"`
}

func (t *RespuestaSolicitud) TableName() string {
	return "respuesta_solicitud"
}

func init() {
	orm.RegisterModel(new(RespuestaSolicitud))
}

// AddRespuestaSolicitud insert a new RespuestaSolicitud into database and returns
// last inserted Id on success.
func AddRespuestaSolicitud(m *RespuestaSolicitud) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRespuestaSolicitudById retrieves RespuestaSolicitud by Id. Returns error if
// Id doesn't exist
func GetRespuestaSolicitudById(id int) (v *RespuestaSolicitud, err error) {
	o := orm.NewOrm()
	v = &RespuestaSolicitud{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRespuestaSolicitud retrieves all RespuestaSolicitud matches certain condition. Returns empty list if
// no records exist
func GetAllRespuestaSolicitud(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, exclude map[string]string) (ml []interface{}, err error) {
	o := orm.NewOrm() // query: k:v,k:v
	qs := o.QueryTable(new(RespuestaSolicitud)).RelatedSel(4)

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}

	// exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Exclude(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "in") {
			arr := strings.Split(v, "|")
			qs = qs.Exclude(k, arr)
		} else {
			qs = qs.Exclude(k, v)
		}
	}

	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RespuestaSolicitud
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRespuestaSolicitud updates RespuestaSolicitud by Id and returns error if
// the record to be updated doesn't exist
func UpdateRespuestaSolicitudById(m *RespuestaSolicitud) (err error) {
	o := orm.NewOrm()
	v := RespuestaSolicitud{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRespuestaSolicitud deletes RespuestaSolicitud by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRespuestaSolicitud(id int) (err error) {
	o := orm.NewOrm()
	v := RespuestaSolicitud{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RespuestaSolicitud{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
