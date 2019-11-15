# recibo_crud
API para gestión de recibos para Campus Virtual

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `recibo_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/recibo_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `RECIBO_CRUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `RECIBO_CRUD__PGUSER`: Usuario de la base de datos
 - `RECIBO_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `RECIBO_CRUD__PGURLS`: Host de conexión
 - `RECIBO_CRUD__PGDB`: Nombre de la base de datos
 - `RECIBO_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
RECIBO_CRUD_HTTP_PORT=8113 RECIBO_CRUD__PGUSER=user RECIBO_CRUD__PGPASS=password RECIBO_CRUD__PGURLS=localhost RECIBO_CRUD__PGDB=bd RECIBO_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/recibo_crud/blob/develop/modelo_recibo_crud.png).
