package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type EstadoRecibo struct {
	Id                int     `orm:"column(id);pk;auto"`
	Nombre            string  `orm:"column(nombre)"`
	Descripcion       string  `orm:"column(descripcion)"`
	CodigoAbreviacion string  `orm:"column(codigo_abreviacion)"`
	NumeroOrden       float64 `orm:"column(numero_orden)"`
	Activo            bool    `orm:"column(activo)"`
	FechaCreacion     string  `orm:"column(fecha_creacion);null"`
	FechaModificacion string  `orm:"column(fecha_modificacion);null"`
}

func (t *EstadoRecibo) TableName() string {
	return "estado_recibo"
}

func init() {
	orm.RegisterModel(new(EstadoRecibo))
}

// AddEstadoRecibo insert a new EstadoRecibo into database and returns
// last inserted Id on success.
func AddEstadoRecibo(m *EstadoRecibo) (id int64, err error) {
	m.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEstadoReciboById retrieves EstadoRecibo by Id. Returns error if
// Id doesn't exist
func GetEstadoReciboById(id int) (v *EstadoRecibo, err error) {
	o := orm.NewOrm()
	v = &EstadoRecibo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEstadoRecibo retrieves all EstadoRecibo matches certain condition. Returns empty list if
// no records exist
func GetAllEstadoRecibo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EstadoRecibo))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
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

	var l []EstadoRecibo
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

// UpdateEstadoRecibo updates EstadoRecibo by Id and returns error if
// the record to be updated doesn't exist
func UpdateEstadoReciboById(m *EstadoRecibo) (err error) {
	o := orm.NewOrm()
	v := EstadoRecibo{Id: m.Id}
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Nombre", "Descripcion", "CodigoAbreviacion", "Activo", "NumeroOrden", "FechaModificacion"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEstadoRecibo deletes EstadoRecibo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEstadoRecibo(id int) (err error) {
	o := orm.NewOrm()
	v := EstadoRecibo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EstadoRecibo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
