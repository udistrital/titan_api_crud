package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNovedad_20211125_134239 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNovedad_20211125_134239{}
	m.Created = "20211125_134239"

	migration.Register("CrearTablaNovedad_20211125_134239", m)
}

// Run the migrations
func (m *CrearTablaNovedad_20211125_134239) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS liquidador.novedad(id serial NOT NULL,contrato_id integer NOT NULL,concepto_nomina_id integer NOT NULL,valor numeric(20,7) NOT NULL,cuotas integer NOT NULL,fecha_inicio timestamp with time zone NOT NULL,fecha_fin timestamp with time zone NOT NULL,activo boolean NOT NULL,fecha_creacion timestamp without time zone NOT NULL,fecha_modificacion timestamp without time zone NOT NULL,CONSTRAINT pk_novedad PRIMARY KEY (id),CONSTRAINT fk_novedad_concepto_nomina FOREIGN KEY (concepto_nomina_id)REFERENCES liquidador.concepto_nomina (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION,CONSTRAINT fk_novedad_contrato FOREIGN KEY (contrato_id)REFERENCES liquidador.contrato (id) MATCH FULL ON UPDATE NO ACTION ON DELETE NO ACTION);")
	m.SQL("ALTER TABLE liquidador.novedad OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE liquidador.novedad IS 'Tabla que almacena las novedades registradas en cada contrato';")
	m.SQL("COMMENT ON COLUMN liquidador.novedad.id IS 'Identificador de la tabla novedad';")
	m.SQL("COMMENT ON COLUMN liquidador.novedad.valor IS 'Almacena el valor de la novedad, si es fijo almacena el valor mensual de la cuota, si es fijo almacena el porcentaje correspondiente';")
	m.SQL("COMMENT ON COLUMN liquidador.novedad.cuotas IS 'Almacena el númeor de cuotas de la novedad';")
	m.SQL("COMMENT ON COLUMN liquidador.novedad.concepto_nomina_id IS 'Id del concepto al cual está asociada la novedad';")

}

// Reverse the migrations
func (m *CrearTablaNovedad_20211125_134239) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS liquidador.novedad;")
}
