CREATE TABLE academica.detalle_trabajo_grado (
	id integer NOT NULL,
	parametro integer NOT NULL,
	valor text NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT detalle_trabajo_grado_pk PRIMARY KEY (id),
	CONSTRAINT trabajo_grado_fk FOREIGN KEY (trabajo_grado) REFERENCES academica.trabajo_grado (id)
);
