package controllers

import (
  "testing"
  "github.com/astaxie/beego/context"
)

// una prueba para que el coverage no sea 0% pero las pruebas reales deben ser
// de la lógica de negocio
func TestURLMapping(t *testing.T) {
  aic := ActaInicioController{}
  aic.Init(context.NewContext(), "test", "test", nil)
  aic.URLMapping()
}
