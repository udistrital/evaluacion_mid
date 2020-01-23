# evaluacion_mid

Api intermediaria entre el cliente de Evaluaciones y las apis necesarios para la gestion de la informacion de proveedores y evaluaciones para estos mismos.

El api principalmente es pensado para darle informacion a [evaluacion_cliente](https://github.com/udistrital/evaluacion_cliente). (esto no limita a su consumo desde el clinete o api que desee consumirle)


Al momento de Crear el presente Readme hace consumo de las siguientes Apis:

- [evaluaciones_crud](https://github.com/udistrital/evaluacion_crud)
- [administrativa_amazon_api](https://github.com/udistrital/administrativa_amazon_api)

## Especificación Técnica

 Para instalar el proyecto de debe relizar lo siguientes pasos:

Ejecutar desde la terminal 'go get repositorio':
```shell
go get github.com/udistrital/evaluacion_mid
```

### Variables de Entorno

- EVALUACIONES_MID_HTTP_PORT=[puerto de ejecucion del api]
- ADMIN_AMAZON_API_URL=[direccion de api administrativa_amazon_api]
- ADMIN_AMAZON_JBPM_URL=[direccion de administrativa_jbpm]
- EVALUACION_CRUD_URL=[direccion del api evaluacion_crud]


### Ejecución del proyecto

- Ejecutar:
```shell
EVALUACIONES_MID_HTTP_PORT=XXX ADMIN_AMAZON_API_URL=XXX ADMIN_AMAZON_JBPM_URL=XXX EVALUACION_CRUD_URL=XXX bee run
```
- O si se quiere ejecutar el swager:
```shell
EVALUACIONES_MID_HTTP_PORT=XXX ADMIN_AMAZON_API_URL=XXX ADMIN_AMAZON_JBPM_URL=XXX EVALUACION_CRUD_URL=XXX bee run -downdoc=true -gendoc=true
```

---
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

### EndPoints

Al ejecutar el swagger se puede tener mayor apreciacion de los diferentes metodos de peticion por cada endpoint cuales son los distinpos endpoint disponibles y como usarlos.



## Licencia

This file is part of evaluacion_mid.

evaluacion_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.
