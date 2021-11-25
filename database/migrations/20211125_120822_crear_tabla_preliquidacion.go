package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPreliquidacion_20211125_120822 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPreliquidacion_20211125_120822{}
	m.Created = "20211125_120822"

	migration.Register("CrearTablaPreliquidacion_20211125_120822", m)
}

// Run the migrations
func (m *CrearTablaPreliquidacion_20211125_120822) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS liquidador.preliquidacion(id serial NOT NULL,descripcion character varying(100),mes integer NOT NULL,ano integer NOT NULL,estado_preliquidacion_id integer NOT NULL,nomina_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp without time zone NOT NULL,fecha_modificacion timestamp without time zone NOT NULL,CONSTRAINT pk_preliquidacion PRIMARY KEY (id),CONSTRAINT uq_periodo_peliquidacion UNIQUE (mes, ano, nomina_id));")
	m.SQL("ALTER TABLE liquidador.preliquidacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE liquidador.preliquidacion IS 'Tabla que guarda los datos relacionados con la preliquidacion, como mes y año';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.id IS 'Identificador unico de la tabla preliquidacion.';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del preliquidacion.';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.ano IS 'Mes de la preliquidación';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.mes IS 'Año de la preliquidación';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.nomina_id IS 'nomina a la que pertenece la preliquidación';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN liquidador.preliquidacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaPreliquidacion_20211125_120822) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS liquidador.preliquidacion")

}
