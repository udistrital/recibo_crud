package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Recibo struct {
	Id                  int           `orm:"column(id);pk;auto"`
	Referencia          int           `orm:"column(referencia);null"`
	TipoReciboId        *TipoRecibo   `orm:"column(tipo_recibo_id);rel(fk)"`
	EstadoReciboId      *EstadoRecibo `orm:"column(estado_recibo_id);rel(fk)"`
	FechaOrdinaria      time.Time     `orm:"column(fecha_ordinaria);type(date)"`
	FechaExtraordinaria time.Time     `orm:"column(fecha_extraordinaria);type(date);null"`
	ValorOrdinario      int           `orm:"column(valor_ordinario)"`
	ValorExtraordinario int           `orm:"column(valor_extraordinario);null"`
	FechaModificacion 	time.Time 	  `orm:"column(fecha_modificacion);type(timestamp with time zone);auto_now;null"`
}

func (t *Recibo) TableName() string {
	return "recibo"
}

func init() {
	orm.RegisterModel(new(Recibo))
}

// AddRecibo insert a new Recibo into database and returns
// last inserted Id on success.
func AddRecibo(m *Recibo) (id int64, err error) {
	m.FechaModificacion = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReciboById retrieves Recibo by Id. Returns error if
// Id doesn't exist
func GetReciboById(id int) (v *Recibo, err error) {
	o := orm.NewOrm()
	v = &Recibo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRecibo retrieves all Recibo matches certain condition. Returns empty list if
// no records exist
func GetAllRecibo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Recibo))
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

	var l []Recibo
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

// UpdateRecibo updates Recibo by Id and returns error if
// the record to be updated doesn't exist
func UpdateReciboById(m *Recibo) (err error) {
	o := orm.NewOrm()
	v := Recibo{Id: m.Id}
	m.FechaModificacion = time.Now()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRecibo deletes Recibo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecibo(id int) (err error) {
	o := orm.NewOrm()
	v := Recibo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Recibo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
