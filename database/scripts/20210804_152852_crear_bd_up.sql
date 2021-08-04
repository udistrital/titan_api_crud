
--Tipo nomina

CREATE TABLE titan_preliquidacion.tipo_nomina
(
    id serial NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    numero_orden numeric(5,2),
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_tipo_nomina PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.tipo_nomina
    IS 'Tabla parametrica que lista los tipos de nomina dentro de la Universidad Distrital. Ejemplo: Funcionarios de planta, docentes de planta o de HC';

-- Creacion tabla estado_preliquidacion


CREATE TABLE titan_preliquidacion.estado_preliquidacion
(
    id serial NOT NULL,
    nombre character varying(30) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    numero_orden numeric(5,2),
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_estado_preliquidacion PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.estado_preliquidacion
    IS 'Tabla que parametriza los diferentes estados que tiene una preliquidacion. Ejemplo: Si está abierta, está cerrada o en solicitud de necesidad';

-- Creacion de tabla nomina 

CREATE TABLE titan_preliquidacion.nomina
(
    id serial NOT NULL,
    descripcion character varying(50),
    activo boolean NOT NULL DEFAULT TRUE,
    fecha_creacion timestamp without time zone NOT NULL,
    fecha_modificacion timestamp without time zone NOT NULL,
    tipo_nomina_id integer,
    CONSTRAINT pk_nomina PRIMARY KEY (id),
    CONSTRAINT uq_tipo_nomina_nomina UNIQUE (tipo_nomina_id),
    CONSTRAINT fk_nomina_tipo_nomina FOREIGN KEY (tipo_nomina_id)
        REFERENCES titan_preliquidacion.tipo_nomina (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

COMMENT ON TABLE titan_preliquidacion.nomina
    IS 'Tabla que contiene las diferentes nominas presentes de la Universidad Distrital y sobre las cuales se realizaran calculos de preliquidacion';

COMMENT ON COLUMN titan_preliquidacion.nomina.descripcion
    IS 'Nombre de la nomina, formado desde la aplicacion utilizando el tipo de vinculacion y de nomina';

COMMENT ON COLUMN titan_preliquidacion.nomina.activo
    IS 'Describe si la nomina se encuentra activa o no';

COMMENT ON COLUMN titan_preliquidacion.nomina.tipo_nomina_id
    IS 'Llave foranea relacionada a la tabla tipo_nomina';
    
--tabla preliquidacion

CREATE TABLE titan_preliquidacion.preliquidacion
(
    id serial NOT NULL,
    descripcion character varying(100),
    mes integer NOT NULL,
    ano integer NOT NULL,
    estado_preliquidacion_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT TRUE,
    nomina_id integer NOT NULL,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_preliquidacion PRIMARY KEY (id),
    CONSTRAINT uq_periodo_peliquidacion UNIQUE (mes, ano, nomina_id),
    CONSTRAINT fk_preliquidacion_estado_preliquidacion FOREIGN KEY (estado_preliquidacion_id)
        REFERENCES titan_preliquidacion.estado_preliquidacion (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_preliquidacion_nomina FOREIGN KEY (nomina_id)
        REFERENCES titan_preliquidacion.nomina (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

COMMENT ON TABLE titan_preliquidacion.preliquidacion
    IS 'Tabla que detalla el mes y el año para el cual se realizaran calculos de pagos a las personas vinculadas contractualmente a la Universidad Distrital';

COMMENT ON COLUMN titan_preliquidacion.preliquidacion.descripcion
    IS 'Campo que describe la preliquidacion, creado a partir de aplicacion';

COMMENT ON COLUMN titan_preliquidacion.preliquidacion.mes
    IS 'Mes al que corresponde la preliquidación';

COMMENT ON COLUMN titan_preliquidacion.preliquidacion.ano
    IS 'Año al que corresponde la preliquidación';

COMMENT ON COLUMN titan_preliquidacion.preliquidacion.fecha_creacion
    IS 'Fecha en la que se realizó la preliquidación';

COMMENT ON COLUMN titan_preliquidacion.preliquidacion.estado_preliquidacion_id
    IS 'Llave foránea a estado_preliquidación';

COMMENT ON COLUMN titan_preliquidacion.preliquidacion.nomina_id
    IS 'Llave foránea a nómina';
    
--Tabla tipo_preliquidacion

CREATE TABLE titan_preliquidacion.tipo_preliquidacion
(
    id serial NOT NULL,
    nombre character varying(30) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    activo boolean NOT NULL DEFAULT TRUE,
    numero_orden numeric(5, 2),
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_tipo_preliquidacion PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.tipo_preliquidacion
    IS 'Corresponde al periodo a liquidar.
0 es la primera quincena, 1 la segunda quincena, 2 el mes completo, 3 junio y 4 diciembre';

-- Crear tabla estado disponibilidad

-- DROP TABLE administrativa.estado_disponibilidad;

CREATE TABLE titan_preliquidacion.estado_disponibilidad
(
    id serial NOT NULL
    nombre character varying(50) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    activo boolean NOT NULL DEFAULT TRUE,
    numero_orden numeric(5,2),
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_estado_disponibilidad PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.estado_disponibilidad
    IS 'Indica si el dinero para pagar dicho concepto se encuentra disponible, según verificación a nivel de api.';
    
-- Tabla naturaleza_concepto_nomina

CREATE TABLE titan_preliquidacion.naturaleza_concepto_nomina
(
    id serial NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    activo boolean NOT NULL DEFAULT TRUE,
    numero_orden numeric(5, 2),
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_naturaleza_concepto_nomina PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.naturaleza_concepto_nomina
    IS 'Describe si el concepto es un devengo o un descuento, o si hace parte de seguridad social.';
    
-- Tabla estado_concepto_nomina

CREATE TABLE titan_preliquidacion.estado_concepto_nomina
(
    id serial NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    activo boolean NOT NULL DEFAULT TRUE,
    numero_orden numeric(5, 2),
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_estado_concepto_nomina PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.estado_concepto_nomina
    IS 'Describe si el concepto está activo o inactivo.';
    
-- Tabla tipo_concepto_nomina

CREATE TABLE titan_preliquidacion.tipo_concepto_nomina
(
    id serial NOT NULL,
    nombre character varying(50) NOT NULL,
    descripcion character varying(100),
    codigo_abreviacion character varying(20),
    activo boolean NOT NULL DEFAULT TRUE,
    numero_orden numeric(5, 2),
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_tipo_concepto_nomina PRIMARY KEY (id)
);

COMMENT ON TABLE titan_preliquidacion.tipo_concepto_nomina
    IS 'Describe si el concepto a la hora de ser calculado corresponde a un valor fijo o porcentual.';

-- Tabla concepto_nomina

CREATE TABLE titan_preliquidacion.concepto_nomina
(
    id serial NOT NULL,
    nombre_concepto character varying(80) NOT NULL,
    alias_concepto character varying(80),
    naturaleza_concepto_nomina_id integer NOT NULL,
    tipo_concepto_nomina_id integer NOT NULL,
    estado_concepto_nomina_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT TRUE,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_concepto_nomina PRIMARY KEY (id),
    CONSTRAINT fk_concepto_nomina_naturaleza_concepto_nomina FOREIGN KEY (naturaleza_concepto_nomina_id)
        REFERENCES titan_preliquidacion.naturaleza_concepto_nomina (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_concepto_nomina_tipo_concepto_nomina FOREIGN KEY (tipo_concepto_nomina_id)
        REFERENCES titan_preliquidacion.tipo_concepto_nomina (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_concepto_nomina_estado_concepto_nomina FOREIGN KEY (estado_concepto_nomina_id)
        REFERENCES titan_preliquidacion.estado_concepto_nomina (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

COMMENT ON TABLE titan_preliquidacion.concepto_nomina
    IS 'Describen a que corresponden los distintos pagos calculados por la nomina';

COMMENT ON COLUMN titan_preliquidacion.concepto_nomina.nombre_concepto
    IS 'Nombre de la regla asociada al concepto';

COMMENT ON COLUMN titan_preliquidacion.concepto_nomina.alias_concepto
    IS 'Nombre del concepto a mostrar en la interfaz';

COMMENT ON COLUMN titan_preliquidacion.concepto_nomina.naturaleza_concepto_nomina_id
    IS 'Llave foranea a naturaleza_concepto_nomina';

COMMENT ON COLUMN titan_preliquidacion.concepto_nomina.tipo_concepto_nomina_id
    IS 'Llave foraena a tabla tipo_concepto_nomina';

COMMENT ON COLUMN titan_preliquidacion.concepto_nomina.estado_concepto_nomina_id
    IS 'Llave foraena a tabla estado_concepto_nomina';
    
-- Tabla detalle preliquidacion

CREATE TABLE titan_preliquidacion.detalle_preliquidacion
(
    id serial NOT NULL,
    valor_calculado numeric(20, 7) NOT NULL,
    numero_contrato character varying NOT NULL,
    vigencia_contrato integer NOT NULL,
    dias_liquidados numeric(2, 0) NOT NULL,
    tipo_preliquidacion_id integer NOT NULL,
    preliquidacion_id integer NOT NULL,
    concepto_nomina_id integer NOT NULL,
    estado_disponibilidad_id integer NOT NULL,
    persona_id integer NOT NULL,
    dependencia_id integer,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    activo boolean NOT NULL DEFAULT TRUE,
    CONSTRAINT pk_detalle_preliquidacion PRIMARY KEY (id),
    CONSTRAINT fk_detalle_preliquidacion_preliquidacion FOREIGN KEY (preliquidacion_id)
        REFERENCES titan_preliquidacion.preliquidacion (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_detalle_preliquidacion_concepto_nomina FOREIGN KEY (concepto_nomina_id)
        REFERENCES titan_preliquidacion.concepto_nomina (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_detalle_preliquidacion_estado_disponibilidad FOREIGN KEY (estado_disponibilidad_id)
        REFERENCES titan_preliquidacion.estado_disponibilidad (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_detalle_preliquidacion_tipo_preliquidacion FOREIGN KEY (tipo_preliquidacion_id)
        REFERENCES titan_preliquidacion.tipo_preliquidacion (id) MATCH FULL
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

COMMENT ON TABLE titan_preliquidacion.detalle_preliquidacion
    IS 'Tabla que detalla los pagos realizados a las personas por preliquidacion';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.valor_calculado
    IS 'Valor pagado a la persona, calculado por reglas de negocio';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.numero_contrato
    IS 'Numero de contrato de persona a la que se le realiza el pago';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.vigencia_contrato
    IS 'igencia del contrato de persona a la que se le realiza el pago';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.dias_liquidados
    IS 'Dias bajo los que fueron calculados los conceptos a la persona';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.tipo_preliquidacion_id
    IS 'Llave foranea a tipo de preliquidacion. Especifica el tipo de preliquidacion para el que corresponde el pago del concepto';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.preliquidacion_id
    IS 'Llave foranea a preliquidacion. Indica a que preliquidacion pertenece cada pago';

COMMENT ON COLUMN titan_preliquidacion.detalle_preliquidacion.concepto_nomina_id
    IS 'Llave foranea a concepto_nomina. Indica bajo que concepto se realiza el pago';