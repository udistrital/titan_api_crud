package controllers

import (
	"encoding/json"
	//	"errors"
	//"strconv"
	//"strings"
	"github.com/udistrital/titan_api_crud/models"

	"fmt"

	"github.com/astaxie/beego"
)

// ConceptoController operations for Concepto
type FuncionarioProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionarioProveedorController) URLMapping() {
	c.Mapping("FuncionarioProveedor", c.ConsultarIDProveedor)

}

func (c *FuncionarioProveedorController) ConsultarIDProveedor() {

	var v models.Nomina
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {


	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()
}
