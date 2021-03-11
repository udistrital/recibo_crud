# recibo_crud

API para gestión de recibos de pago de matrícula y otros para el Nuevo Sistema de Gestión Académica SGA

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)

### Variables de Entorno
```shell
RECIBO_CRUD_HTTP_PORT=[Puerto asignado para la ejecución del API]
RECIBO_CRUD__PGUSER=[Usuario de la base de datos]
RECIBO_CRUD__PGPASS=[Clave del usuario para la conexión a la base de datos]
RECIBO_CRUD__PGURLS=[Host de conexión]
RECIBO_CRUD__PGDB=[Nombre de la base de datos]
RECIBO_CRUD__SCHEMA=[Esquema a utilizar en la base de datos] bee run
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con RECIBO_CRUD__...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/recibo_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/recibo_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
RECIBO_CRUD_HTTP_PORT=8113 RECIBO_CRUD__PGUSER=user RECIBO_CRUD__PGPASS=password RECIBO_CRUD__PGURLS=localhost RECIBO_CRUD__PGDB=bd RECIBO_CRUD__SCHEMA=schema_new bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=recibo_crud . --no-cache
# docker run -p 80:80 recibo_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/recibo_crud

#2. Moverse a la carpeta del repositorio
cd recibo_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```

## Estado CI

| Develop | Release | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/recibo_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/recibo_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/recibo_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/recibo_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/recibo_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/recibo_crud/) |

## Modelo de Datos
![image](https://github.com/planesticud/recibo_crud/blob/develop/modelo_recibo_crud.png).


## Licencia

This file is part of recibo_crud.

recibo_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

recibo_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with recibo_crud. If not, see https://www.gnu.org/licenses/.

