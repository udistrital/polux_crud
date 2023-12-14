ALTER TABLE academica.correccion
ADD COLUMN IF NOT EXISTS enlace_documento VARCHAR(36) DEFAULT '';

ALTER TABLE academica.correccion
ALTER COLUMN enlace_documento DROP DEFAULT;
