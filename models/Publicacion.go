package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Publicacion struct {
	Id                    int                `orm:"column(id_publicacion);pk;auto"`
	IdUsuario             int                `orm:"column(id_usuario)"`
	IdEstado              *EstadoPublicacion `orm:"column(id_estado);rel(fk)"`
	Titulo                string             `orm:"column(titulo)"`
	Descripcion           string             `orm:"column(descripcion)"`
	Tipo                  string             `orm:"column(tipo)"`
	Visibilidad           string             `orm:"column(visibilidad);null"`
	Departamento          string             `orm:"column(departamento)"`
	Municipio             string             `orm:"column(municipio)"`
	Barrio                string             `orm:"column(barrio);null"`
	DisponibleParaTrueque bool               `orm:"column(disponible_para_trueque);null"`
	CantidadDisponible    int                `orm:"column(cantidad_disponible);null"`
	Prioridad             int                `orm:"column(prioridad);null"`
	MensajeContacto       string             `orm:"column(mensaje_contacto);null"`
	Vistas                int                `orm:"column(vistas);null"`
	Favoritos             int                `orm:"column(favoritos);null"`
	Activo                bool               `orm:"column(activo)"`
	FechaCreacion         time.Time          `orm:"column(fecha_creacion);type(timestamp without time zone);null;auto_now_add"`
	FechaModificacion     time.Time          `orm:"column(fecha_modificacion);type(timestamp without time zone);null;auto_now"`
}

func (t *Publicacion) TableName() string {
	return "Publicacion"
}

func init() {
	orm.RegisterModel(new(Publicacion))
}

// AddPublicacion insert a new Publicacion into database and returns
// last inserted Id on success.
func AddPublicacion(m *Publicacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPublicacionById retrieves Publicacion by Id. Returns error if
// Id doesn't exist
func GetPublicacionById(id int) (v *Publicacion, err error) {
	o := orm.NewOrm()
	v = &Publicacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPublicacion retrieves all Publicacion matches certain condition. Returns empty list if
// no records exist
func GetAllPublicacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Publicacion))
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

	var l []Publicacion
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

// UpdatePublicacion updates Publicacion by Id and returns error if
// the record to be updated doesn't exist
func UpdatePublicacionById(m *Publicacion) (err error) {
	o := orm.NewOrm()
	v := Publicacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePublicacion deletes Publicacion by Id and returns error if
// the record to be deleted doesn't exist
func DeletePublicacion(id int) (err error) {
	o := orm.NewOrm()
	v := Publicacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Publicacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
