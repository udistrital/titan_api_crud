CREATE SCHEMA liquidador

CREATE TABLE IF NOT EXISTS liquidador.preliquidacion
(
    id serial,
    descripcion character varying(100),
    mes integer NOT NULL,
    ano integer NOT NULL,
    estado_preliquidacion_id integer NOT NULL,
    nomina_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_preliquidacion PRIMARY KEY (id),
    CONSTRAINT uq_periodo_peliquidacion UNIQUE (mes, ano, nomina_id)
);

CREATE TABLE IF NOT EXISTS liquidador.contrato
(
    id serial,
    numero_contrato character varying(20) NOT NULL,
    vigencia integer NOT NULL,
    nombre_completo character varying(100) NOT NULL,
    documento character varying NOT NULL,
    persona_id integer,
    tipo_nomina_id integer NOT NULL,
    fecha_inicio timestamp with time zone NOT NULL,
    fecha_fin timestamp with time zone NOT NULL,
    valor_contrato numeric(20,7) NOT NULL,
    dependencia_id integer,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_contrato PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS liquidador.contrato_preliquidacion
(
    id serial,
    contrato_id integer NOT NULL,
    preliquidacion_id integer NOT NULL,
    cumplido boolean DEFAULT false,
    preliquidado boolean DEFAULT false,
    responsable_iva boolean,
    dependientes boolean,
    intereses_vivienda numeric(20,7),
    afc numeric(20,7),
    medicina_prepagada_uvt numeric(20,7),
    pension_voluntaria numeric(20,7),
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_contrato_preliquidacion PRIMARY KEY (id),
    CONSTRAINT uq_contrato_id_preliquidacion_id UNIQUE (contrato_id, preliquidacion_id),
    CONSTRAINT fk_contrato_preliquidacion_contrato FOREIGN KEY (contrato_id)
        REFERENCES liquidador.contrato (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_contrato_preliquidacion_preliquidacion FOREIGN KEY (preliquidacion_id)
        REFERENCES liquidador.preliquidacion (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS liquidador.concepto_nomina
(
    id serial NOT NULL,
    nombre_concepto character varying(80) NOT NULL,
    alias_concepto character varying(80),
    naturaleza_concepto_nomina_id integer NOT NULL,
    tipo_concepto_nomina_id integer NOT NULL,
    estado_concepto_nomina_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT TRUE,
    fecha_creacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_concepto_nomina PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS liquidador.detalle_preliquidacion
(
    id serial,
    contrato_preliquidacion_id integer NOT NULL,
    valor_calculado numeric(20,7) NOT NULL,
    dias_liquidados numeric(2,0) NOT NULL,
    dias_especificos character varying(50),
    tipo_preliquidacion_id integer NOT NULL,
    concepto_nomina_id integer NOT NULL,
    estado_disponibilidad_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_detalle_preliquidacion PRIMARY KEY (id),
    CONSTRAINT fk_detalle_preliquidacion_concepto_nomina FOREIGN KEY (concepto_nomina_id)
        REFERENCES liquidador.concepto_nomina (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_detalle_preliquidacion_cotrato_preliquidacion FOREIGN KEY (contrato_preliquidacion_id)
        REFERENCES liquidador.contrato_preliquidacion (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


