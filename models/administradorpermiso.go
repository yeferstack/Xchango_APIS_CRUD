package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Administradorpermiso struct {
	Id                int           `orm:"column(id);pk"`
	IdAdmin           *Administrador `orm:"column(id_admin);rel(fk)"`
	IdPermiso         *Permiso      `orm:"column(id_permiso);rel(fk)"`
	Activo            bool          `orm:"column(activo)"`
	FechaCreacion     time.Time     `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion time.Time     `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Administradorpermiso) TableName() string {
	return "administradorpermiso"
}

func init() {
	orm.RegisterModel(new(Administradorpermiso))
}

func AddAdministradorpermiso(m *Administradorpermiso) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetAdministradorpermisoById(id int) (v *Administradorpermiso, err error) {
	o := orm.NewOrm()
	v = &Administradorpermiso{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IdAdmin")
		o.LoadRelated(v, "IdPermiso")
		return v, nil
	}
	return nil, err
}

func GetAllAdministradorpermiso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Administradorpermiso)).RelatedSel()

	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}

	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
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

	var l []Administradorpermiso
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
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

func UpdateAdministradorpermisoById(m *Administradorpermiso) (err error) {
	o := orm.NewOrm()
	v := Administradorpermiso{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteAdministradorpermiso(id int) (err error) {
	o := orm.NewOrm()
	v := Administradorpermiso{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Administradorpermiso{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}