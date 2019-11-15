
-- object: recibo | type: SCHEMA --
-- DROP SCHEMA IF EXISTS recibo CASCADE;
CREATE SCHEMA recibo;
-- ddl-end --

SET search_path TO pg_catalog,public,recibo;
-- ddl-end --

-- object: recibo.estado_recibo | type: TABLE --
-- DROP TABLE IF EXISTS recibo.estado_recibo CASCADE;
CREATE TABLE recibo.estado_recibo (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_estado_recibo PRIMARY KEY (id)

);
-- ddl-end --

-- object: recibo.tipo_recibo | type: TABLE --
-- DROP TABLE IF EXISTS recibo.tipo_recibo CASCADE;
CREATE TABLE recibo.tipo_recibo (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	fecha_modificacion timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT pk_tipo_recibo PRIMARY KEY (id)

);
-- ddl-end --

-- object: recibo.tipo_pago | type: TABLE --
-- DROP TABLE IF EXISTS recibo.tipo_pago CASCADE;
CREATE TABLE recibo.tipo_pago (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_pago PRIMARY KEY (id)

);
-- ddl-end --

-- object: recibo.recibo | type: TABLE --
-- DROP TABLE IF EXISTS recibo.recibo CASCADE;
CREATE TABLE recibo.recibo (
	id serial NOT NULL,
	referencia integer,
	valor_ordinario integer NOT NULL,
	valor_extraordinario integer,
	fecha_ordinaria date NOT NULL,
	fecha_extraordinaria date,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	estado_recibo_id integer NOT NULL,
	tipo_recibo_id integer NOT NULL,
	CONSTRAINT pk_recibo PRIMARY KEY (id)

);
-- ddl-end --

-- object: recibo.pago_recibo | type: TABLE --
-- DROP TABLE IF EXISTS recibo.pago_recibo CASCADE;
CREATE TABLE recibo.pago_recibo (
	id serial NOT NULL,
	fecha_pago date NOT NULL,
	aprobado boolean NOT NULL DEFAULT false,
	comprobante integer,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	recibo_id integer NOT NULL,
	tipo_pago_id integer NOT NULL,
	CONSTRAINT pk_pago_recibo PRIMARY KEY (id)

);
-- ddl-end --

-- object: fk_recibo_estado_recibo | type: CONSTRAINT --
-- ALTER TABLE recibo.recibo DROP CONSTRAINT IF EXISTS fk_recibo_estado_recibo CASCADE;
ALTER TABLE recibo.recibo ADD CONSTRAINT fk_recibo_estado_recibo FOREIGN KEY (estado_recibo_id)
REFERENCES recibo.estado_recibo (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_recibo_tipo_recibo | type: CONSTRAINT --
-- ALTER TABLE recibo.recibo DROP CONSTRAINT IF EXISTS fk_recibo_tipo_recibo CASCADE;
ALTER TABLE recibo.recibo ADD CONSTRAINT fk_recibo_tipo_recibo FOREIGN KEY (tipo_recibo_id)
REFERENCES recibo.tipo_recibo (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_pago_recibo_tipo_pago | type: CONSTRAINT --
-- ALTER TABLE recibo.pago_recibo DROP CONSTRAINT IF EXISTS fk_pago_recibo_tipo_pago CASCADE;
ALTER TABLE recibo.pago_recibo ADD CONSTRAINT fk_pago_recibo_tipo_pago FOREIGN KEY (tipo_pago_id)
REFERENCES recibo.tipo_pago (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_pago_recibo_recibo | type: CONSTRAINT --
-- ALTER TABLE recibo.pago_recibo DROP CONSTRAINT IF EXISTS fk_pago_recibo_recibo CASCADE;
ALTER TABLE recibo.pago_recibo ADD CONSTRAINT fk_pago_recibo_recibo FOREIGN KEY (recibo_id)
REFERENCES recibo.recibo (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA proyecto_academico TO desarrollooas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA proyecto_academico TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA proyecto_academico TO desarrollooas;
