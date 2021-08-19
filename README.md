# 57blocks Movies

instale el repositorio

`git clone https://github.com/jeffleon/57blocks`

## API REST
para que la api rest funcione por favor ubicate en la carpeta Microservices

la aplicacion cuenta con dos microservicios con sus respectivas bases de datos y con su docker file

para correr el proyecto se debe correr el docker-compose

`docker-compose up --force-recreate -d`

* la api corre en los puertos 8000 y 9000

si se quiere bajar los contenedores

`docker-compose down`

### EndPoints

### User Service

* __POST__  _/_
* __POST__  _/login_
* __GET__  _/user_
* __POST__  _/logout_
* __POST__  _/user/:id_
* __GET__  _/user/:id_

### Movies Service

* __GET__  _/user/:id/movies_
* __GET__  _/user/movie/:id_
* __POST__  _/user/movie_
* __POST__  _/user/movie/:id_
* __DELETE__  _/user/movie/:id_

### Depenpendencias usadas

si por alguna razon la api no corre se utilizaron las siguientes dependencias

* `go get github.com/gofiber/fiber/v2`

* `go get github.com/jinzhu/gorm`

* `go get gorm.io/driver/mysql`

* `go get database/sql`
