#POLUX_CRUD
API CRUD para el consumo de datos del módulo UDISTRITAL_CORE de la BD, referente al modelo de negocio de Academica de la Universidad Distrital.

##Especificaciones Técnicas
Tecnologías Implementadas y Versiones
*.[Golang]_(https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
*.[Beego]_(https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
*.[Docker]_(https://docs.docker.com/engine/install/ubuntu/)
*.[Docker_compose]_(https://docs.docker.com/compose/)

##Variables de Entorno
´´caparazon

httpport = [Puerto asignado para la ejecución del API]
PGuser = [Usuario de la base de datos]
PGpass = [Clave del usuario para la conexión a la base de datos]
PGurls = [Host de conexión]
PGdb   =  [Nombre de la base de datos]
PGschemas = [Esquema a utilizar en la base de datos]
logPath = [Archivo para logs]
RUN_MODE = [Modo de ejecución del api]

# Ejemplo
httpport = 8001
PGuser = postgres
PGpass = *******
PGurls = 127.0.0.1
PGdb   = udistrital_core_db
PGschemas = academica
logPath = ${POLUX_CRUD__LOG}
RUN_MODE = dev
''
NOTA: Las variables se pueden ver en el fichero conf / app.conf y están identificadas con POLUX_CRUD__DB
Para definir puertos, dns y configuraciones internas dentro del archivo .env
Para definir conexiones externas a otros apis se debe crear el archivo custom.env en la raiz del proyecto

Ejecución del Proyecto
´´caparazon
# 1. Obtener el repositorio con Go
ve a github.com/udistrital/polux_crud

# 2. Moverse a la carpeta del repositorio 
cd  $ GOPATH /src/github.com/udistrital/polux_crud

# 3. Moverse a la rama ** desarrollar ** 
git pull origin desarrollar && git checkout desarrollar

# 4. alimentar todas las variables de entorno que utiliza el proyecto. 
AWS_DEFAULT_REGION=us-east-1 POLUX_CRUD__DB_USER=desarrollooas POLUX_CRUD__DB_PASS=W7Sz1lbWFwfE798b PARAMETER_STORE=preprod POLUX_CRUD__DB_NAME=udistrital_core_db POLUX_CRUD__DB_SCHEMA=academica POLUX_CRUD__DB_URL=pgtst.udistritaloas.edu.co POLUX_CRUD__HTTP_PORT=8001 bee run
###Ejecución Dockerfile
´´caparazon
# docker build --tag = polux_crud. --no-cache 
# docker run -p 80:80 polux_crud
''
###Ejecución docker-compose
´´caparazon
# 1. Clonar el repositorio
git clone -b desarrollar https://github.com/udistrital/polux_crud

# 2. Moverse a la carpeta del repositorio 
cd polux_crud

# 3. Crear un fichero con el nombre ** custom.env ** 
# En windows ejecutar: * `ni custom.env`
toque custom.env

# 4. Crear la red ** back_end ** para los contenedores
red acoplable crear back_end

# 5. Ejecutar el compose del contenedor
docker-compose up --build

# 6. Comprobar que los contenedores estén en ejecución 
docker ps
''
###Ejecución Pruebas

Pruebas unitarias
´´caparazon
# Sin datos
''

##Estado CI

| develop | release 0.0.1 | master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/polux_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/polux_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/polux_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/polux_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/polux_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/polux_crud) |

## Modelo de Datos

[SVG](database/model.svg)

Licencia
Este archivo es parte de polux_crud.

polux_crud es software libre: puede redistribuirlo y / o modificarlo según los términos de la Licencia Pública General GNU publicada por la Free Software Foundation, ya sea la versión 3 de la Licencia o (a su elección) cualquier versión posterior.

polux_crud se distribuye con la esperanza de que sea útil, pero SIN NINGUNA GARANTÍA; incluso sin la garantía implícita de COMERCIABILIDAD o APTITUD PARA UN PROPÓSITO PARTICULAR. Consulte la Licencia pública general de GNU para obtener más detalles.

Debería haber recibido una copia de la Licencia Pública General GNU junto con polux_crud. De lo contrario, consulte https://www.gnu.org/licenses/ .
