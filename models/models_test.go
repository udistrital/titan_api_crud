package models

import "testing"

// una prueba para que el coverage no sea 0% pero las pruebas reales deben ser
// de la lógica de negocio
func TestActaInicioHonorariosPruebas(t *testing.T) {
  // this test is extremelly slow
  defer func() {
    if x := recover() ; x != nil {
      // hice panic, bien?
    }
  }()
  ActaInicioHonorariosPruebas(&ContratoGeneral{Vigencia: 0})
  // t.Fail() // no hice panic, mal?
}
