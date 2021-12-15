package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaConceptoNomina_20211125_123113 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaConceptoNomina_20211125_123113{}
	m.Created = "20211125_123113"

	migration.Register("CrearTablaConceptoNomina_20211125_123113", m)
}

// Run the migrations
func (m *CrearTablaConceptoNomina_20211125_123113) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS liquidador.concepto_nomina(id serial NOT NULL,nombre_concepto character varying(80) NOT NULL,alias_concepto character varying(80),naturaleza_concepto_nomina_id integer NOT NULL,tipo_concepto_nomina_id integer NOT NULL,estado_concepto_nomina_id integer NOT NULL,activo boolean NOT NULL DEFAULT TRUE,fecha_creacion timestamp without time zone NOT NULL,fecha_modificacion timestamp without time zone NOT NULL,CONSTRAINT pk_concepto_nomina PRIMARY KEY (id));")
	m.SQL("ALTER TABLE liquidador.concepto_nomina OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE liquidador.concepto_nomina IS 'Tabla que almacena los distitnos conceptos por los cuales se paga al al empleado';")
	m.SQL("COMMENT ON COLUMN liquidador.concepto_nomina.id IS 'Identificador unico de la tabla concepto_nomina.';")
	m.SQL("COMMENT ON COLUMN liquidador.concepto_nomina.nombre_concepto IS 'Nombre del concepto';")
	m.SQL("COMMENT ON COLUMN liquidador.concepto_nomina.alias_concepto IS 'Alias del concepto';")
	m.SQL("COMMENT ON COLUMN liquidador.concepto_nomina.naturaleza_concepto_nomina_id IS 'Indica la naturaleza del concepto, i es descuento, devendo o de seguridad social';")
	m.SQL("COMMENT ON COLUMN liquidador.concepto_nomina.tipo_concepto_nomina_id IS 'Indica si es de valor fijo, porcental o de segurirdad social';")
	m.SQL("COMMENT ON COLUMN liquidador.concepto_nomina.estado_concepto_nomina_id IS 'Indica si el concepto se encuentra activo o inactivo';")
}

// Reverse the migrations
func (m *CrearTablaConceptoNomina_20211125_123113) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS liquidador.concepto_nomina")
}
