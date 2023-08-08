-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.3
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file.
-- These commands were put in this file only as a convenience.
--
-- object: postgres | type: DATABASE --
-- DROP DATABASE IF EXISTS postgres;
CREATE DATABASE postgres
	ENCODING = 'UTF8'
	LC_COLLATE = 'en_US.UTF-8'
	LC_CTYPE = 'en_US.UTF-8'
	TABLESPACE = pg_default
	OWNER = postgres;
-- ddl-end --
COMMENT ON DATABASE postgres IS E'default administrative connection database';
-- ddl-end --


-- object: academica | type: SCHEMA --
-- DROP SCHEMA IF EXISTS academica CASCADE;
CREATE SCHEMA academica;
-- ddl-end --
ALTER SCHEMA academica OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,academica;
-- ddl-end --

-- object: academica.area_conocimiento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.area_conocimiento_id_seq CASCADE;
CREATE SEQUENCE academica.area_conocimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.area_conocimiento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.area_conocimiento | type: TABLE --
-- DROP TABLE IF EXISTS academica.area_conocimiento CASCADE;
CREATE TABLE academica.area_conocimiento (
	id integer NOT NULL DEFAULT nextval('academica.area_conocimiento_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	snies_area integer NOT NULL,
	CONSTRAINT pk_area_conocimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.area_conocimiento IS E'Entidad que almacena la información de las áreas de conocimiento';
-- ddl-end --
COMMENT ON COLUMN academica.area_conocimiento.id IS E'Identificador del área de conocimiento';
-- ddl-end --
COMMENT ON COLUMN academica.area_conocimiento.nombre IS E'Nombre que identifica a un área de conocimiento';
-- ddl-end --
COMMENT ON COLUMN academica.area_conocimiento.descripcion IS E'Descripción de un área de conocimiento';
-- ddl-end --
COMMENT ON COLUMN academica.area_conocimiento.codigo_abreviacion IS E'Abreviación del nombre del área de conocimiento';
-- ddl-end --
COMMENT ON COLUMN academica.area_conocimiento.activo IS E'Indica si el área de conocimiento se encuentra activa o no';
-- ddl-end --
COMMENT ON COLUMN academica.area_conocimiento.snies_area IS E'Area de SNIES a la que pertenece el área de conocimiento, referencia la tabla del core snies_area';
-- ddl-end --
COMMENT ON CONSTRAINT pk_area_conocimiento ON academica.area_conocimiento  IS E'Constraint primary key area_conocimiento';
-- ddl-end --
ALTER TABLE academica.area_conocimiento OWNER TO postgres;
-- ddl-end --

-- object: academica.areas_docente_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.areas_docente_id_seq CASCADE;
CREATE SEQUENCE academica.areas_docente_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.areas_docente_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.areas_docente | type: TABLE --
-- DROP TABLE IF EXISTS academica.areas_docente CASCADE;
CREATE TABLE academica.areas_docente (
	id integer NOT NULL DEFAULT nextval('academica.areas_docente_id_seq'::regclass),
	docente integer NOT NULL,
	area_conocimiento integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_areas_docente PRIMARY KEY (id),
	CONSTRAINT uq_docente_area_conocimiento UNIQUE (docente,area_conocimiento)

);
-- ddl-end --
COMMENT ON TABLE academica.areas_docente IS E'Entidad que almacena las áreas de conocimiento que maneja un docente';
-- ddl-end --
COMMENT ON COLUMN academica.areas_docente.id IS E'Identificador de la relación entre el área de conocimiento y el docente';
-- ddl-end --
COMMENT ON COLUMN academica.areas_docente.docente IS E'Hace referencia al docente relacionado con el área de conocimiento, se referencia de la tabla um_user del modelo de WSO2';
-- ddl-end --
COMMENT ON COLUMN academica.areas_docente.area_conocimiento IS E'Identificador de la tabla area_conocimiento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_areas_docente ON academica.areas_docente  IS E'Constraint primary key areas_docente';
-- ddl-end --
COMMENT ON CONSTRAINT uq_docente_area_conocimiento ON academica.areas_docente  IS E'Constraint unique uq_docente_area_conocimiento';
-- ddl-end --
ALTER TABLE academica.areas_docente OWNER TO postgres;
-- ddl-end --

-- object: academica.areas_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.areas_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.areas_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.areas_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.areas_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.areas_trabajo_grado CASCADE;
CREATE TABLE academica.areas_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.areas_trabajo_grado_id_seq'::regclass),
	area_conocimiento integer NOT NULL,
	trabajo_grado integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_areas_trabajo_grado PRIMARY KEY (id),
	CONSTRAINT uq_trabajo_grado_area_conocimiento UNIQUE (trabajo_grado,area_conocimiento)

);
-- ddl-end --
COMMENT ON TABLE academica.areas_trabajo_grado IS E'Tabla que relaciona las áreas de conocimiento con un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.areas_trabajo_grado.id IS E'Identificador de la relación entre un área de conocimiento y un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.areas_trabajo_grado.area_conocimiento IS E'Area de conocimiento que tiene relación con el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.areas_trabajo_grado.trabajo_grado IS E'Identificador del trabajo de grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_areas_trabajo_grado ON academica.areas_trabajo_grado  IS E'Constraint primary key areas_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT uq_trabajo_grado_area_conocimiento ON academica.areas_trabajo_grado  IS E'Constraint unique uq_trabajo_grado_area_conocimiento';
-- ddl-end --
ALTER TABLE academica.areas_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.asignatura_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.asignatura_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.asignatura_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.asignatura_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.asignatura_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.asignatura_trabajo_grado CASCADE;
CREATE TABLE academica.asignatura_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.asignatura_trabajo_grado_id_seq'::regclass),
	codigo_asignatura integer NOT NULL,
	periodo numeric(1,0) NOT NULL,
	anio numeric(4,0) NOT NULL,
	calificacion double precision,
	trabajo_grado integer NOT NULL,
	estado_asignatura_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_asignatura_inscrita PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.asignatura_trabajo_grado IS E'Guarda la asignatura inscrita para la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.id IS E'Identificador de la asignatura inscrita en el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.codigo_asignatura IS E'Código de la asignatura (TGI, TGII, TG Tecnologico), (academica: ACASI)';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.periodo IS E'Periodo en el cual se cursa la asignatura';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.anio IS E'Permite almaceñar el año en el cual se cursa la asignatura';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.calificacion IS E'Calificación de la asignatura inscrita';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.trabajo_grado IS E'Identificador del trabajo de grado relacionado con la asignatura';
-- ddl-end --
COMMENT ON COLUMN academica.asignatura_trabajo_grado.estado_asignatura_trabajo_grado IS E'Estado de la asignatura de trabajo de grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_asignatura_inscrita ON academica.asignatura_trabajo_grado  IS E'Constraint primary key asignatura_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.asignatura_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.carrera_elegible_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.carrera_elegible_id_seq CASCADE;
CREATE SEQUENCE academica.carrera_elegible_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.carrera_elegible_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.carrera_elegible | type: TABLE --
-- DROP TABLE IF EXISTS academica.carrera_elegible CASCADE;
CREATE TABLE academica.carrera_elegible (
	id integer NOT NULL DEFAULT nextval('academica.carrera_elegible_id_seq'::regclass),
	codigo_carrera integer NOT NULL,
	cupos_excelencia numeric(2,0),
	cupos_adicionales numeric(2,0),
	periodo numeric(1,0) NOT NULL,
	anio numeric(4,0) NOT NULL,
	codigo_pensum numeric(3,0) NOT NULL,
	nivel character varying(10) NOT NULL,
	CONSTRAINT pk_carrera_elegible PRIMARY KEY (id),
	CONSTRAINT uq_codigo_carrera_periodo_anio_codigo_pensum UNIQUE (codigo_carrera,periodo,anio,codigo_pensum)

);
-- ddl-end --
COMMENT ON TABLE academica.carrera_elegible IS E'Carreras que son elegidas para cursar como modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.id IS E'Identificador de la tabla carrera_elegible';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.codigo_carrera IS E'Campo en el que se indica la carrera que se puede cursar en la modalidad de trabajo de grado (ACPEN: PEN_CRA_COD)';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.cupos_excelencia IS E'Número de admitidos por excelencia académica y exentos de pago';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.cupos_adicionales IS E'Número de admitidos admitidos por condiciones económicas y calidades académicas, con pago';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.periodo IS E'Campo en el que se registra el periodo academico de una asignatura';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.anio IS E'Campo en el que se indica el año del periodo académico (ACASPERI: APE_ANO)';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.codigo_pensum IS E'Campo en el que se registra el código del pensum al cual pertenece la asignatura (ACPEN: PEN_NRO)';
-- ddl-end --
COMMENT ON COLUMN academica.carrera_elegible.nivel IS E'Nivel al que pertenece la carrera (POSGRADO o PREGRADO)';
-- ddl-end --
COMMENT ON CONSTRAINT pk_carrera_elegible ON academica.carrera_elegible  IS E'Constraint primary key carrera_elegible';
-- ddl-end --
COMMENT ON CONSTRAINT uq_codigo_carrera_periodo_anio_codigo_pensum ON academica.carrera_elegible  IS E'Constraint unique codigo_carrera_periodo_anio_codigo_pensum';
-- ddl-end --
ALTER TABLE academica.carrera_elegible OWNER TO postgres;
-- ddl-end --

-- object: academica.comentario_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.comentario_id_seq CASCADE;
CREATE SEQUENCE academica.comentario_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.comentario_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.comentario | type: TABLE --
-- DROP TABLE IF EXISTS academica.comentario CASCADE;
CREATE TABLE academica.comentario (
	id integer NOT NULL DEFAULT nextval('academica.comentario_id_seq'::regclass),
	comentario text NOT NULL,
	fecha timestamp NOT NULL,
	autor character varying(80) NOT NULL,
	correccion integer NOT NULL,
	CONSTRAINT pk_comentario PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.comentario IS E'Comentarios que se realizan sobre una corrección de una revisión de un documento';
-- ddl-end --
COMMENT ON COLUMN academica.comentario.id IS E'Identificador de la tabla comentario';
-- ddl-end --
COMMENT ON COLUMN academica.comentario.comentario IS E'Campo en el que se registra el comentario de una persona sobre una corrección';
-- ddl-end --
COMMENT ON COLUMN academica.comentario.fecha IS E'Fecha del comentario';
-- ddl-end --
COMMENT ON COLUMN academica.comentario.autor IS E'Persona que realiza el comentario';
-- ddl-end --
COMMENT ON COLUMN academica.comentario.correccion IS E'Identificador de la tabla correccion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_comentario ON academica.comentario  IS E'Constraint primary key comentario';
-- ddl-end --
ALTER TABLE academica.comentario OWNER TO postgres;
-- ddl-end --

-- object: academica.correccion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.correccion_id_seq CASCADE;
CREATE SEQUENCE academica.correccion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.correccion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.correccion | type: TABLE --
-- DROP TABLE IF EXISTS academica.correccion CASCADE;
CREATE TABLE academica.correccion (
	id integer NOT NULL DEFAULT nextval('academica.correccion_id_seq'::regclass),
	observacion text NOT NULL,
	pagina numeric(4,0),
	revision_trabajo_grado integer NOT NULL,
	documento boolean,
	CONSTRAINT pk_correcion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.correccion IS E'Entidad para almacenar las observaciones, correcciones o modificaciones sobre un documento';
-- ddl-end --
COMMENT ON COLUMN academica.correccion.id IS E'Identificador de la tabla correccion';
-- ddl-end --
COMMENT ON COLUMN academica.correccion.observacion IS E'Corrección que se realiza en una revisión a un documento';
-- ddl-end --
COMMENT ON COLUMN academica.correccion.pagina IS E'Página del documento en la que se debe hacer la correción';
-- ddl-end --
COMMENT ON COLUMN academica.correccion.revision_trabajo_grado IS E'Identificador de la tabla revision_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.correccion.documento IS E'Campo que indica si la corrección es un documento o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_correcion ON academica.correccion  IS E'Constraint primary key correccion';
-- ddl-end --
ALTER TABLE academica.correccion OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.detalle_id_seq CASCADE;
CREATE SEQUENCE academica.detalle_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.detalle_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle | type: TABLE --
-- DROP TABLE IF EXISTS academica.detalle CASCADE;
CREATE TABLE academica.detalle (
	id integer NOT NULL DEFAULT nextval('academica.detalle_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	enunciado character varying(250) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	tipo_detalle integer NOT NULL,
	CONSTRAINT pk_detalle PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.detalle IS E'Tabla que permite almacenar los diferentes detalles de una solicitud, es decir cada uno de los campos que deberá diligenciar un usuario al momento de presentar una solicitud.';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.id IS E'Identificador que permite diferenciar los detalles de una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.nombre IS E'Nombre del detalle o campo que se debe diligenciar al realizar una solicitud de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.enunciado IS E'Enunciado con el cual se solicita al usuario diligenciar el detalle.';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.descripcion IS E'Descripción breve que define la funciónalidad del detalle';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.codigo_abreviacion IS E'Abreviación del nombre del detalle';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.activo IS E'Campo que permite identificar si el detalle se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN academica.detalle.tipo_detalle IS E'Identificador del tipo de detalle que tiene la solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle ON academica.detalle  IS E'Constraint primary key tabla detalle';
-- ddl-end --
ALTER TABLE academica.detalle OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_pasantia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.detalle_pasantia_id_seq CASCADE;
CREATE SEQUENCE academica.detalle_pasantia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.detalle_pasantia_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_pasantia | type: TABLE --
-- DROP TABLE IF EXISTS academica.detalle_pasantia CASCADE;
CREATE TABLE academica.detalle_pasantia (
	id integer NOT NULL DEFAULT nextval('academica.detalle_pasantia_id_seq'::regclass),
	empresa integer NOT NULL,
	horas integer,
	objeto_contrato text,
	observaciones text,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_detalle_pasantia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.detalle_pasantia IS E'Tabla que almacena la información de la pasantía de un estudiante';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_pasantia.id IS E'Identificador de la pasantía';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_pasantia.empresa IS E'Empresa donde el estudiante realiza la pasantía, referencia a la tabla um_tenam del modelo de WSO2';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_pasantia.horas IS E'Número de horas dedicadas por el estudiante a la pasantía';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_pasantia.objeto_contrato IS E'Objeto de contrato del estudiante en la pasantía';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_pasantia.observaciones IS E'Observaciones hechas por la empresa';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_pasantia.trabajo_grado IS E'Identificador del trabajo de grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle_pasantia ON academica.detalle_pasantia  IS E'Llave primaria de la tabla detalle_pasantia';
-- ddl-end --
ALTER TABLE academica.detalle_pasantia OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.detalle_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.detalle_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.detalle_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.detalle_solicitud CASCADE;
CREATE TABLE academica.detalle_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.detalle_solicitud_id_seq'::regclass),
	descripcion text NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	detalle_tipo_solicitud integer NOT NULL,
	CONSTRAINT pk_detalle_solicitud PRIMARY KEY (id),
	CONSTRAINT uq_solicitud_trabajo_grado_detalle_tipo_solicitud UNIQUE (solicitud_trabajo_grado,detalle_tipo_solicitud)

);
-- ddl-end --
COMMENT ON TABLE academica.detalle_solicitud IS E'Tabla que almacena la respuesta que el usuario dio al detalle de una solicitud.';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_solicitud.id IS E'Identificador de la respuesta que da el usuario al detalle requerido en la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_solicitud.descripcion IS E'Respuesta que da el usuario al detalle requerido en la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_solicitud.solicitud_trabajo_grado IS E'Identificador de la solicitud a la que pertenecen los detalles';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_solicitud.detalle_tipo_solicitud IS E'Detalle que se relaciona con la solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle_solicitud ON academica.detalle_solicitud  IS E'Constraint primary key tabla detalle_solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT uq_solicitud_trabajo_grado_detalle_tipo_solicitud ON academica.detalle_solicitud  IS E'Constraint unique uq_solicitud_trabajo_grado_detalle_tipo_solicitud';
-- ddl-end --
ALTER TABLE academica.detalle_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_tipo_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.detalle_tipo_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.detalle_tipo_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.detalle_tipo_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.detalle_tipo_solicitud CASCADE;
CREATE TABLE academica.detalle_tipo_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.detalle_tipo_solicitud_id_seq'::regclass),
	detalle integer NOT NULL,
	modalidad_tipo_solicitud integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	requerido boolean NOT NULL,
	numero_orden integer NOT NULL,
	CONSTRAINT pk_detalle_tipo_solicitud PRIMARY KEY (id),
	CONSTRAINT uq_detalle_modalidad_tipo_solicitud UNIQUE (detalle,modalidad_tipo_solicitud)

);
-- ddl-end --
COMMENT ON TABLE academica.detalle_tipo_solicitud IS E'Tabla que relaciona cada uno de los detalles que puede tener una solicitud que se realiza dentro de una modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_tipo_solicitud.id IS E'Identificador de la relación entre el detalle la modalidad y el tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_tipo_solicitud.detalle IS E'Detalle relacionado con el tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_tipo_solicitud.modalidad_tipo_solicitud IS E'Identificador de la relación de la modalidad con el tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_tipo_solicitud.activo IS E'Campo que identifica si un detalle se encuentra en vigencia o es requerido para un tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_tipo_solicitud.requerido IS E'Indica si el campo es obligatorio para la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.detalle_tipo_solicitud.numero_orden IS E'Número en el que aparecera el detalle en el formulario.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle_tipo_solicitud ON academica.detalle_tipo_solicitud  IS E'Constraint primary key detalle_tipo_solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT uq_detalle_modalidad_tipo_solicitud ON academica.detalle_tipo_solicitud  IS E'Constraint unique uq_detalle_modalidad_tipo_solicitud';
-- ddl-end --
ALTER TABLE academica.detalle_tipo_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.detalle_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.detalle_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.detalle_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 99999
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.detalle_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.distincion_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.distincion_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.distincion_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.distincion_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.distincion_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.distincion_trabajo_grado CASCADE;
CREATE TABLE academica.distincion_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.distincion_trabajo_grado_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_distincion_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.distincion_trabajo_grado IS E'Tabla que permite almacenar los diferentes tipos de distinciones que puede tener un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.distincion_trabajo_grado.id IS E'Identificador de la tabla distincion';
-- ddl-end --
COMMENT ON COLUMN academica.distincion_trabajo_grado.nombre IS E'Nombre de la distinción que recibe el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.distincion_trabajo_grado.descripcion IS E'Campo que permite almacenar una pequeña descripción de la distinción que se otorga';
-- ddl-end --
COMMENT ON COLUMN academica.distincion_trabajo_grado.codigo_abreviacion IS E'Abreviación del nombre de la distinción';
-- ddl-end --
COMMENT ON COLUMN academica.distincion_trabajo_grado.activo IS E'Permite saber si la distinción esta activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_distincion_trabajo_grado ON academica.distincion_trabajo_grado  IS E'Constraint primary key distincion';
-- ddl-end --
ALTER TABLE academica.distincion_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_entidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.documento_entidad_id_seq CASCADE;
CREATE SEQUENCE academica.documento_entidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.documento_entidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_entidad | type: TABLE --
-- DROP TABLE IF EXISTS academica.documento_entidad CASCADE;
CREATE TABLE academica.documento_entidad (
	id integer NOT NULL DEFAULT nextval('academica.documento_entidad_id_seq'::regclass),
	entidad integer NOT NULL,
	documento_escrito integer NOT NULL,
	CONSTRAINT pk_documento_entidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.documento_entidad IS E'Tabla que almacena la información de la relación entre una entidad y un documento';
-- ddl-end --
COMMENT ON COLUMN academica.documento_entidad.id IS E'Identificador de la tabla documento_entidad';
-- ddl-end --
COMMENT ON COLUMN academica.documento_entidad.entidad IS E'Identificador de la entidad a la que pertenece el documento, se referencia de la tabla um_tenam del modelo de WSO2';
-- ddl-end --
COMMENT ON COLUMN academica.documento_entidad.documento_escrito IS E'Identificador del documento relacionado con la entidad';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento_entidad ON academica.documento_entidad  IS E'Constraint primary key documento_entidad';
-- ddl-end --
ALTER TABLE academica.documento_entidad OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_escrito_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.documento_escrito_id_seq CASCADE;
CREATE SEQUENCE academica.documento_escrito_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.documento_escrito_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_escrito | type: TABLE --
-- DROP TABLE IF EXISTS academica.documento_escrito CASCADE;
CREATE TABLE academica.documento_escrito (
	id integer NOT NULL DEFAULT nextval('academica.documento_escrito_id_seq'::regclass),
	titulo text NOT NULL,
	enlace character varying(100) NOT NULL,
	resumen text,
	tipo_documento_escrito integer NOT NULL,
	CONSTRAINT pk_documento_escrito PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.documento_escrito IS E'Tabla que almacena los documentos registrados en el sistema';
-- ddl-end --
COMMENT ON COLUMN academica.documento_escrito.id IS E'Identificador del documento';
-- ddl-end --
COMMENT ON COLUMN academica.documento_escrito.titulo IS E'Título del documento asociado a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.documento_escrito.enlace IS E'Enlace del documento (NUXEO)';
-- ddl-end --
COMMENT ON COLUMN academica.documento_escrito.resumen IS E'Resumen del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.documento_escrito.tipo_documento_escrito IS E'Tipo de documento, se referencia de la tabla tipo_documento_escrito del core';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento_escrito ON academica.documento_escrito  IS E'Constraint primary key documento_escrito';
-- ddl-end --
ALTER TABLE academica.documento_escrito OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.documento_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.documento_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.documento_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.documento_solicitud CASCADE;
CREATE TABLE academica.documento_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.documento_solicitud_id_seq'::regclass),
	documento_escrito integer NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_documento_solicitud PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.documento_solicitud IS E'Tabla que relaciona un documento con la solicitud a la que pertenece';
-- ddl-end --
COMMENT ON COLUMN academica.documento_solicitud.id IS E'Identificador de la tabla documento_solicitud_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.documento_solicitud.documento_escrito IS E'Documento asociado a la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.documento_solicitud.solicitud_trabajo_grado IS E'Identificador de la solicitud de trabajo de grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento_solicitud ON academica.documento_solicitud  IS E'Constraint primary key tabla documento_solicitud';
-- ddl-end --
ALTER TABLE academica.documento_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.documento_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.documento_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.documento_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.documento_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.documento_trabajo_grado CASCADE;
CREATE TABLE academica.documento_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.documento_trabajo_grado_id_seq'::regclass),
	trabajo_grado integer NOT NULL,
	documento_escrito integer NOT NULL,
	CONSTRAINT pk_documento_trabajo_grado PRIMARY KEY (id),
	CONSTRAINT uq_trabajo_grado_documento UNIQUE (trabajo_grado,documento_escrito)

);
-- ddl-end --
COMMENT ON TABLE academica.documento_trabajo_grado IS E'Almacena la informacion de un documento digital pdf relacionado con un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.documento_trabajo_grado.id IS E'Identificador documento_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.documento_trabajo_grado.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.documento_trabajo_grado.documento_escrito IS E'Identificador de la tabla documento_escrito';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento_trabajo_grado ON academica.documento_trabajo_grado  IS E'Constraint primary key de la tabla documento_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT uq_trabajo_grado_documento ON academica.documento_trabajo_grado  IS E'Constraint unique uq_trabajo_grado_documento';
-- ddl-end --
ALTER TABLE academica.documento_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.espacio_academico_inscrito_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.espacio_academico_inscrito_id_seq CASCADE;
CREATE SEQUENCE academica.espacio_academico_inscrito_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.espacio_academico_inscrito_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.espacio_academico_inscrito | type: TABLE --
-- DROP TABLE IF EXISTS academica.espacio_academico_inscrito CASCADE;
CREATE TABLE academica.espacio_academico_inscrito (
	id integer NOT NULL DEFAULT nextval('academica.espacio_academico_inscrito_id_seq'::regclass),
	nota numeric(3,2),
	espacios_academicos_elegibles integer NOT NULL,
	estado_espacio_academico_inscrito integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_espacio_academico_inscrito PRIMARY KEY (id),
	CONSTRAINT uq_espacios_academicos_elegibles_trabajo_grado UNIQUE (espacios_academicos_elegibles,trabajo_grado)

);
-- ddl-end --
COMMENT ON TABLE academica.espacio_academico_inscrito IS E'Entidad que guarda las asignaturas solicitadas por un estudiante y las asignadas';
-- ddl-end --
COMMENT ON COLUMN academica.espacio_academico_inscrito.id IS E'Identificador del espacio academico inscrito';
-- ddl-end --
COMMENT ON COLUMN academica.espacio_academico_inscrito.nota IS E'Nota de la asignatura vista';
-- ddl-end --
COMMENT ON COLUMN academica.espacio_academico_inscrito.espacios_academicos_elegibles IS E'Identificador de la tabla espacios_academicos_elegibles';
-- ddl-end --
COMMENT ON COLUMN academica.espacio_academico_inscrito.estado_espacio_academico_inscrito IS E'Identificador de la tabla estado_espacio_academico_inscrito';
-- ddl-end --
COMMENT ON COLUMN academica.espacio_academico_inscrito.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_espacio_academico_inscrito ON academica.espacio_academico_inscrito  IS E'Constraint primary key espacio_academico_inscrito';
-- ddl-end --
COMMENT ON CONSTRAINT uq_espacios_academicos_elegibles_trabajo_grado ON academica.espacio_academico_inscrito  IS E'Constraint unique uq_espacios_academicos_elegibles_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.espacio_academico_inscrito OWNER TO postgres;
-- ddl-end --

-- object: academica.espacios_academicos_elegibles_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.espacios_academicos_elegibles_id_seq CASCADE;
CREATE SEQUENCE academica.espacios_academicos_elegibles_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.espacios_academicos_elegibles_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.espacios_academicos_elegibles | type: TABLE --
-- DROP TABLE IF EXISTS academica.espacios_academicos_elegibles CASCADE;
CREATE TABLE academica.espacios_academicos_elegibles (
	id integer NOT NULL DEFAULT nextval('academica.espacios_academicos_elegibles_id_seq'::regclass),
	codigo_asignatura integer NOT NULL,
	activo boolean NOT NULL DEFAULT false,
	carrera_elegible integer NOT NULL,
	CONSTRAINT pk_espacios_academicos_elegibles PRIMARY KEY (id),
	CONSTRAINT uq_codigo_asignatura_carrera_elegible UNIQUE (codigo_asignatura,carrera_elegible)

);
-- ddl-end --
COMMENT ON TABLE academica.espacios_academicos_elegibles IS E'Asignaturas elegibles para optar por la modalidad de espacios academicos de posgrado o profundización';
-- ddl-end --
COMMENT ON COLUMN academica.espacios_academicos_elegibles.id IS E'Identificador de la tabla asignaturas_elegibles';
-- ddl-end --
COMMENT ON COLUMN academica.espacios_academicos_elegibles.codigo_asignatura IS E'Código de la asignatura que puede ser elegible por el estudiante que opte por la modalidad de espacios académicos de posgrado o profundización';
-- ddl-end --
COMMENT ON COLUMN academica.espacios_academicos_elegibles.activo IS E'Permite saber si la carrera se encuentra elegible aun o no';
-- ddl-end --
COMMENT ON COLUMN academica.espacios_academicos_elegibles.carrera_elegible IS E'Identificador de la tabla carrera_elegible';
-- ddl-end --
COMMENT ON CONSTRAINT pk_espacios_academicos_elegibles ON academica.espacios_academicos_elegibles  IS E'Constraint primary key de la tabla espacios_academicos_elegibles';
-- ddl-end --
COMMENT ON CONSTRAINT uq_codigo_asignatura_carrera_elegible ON academica.espacios_academicos_elegibles  IS E'Constraint unique uq_codigo_asignatura_carrera_elegible';
-- ddl-end --
ALTER TABLE academica.espacios_academicos_elegibles OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_asignatura_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estado_asignatura_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.estado_asignatura_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estado_asignatura_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_asignatura_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.estado_asignatura_trabajo_grado CASCADE;
CREATE TABLE academica.estado_asignatura_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.estado_asignatura_trabajo_grado_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_asignatura_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.estado_asignatura_trabajo_grado IS E'Tabla que parametriza los estados de las asignaturas de TG I y TG II';
-- ddl-end --
COMMENT ON COLUMN academica.estado_asignatura_trabajo_grado.id IS E'Identificador del estado de estados de las asignaturas de TG I y TG II';
-- ddl-end --
COMMENT ON COLUMN academica.estado_asignatura_trabajo_grado.nombre IS E'Nombre del estado de la asignatura de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_asignatura_trabajo_grado.descripcion IS E'Permite describir cuando una asignatura se encuentra en este estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_asignatura_trabajo_grado.codigo_abreviacion IS E'Abreviación del nombre del estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_asignatura_trabajo_grado.activo IS E'Permite identificar si el estado esta activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_asignatura_trabajo_grado ON academica.estado_asignatura_trabajo_grado  IS E'Constraint primary key estado espacio academico inscrito';
-- ddl-end --
ALTER TABLE academica.estado_asignatura_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_espacio_academico_inscrito_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estado_espacio_academico_inscrito_id_seq CASCADE;
CREATE SEQUENCE academica.estado_espacio_academico_inscrito_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estado_espacio_academico_inscrito_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_espacio_academico_inscrito | type: TABLE --
-- DROP TABLE IF EXISTS academica.estado_espacio_academico_inscrito CASCADE;
CREATE TABLE academica.estado_espacio_academico_inscrito (
	id integer NOT NULL DEFAULT nextval('academica.estado_espacio_academico_inscrito_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_espacio_academico_inscrito PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.estado_espacio_academico_inscrito IS E'Tabla que parametriza los estados de los espacios académicos inscritos';
-- ddl-end --
COMMENT ON COLUMN academica.estado_espacio_academico_inscrito.id IS E'Identificador del estado del espacio académico inscrito';
-- ddl-end --
COMMENT ON COLUMN academica.estado_espacio_academico_inscrito.nombre IS E'Nombre del estado del espacio académico inscrito';
-- ddl-end --
COMMENT ON COLUMN academica.estado_espacio_academico_inscrito.descripcion IS E'Permite describir cuando un espacio académico se encuentra en este estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_espacio_academico_inscrito.codigo_abreviacion IS E'Abreviación del nombre del estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_espacio_academico_inscrito.activo IS E'Permite identificar si el estado esta activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_espacio_academico_inscrito ON academica.estado_espacio_academico_inscrito  IS E'Constraint primary key estado espacio academico inscrito';
-- ddl-end --
ALTER TABLE academica.estado_espacio_academico_inscrito OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_estudiante_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estado_estudiante_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.estado_estudiante_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estado_estudiante_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_estudiante_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.estado_estudiante_trabajo_grado CASCADE;
CREATE TABLE academica.estado_estudiante_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.estado_estudiante_trabajo_grado_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_estado_estudiante_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.estado_estudiante_trabajo_grado IS E'Tabla que permite parametrizar los diferentes estados en los que se puede encontrar un estudiante respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_estudiante_trabajo_grado.id IS E'Identificador que diferencia cada uno de los estados en los que se puede encontrar un estudiante respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_estudiante_trabajo_grado.nombre IS E'Nombre del estado en el que se puede encontrar un estudiante con respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_estudiante_trabajo_grado.descripcion IS E'Campo que permite describir los estados en los que se puede encontrar un estudiante con respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_estudiante_trabajo_grado.codigo_abreviacion IS E'Código de abreviación del estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_estudiante_trabajo_grado.activo IS E'Permite identificar si un tipo de estado esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_estudiante_trabajo_grado ON academica.estado_estudiante_trabajo_grado  IS E'Constraint primary keb tabla estado_estudiante_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.estado_estudiante_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_revision_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estado_revision_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.estado_revision_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estado_revision_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_revision_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.estado_revision_trabajo_grado CASCADE;
CREATE TABLE academica.estado_revision_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.estado_revision_trabajo_grado_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_estado_revision_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.estado_revision_trabajo_grado IS E'Tabla que permite parametrizar los diferentes estados en los que se puede encontrar una revisión';
-- ddl-end --
COMMENT ON COLUMN academica.estado_revision_trabajo_grado.id IS E'Identificador que diferencia cada uno de los estados en los que se puede encontrar una revisión';
-- ddl-end --
COMMENT ON COLUMN academica.estado_revision_trabajo_grado.nombre IS E'Nombre del tipo de estado de la revisión';
-- ddl-end --
COMMENT ON COLUMN academica.estado_revision_trabajo_grado.descripcion IS E'Campo que permite describir los estados en los que se puede encontrar una revisión';
-- ddl-end --
COMMENT ON COLUMN academica.estado_revision_trabajo_grado.codigo_abreviacion IS E'Código de abreviación del nombre del estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_revision_trabajo_grado.activo IS E'Permite identificar si un tipo de estado esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_revision_trabajo_grado ON academica.estado_revision_trabajo_grado  IS E'Constraint primary key tabla estado_revision_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.estado_revision_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estado_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.estado_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estado_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.estado_solicitud CASCADE;
CREATE TABLE academica.estado_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.estado_solicitud_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_estado_solicitud PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.estado_solicitud IS E'Tabla que permite parametrizar los diferentes estados en los que se puede encontrar una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.estado_solicitud.id IS E'Identificador que diferencia cada uno de los estados en los que se puede encontrar una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.estado_solicitud.nombre IS E'Nombre del tipo de estado de la solicitud.';
-- ddl-end --
COMMENT ON COLUMN academica.estado_solicitud.descripcion IS E'Campo que permite describir los estados en los que se puede encontrar una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.estado_solicitud.codigo_abreviacion IS E'Código de abreviación del nombre del tipo de estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_solicitud.activo IS E'Permite identificar si un tipo de estado esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_solicitud ON academica.estado_solicitud  IS E'Constraint primary key tabla estado_solicitud';
-- ddl-end --
ALTER TABLE academica.estado_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estado_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.estado_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estado_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estado_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.estado_trabajo_grado CASCADE;
CREATE TABLE academica.estado_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.estado_trabajo_grado_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.estado_trabajo_grado IS E'Tabla que permite parametrizar los diferentes estados o etapas en los que se puede encontrar un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_trabajo_grado.id IS E'Identificador de la tabla estado_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_trabajo_grado.nombre IS E'Nombre del estado del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_trabajo_grado.descripcion IS E'Breve descripción donde se explica cuando un trabajo de grado se encuentra en este estado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_trabajo_grado.codigo_abreviacion IS E'Abreviación del nombre del estado del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estado_trabajo_grado.activo IS E'Permite identificar si el estado se encuentra activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_trabajo_grado ON academica.estado_trabajo_grado  IS E'Constraint primary key estado_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.estado_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.estructura_investigacion_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estructura_investigacion_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.estructura_investigacion_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estructura_investigacion_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estructura_investigacion_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.estructura_investigacion_trabajo_grado CASCADE;
CREATE TABLE academica.estructura_investigacion_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.estructura_investigacion_trabajo_grado_id_seq'::regclass),
	codigo_estructura_investigacion integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_estructura_investigacion_trabajo_grado PRIMARY KEY (id),
	CONSTRAINT uq_trabajo_grado_estructura_investigacion UNIQUE (codigo_estructura_investigacion,trabajo_grado)

);
-- ddl-end --
COMMENT ON TABLE academica.estructura_investigacion_trabajo_grado IS E'Tabla que almacena las estructuras de investigación (instituto, grupo o semillero) pertenecientes a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estructura_investigacion_trabajo_grado.id IS E'Tabla que almacena las estructuras de investigación (instituto, grupo o semillero) pertenecientes a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estructura_investigacion_trabajo_grado.codigo_estructura_investigacion IS E'Código asociado a la estructura de investigación, se obtiene del CIDC';
-- ddl-end --
COMMENT ON COLUMN academica.estructura_investigacion_trabajo_grado.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estructura_investigacion_trabajo_grado ON academica.estructura_investigacion_trabajo_grado  IS E'Constraint primary key estructura_investigacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_trabajo_grado_estructura_investigacion ON academica.estructura_investigacion_trabajo_grado  IS E'Constraint unique uq_trabajo_grado_estructura_investigacion';
-- ddl-end --
ALTER TABLE academica.estructura_investigacion_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.estudiante_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.estudiante_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.estudiante_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.estudiante_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.estudiante_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.estudiante_trabajo_grado CASCADE;
CREATE TABLE academica.estudiante_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.estudiante_trabajo_grado_id_seq'::regclass),
	estudiante character varying(15) NOT NULL,
	trabajo_grado integer NOT NULL,
	estado_estudiante_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_estudiante_trabajo_grado PRIMARY KEY (id),
	CONSTRAINT uq_trabajo_grado_codigo_estudiante UNIQUE (estudiante,trabajo_grado)

);
-- ddl-end --
COMMENT ON TABLE academica.estudiante_trabajo_grado IS E'Tabla en la que se relaciona uno o varios estudiantes a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.estudiante_trabajo_grado.id IS E'Id de la tabla estudiante_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.estudiante_trabajo_grado.estudiante IS E'Identificador del estudiante relacionado con el trabajo de grado, se referencia del modelo de WSO2 tabla um_user';
-- ddl-end --
COMMENT ON COLUMN academica.estudiante_trabajo_grado.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.estudiante_trabajo_grado.estado_estudiante_trabajo_grado IS E'Identificador de la tabla estado_estudiante_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estudiante_trabajo_grado ON academica.estudiante_trabajo_grado  IS E'Constraint primary key estudiante_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.estudiante_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.evaluacion_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.evaluacion_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.evaluacion_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.evaluacion_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.evaluacion_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.evaluacion_trabajo_grado CASCADE;
CREATE TABLE academica.evaluacion_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.evaluacion_trabajo_grado_id_seq'::regclass),
	nota double precision NOT NULL,
	vinculacion_trabajo_grado integer NOT NULL,
	formato_evaluacion_carrera integer,
	socializacion integer,
	aprobado boolean DEFAULT true,
	observaciones text DEFAULT '',
	activo boolean DEFAULT true,
	CONSTRAINT pk_evaluacion_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.evaluacion_trabajo_grado IS E'Almacena la evaluación de la sustentación de un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.evaluacion_trabajo_grado.id IS E'Identificador de la evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.evaluacion_trabajo_grado.nota IS E'Almacena la nota final de la socialización';
-- ddl-end --
COMMENT ON COLUMN academica.evaluacion_trabajo_grado.vinculacion_trabajo_grado IS E'Identificador de la tabla vinculacion_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.evaluacion_trabajo_grado.formato_evaluacion_carrera IS E'Identificador de la taabla formato_evaluacion_carrera';
-- ddl-end --
COMMENT ON COLUMN academica.evaluacion_trabajo_grado.socializacion IS E'Identificador de la tabla socializacion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_evaluacion_trabajo_grado ON academica.evaluacion_trabajo_grado  IS E'Constraint primary key evaluacion';
-- ddl-end --
ALTER TABLE academica.evaluacion_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.formato_evaluacion_carrera_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.formato_evaluacion_carrera_id_seq CASCADE;
CREATE SEQUENCE academica.formato_evaluacion_carrera_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.formato_evaluacion_carrera_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.formato_evaluacion_carrera | type: TABLE --
-- DROP TABLE IF EXISTS academica.formato_evaluacion_carrera CASCADE;
CREATE TABLE academica.formato_evaluacion_carrera (
	id integer NOT NULL DEFAULT nextval('academica.formato_evaluacion_carrera_id_seq'::regclass),
	activo boolean NOT NULL DEFAULT false,
	codigo_proyecto integer NOT NULL,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	modalidad integer NOT NULL,
	formato integer NOT NULL,
	CONSTRAINT pk_formato_evaluacion_carrera PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.formato_evaluacion_carrera IS E'Tabla que relaciona los registros del formato de evaluacion dependiendo del proyecto curricular';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.id IS E'Identificador de la tabla formato_evaluacion_carrera';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.activo IS E'Estado(activo,inactivo)';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.codigo_proyecto IS E'Código del proyecto que viene del web service';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.fecha_inicio IS E'Fecha de inicio del formato';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.fecha_fin IS E'Fecha de fin del formato';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.modalidad IS E'Modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.formato_evaluacion_carrera.formato IS E'Identificador del formato de evaluación relacionado con la carrera';
-- ddl-end --
COMMENT ON CONSTRAINT pk_formato_evaluacion_carrera ON academica.formato_evaluacion_carrera  IS E'Constraint primary key formato_evaluacion_carrera';
-- ddl-end --
ALTER TABLE academica.formato_evaluacion_carrera OWNER TO postgres;
-- ddl-end --

-- object: academica.formato_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.formato_id_seq CASCADE;
CREATE SEQUENCE academica.formato_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.formato_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.formato | type: TABLE --
-- DROP TABLE IF EXISTS academica.formato CASCADE;
CREATE TABLE academica.formato (
	id integer NOT NULL DEFAULT nextval('academica.formato_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	introduccion character varying(500),
	fecha_realizacion timestamp NOT NULL,
	CONSTRAINT pk_formato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.formato IS E'Tabla que almacena el formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.formato.id IS E'Identificador del formato';
-- ddl-end --
COMMENT ON COLUMN academica.formato.nombre IS E'Nombre del formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.formato.introduccion IS E'Introducción del formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.formato.fecha_realizacion IS E'Fecha de realización del formato de evaluación';
-- ddl-end --
COMMENT ON CONSTRAINT pk_formato ON academica.formato  IS E'Constraint primary key formato';
-- ddl-end --
ALTER TABLE academica.formato OWNER TO postgres;
-- ddl-end --

-- object: academica.modalidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.modalidad_id_seq CASCADE;
CREATE SEQUENCE academica.modalidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.modalidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.modalidad | type: TABLE --
-- DROP TABLE IF EXISTS academica.modalidad CASCADE;
CREATE TABLE academica.modalidad (
	id integer NOT NULL DEFAULT nextval('academica.modalidad_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activa boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_modalidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.modalidad IS E'Tabla que almacena las modalidades de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad.id IS E'Identificador de la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad.nombre IS E'Nombre de la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad.descripcion IS E'Descripción corta de la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad.codigo_abreviacion IS E'Abreviación del nombre de la modalidad';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad.activa IS E'Permite saber si una modalidad de se encuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_modalidad ON academica.modalidad  IS E'Constraint primary key para modalidad';
-- ddl-end --
ALTER TABLE academica.modalidad OWNER TO postgres;
-- ddl-end --

-- object: academica.modalidad_tipo_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.modalidad_tipo_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.modalidad_tipo_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.modalidad_tipo_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.modalidad_tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.modalidad_tipo_solicitud CASCADE;
CREATE TABLE academica.modalidad_tipo_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.modalidad_tipo_solicitud_id_seq'::regclass),
	tipo_solicitud integer NOT NULL,
	modalidad integer NOT NULL,
	CONSTRAINT pk_modalidad_tipo_solicitud PRIMARY KEY (id),
	CONSTRAINT uq_tipo_solicitud_modalidad UNIQUE (tipo_solicitud,modalidad)

);
-- ddl-end --
COMMENT ON TABLE academica.modalidad_tipo_solicitud IS E'Tabla que relaciona las modalidades de trabajo de grado con los tipos de solicitud.';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad_tipo_solicitud.id IS E'Identificador de la tabla modalidad_tipo_solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad_tipo_solicitud.tipo_solicitud IS E'Identificador del tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.modalidad_tipo_solicitud.modalidad IS E'Modalidad asociada al tipo de solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_modalidad_tipo_solicitud ON academica.modalidad_tipo_solicitud  IS E'Constraint primary key modalidad_tipo_solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT uq_tipo_solicitud_modalidad ON academica.modalidad_tipo_solicitud  IS E'Constraint unique uq_tipo_solicitud_modalidad';
-- ddl-end --
ALTER TABLE academica.modalidad_tipo_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.pregunta_formato_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.pregunta_formato_id_seq CASCADE;
CREATE SEQUENCE academica.pregunta_formato_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.pregunta_formato_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.pregunta_formato | type: TABLE --
-- DROP TABLE IF EXISTS academica.pregunta_formato CASCADE;
CREATE TABLE academica.pregunta_formato (
	id integer NOT NULL DEFAULT nextval('academica.pregunta_formato_id_seq'::regclass),
	activo boolean NOT NULL DEFAULT true,
	orden numeric(2,0),
	valoracion double precision,
	pregunta integer NOT NULL,
	formato integer NOT NULL,
	tipo_pregunta integer NOT NULL,
	CONSTRAINT pk_pregunta_formato PRIMARY KEY (id),
	CONSTRAINT uq_pregunta_formato UNIQUE (pregunta,formato)

);
-- ddl-end --
COMMENT ON TABLE academica.pregunta_formato IS E'Entidad que almacena las preguntas del formato';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.id IS E'Identificador formato de preguntas';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.activo IS E'Indica si la pregunta esta activa o no';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.orden IS E'Orden de la pregunta en el formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.valoracion IS E'Determina el valor de una pregunta dentro del formato';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.pregunta IS E'Identificador de la tabla pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.formato IS E'Identificador de la tabla formato';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta_formato.tipo_pregunta IS E'Identificador de la tabla tipo_pregunta';
-- ddl-end --
COMMENT ON CONSTRAINT pk_pregunta_formato ON academica.pregunta_formato  IS E'Constraint primary key pregunta_formato';
-- ddl-end --
ALTER TABLE academica.pregunta_formato OWNER TO postgres;
-- ddl-end --

-- object: academica.pregunta_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.pregunta_id_seq CASCADE;
CREATE SEQUENCE academica.pregunta_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.pregunta_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.pregunta | type: TABLE --
-- DROP TABLE IF EXISTS academica.pregunta CASCADE;
CREATE TABLE academica.pregunta (
	id integer NOT NULL DEFAULT nextval('academica.pregunta_id_seq'::regclass),
	enunciado character varying(1000) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_pregunta PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.pregunta IS E'Entidad que almacena las preguntas para un formato evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta.id IS E'Identificador de la pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta.enunciado IS E'Enunciado de la pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta.descripcion IS E'Breve descripción de la funcionalidad de la pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta.codigo_abreviacion IS E'Abreviación del nombre de la pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.pregunta.activo IS E'Permite saber si la pregunta se encuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_pregunta ON academica.pregunta  IS E'Constraint primary key pregunta';
-- ddl-end --
ALTER TABLE academica.pregunta OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_evaluacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.respuesta_evaluacion_id_seq CASCADE;
CREATE SEQUENCE academica.respuesta_evaluacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.respuesta_evaluacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS academica.respuesta_evaluacion CASCADE;
CREATE TABLE academica.respuesta_evaluacion (
	id integer NOT NULL DEFAULT nextval('academica.respuesta_evaluacion_id_seq'::regclass),
	justificacion character varying(500),
	respuesta_formato integer NOT NULL,
	evaluacion_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_respuesta_evaluacion PRIMARY KEY (id),
	CONSTRAINT uq_respuesta_formato_evaluacion UNIQUE (respuesta_formato,evaluacion_trabajo_grado)

);
-- ddl-end --
COMMENT ON TABLE academica.respuesta_evaluacion IS E'Entidad que almacena las relaciones entre la evaluación y el formato de las respuestas';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_evaluacion.id IS E'Identificador de la tabla respuesta_evaluacion';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_evaluacion.justificacion IS E'Campo que contiene la justificación a una respuesta de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_evaluacion.respuesta_formato IS E'Identificador de la tabla respuesta_formato';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_evaluacion.evaluacion_trabajo_grado IS E'Identificador de la tabla evaluacion_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuesta_evaluacion ON academica.respuesta_evaluacion  IS E'Constraint primary key tabla respuesta_evaluacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_respuesta_formato_evaluacion ON academica.respuesta_evaluacion  IS E'Constraint unique uq_respuesta_formato_evaluacion';
-- ddl-end --
ALTER TABLE academica.respuesta_evaluacion OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_formato_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.respuesta_formato_id_seq CASCADE;
CREATE SEQUENCE academica.respuesta_formato_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.respuesta_formato_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_formato | type: TABLE --
-- DROP TABLE IF EXISTS academica.respuesta_formato CASCADE;
CREATE TABLE academica.respuesta_formato (
	id integer NOT NULL DEFAULT nextval('academica.respuesta_formato_id_seq'::regclass),
	orden numeric(2,0),
	valoracion double precision,
	respuesta integer NOT NULL,
	pregunta_formato integer NOT NULL,
	CONSTRAINT pk_respuestas_formato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.respuesta_formato IS E'Entidad que guarda las respuestas de un formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_formato.id IS E'Identificador';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_formato.orden IS E'Representa el orden en que se muestran las respuestas en un formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_formato.valoracion IS E'Determina el valor de una respuesta';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_formato.respuesta IS E'Identificador de la tabla respuesta';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_formato.pregunta_formato IS E'Identificador de la tabla pregunta_formato';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuestas_formato ON academica.respuesta_formato  IS E'Constraint primary key respuestas_formato';
-- ddl-end --
ALTER TABLE academica.respuesta_formato OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.respuesta_id_seq CASCADE;
CREATE SEQUENCE academica.respuesta_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.respuesta_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta | type: TABLE --
-- DROP TABLE IF EXISTS academica.respuesta CASCADE;
CREATE TABLE academica.respuesta (
	id integer NOT NULL DEFAULT nextval('academica.respuesta_id_seq'::regclass),
	descripcion character varying(250),
	nombre character varying(1000) NOT NULL,
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_respuesta PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.respuesta IS E'Entidad que almacena las respuestas para un formato';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta.id IS E'Identificador de la respuesta';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta.descripcion IS E'Breve descripción de la respuesta';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta.nombre IS E'Permite identificar a la respuesta';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta.codigo_abreviacion IS E'Abreviación del nombre de la respuesta';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta.activo IS E'Permite identificar cuando una respuesta se encuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuesta ON academica.respuesta  IS E'Identificador respuesta';
-- ddl-end --
ALTER TABLE academica.respuesta OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.respuesta_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.respuesta_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.respuesta_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.respuesta_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.respuesta_solicitud CASCADE;
CREATE TABLE academica.respuesta_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.respuesta_solicitud_id_seq'::regclass),
	fecha timestamp NOT NULL,
	justificacion text,
	ente_responsable integer,
	usuario integer,
	estado_solicitud integer NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_respuesta_solicitud PRIMARY KEY (id),
	CONSTRAINT uq_solicitud_trabajo_grado UNIQUE (solicitud_trabajo_grado,estado_solicitud)

);
-- ddl-end --
COMMENT ON TABLE academica.respuesta_solicitud IS E'Tabla que permite identificar el estado y la respuesta dada a una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.id IS E'Identificador de la tabla respuesta_solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.fecha IS E'Fecha en la cual cambia el estado de una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.justificacion IS E'Almacena la justificación con la cual se hace un cambio al estado de una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.ente_responsable IS E'Identificador que relaciona a la organización que se encarga de evaluar y dar respuesta a la solicitud, este campo se va a consumir de un servicio.';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.usuario IS E'Usuario que se encarga de cambiar el estado de la solicitud y de informar lo decidido por la organización responsable (Hace referencia a la tabla um_user del modelo de WSO2)';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.estado_solicitud IS E'Identificador de la tabla estado_solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.respuesta_solicitud.solicitud_trabajo_grado IS E'Identificador de la tabla solicitud_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuesta_solicitud ON academica.respuesta_solicitud  IS E'Constraint primary key tabla respuesta_solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT uq_solicitud_trabajo_grado ON academica.respuesta_solicitud  IS E'Constraint unique uq_solicitud_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.respuesta_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.revision_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.revision_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.revision_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.revision_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.revision_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.revision_trabajo_grado CASCADE;
CREATE TABLE academica.revision_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.revision_trabajo_grado_id_seq'::regclass),
	numero_revision numeric(2,0) NOT NULL,
	fecha_recepcion timestamp NOT NULL,
	fecha_revision timestamp,
	estado_revision_trabajo_grado integer NOT NULL,
	documento_trabajo_grado integer NOT NULL,
	vinculacion_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_revision_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.revision_trabajo_grado IS E'Entidad que aloja el número de revisiones hechas a un documento';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.id IS E'Identificador para la tabla revision';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.numero_revision IS E'Indica el número de la revisión realizada sobre un documento';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.fecha_recepcion IS E'Campo que indica la fecha en que la revisión es solicitada';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.fecha_revision IS E'Fecha de realización de una revisión';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.estado_revision_trabajo_grado IS E'Identificador de la tabla estado_revision_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.documento_trabajo_grado IS E'Identificador de la tabla documento_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.revision_trabajo_grado.vinculacion_trabajo_grado IS E'Identificador de la tabla vinculacion_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_revision_trabajo_grado ON academica.revision_trabajo_grado  IS E'Constraint primary key revision_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.revision_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.rol_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.rol_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.rol_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.rol_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.rol_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.rol_trabajo_grado CASCADE;
CREATE TABLE academica.rol_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.rol_trabajo_grado_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_rol_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.rol_trabajo_grado IS E'Tabla que permite parametrizar el rol que cumple una persona u entidad dentro del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.rol_trabajo_grado.id IS E'Identificador de la tabla trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.rol_trabajo_grado.nombre IS E'Nombre del rol que describe la actividad que se realiza dentro del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.rol_trabajo_grado.descripcion IS E'Descripción de la labor que se realiza dentro del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.rol_trabajo_grado.codigo_abreviacion IS E'Abreviación del nombre';
-- ddl-end --
COMMENT ON COLUMN academica.rol_trabajo_grado.activo IS E'Permite determinar si el rol se encuentra activo o no en el momento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_rol_trabajo_grado ON academica.rol_trabajo_grado  IS E'Constraint primary key rol_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.rol_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.socializacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.socializacion_id_seq CASCADE;
CREATE SEQUENCE academica.socializacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.socializacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.socializacion | type: TABLE --
-- DROP TABLE IF EXISTS academica.socializacion CASCADE;
CREATE TABLE academica.socializacion (
	id integer NOT NULL DEFAULT nextval('academica.socializacion_id_seq'::regclass),
	fecha timestamp NOT NULL,
	lugar integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_socializacion PRIMARY KEY (id),
	CONSTRAINT uq_trabajo_grado UNIQUE (trabajo_grado)

);
-- ddl-end --
COMMENT ON TABLE academica.socializacion IS E'Entidad que almacena la sustentación del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.socializacion.id IS E'Identificador de la socialización';
-- ddl-end --
COMMENT ON COLUMN academica.socializacion.fecha IS E'Fecha de realización de la socialización';
-- ddl-end --
COMMENT ON COLUMN academica.socializacion.lugar IS E'Lugar del evento, se referencia del core, tabla espacio_fisico';
-- ddl-end --
COMMENT ON COLUMN academica.socializacion.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_socializacion ON academica.socializacion  IS E'Constraint primary key tabla socializacion';
-- ddl-end --
ALTER TABLE academica.socializacion OWNER TO postgres;
-- ddl-end --

-- object: academica.solicitud_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.solicitud_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.solicitud_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.solicitud_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.solicitud_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.solicitud_trabajo_grado CASCADE;
CREATE TABLE academica.solicitud_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.solicitud_trabajo_grado_id_seq'::regclass),
	fecha timestamp NOT NULL,
	modalidad_tipo_solicitud integer NOT NULL,
	trabajo_grado integer,
	periodo_academico character varying(6),
	CONSTRAINT pk_solicitud_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.solicitud_trabajo_grado IS E'Tabla que permite almacenar las diferentes solicitudes de trabajo de grado que se presenten en el sistema';
-- ddl-end --
COMMENT ON COLUMN academica.solicitud_trabajo_grado.id IS E'Identificador de la tabla solicitud_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.solicitud_trabajo_grado.fecha IS E'Almacena la fecha en la cual se diligencia y presenta la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.solicitud_trabajo_grado.modalidad_tipo_solicitud IS E'Identificador de la tabla modalidad_tipo_solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.solicitud_trabajo_grado.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.solicitud_trabajo_grado.periodo_academico IS E'Periodo academico en el que se realiza la solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_solicitud_trabajo_grado ON academica.solicitud_trabajo_grado  IS E'Constraint primary key tabla solicitud_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.solicitud_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.tipo_detalle_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.tipo_detalle_id_seq CASCADE;
CREATE SEQUENCE academica.tipo_detalle_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.tipo_detalle_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.tipo_detalle | type: TABLE --
-- DROP TABLE IF EXISTS academica.tipo_detalle CASCADE;
CREATE TABLE academica.tipo_detalle (
	id integer NOT NULL DEFAULT nextval('academica.tipo_detalle_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_detalle PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.tipo_detalle IS E'Tabla que permite almacenar los diferentes tipos de detalle que puede tener una solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_detalle.id IS E'Identificador del tipo de detalle';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_detalle.nombre IS E'Nombre que permite diferenciar a que tipo de detalle se esta refiriendo';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_detalle.descripcion IS E'Descripcion del tipo del detalle';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_detalle.codigo_abreviacion IS E'Código de abreviación del nombre que identifica el tipo de detalle';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_detalle.activo IS E'Campo que permite saber si un tipo de detalle se encuentra vigente o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_detalle ON academica.tipo_detalle  IS E'Constraint primary key tabla tipo_detalle';
-- ddl-end --
ALTER TABLE academica.tipo_detalle OWNER TO postgres;
-- ddl-end --

-- object: academica.tipo_pregunta_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.tipo_pregunta_id_seq CASCADE;
CREATE SEQUENCE academica.tipo_pregunta_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.tipo_pregunta_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.tipo_pregunta | type: TABLE --
-- DROP TABLE IF EXISTS academica.tipo_pregunta CASCADE;
CREATE TABLE academica.tipo_pregunta (
	id integer NOT NULL DEFAULT nextval('academica.tipo_pregunta_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_regunta PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.tipo_pregunta IS E'Tabla que permite parametrizar los diferentes tipos de pregunta que se pueden encontrar en un formato';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_pregunta.id IS E'Identificador que diferencia cada uno de los tipos de preguntas';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_pregunta.nombre IS E'Nombre del tipo de pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_pregunta.descripcion IS E'Campo que permite describir los tipos de preguntas que puede tener un formato';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_pregunta.codigo_abreviacion IS E'Código de abreviación del nombre del tipo de pregunta';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_pregunta.activo IS E'Permite identificar si un tipo de pregunta esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_regunta ON academica.tipo_pregunta  IS E'Constraint primary key tabla tipo_pregunta';
-- ddl-end --
ALTER TABLE academica.tipo_pregunta OWNER TO postgres;
-- ddl-end --

-- object: academica.tipo_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.tipo_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.tipo_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.tipo_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.tipo_solicitud CASCADE;
CREATE TABLE academica.tipo_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.tipo_solicitud_id_seq'::regclass),
	nombre character varying(200) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_solicitud PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.tipo_solicitud IS E'Tabla que almacena los tipos de solicitud que se puedan presentar dentro de cada una de las modalidades de trabajo de grado.';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_solicitud.id IS E'Identificador del tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_solicitud.nombre IS E'Nombre que identifica y deferencia una solicitud de otra';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_solicitud.descripcion IS E'Describe la función de la solicitud y cada una de las caracteristicas de esta asi como tambien a quien va dirigida';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_solicitud.codigo_abreviacion IS E'Abreviación del nombre que identifica un tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.tipo_solicitud.activo IS E'Permite definir si un tipo de solicitud se encuentra activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_solicitud ON academica.tipo_solicitud  IS E'Constraint primary key tipo_solicitud';
-- ddl-end --
ALTER TABLE academica.tipo_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.trabajo_grado CASCADE;
CREATE TABLE academica.trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.trabajo_grado_id_seq'::regclass),
	titulo text NOT NULL,
	modalidad integer NOT NULL,
	estado_trabajo_grado integer NOT NULL,
	distincion_trabajo_grado integer,
	periodo_academico character varying(6),
	objetivo text,
	CONSTRAINT pk_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.trabajo_grado IS E'Tabla sobre la cual se relacionan los ítems de un trabajo de grado con sus respectivas características';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.id IS E'Identificador de la tabla trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.titulo IS E'Título del trabajo de grado a realizar';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.modalidad IS E'Identificador de la tabla modalidad';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.estado_trabajo_grado IS E'Identificador de la tabla estado_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.distincion_trabajo_grado IS E'Identificador de la tabla distincion_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.periodo_academico IS E'Periodo academico en el que se comienza a realizar el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.trabajo_grado.objetivo IS E'Atributo que guarda objetivos de las propuestas';
-- ddl-end --
COMMENT ON CONSTRAINT pk_trabajo_grado ON academica.trabajo_grado  IS E'Constraint primary key para la tabla trabajo_grado';
-- ddl-end --
ALTER TABLE academica.trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: academica.usuario_solicitud_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.usuario_solicitud_id_seq CASCADE;
CREATE SEQUENCE academica.usuario_solicitud_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.usuario_solicitud_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.usuario_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS academica.usuario_solicitud CASCADE;
CREATE TABLE academica.usuario_solicitud (
	id integer NOT NULL DEFAULT nextval('academica.usuario_solicitud_id_seq'::regclass),
	usuario character varying(15) NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_usuario_solicitud PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE academica.usuario_solicitud IS E'Tabla que permite relacionar una solicitud con el usuario que emite la solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.usuario_solicitud.id IS E'Identificador de la tabla usuario_solicitud';
-- ddl-end --
COMMENT ON COLUMN academica.usuario_solicitud.usuario IS E'Código que identifica al usuario que realiza la solicitud, se referencia del modelo de WSO2 tabla um_user';
-- ddl-end --
COMMENT ON COLUMN academica.usuario_solicitud.solicitud_trabajo_grado IS E'Identificador de la tabla solicitud_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_usuario_solicitud ON academica.usuario_solicitud  IS E'Constraint primary key tabla usuario_solicitud';
-- ddl-end --
ALTER TABLE academica.usuario_solicitud OWNER TO postgres;
-- ddl-end --

-- object: academica.vinculacion_trabajo_grado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS academica.vinculacion_trabajo_grado_id_seq CASCADE;
CREATE SEQUENCE academica.vinculacion_trabajo_grado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE academica.vinculacion_trabajo_grado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: academica.vinculacion_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS academica.vinculacion_trabajo_grado CASCADE;
CREATE TABLE academica.vinculacion_trabajo_grado (
	id integer NOT NULL DEFAULT nextval('academica.vinculacion_trabajo_grado_id_seq'::regclass),
	usuario integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	rol_trabajo_grado integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_vinculacion_trabajo_grado PRIMARY KEY (id),
	CONSTRAINT uq_identificacion_docente_rol_trabajo_grado UNIQUE (usuario,rol_trabajo_grado,trabajo_grado)

);
-- ddl-end --
COMMENT ON TABLE academica.vinculacion_trabajo_grado IS E'Entidad que almacena los tipos de vinculación del docente con el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.id IS E'Identificación de la vinculacion docente';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.usuario IS E'Identificación del usuario relacionado con el trabajo de grado (Se trae del modelo de autenticación unica, tabla um_user)';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.activo IS E'En caso de que se presente cambio de director';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.fecha_inicio IS E'Fecha de inicio de la vinculación';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.fecha_fin IS E'Fecha de finalización de la vinculación';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.rol_trabajo_grado IS E'Identificador de la tabla rol_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN academica.vinculacion_trabajo_grado.trabajo_grado IS E'Identificador de la tabla trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_vinculacion_trabajo_grado ON academica.vinculacion_trabajo_grado  IS E'Constraint de primary key vinculacion_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT uq_identificacion_docente_rol_trabajo_grado ON academica.vinculacion_trabajo_grado  IS E'Constraint unique uq_identificacion_docente_rol_trabajo_grado';
-- ddl-end --
ALTER TABLE academica.vinculacion_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: trabajo_grado_titulo | type: INDEX --
-- DROP INDEX IF EXISTS academica.trabajo_grado_titulo CASCADE;
CREATE INDEX trabajo_grado_titulo ON academica.trabajo_grado
	USING btree
	(
	  titulo
	)
	WITH (FILLFACTOR = 90);
-- ddl-end --
COMMENT ON INDEX academica.trabajo_grado_titulo IS E'Indice que facilita las busquedas por titulos de trabajos de grado';
-- ddl-end --

-- object: fk_areas_docente_area_conocimiento | type: CONSTRAINT --
-- ALTER TABLE academica.areas_docente DROP CONSTRAINT IF EXISTS fk_areas_docente_area_conocimiento CASCADE;
ALTER TABLE academica.areas_docente ADD CONSTRAINT fk_areas_docente_area_conocimiento FOREIGN KEY (area_conocimiento)
REFERENCES academica.area_conocimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_areas_docente_area_conocimiento ON academica.areas_docente  IS E'Constraint foreign key fk_areas_docente_area_conocimiento';
-- ddl-end --


-- object: fk_areas_trabajo_grado_area_conocimiento | type: CONSTRAINT --
-- ALTER TABLE academica.areas_trabajo_grado DROP CONSTRAINT IF EXISTS fk_areas_trabajo_grado_area_conocimiento CASCADE;
ALTER TABLE academica.areas_trabajo_grado ADD CONSTRAINT fk_areas_trabajo_grado_area_conocimiento FOREIGN KEY (area_conocimiento)
REFERENCES academica.area_conocimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_areas_trabajo_grado_area_conocimiento ON academica.areas_trabajo_grado  IS E'Constraint foreign key fk_areas_trabajo_grado_area_conocimiento';
-- ddl-end --


-- object: fk_areas_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.areas_trabajo_grado DROP CONSTRAINT IF EXISTS fk_areas_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.areas_trabajo_grado ADD CONSTRAINT fk_areas_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_areas_trabajo_grado_trabajo_grado ON academica.areas_trabajo_grado  IS E'Constraint foreign key fk_areas_trabajo_grado_trabajo_grado';
-- ddl-end --


-- object: fk_asignatura_trabajo_grado_estado_asignatura_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.asignatura_trabajo_grado DROP CONSTRAINT IF EXISTS fk_asignatura_trabajo_grado_estado_asignatura_trabajo_grado CASCADE;
ALTER TABLE academica.asignatura_trabajo_grado ADD CONSTRAINT fk_asignatura_trabajo_grado_estado_asignatura_trabajo_grado FOREIGN KEY (estado_asignatura_trabajo_grado)
REFERENCES academica.estado_asignatura_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_asignatura_trabajo_grado_estado_asignatura_trabajo_grado ON academica.asignatura_trabajo_grado  IS E'Constraint foreign key fk_asignatura_trabajo_grado_estado_asignatura_trabajo_grado';
-- ddl-end --


-- object: fk_asignatura_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.asignatura_trabajo_grado DROP CONSTRAINT IF EXISTS fk_asignatura_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.asignatura_trabajo_grado ADD CONSTRAINT fk_asignatura_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_asignatura_trabajo_grado_trabajo_grado ON academica.asignatura_trabajo_grado  IS E'Constraint foreign key fk_asignatura_trabajo_grado_trabajo_grado';
-- ddl-end --


-- object: fk_correccion_revision_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.correccion DROP CONSTRAINT IF EXISTS fk_correccion_revision_trabajo_grado CASCADE;
ALTER TABLE academica.correccion ADD CONSTRAINT fk_correccion_revision_trabajo_grado FOREIGN KEY (revision_trabajo_grado)
REFERENCES academica.revision_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_correccion_revision_trabajo_grado ON academica.correccion  IS E'Constraint foreign key fk_correccion_revision_trabajo_grado';
-- ddl-end --


-- object: fk_detalle_tipo_detalle | type: CONSTRAINT --
-- ALTER TABLE academica.detalle DROP CONSTRAINT IF EXISTS fk_detalle_tipo_detalle CASCADE;
ALTER TABLE academica.detalle ADD CONSTRAINT fk_detalle_tipo_detalle FOREIGN KEY (tipo_detalle)
REFERENCES academica.tipo_detalle (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_detalle_tipo_detalle ON academica.detalle  IS E'Constraint foreign key fk_detalle_tipo_detalle';
-- ddl-end --


-- object: fk_detalle_pasantia_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.detalle_pasantia DROP CONSTRAINT IF EXISTS fk_detalle_pasantia_trabajo_grado CASCADE;
ALTER TABLE academica.detalle_pasantia ADD CONSTRAINT fk_detalle_pasantia_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_detalle_pasantia_trabajo_grado ON academica.detalle_pasantia  IS E'Constraint foreign key fk_detalle_pasantia_trabajo_grado';
-- ddl-end --


-- object: fk_detalle_solicitud_detalle_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE academica.detalle_solicitud DROP CONSTRAINT IF EXISTS fk_detalle_solicitud_detalle_tipo_solicitud CASCADE;
ALTER TABLE academica.detalle_solicitud ADD CONSTRAINT fk_detalle_solicitud_detalle_tipo_solicitud FOREIGN KEY (detalle_tipo_solicitud)
REFERENCES academica.detalle_tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_detalle_solicitud_detalle_tipo_solicitud ON academica.detalle_solicitud  IS E'Constraint foreign key fk_detalle_solicitud_detalle_tipo_solicitud';
-- ddl-end --


-- object: fk_detalle_solicitud_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.detalle_solicitud DROP CONSTRAINT IF EXISTS fk_detalle_solicitud_solicitud_trabajo_grado CASCADE;
ALTER TABLE academica.detalle_solicitud ADD CONSTRAINT fk_detalle_solicitud_solicitud_trabajo_grado FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES academica.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_detalle_solicitud_solicitud_trabajo_grado ON academica.detalle_solicitud  IS E'Constraint foreign key fk_detalle_solicitud_solicitud_trabajo_grado';
-- ddl-end --


-- object: fk_detalle_tipo_solicitud_detalle | type: CONSTRAINT --
-- ALTER TABLE academica.detalle_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_detalle_tipo_solicitud_detalle CASCADE;
ALTER TABLE academica.detalle_tipo_solicitud ADD CONSTRAINT fk_detalle_tipo_solicitud_detalle FOREIGN KEY (detalle)
REFERENCES academica.detalle (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_detalle_tipo_solicitud_detalle ON academica.detalle_tipo_solicitud  IS E'Constraint foreign key fk_detalle_tipo_solicitud_detalle';
-- ddl-end --


-- object: fk_detalle_tipo_solicitud_modalidad_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE academica.detalle_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_detalle_tipo_solicitud_modalidad_tipo_solicitud CASCADE;
ALTER TABLE academica.detalle_tipo_solicitud ADD CONSTRAINT fk_detalle_tipo_solicitud_modalidad_tipo_solicitud FOREIGN KEY (modalidad_tipo_solicitud)
REFERENCES academica.modalidad_tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_detalle_tipo_solicitud_modalidad_tipo_solicitud ON academica.detalle_tipo_solicitud  IS E'Constraint foreign key fk_detalle_tipo_solicitud_modalidad_tipo_solicitud';
-- ddl-end --


-- object: fk_documento_entidad_documento_escrito | type: CONSTRAINT --
-- ALTER TABLE academica.documento_entidad DROP CONSTRAINT IF EXISTS fk_documento_entidad_documento_escrito CASCADE;
ALTER TABLE academica.documento_entidad ADD CONSTRAINT fk_documento_entidad_documento_escrito FOREIGN KEY (documento_escrito)
REFERENCES academica.documento_escrito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_documento_entidad_documento_escrito ON academica.documento_entidad  IS E'Constraint foreign key fk_documento_entidad_documento_escrito';
-- ddl-end --


-- object: fk_documento_solicitud_documento_escrito | type: CONSTRAINT --
-- ALTER TABLE academica.documento_solicitud DROP CONSTRAINT IF EXISTS fk_documento_solicitud_documento_escrito CASCADE;
ALTER TABLE academica.documento_solicitud ADD CONSTRAINT fk_documento_solicitud_documento_escrito FOREIGN KEY (documento_escrito)
REFERENCES academica.documento_escrito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_documento_solicitud_documento_escrito ON academica.documento_solicitud  IS E'Constraint foreign key fk_documento_solicitud_documento_escrito';
-- ddl-end --


-- object: fk_documento_solicitud_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.documento_solicitud DROP CONSTRAINT IF EXISTS fk_documento_solicitud_solicitud_trabajo_grado CASCADE;
ALTER TABLE academica.documento_solicitud ADD CONSTRAINT fk_documento_solicitud_solicitud_trabajo_grado FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES academica.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_documento_solicitud_solicitud_trabajo_grado ON academica.documento_solicitud  IS E'Constraint foreign key fk_documento_solicitud_solicitud_trabajo_grado';
-- ddl-end --


-- object: fk_documento_trabajo_grado_documento_escrito | type: CONSTRAINT --
-- ALTER TABLE academica.documento_trabajo_grado DROP CONSTRAINT IF EXISTS fk_documento_trabajo_grado_documento_escrito CASCADE;
ALTER TABLE academica.documento_trabajo_grado ADD CONSTRAINT fk_documento_trabajo_grado_documento_escrito FOREIGN KEY (documento_escrito)
REFERENCES academica.documento_escrito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_documento_trabajo_grado_documento_escrito ON academica.documento_trabajo_grado  IS E'Constraint foreign key fk_documento_trabajo_grado_documento_escrito';
-- ddl-end --


-- object: fk_documento_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.documento_trabajo_grado DROP CONSTRAINT IF EXISTS fk_documento_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.documento_trabajo_grado ADD CONSTRAINT fk_documento_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_documento_trabajo_grado_trabajo_grado ON academica.documento_trabajo_grado  IS E'Constraint foreign key fk_documento_trabajo_grado_trabajo_grado';
-- ddl-end --


-- object: fk_espacio_academico_inscrito_espacios_academicos_elegibles | type: CONSTRAINT --
-- ALTER TABLE academica.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS fk_espacio_academico_inscrito_espacios_academicos_elegibles CASCADE;
ALTER TABLE academica.espacio_academico_inscrito ADD CONSTRAINT fk_espacio_academico_inscrito_espacios_academicos_elegibles FOREIGN KEY (espacios_academicos_elegibles)
REFERENCES academica.espacios_academicos_elegibles (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_espacio_academico_inscrito_espacios_academicos_elegibles ON academica.espacio_academico_inscrito  IS E'Constraint foreign key fk_espacio_academico_inscrito_espacios_academicos_elegibles';
-- ddl-end --


-- object: fk_espacio_academico_inscrito_estado_espacio_academico_inscrito | type: CONSTRAINT --
-- ALTER TABLE academica.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS fk_espacio_academico_inscrito_estado_espacio_academico_inscrito CASCADE;
ALTER TABLE academica.espacio_academico_inscrito ADD CONSTRAINT fk_espacio_academico_inscrito_estado_espacio_academico_inscrito FOREIGN KEY (estado_espacio_academico_inscrito)
REFERENCES academica.estado_espacio_academico_inscrito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_espacio_academico_inscrito_estado_espacio_academico_inscrito ON academica.espacio_academico_inscrito  IS E'Constraint foreign key fk_espacio_academico_inscrito_estado_espacio_academico_inscrito';
-- ddl-end --


-- object: fk_espacio_academico_inscrito_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS fk_espacio_academico_inscrito_trabajo_grado CASCADE;
ALTER TABLE academica.espacio_academico_inscrito ADD CONSTRAINT fk_espacio_academico_inscrito_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_espacio_academico_inscrito_trabajo_grado ON academica.espacio_academico_inscrito  IS E'Constraint foreign key fk_espacio_academico_inscrito_trabajo_grado';
-- ddl-end --


-- object: fk_espacios_academicos_elegibles_carrera_elegible | type: CONSTRAINT --
-- ALTER TABLE academica.espacios_academicos_elegibles DROP CONSTRAINT IF EXISTS fk_espacios_academicos_elegibles_carrera_elegible CASCADE;
ALTER TABLE academica.espacios_academicos_elegibles ADD CONSTRAINT fk_espacios_academicos_elegibles_carrera_elegible FOREIGN KEY (carrera_elegible)
REFERENCES academica.carrera_elegible (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_espacios_academicos_elegibles_carrera_elegible ON academica.espacios_academicos_elegibles  IS E'Constraint foreign key fk_espacios_academicos_elegibles_carrera_elegible';
-- ddl-end --


-- object: fk_estructura_investigacion_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.estructura_investigacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_estructura_investigacion_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.estructura_investigacion_trabajo_grado ADD CONSTRAINT fk_estructura_investigacion_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_estructura_investigacion_trabajo_grado_trabajo_grado ON academica.estructura_investigacion_trabajo_grado  IS E'Constraint foreign key fk_estructura_investigacion_trabajo_grado_trabajo_grado';
-- ddl-end --


-- object: fk_estudiante_trabajo_grado_estado_estudiante_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.estudiante_trabajo_grado DROP CONSTRAINT IF EXISTS fk_estudiante_trabajo_grado_estado_estudiante_trabajo_grado CASCADE;
ALTER TABLE academica.estudiante_trabajo_grado ADD CONSTRAINT fk_estudiante_trabajo_grado_estado_estudiante_trabajo_grado FOREIGN KEY (estado_estudiante_trabajo_grado)
REFERENCES academica.estado_estudiante_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_estudiante_trabajo_grado_estado_estudiante_trabajo_grado ON academica.estudiante_trabajo_grado  IS E'Constraint foreign key fk_estudiante_trabajo_grado_estado_estudiante_trabajo_grado';
-- ddl-end --


-- object: fk_estudiante_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.estudiante_trabajo_grado DROP CONSTRAINT IF EXISTS fk_estudiante_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.estudiante_trabajo_grado ADD CONSTRAINT fk_estudiante_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_estudiante_trabajo_grado_trabajo_grado ON academica.estudiante_trabajo_grado  IS E'Constraint foreign key fk_estudiante_trabajo_grado_trabajo_grado';
-- ddl-end --


-- object: fk_evaluacion_trabajo_grado_formato_evaluacion_carrera | type: CONSTRAINT --
-- ALTER TABLE academica.evaluacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_evaluacion_trabajo_grado_formato_evaluacion_carrera CASCADE;
ALTER TABLE academica.evaluacion_trabajo_grado ADD CONSTRAINT fk_evaluacion_trabajo_grado_formato_evaluacion_carrera FOREIGN KEY (formato_evaluacion_carrera)
REFERENCES academica.formato_evaluacion_carrera (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_evaluacion_trabajo_grado_formato_evaluacion_carrera ON academica.evaluacion_trabajo_grado  IS E'Constraint foreign key fk_evaluacion_trabajo_grado_formato_evaluacion_carrera';
-- ddl-end --


-- object: fk_evaluacion_trabajo_grado_socializacion | type: CONSTRAINT --
-- ALTER TABLE academica.evaluacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_evaluacion_trabajo_grado_socializacion CASCADE;
ALTER TABLE academica.evaluacion_trabajo_grado ADD CONSTRAINT fk_evaluacion_trabajo_grado_socializacion FOREIGN KEY (socializacion)
REFERENCES academica.socializacion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_evaluacion_trabajo_grado_socializacion ON academica.evaluacion_trabajo_grado  IS E'Constraint foreign key fk_evaluacion_trabajo_grado_socializacion';
-- ddl-end --


-- object: fk_evaluacion_trabajo_grado_vinculacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.evaluacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_evaluacion_trabajo_grado_vinculacion_trabajo_grado CASCADE;
ALTER TABLE academica.evaluacion_trabajo_grado ADD CONSTRAINT fk_evaluacion_trabajo_grado_vinculacion_trabajo_grado FOREIGN KEY (vinculacion_trabajo_grado)
REFERENCES academica.vinculacion_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_evaluacion_trabajo_grado_vinculacion_trabajo_grado ON academica.evaluacion_trabajo_grado  IS E'Constraint foreign key fk_evaluacion_trabajo_grado_vinculacion_trabajo_grado';
-- ddl-end --


-- object: fk_formato_evaluacion_carrera_formato | type: CONSTRAINT --
-- ALTER TABLE academica.formato_evaluacion_carrera DROP CONSTRAINT IF EXISTS fk_formato_evaluacion_carrera_formato CASCADE;
ALTER TABLE academica.formato_evaluacion_carrera ADD CONSTRAINT fk_formato_evaluacion_carrera_formato FOREIGN KEY (formato)
REFERENCES academica.formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_formato_evaluacion_carrera_formato ON academica.formato_evaluacion_carrera  IS E'Constraint foreign key fk_formato_evaluacion_carrera_formato';
-- ddl-end --


-- object: fk_formato_evaluacion_carrera_modalidad | type: CONSTRAINT --
-- ALTER TABLE academica.formato_evaluacion_carrera DROP CONSTRAINT IF EXISTS fk_formato_evaluacion_carrera_modalidad CASCADE;
ALTER TABLE academica.formato_evaluacion_carrera ADD CONSTRAINT fk_formato_evaluacion_carrera_modalidad FOREIGN KEY (modalidad)
REFERENCES academica.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_formato_evaluacion_carrera_modalidad ON academica.formato_evaluacion_carrera  IS E'Constraint foreign key fk_formato_evaluacion_carrera_modalidad';
-- ddl-end --


-- object: fk_modalidad_tipo_solicitud_modalidad | type: CONSTRAINT --
-- ALTER TABLE academica.modalidad_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_modalidad_tipo_solicitud_modalidad CASCADE;
ALTER TABLE academica.modalidad_tipo_solicitud ADD CONSTRAINT fk_modalidad_tipo_solicitud_modalidad FOREIGN KEY (modalidad)
REFERENCES academica.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_modalidad_tipo_solicitud_modalidad ON academica.modalidad_tipo_solicitud  IS E'Constraint foreign key fk_modalidad_tipo_solicitud_modalidad';
-- ddl-end --


-- object: fk_modalidad_tipo_solicitud_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE academica.modalidad_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_modalidad_tipo_solicitud_tipo_solicitud CASCADE;
ALTER TABLE academica.modalidad_tipo_solicitud ADD CONSTRAINT fk_modalidad_tipo_solicitud_tipo_solicitud FOREIGN KEY (tipo_solicitud)
REFERENCES academica.tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_modalidad_tipo_solicitud_tipo_solicitud ON academica.modalidad_tipo_solicitud  IS E'Constraint foreign key fk_modalidad_tipo_solicitud_tipo_solicitud';
-- ddl-end --


-- object: fk_pregunta_formato_formato | type: CONSTRAINT --
-- ALTER TABLE academica.pregunta_formato DROP CONSTRAINT IF EXISTS fk_pregunta_formato_formato CASCADE;
ALTER TABLE academica.pregunta_formato ADD CONSTRAINT fk_pregunta_formato_formato FOREIGN KEY (formato)
REFERENCES academica.formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_pregunta_formato_formato ON academica.pregunta_formato  IS E'Constraint foreign key fk_pregunta_formato_formato';
-- ddl-end --


-- object: fk_pregunta_formato_pregunta | type: CONSTRAINT --
-- ALTER TABLE academica.pregunta_formato DROP CONSTRAINT IF EXISTS fk_pregunta_formato_pregunta CASCADE;
ALTER TABLE academica.pregunta_formato ADD CONSTRAINT fk_pregunta_formato_pregunta FOREIGN KEY (pregunta)
REFERENCES academica.pregunta (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_pregunta_formato_pregunta ON academica.pregunta_formato  IS E'Constraint foreign key fk_pregunta_formato_pregunta';
-- ddl-end --


-- object: fk_pregunta_formato_tipo_pregunta | type: CONSTRAINT --
-- ALTER TABLE academica.pregunta_formato DROP CONSTRAINT IF EXISTS fk_pregunta_formato_tipo_pregunta CASCADE;
ALTER TABLE academica.pregunta_formato ADD CONSTRAINT fk_pregunta_formato_tipo_pregunta FOREIGN KEY (tipo_pregunta)
REFERENCES academica.tipo_pregunta (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_pregunta_formato_tipo_pregunta ON academica.pregunta_formato  IS E'Constraint foreign key fk_pregunta_formato_tipo_pregunta';
-- ddl-end --


-- object: fk_respuesta_evaluacion_evaluacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.respuesta_evaluacion DROP CONSTRAINT IF EXISTS fk_respuesta_evaluacion_evaluacion_trabajo_grado CASCADE;
ALTER TABLE academica.respuesta_evaluacion ADD CONSTRAINT fk_respuesta_evaluacion_evaluacion_trabajo_grado FOREIGN KEY (evaluacion_trabajo_grado)
REFERENCES academica.evaluacion_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_respuesta_evaluacion_evaluacion_trabajo_grado ON academica.respuesta_evaluacion  IS E'Constraint foreign key fk_respuesta_evaluacion_evaluacion_trabajo_grado';
-- ddl-end --


-- object: fk_respuesta_evaluacion_respuesta_formato | type: CONSTRAINT --
-- ALTER TABLE academica.respuesta_evaluacion DROP CONSTRAINT IF EXISTS fk_respuesta_evaluacion_respuesta_formato CASCADE;
ALTER TABLE academica.respuesta_evaluacion ADD CONSTRAINT fk_respuesta_evaluacion_respuesta_formato FOREIGN KEY (respuesta_formato)
REFERENCES academica.respuesta_formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_respuesta_evaluacion_respuesta_formato ON academica.respuesta_evaluacion  IS E'Constraint foreign key fk_respuesta_evaluacion_respuesta_formato';
-- ddl-end --


-- object: fk_respuesta_formato_pregunta_formato | type: CONSTRAINT --
-- ALTER TABLE academica.respuesta_formato DROP CONSTRAINT IF EXISTS fk_respuesta_formato_pregunta_formato CASCADE;
ALTER TABLE academica.respuesta_formato ADD CONSTRAINT fk_respuesta_formato_pregunta_formato FOREIGN KEY (pregunta_formato)
REFERENCES academica.pregunta_formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_respuesta_formato_pregunta_formato ON academica.respuesta_formato  IS E'Constraint foreign key fk_respuesta_formato_pregunta_formato';
-- ddl-end --


-- object: fk_respuesta_formato_respuesta | type: CONSTRAINT --
-- ALTER TABLE academica.respuesta_formato DROP CONSTRAINT IF EXISTS fk_respuesta_formato_respuesta CASCADE;
ALTER TABLE academica.respuesta_formato ADD CONSTRAINT fk_respuesta_formato_respuesta FOREIGN KEY (respuesta)
REFERENCES academica.respuesta (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_respuesta_formato_respuesta ON academica.respuesta_formato  IS E'Constraint foreign key fk_respuesta_formato_respuesta';
-- ddl-end --


-- object: fk_respuesta_solicitud_estado_solicitud | type: CONSTRAINT --
-- ALTER TABLE academica.respuesta_solicitud DROP CONSTRAINT IF EXISTS fk_respuesta_solicitud_estado_solicitud CASCADE;
ALTER TABLE academica.respuesta_solicitud ADD CONSTRAINT fk_respuesta_solicitud_estado_solicitud FOREIGN KEY (estado_solicitud)
REFERENCES academica.estado_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_respuesta_solicitud_estado_solicitud ON academica.respuesta_solicitud  IS E'Constraint foreign key fk_respuesta_solicitud_estado_solicitud';
-- ddl-end --


-- object: fk_respuesta_solicitud_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.respuesta_solicitud DROP CONSTRAINT IF EXISTS fk_respuesta_solicitud_solicitud_trabajo_grado CASCADE;
ALTER TABLE academica.respuesta_solicitud ADD CONSTRAINT fk_respuesta_solicitud_solicitud_trabajo_grado FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES academica.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_respuesta_solicitud_solicitud_trabajo_grado ON academica.respuesta_solicitud  IS E'Constraint foreign key fk_respuesta_solicitud_solicitud_trabajo_grado';
-- ddl-end --


-- object: fk_revision_trabajo_grado_estado_revision_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.revision_trabajo_grado DROP CONSTRAINT IF EXISTS fk_revision_trabajo_grado_estado_revision_trabajo_grado CASCADE;
ALTER TABLE academica.revision_trabajo_grado ADD CONSTRAINT fk_revision_trabajo_grado_estado_revision_trabajo_grado FOREIGN KEY (estado_revision_trabajo_grado)
REFERENCES academica.estado_revision_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_revision_trabajo_grado_estado_revision_trabajo_grado ON academica.revision_trabajo_grado  IS E'Constraint foreign key fk_revision_trabajo_grado_estado_revision_trabajo_grado';
-- ddl-end --


-- object: fk_revision_trabajo_grado_vinculacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.revision_trabajo_grado DROP CONSTRAINT IF EXISTS fk_revision_trabajo_grado_vinculacion_trabajo_grado CASCADE;
ALTER TABLE academica.revision_trabajo_grado ADD CONSTRAINT fk_revision_trabajo_grado_vinculacion_trabajo_grado FOREIGN KEY (vinculacion_trabajo_grado)
REFERENCES academica.vinculacion_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_revision_trabajo_grado_vinculacion_trabajo_grado ON academica.revision_trabajo_grado  IS E'Constraint foreign key fk_revision_trabajo_grado_vinculacion_trabajo_grado';
-- ddl-end --


-- object: fk_socializacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.socializacion DROP CONSTRAINT IF EXISTS fk_socializacion_trabajo_grado CASCADE;
ALTER TABLE academica.socializacion ADD CONSTRAINT fk_socializacion_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_socializacion_trabajo_grado ON academica.socializacion  IS E'Constraint foreign key fk_socializacion_trabajo_grado';
-- ddl-end --


-- object: fk_solicitud_trabajo_grado_modalidad_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE academica.solicitud_trabajo_grado DROP CONSTRAINT IF EXISTS fk_solicitud_trabajo_grado_modalidad_tipo_solicitud CASCADE;
ALTER TABLE academica.solicitud_trabajo_grado ADD CONSTRAINT fk_solicitud_trabajo_grado_modalidad_tipo_solicitud FOREIGN KEY (modalidad_tipo_solicitud)
REFERENCES academica.modalidad_tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_solicitud_trabajo_grado_modalidad_tipo_solicitud ON academica.solicitud_trabajo_grado  IS E'Constraint foreign key fk_solicitud_trabajo_grado_modalidad_tipo_solicitud';
-- ddl-end --


-- object: fk_solicitud_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.solicitud_trabajo_grado DROP CONSTRAINT IF EXISTS fk_solicitud_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.solicitud_trabajo_grado ADD CONSTRAINT fk_solicitud_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_solicitud_trabajo_grado_trabajo_grado ON academica.solicitud_trabajo_grado  IS E'Constraint foreign key fk_solicitud_trabajo_grado_trabajo_grado';
-- ddl-end --


-- object: fk_trabajo_grado_distincion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_distincion_trabajo_grado CASCADE;
ALTER TABLE academica.trabajo_grado ADD CONSTRAINT fk_trabajo_grado_distincion_trabajo_grado FOREIGN KEY (distincion_trabajo_grado)
REFERENCES academica.distincion_trabajo_grado (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_trabajo_grado_distincion_trabajo_grado ON academica.trabajo_grado  IS E'Constraint foreign key fk_trabajo_grado_distincion_trabajo_grado';
-- ddl-end --


-- object: fk_trabajo_grado_estado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_estado_trabajo_grado CASCADE;
ALTER TABLE academica.trabajo_grado ADD CONSTRAINT fk_trabajo_grado_estado_trabajo_grado FOREIGN KEY (estado_trabajo_grado)
REFERENCES academica.estado_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_trabajo_grado_estado_trabajo_grado ON academica.trabajo_grado  IS E'Constraint foreign key fk_trabajo_grado_estado_trabajo_grado';
-- ddl-end --


-- object: fk_trabajo_grado_modalidad | type: CONSTRAINT --
-- ALTER TABLE academica.trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_modalidad CASCADE;
ALTER TABLE academica.trabajo_grado ADD CONSTRAINT fk_trabajo_grado_modalidad FOREIGN KEY (modalidad)
REFERENCES academica.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_trabajo_grado_modalidad ON academica.trabajo_grado  IS E'Constraint foreign key fk_trabajo_grado_modalidad';
-- ddl-end --


-- object: fk_usuario_solicitud_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.usuario_solicitud DROP CONSTRAINT IF EXISTS fk_usuario_solicitud_solicitud_trabajo_grado CASCADE;
ALTER TABLE academica.usuario_solicitud ADD CONSTRAINT fk_usuario_solicitud_solicitud_trabajo_grado FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES academica.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_usuario_solicitud_solicitud_trabajo_grado ON academica.usuario_solicitud  IS E'Constraint foreign key fk_usuario_solicitud_solicitud_trabajo_grado';
-- ddl-end --


-- object: fk_vinculacion_trabajo_grado_rol_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.vinculacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_vinculacion_trabajo_grado_rol_trabajo_grado CASCADE;
ALTER TABLE academica.vinculacion_trabajo_grado ADD CONSTRAINT fk_vinculacion_trabajo_grado_rol_trabajo_grado FOREIGN KEY (rol_trabajo_grado)
REFERENCES academica.rol_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_vinculacion_trabajo_grado_rol_trabajo_grado ON academica.vinculacion_trabajo_grado  IS E'Constraint foreign key fk_vinculacion_trabajo_grado_rol_trabajo_grado';
-- ddl-end --


-- object: fk_vinculacion_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE academica.vinculacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_vinculacion_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE academica.vinculacion_trabajo_grado ADD CONSTRAINT fk_vinculacion_trabajo_grado_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES academica.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
COMMENT ON CONSTRAINT fk_vinculacion_trabajo_grado_trabajo_grado ON academica.vinculacion_trabajo_grado  IS E'Constraint foreign key fk_vinculacion_trabajo_grado_trabajo_grado';
-- ddl-end --



