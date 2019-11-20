# EVALUACION_MID

Api intermediaria entre el cliente de Evaluaciones y las apis necesarios para la gestion de la informacion de proveedores y evaluaciones para estos mismos.

El api principalmente es pensado para darle informacion a [evaluacion_cliente](https://github.com/udistrital/evaluacion_cliente). (esto no limita a su consumo desde el clinete o api que desee consumirle)


Al momento de Crear el presente Readme hace consumo de las siguientes Apis:

- [evaluaciones_crud](https://github.com/udistrital/evaluacion_crud)
- [administrativa_amazon_api](https://github.com/udistrital/administrativa_amazon_api)

# Instalación
Para instalar el proyecto de debe relizar lo siguientes pasos:

Ejecutar desde la terminal 'go get repositorio':
```shell 
go get github.com/udistrital/evaluacion_mid
```

# Ejecución del proyecto


- Ejecutar: 
```shell 
bee run
```
- O si se quiere ejecutar el swager:

```shell 
bee run -downdoc=true -gendoc=true
```


---

## DIAGRAMAS

ac continuacion se visualizaran los diagramas inicialmente planteados para el manejo de informacion y funciones principales.

<details>
    <summary><b>Ppost plantilla</b></summary>

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