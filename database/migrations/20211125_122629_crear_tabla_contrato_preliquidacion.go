package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaContratoPreliquidacion_20211125_122629 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaContratoPreliquidacion_20211125_122629{}
	m.Created = "20211125_122629"

	migration.Register("CrearTablaContratoPreliquidacion_20211125_122629", m)
}

// Run the migrations
func (m *CrearTablaContratoPreliquidacion_20211125_122629) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS liquidador.contrato_preliquidacion(id serial NOT NULL,contrato_id integer NOT NULL,preliquidacion_id integer NOT NULL,cumplido boolean DEFAULT false,preliquidado boolean DEFAULT false,responsable_iva boolean,dependientes boolean,pensionado boolean,intereses_vivienda numeric(20,7),afc numeric(20,7),medicina_prepagada_uvt numeric(20,7),pension_voluntaria numeric(20,7),activo boolean NOT NULL DEFAULT true,fecha_creacion timestamp without time zone NOT NULL,fecha_modificacion timestamp without time zone NOT NULL,CONSTRAINT pk_contrato_preliquidacion PRIMARY KEY (id),CONSTRAINT uq_contrato_id_preliquidacion_id UNIQUE (contrato_id, preliquidacion_id),CONSTRAINT fk_contrato_preliquidacion_contrato FOREIGN KEY (contrato_id)REFERENCES liquidador.contrato (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION,CONSTRAINT fk_contrato_preliquidacion_preliquidacion FOREIGN KEY (preliquidacion_id)REFERENCES liquidador.preliquidacion (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION);")
	m.SQL("ALTER TABLE liquidador.contrato_preliquidacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE liquidador.contrato_preliquidacion IS 'Tabla que almacena los datos específicos que se tuvieron en cuenta para un contrato en la preliquidación, como los alivios y el cumplido';")
	m.SQL("COMMENT ON COLUMN liquidador.contrato_preliquidacion.id IS 'Identificador unico de la tabla contrato_preliquidacion.';")
}

// Reverse the migrations
func (m *CrearTablaContratoPreliquidacion_20211125_122629) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS liquidador.contrato_preliquidacion")
}
