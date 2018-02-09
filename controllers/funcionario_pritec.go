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
type FuncionarioPritecController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionarioPritecController) URLMapping() {
	c.Mapping("FuncionarioPritec", c.ConsultarPrimaTecnica)

}

// FuncionarioPritecController ...
// @Title ConsultarPrimaTecnica
// @Description Consulta prima técnica del funcionario
// @Param	  body		body 	int	true		"body int"
// @Success 201 {object} int
// @Failure 403 body is empty
// @router / [post]
func (c *FuncionarioPritecController) ConsultarPrimaTecnica() {
  var v int
  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

    respuesta := models.GetPorcentajePT(v)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = respuesta


  } else {
    c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
  }
  c.ServeJSON()
  }
