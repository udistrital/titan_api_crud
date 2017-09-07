package controllers

import (
  	"encoding/json"
//	"errors"
	//"strconv"
	//"strings"
	"github.com/udistrital/titan_api_crud2/models"

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
