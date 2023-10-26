INSERT INTO academica.detalle_tipo_solicitud(detalle, modalidad_tipo_solicitud, activo, requerido, numero_orden)
VALUES (80, 13, true, true, 2);

UPDATE academica.detalle_tipo_solicitud
SET numero_orden=3
WHERE id=309;