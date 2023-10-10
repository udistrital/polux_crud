ALTER TABLE academica.detalle_trabajo_grado ADD COLUMN if not exists activo boolean;
ALTER TABLE academica.detalle_trabajo_grado ADD COLUMN if not exists fecha_creacion timestamp;
ALTER TABLE academica.detalle_trabajo_grado ADD COLUMN if not exists fecha_modificacion timestamp;

COMMENT ON COLUMN academica.detalle_trabajo_grado.activo IS 'Campo parametrico que indica si el registro está activo';
COMMENT ON COLUMN academica.detalle_trabajo_grado.fecha_creacion IS 'Campo parametrico que indica la fecha de creación del registro';
COMMENT ON COLUMN academica.detalle_trabajo_grado.fecha_modificacion IS 'Campo parametrico que indica la fecha de la ultima modificación del registro';

ALTER TABLE academica.asignatura_trabajo_grado ADD COLUMN if not exists activo boolean;
ALTER TABLE academica.asignatura_trabajo_grado ADD COLUMN if not exists fecha_creacion timestamp;
ALTER TABLE academica.asignatura_trabajo_grado ADD COLUMN if not exists fecha_modificacion timestamp;

COMMENT ON COLUMN academica.asignatura_trabajo_grado.activo IS 'Campo parametrico que indica si el registro está activo';
COMMENT ON COLUMN academica.asignatura_trabajo_grado.fecha_creacion IS 'Campo parametrico que indica la fecha de creación del registro';
COMMENT ON COLUMN academica.asignatura_trabajo_grado.fecha_modificacion IS 'Campo parametrico que indica la fecha de la ultima modificación del registro';