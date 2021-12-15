package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDetallePreliquidacion_20211125_124020 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDetallePreliquidacion_20211125_124020{}
	m.Created = "20211125_124020"

	migration.Register("CrearTablaDetallePreliquidacion_20211125_124020", m)
}

// Run the migrations
func (m *CrearTablaDetallePreliquidacion_20211125_124020) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS liquidador.detalle_preliquidacion(id serial NOT NULL,contrato_preliquidacion_id integer NOT NULL,valor_calculado numeric(20,7) NOT NULL,dias_liquidados numeric(2,0) NOT NULL,dias_especificos character varying(50),tipo_preliquidacion_id integer NOT NULL,concepto_nomina_id integer NOT NULL,estado_disponibilidad_id integer NOT NULL,activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp without time zone NOT NULL,fecha_modificacion timestamp without time zone NOT NULL,CONSTRAINT pk_detalle_preliquidacion PRIMARY KEY (id),CONSTRAINT fk_detalle_preliquidacion_concepto_nomina FOREIGN KEY (concepto_nomina_id) REFERENCES liquidador.concepto_nomina (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION,CONSTRAINT fk_detalle_preliquidacion_cotrato_preliquidacion FOREIGN KEY (contrato_preliquidacion_id) REFERENCES liquidador.contrato_preliquidacion (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION);")
	m.SQL("ALTER TABLE liquidador.detalle_preliquidacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE liquidador.detalle_preliquidacion IS 'Tabla que almacena los valores de cada concepto que se paga a acada contrato en cada preliquidacion';")
	m.SQL("COMMENT ON COLUMN liquidador.detalle_preliquidacion.id IS 'Identificador de la tabla detalle_preliquidacion';")
	m.SQL("COMMENT ON COLUMN liquidador.detalle_preliquidacion.concepto_nomina_id IS 'Asocia el detalle al concepto al que pertenece';")
	m.SQL("COMMENT ON COLUMN liquidador.detalle_preliquidacion.valor_calculado IS 'El valor calculado para el concepto';")
}

// Reverse the migrations
func (m *CrearTablaDetallePreliquidacion_20211125_124020) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS liquidador.detalle_preliquidacion;")
}
