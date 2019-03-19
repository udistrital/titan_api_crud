package controllers

import (
  	"encoding/json"
//	"errors"
	//"strconv"
	//"strings"
	"github.com/udistrital/titan_api_crud/models"

	"github.com/astaxie/beego"
  "fmt"
)

// ConceptoController operations for Concepto
type FuncionarioCargoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionarioCargoController) URLMapping() {
	c.Mapping("FuncionarioCargo", c.ConsultarAsignacionBasicaFuncionario)

}

// FuncionarioCargoController ...
// @Title ConsultarAsignacionBasicaFuncionario
// @Description create Consultar asignación básica según cargo
// @Param	  body		body 	models..Funcionario_x_Cargo	true		"body for models..Funcionario_x_Cargo"
// @Success 201 {object} []Docente_x_Cargo
// @Failure 403 body is empty
// @router /get_asignacion_basica [post]
func (c *FuncionarioCargoController) ConsultarAsignacionBasicaFuncionario() {
  var v models.Funcionario_x_Cargo
    fmt.Println("vvvvv111v")
  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
    fmt.Println("vvvvvv", v.Id)
    respuesta := models.GetAsignacionBasica(v.Id)
    c.Ctx.Output.SetStatus(201)

    c.Data["json"] = respuesta


  } else {
    c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
  }
  c.ServeJSON()
  }
