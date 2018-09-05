package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/titan_api_crud/models"
)

// ConceptoNominaPorPersonaController operations for ConceptoNominaPorPersona
type ConceptoNominaPorPersonaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ConceptoNominaPorPersonaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

}

// Post ...
// @Title Post
// @Description create ConceptoNominaPorPersona
// @Param	body		body 	models.ConceptoNominaPorPersona	true		"body for ConceptoNominaPorPersona content"
// @Success 201 {int} models.ConceptoNominaPorPersona
// @Failure 403 body is empty
// @router / [post]
func (c *ConceptoNominaPorPersonaController) Post() {
	var v models.ConceptoNominaPorPersona
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddConceptoNominaPorPersona(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ConceptoNominaPorPersona by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ConceptoNominaPorPersona
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ConceptoNominaPorPersonaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetConceptoNominaPorPersonaById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ConceptoNominaPorPersona
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ConceptoNominaPorPersona
// @Failure 403
// @router / [get]
func (c *ConceptoNominaPorPersonaController) GetAll() {
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
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllConceptoNominaPorPersona(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ConceptoNominaPorPersona
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ConceptoNominaPorPersona	true		"body for ConceptoNominaPorPersona content"
// @Success 200 {object} models.ConceptoNominaPorPersona
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ConceptoNominaPorPersonaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ConceptoNominaPorPersona{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateConceptoNominaPorPersonaById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ConceptoNominaPorPersona
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ConceptoNominaPorPersonaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteConceptoNominaPorPersona(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post ...
// @Title TrConceptosPorPersona
// @Description create ConceptoNominaPorPersona
// @Param	body		body 	models.TrConceptosNomPersona	true		"body for TrConceptosNomPersona content"
// @Success 201 {alert} models.Alerta
// @Failure 403 body is empty
// @router /TrConceptosPorPersona [post]
func (c *ConceptoNominaPorPersonaController) TrConceptosPorPersona() {
	var v models.TrConceptosNomPersona
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		beego.Info(v)
		if alerta, err := models.RegistrarConceptos(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
