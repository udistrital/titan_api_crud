CREATE TABLE IF NOT EXISTS titan.preliquidacion
(
    id serial NOT NULL,
    descripcion character varying(100),
    mes integer NOT NULL,
    ano integer NOT NULL,
    estado_preliquidacion_id integer NOT NULL,
    nomina_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT TRUE,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_preliquidacion PRIMARY KEY (id),
    CONSTRAINT uq_periodo_peliquidacion UNIQUE (mes, ano, nomina_id)
);

CREATE TABLE IF NOT EXISTS titan.contrato
(
    id serial NOT NULL,
    numero_contrato character varying(20) COLLATE pg_catalog."default" NOT NULL,
    vigencia integer NOT NULL,
    nombre_completo character varying(100) COLLATE pg_catalog."default" NOT NULL,
    documento character varying COLLATE pg_catalog."default" NOT NULL,
    persona_id integer,
    cumplido boolean NOT NULL,
    preliquidado boolean NOT NULL,
    fecha_inicio timestamp with time zone NOT NULL,
    fecha_fin timestamp with time zone NOT NULL,
    valor_contrato numeric(20,7) NOT NULL,
    intereses_vivienda numeric(20,7) NOT NULL,
    medicina_prepagada_uvt numeric(20,7) NOT NULL,
    pension_voluntaria numeric(20,7) NOT NULL,
    responsable_iva boolean NOT NULL DEFAULT false,
    "AFC" numeric(20,7) NOT NULL,
    dependientes boolean NOT NULL DEFAULT false,
    dependencia_id integer,
    preliquidacion_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_contrato PRIMARY KEY (id),
    CONSTRAINT fk_contrato_preliquidacion FOREIGN KEY (preliquidacion_id)
        REFERENCES titan.preliquidacion (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS titan.detalle_preliquidacion
(
    id integer NOT NULL DEFAULT nextval('titan.detalle_preliquidacion_id_seq'::regclass),
    valor_calculado numeric(20,7) NOT NULL,
    contrato_id integer NOT NULL,
    numero_contrato character varying(20) COLLATE pg_catalog."default" NOT NULL,
    vigencia_contrato integer NOT NULL,
    dias_liquidados numeric(2,0) NOT NULL,
    dias_especificos character varying(50) COLLATE pg_catalog."default",
    tipo_preliquidacion_id integer NOT NULL,
    concepto_nomina_id integer NOT NULL,
    estado_disponibilidad_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_detalle_preliquidacion PRIMARY KEY (id),
    CONSTRAINT fk_detalle_preliquidacion_contrato FOREIGN KEY (contrato_id)
        REFERENCES titan.contrato (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
