package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/udistrital/titan_api_crud/models"

	"github.com/astaxie/beego"
)

// operations for DocenteCargo
type DocenteCargoController struct {
	beego.Controller
}

func (c *DocenteCargoController) URLMapping() {
	c.Mapping("DocenteCargo", c.ConsultarAsignacionBasica)
	c.Mapping("DocenteCedula", c.ConsultarCedulaDocente)
}

// DocenteCargoController ...
// @Title ConsultarAsignacionBasica
// @Description create Consultar asignación básica según cargo
// @Param	  body		body 	models.Docente_x_Cargo	true		"body for models.Docente_x_Cargo"
// @Success 201 {object} []Docente_x_Cargo
// @Failure 403 body is empty
// @router / [post]
func (c *DocenteCargoController) ConsultarAsignacionBasica() {
	var v models.Docente_x_Cargo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta := models.GetAsignacionBasicaDocente(v.Id)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()
}

// DocenteCargoController ...
// @Title ConsultarCedulaDocente
// @Description create Consultar cédula del docente
// @Param	  body		body 	string	true		"body string"
// @Success 201 {object} []CedulaDocente
// @Failure 403 body is empty
// @router /consultarCedulaDocente/ [post]
func (c *DocenteCargoController) ConsultarCedulaDocente() {
	var v string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta := models.GetCedulaDocente(v)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()

}
