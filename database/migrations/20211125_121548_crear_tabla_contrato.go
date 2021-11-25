package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaContrato_20211125_121548 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaContrato_20211125_121548{}
	m.Created = "20211125_121548"

	migration.Register("CrearTablaContrato_20211125_121548", m)
}

// Run the migrations
func (m *CrearTablaContrato_20211125_121548) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS liquidador.contrato(id serial NOT NULL,numero_contrato character varying(20) NOT NULL,nombre_completo character varying(100) NOT NULL,documento character varying NOT NULL,persona_id integer,tipo_nomina_id integer NOT NULL,fecha_inicio timestamp without time zone NOT NULL,fecha_fin timestamp without time zone NOT NULL,valor_contrato numeric(20,7) NOT NULL,dependencia_id integer,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp without time zone NOT NULL,fecha_modificacion timestamp without time zone NOT NULL,CONSTRAINT pk_contrato PRIMARY KEY (id));")
	m.SQL("ALTER TABLE liquidador.contrato OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE liquidador.contrato IS 'Tabla que guarda los datos relacionados con la contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.id IS 'Identificador unico de la tabla contrato.';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.numero_contrato IS 'Número del contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.vigencia IS 'Vigencia del contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.nombre_completo IS 'Nombre completo de la persona asociada al contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.fecha_inicio IS 'Fecha inicio de contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.fecha_fin IS 'Fecha fin del contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.documento IS 'Documento del proveedor';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.valor_contrato IS 'Valor del contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.dependencia_id IS 'Indica la dependencia del contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaContrato_20211125_121548) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS liquidador.contrato")
}
