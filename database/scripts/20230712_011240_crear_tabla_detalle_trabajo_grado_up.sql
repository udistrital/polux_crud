CREATE TABLE IF NOT EXISTS academica.detalle_trabajo_grado (
	id integer NOT NULL,
	parametro integer NOT NULL,
	valor text NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT detalle_trabajo_grado_pk PRIMARY KEY (id),
	CONSTRAINT trabajo_grado_fk FOREIGN KEY (trabajo_grado) REFERENCES academica.trabajo_grado (id)
);

COMMENT ON TABLE academica.detalle_trabajo_grado IS 'Tabla para almacenar los detalles del trabajo de grado de todas las modalidades';
COMMENT ON COLUMN academica.detalle_trabajo_grado.id IS 'Identificador de la tabla detalle_trabajo_grado';
COMMENT ON COLUMN academica.detalle_trabajo_grado.parametro IS 'parametro al cual se hará referencia, ya que guardará el id de la tabla parametro en la base de datos udistrital_core';
COMMENT ON COLUMN academica.detalle_trabajo_grado.valor IS 'El valor que llevara dicho parametro';
COMMENT ON COLUMN academica.detalle_trabajo_grado.trabajo_grado IS 'Llave foranea que hará referencia al trabajo de grado relacionado';
