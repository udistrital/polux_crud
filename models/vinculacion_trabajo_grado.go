package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type VinculacionTrabajoGrado struct {
	Id              int           `orm:"column(id);pk;auto"`
	Usuario         int           `orm:"column(usuario)"`
	Activo          bool          `orm:"column(activo)"`
	FechaInicio     string     	  `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin        *string       `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	RolTrabajoGrado int           `orm:"column(rol_trabajo_grado)"`
	TrabajoGrado    *TrabajoGrado `orm:"column(trabajo_grado);rel(fk)"`
	CodigoAbreviacion string	  `orm:"column(codigo_abreviacion)"`
}

func (t *VinculacionTrabajoGrado) TableName() string {
	return "vinculacion_trabajo_grado"
}

func init() {
	orm.RegisterModel(new(VinculacionTrabajoGrado))
}

// AddVinculacionTrabajoGrado insert a new VinculacionTrabajoGrado into database and returns
// last inserted Id on success.
func AddVinculacionTrabajoGrado(m *VinculacionTrabajoGrado) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVinculacionTrabajoGradoById retrieves VinculacionTrabajoGrado by Id. Returns error if
// Id doesn't exist
func GetVinculacionTrabajoGradoById(id int) (v *VinculacionTrabajoGrado, err error) {
	o := orm.NewOrm()
	v = &VinculacionTrabajoGrado{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVinculacionTrabajoGrado retrieves all VinculacionTrabajoGrado matches certain condition. Returns empty list if
// no records exist
func GetAllVinculacionTrabajoGrado(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VinculacionTrabajoGrado)).RelatedSel()
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

	var l []VinculacionTrabajoGrado
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

// UpdateVinculacionTrabajoGrado updates VinculacionTrabajoGrado by Id and returns error if
// the record to be updated doesn't exist
func UpdateVinculacionTrabajoGradoById(m *VinculacionTrabajoGrado) (err error) {
	o := orm.NewOrm()
	v := VinculacionTrabajoGrado{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVinculacionTrabajoGrado deletes VinculacionTrabajoGrado by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVinculacionTrabajoGrado(id int) (err error) {
	o := orm.NewOrm()
	v := VinculacionTrabajoGrado{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&VinculacionTrabajoGrado{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
