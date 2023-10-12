ALTER TABLE academica.detalle_trabajo_grado DROP COLUMN IF EXISTS activo;
ALTER TABLE academica.detalle_trabajo_grado DROP COLUMN IF EXISTS fecha_creacion;
ALTER TABLE academica.detalle_trabajo_grado DROP COLUMN IF EXISTS fecha_modificacion;

ALTER TABLE academica.asignatura_trabajo_grado DROP COLUMN IF EXISTS activo;
ALTER TABLE academica.asignatura_trabajo_grado DROP COLUMN IF EXISTS fecha_creacion;
ALTER TABLE academica.asignatura_trabajo_grado DROP COLUMN IF EXISTS fecha_modificacion;