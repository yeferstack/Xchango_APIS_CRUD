package controllers

import (
	"Xchango_APIS_CRUD/models"
	"encoding/json"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// IntentoLoginController operations for IntentoLogin
type IntentoLoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *IntentoLoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create IntentoLogin
// @Param	body		body 	models.IntentoLogin	true		"body for IntentoLogin content"
// @Success 201 {int} models.IntentoLogin
// @Failure 403 body is empty
// @router / [post]
func (c *IntentoLoginController) Post() {

	var v models.IntentoLogin

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if _, err := models.AddIntentoLogin(&v); err == nil {

			c.Ctx.Output.SetStatus(201)

			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  201,
				"Message": "Registro creado exitosamente",
				"data":    v,
			}

		} else {

			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error en el servicio Post",
				"error":   err.Error(),
			}
		}

	} else {

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error al leer el body",
			"error":   err.Error(),
		}
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get IntentoLogin by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.IntentoLogin
// @Failure 403 :id is empty
// @router /:id [get]
func (c *IntentoLoginController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetIntentoLoginById(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servicio GetOne: La solicitud contiene un parametro incorrecto o no existe ningun registro"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get IntentoLogin
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.IntentoLogin
// @Failure 403
// @router / [get]
func (c *IntentoLoginController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"status":  400,
					"Message": "Error en el servicio GetAll",
					"error":   "El parámetro query no tiene el formato correcto, debe ser clave:valor",
				}  // <- falta este cierre del map
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllIntentoLogin(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servicio GetAll: La solicitud contiene un parametro incorrecto o no existe ningun registro"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the IntentoLogin
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.IntentoLogin	true		"body for IntentoLogin content"
// @Success 200 {object} models.IntentoLogin
// @Failure 403 :id is not int
// @router /:id [put]
func (c *IntentoLoginController) Put() {

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	v := models.IntentoLogin{Id: id}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if err := models.UpdateIntentoLoginById(&v); err == nil {

			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"Message": "Registro actualizado exitosamente",
				"data":    v,
			}

		} else {

			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"Message": "Error en el servicio Put",
				"error":   err.Error(),
			}
		}

	} else {

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error al leer el body",
			"error":   err.Error(),
		}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the IntentoLogin
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *IntentoLoginController) Delete() {

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if err := models.DeleteIntentoLogin(id); err == nil {

		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"Message": "Registro eliminado exitosamente",
		}

	} else {

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"Message": "Error en el servicio Delete",
			"error":   err.Error(),
		}
	}

	c.ServeJSON()
}
