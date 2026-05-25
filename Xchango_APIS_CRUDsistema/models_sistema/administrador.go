package models_sistema

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Administrador struct {
	Id                int          `orm:"column(id_admin);pk;auto"`
	IdUsuario         int          `orm:"column(id_usuario)"`
	IdNivelacceso     *NivelAcceso `orm:"column(id_nivelacceso);rel(fk)"`
	Activo            bool         `orm:"column(activo)"`
	FechaAsignacion   time.Time    `orm:"column(fecha_asignacion);type(timestamp without time zone)"`
	FechaCreacion     time.Time    `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion time.Time    `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Administrador) TableName() string {
	return "administrador"
}

func init() {
	orm.RegisterModel(new(Administrador))
}

func AddAdministrador(m *Administrador) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdministradorById: 
func GetAdministradorById(id int) (v *Administrador, err error) {
	o := orm.NewOrm()
	v = &Administrador{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IdNivelacceso") 
		return v, nil
	}
	return nil, err
}

// GetAllAdministrador: 
func GetAllAdministrador(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	// <-- LO ÚNICO AGREGADO AQUÍ (.RelatedSel())
	qs := o.QueryTable(new(Administrador)).RelatedSel() 
    
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

	var l []Administrador
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

func UpdateAdministradorById(m *Administrador) (err error) {
	o := orm.NewOrm()
	v := Administrador{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteAdministrador(id int) (err error) {
	o := orm.NewOrm()
	v := Administrador{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Administrador{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}