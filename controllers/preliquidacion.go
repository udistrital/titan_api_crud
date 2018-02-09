package controllers

import (
	"encoding/json"
	"errors"
	"github.com/udistrital/titan_api_crud/models"
	"strconv"
	"strings"
	"fmt"
	"github.com/astaxie/beego"
)

// PreliquidacionController operations for Preliquidacion
type PreliquidacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *PreliquidacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Resumen", c.Resumen)
	c.Mapping("Contratos_x_preliquidacion", c.Contratos_x_preliquidacion)
	c.Mapping("Contratos_x_preliquidacion_cerrada", c.Contratos_x_preliquidacion_cerrada)
	c.Mapping("Totales_ss_x_preliquidacion", c.Totales_ss_x_preliquidacion)

}

// Post ...
// @Title Post
// @Description create Preliquidacion
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 201 {int} models.Preliquidacion
// @Failure 403 body is empty
// @router / [post]
func (c *PreliquidacionController) Post() {
	var v models.Preliquidacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPreliquidacion(&v); err == nil {
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
// @Description get Preliquidacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Preliquidacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PreliquidacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPreliquidacionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Preliquidacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Preliquidacion
// @Failure 403
// @router / [get]
func (c *PreliquidacionController) GetAll() {
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

	l, err := models.GetAllPreliquidacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Preliquidacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 200 {object} models.Preliquidacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PreliquidacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Preliquidacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdatePreliquidacionById(&v); err == nil {
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
// @Description delete the Preliquidacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PreliquidacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePreliquidacion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// PreliquidacionController ...
// @Title Resumen
// @Description create Resumen
// @Param	  body		body 	models.Preliquidacion	true		"body for Preliquidacion content"
// @Success 201 {object} models.InformePreliquidacion
// @Failure 403 body is empty
// @router /resumen/ [post]
func (c *PreliquidacionController) Resumen() {

	var v models.Preliquidacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if res , err := models.ResumenPreliquidacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = res
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Contratos_x_preliquidacion ...
// @Title Contratos_x_preliquidacion
// @Description Agrupa los contratos de una preliquidacion segun mes, año y nomina para preliquidaicones en estado OP
// @Param idNomina query string false "nomina a listar"
// @Param mesLiquidacion query string false "mes de la liquidacion a listar"
// @Param anioLiquidacion query string false "anio de la liquidacion a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /contratos_x_preliquidacion [get]
func (c *PreliquidacionController) Contratos_x_preliquidacion() {

	idNomina, err1 := c.GetInt("idNomina")
	mesLiquidacion, err2 := c.GetInt("mesLiquidacion")
	anioLiquidacion, err3 := c.GetInt("anioLiquidacion")
	if err1 == nil && err2 == nil && err3 == nil {
	fmt.Println(idNomina)
	fmt.Println(mesLiquidacion)
	fmt.Println(anioLiquidacion)
		if res , err := models.Contratos_x_preliquidacion(idNomina, mesLiquidacion, anioLiquidacion); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = res
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err1)
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

// Totales_ss_x_preliquidacion ...
// @Title Totales_ss_x_preliquidacion
// @Description Retorna los totales para los descuentos de salud, pension y fondo de solidaridad segun el tipo de nomina, el mes y el año
// @Param idNomina query string false "nomina a listar"
// @Param mesLiquidacion query string false "mes de la liquidacion a listar"
// @Param anioLiquidacion query string false "anio de la liquidacion a listar"
// @Success 201 {object} models.Totales_x_preliq
// @Failure 403 body is empty
// @router /totales_x_preliq [get]
func (c *PreliquidacionController) Totales_ss_x_preliquidacion() {

	idNomina, err1 := c.GetInt("idNomina")
	mesLiquidacion, err2 := c.GetInt("mesLiquidacion")
	anioLiquidacion, err3 := c.GetInt("anioLiquidacion")
	if err1 == nil && err2 == nil && err3 == nil {
	fmt.Println(idNomina)
	fmt.Println(mesLiquidacion)
	fmt.Println(anioLiquidacion)
		if res , err := models.Totales_ss_x_preliquidacion(idNomina, mesLiquidacion, anioLiquidacion); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = res
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err1)
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

// Contratos_x_preliquidacion_cerrada ...
// @Title Contratos_x_preliquidacion_cerrada
// @Description Agrupa los contratos de una preliquidacion segun mes, año y nomina para preliquidaicones en estado CERRADA
// @Param idNomina query string false "nomina a listar"
// @Param mesLiquidacion query string false "mes de la liquidacion a listar"
// @Param anioLiquidacion query string false "anio de la liquidacion a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /contratos_x_preliquidacion_cerrada [get]
func (c *PreliquidacionController) Contratos_x_preliquidacion_cerrada() {

	idNomina, err1 := c.GetInt("idNomina")
	mesLiquidacion, err2 := c.GetInt("mesLiquidacion")
	anioLiquidacion, err3 := c.GetInt("anioLiquidacion")
	if err1 == nil && err2 == nil && err3 == nil {
	fmt.Println(idNomina)
	fmt.Println(mesLiquidacion)
	fmt.Println(anioLiquidacion)
		if res , err := models.Contratos_x_preliquidacion_cerrada(idNomina, mesLiquidacion, anioLiquidacion); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = res
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err1)
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}
