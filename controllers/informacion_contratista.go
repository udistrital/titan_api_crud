package controllers

import (
	"encoding/json"
	"fmt"
	"titan_api_crud/models"

	"github.com/astaxie/beego"
)

// operations for DocenteCargo
type InformacionContratistaController struct {
	beego.Controller
}

func (c *InformacionContratistaController) URLMapping() {
	c.Mapping("Contratista_datos", c.Contratista_datos)
}

func (c *InformacionContratistaController) Contratista_datos() {
	var v models.Contratista
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		respuesta := models.GetContratista(v.Id)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()

}
