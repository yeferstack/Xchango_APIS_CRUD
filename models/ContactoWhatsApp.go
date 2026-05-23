package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ContactoWhatsApp struct {
	Id                   int       `orm:"column(id_contacto);pk;auto"`
	IdPublicacion        int       `orm:"column(id_publicacion)"`
	IdUsuarioInteresado  int       `orm:"column(id_usuario_interesado)"`
	IdUsuarioPropietario int       `orm:"column(id_usuario_propietario)"`
	Activo               bool      `orm:"column(activo)"`
	FechaCreacion        time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion    time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *ContactoWhatsApp) TableName() string {
	return "ContactoWhatsApp"
}

func init() {
	orm.RegisterModel(new(ContactoWhatsApp))
}

// AddContactoWhatsApp insert a new ContactoWhatsApp into database and returns
// last inserted Id on success.
func AddContactoWhatsApp(m *ContactoWhatsApp) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContactoWhatsAppById retrieves ContactoWhatsApp by Id. Returns error if
// Id doesn't exist
func GetContactoWhatsAppById(id int) (v *ContactoWhatsApp, err error) {
	o := orm.NewOrm()
	v = &ContactoWhatsApp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContactoWhatsApp retrieves all ContactoWhatsApp matches certain condition. Returns empty list if
// no records exist
func GetAllContactoWhatsApp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ContactoWhatsApp))
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

	var l []ContactoWhatsApp
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

// UpdateContactoWhatsAppById updates ContactoWhatsApp by Id and returns error if
// the record to be updated doesn't exist
func UpdateContactoWhatsAppById(m *ContactoWhatsApp) (err error) {
	o := orm.NewOrm()
	v := ContactoWhatsApp{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContactoWhatsApp deletes ContactoWhatsApp by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContactoWhatsApp(id int) (err error) {
	o := orm.NewOrm()
	v := ContactoWhatsApp{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ContactoWhatsApp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
