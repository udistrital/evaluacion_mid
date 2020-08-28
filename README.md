# evaluacion_mid

Api intermediaria entre el cliente de Evaluaciones y las apis necesarios para la gestión de la información de proveedores y evaluaciones para estos mismos.  
El api principalmente es pensado para darle informacion a [evaluacion_cliente](https://github.com/udistrital/evaluacion_cliente). (esto no limita a su consumo desde el clinete o api que desee consumirle)  


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
# Ejemplo que se debe actualizar acorde al proyecto
EVALUACIONES_MID_DB_HOST = [descripción]
EVALUACIONES_MID_DB_NAME = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con EVALUACIONES_MID_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/evaluacion_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/evaluacion_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
EVALUACIONES_MID_PORT=8080 EVALUACIONES_MID_DB_HOST=127.0.0.1:27017 EVALUACIONES_MID_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=evaluacion_mid . --no-cache
# docker run -p 80:80 evaluacion_mid
```


### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/evaluacion_mid

#2. Moverse a la carpeta del repositorio
cd evaluacion_mid

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
# Not Data
```

### Diagramas

ac continuacion se visualizaran los diagramas inicialmente planteados para el manejo de informacion y funciones principales.

<details>
    <summary><b>Post plantilla</b></summary>

![diagrama para plantillas-post_plantilla](https://user-images.githubusercontent.com/28914781/69213926-891f7800-0b33-11ea-81ee-fc63c0de7c60.png)


</details>

<details>
    <summary><b>Get plantilla</b></summary>

![diagrama para plantillas-get_plantilla](https://user-images.githubusercontent.com/28914781/69214069-ea474b80-0b33-11ea-8214-83063252521e.png)


</details>

<details>
    <summary><b>Filtro por contratista</b></summary>


![flujos Agora-filtro-contratista](https://user-images.githubusercontent.com/28914781/69214107-0945dd80-0b34-11ea-96e0-874996dd9d3d.png)


</details>

<details>
    <summary><b>Filtro por contrato y vigencia (vigencia opcinal)</b></summary>


![flujos Agora-filtro-contrato](https://user-images.githubusercontent.com/28914781/69214152-25497f00-0b34-11ea-9801-532125f2b20d.png)


</details>

<details>
    <summary><b>Filtro por contratista, contrato y vigencia (vigencia opcional)</b></summary>

![flujos Agora-filtro-contratista-contrato](https://user-images.githubusercontent.com/28914781/69214179-34303180-0b34-11ea-90ef-1fb375602e18.png)

</details>


## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/evaluacion_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/evaluacion_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/evaluacion_mid/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/evaluacion_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/evaluacion_mid/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/evaluacion_mid) |


## Licencia

This file is part of evaluacion_mid.

evaluacion_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

evaluacion_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with evaluacion_mid. If not, see https://www.gnu.org/licenses/.
