package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type IntentoLogin struct {
	Id                int       `orm:"column(id_intento);pk;auto"`
	IdUsuario         *Usuario  `orm:"column(id_usuario);rel(fk)"`
	EmailIngresado    string    `orm:"column(email_ingresado);null"`
	Exitoso           bool      `orm:"column(exitoso)"`
	MotivoFallo       string    `orm:"column(motivo_fallo);null"`
	IpOrigen          string    `orm:"column(ip_origen);null"`
	Fecha             time.Time `orm:"column(fecha);type(timestamp without time zone)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *IntentoLogin) TableName() string {
	return "intento_login"
}

func init() {
	orm.RegisterModel(new(IntentoLogin))
}

// AddIntentoLogin insert a new IntentoLogin into database and returns
// last inserted Id on success.
func AddIntentoLogin(m *IntentoLogin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetIntentoLoginById retrieves IntentoLogin by Id. Returns error if
// Id doesn't exist
func GetIntentoLoginById(id int) (v *IntentoLogin, err error) {
	o := orm.NewOrm()
	v = &IntentoLogin{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IdUsuario")
		return v, nil
	}
	return nil, err
}

// GetAllIntentoLogin retrieves all IntentoLogin matches certain condition. Returns empty list if
// no records exist
func GetAllIntentoLogin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(IntentoLogin)).RelatedSel()
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

	var l []IntentoLogin
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

// UpdateIntentoLogin updates IntentoLogin by Id and returns error if
// the record to be updated doesn't exist
func UpdateIntentoLoginById(m *IntentoLogin) (err error) {
	o := orm.NewOrm()
	v := IntentoLogin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIntentoLogin deletes IntentoLogin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIntentoLogin(id int) (err error) {
	o := orm.NewOrm()
	v := IntentoLogin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&IntentoLogin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
