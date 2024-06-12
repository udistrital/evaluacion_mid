package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/evaluacion_mid/models"
)

// Dado un numero de documento busca la informacion de un docente
func InformacionDve(numeroDocumento string) (docente models.InformacionDVE, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener los datos del docente",
				"Funcion": "InformacionDve",
			}
			panic(outputError)
		}
	}()

	var respuesta_peticion []map[string]interface{}
	var informacion_proveedor models.InformacionProveedor

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numeroDocumento, &respuesta_peticion); err == nil && response == 200 {

		if len(respuesta_peticion) == 0 {
			return docente, map[string]interface{}{
				"Succes":  false,
				"Status":  204,
				"Message": "No se encontro el docente",
				"Funcion": "InformacionDve",
			}
		} else {

			proveedorJson, err := json.Marshal(respuesta_peticion[0])
			if err != nil {
				return docente, map[string]interface{}{
					"Succes":  false,
					"Status":  502,
					"Message": "Error al obtener los datos del docente",
					"Funcion": "InformacionDve",
				}
			} else {
				err = json.Unmarshal(proveedorJson, &informacion_proveedor)
				if err != nil {
					return docente, map[string]interface{}{
						"Succes":  false,
						"Status":  502,
						"Message": "Error al obtener los datos del docente",
						"Funcion": "InformacionDve",
					}
				} else {
					docente = models.InformacionDVE{
						Activo:             "Activo",
						NombreDocente:      informacion_proveedor.NomProveedor,
						NumeroDocumento:    informacion_proveedor.NumDocumento,
						NivelAcademico:     "Titular",
						Facultad:           "Tecnologica",
						ProyectoCurricular: "Ingenieria en telemática",
					}
				}
			}
		}

	} else {
		return docente, map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al obtener los datos del docente",
			"Funcion": "InformacionDve",
		}
	}
	return
}

func IntensidadHorariaDve(numeroDocumento []string, periodoInicial []string, periodoFinal []string, vinculaciones []string) (intensidadHoraria []models.IntensidadHorariaDVE, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener la intensidad horaria del docente",
				"Funcion": "IntensidadHorariaDve",
			}
			panic(outputError)
		}
	}()

	intensidad1 := models.IntensidadHorariaDVE{
		Ano:              "2002",
		Periodo:          "II",
		NombreAsignatura: "Tecnologia en mecanica",
		HorasSemana:      "16",
		NumeroSemanas:    "18",
		HorasSemestrales: "288",
		SalarioDocente:   "2000000",
	}

	intensidad2 := models.IntensidadHorariaDVE{
		Ano:              "2003",
		Periodo:          "I",
		NombreAsignatura: "Tecnologia en mecanica",
		HorasSemana:      "16",
		NumeroSemanas:    "18",
		HorasSemestrales: "288",
		SalarioDocente:   "2000000",
	}

	intensidad3 := models.IntensidadHorariaDVE{
		Ano:              "2003",
		Periodo:          "II",
		NombreAsignatura: "Tecnologia en mecanica",
		HorasSemana:      "16",
		NumeroSemanas:    "18",
		HorasSemestrales: "288",
		SalarioDocente:   "2000000",
	}

	intensidadHoraria = append(intensidadHoraria, intensidad1)
	intensidadHoraria = append(intensidadHoraria, intensidad2)
	intensidadHoraria = append(intensidadHoraria, intensidad3)

	return intensidadHoraria, nil
}

//helper que retorna la informacion de un docente y el listado de intensidades horarias en un solo objeto
func InformacionCertificacionDve(numeroDocumento []string, periodoInicial []string, periodoFinal []string, vinculaciones []string) (certificacion models.InformacionCertificacionDve, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener la informacion de la certificacion del docente",
				"Funcion": "InformacionCertificacionDve",
			}
			panic(outputError)
		}
	}()

	fmt.Println("Numero de documento: ", numeroDocumento[0])
	docente, errorDocente := InformacionDve(numeroDocumento[0])
	fmt.Println("Docente: ", docente)
	if errorDocente != nil {
		return certificacion, errorDocente
	}

	intensidadHoraria, errorIntensidadHoraria := IntensidadHorariaDve(numeroDocumento, periodoInicial, periodoFinal, vinculaciones)
	if errorIntensidadHoraria != nil {
		return certificacion, errorIntensidadHoraria
	}

	certificacion = models.InformacionCertificacionDve{
		InformacionDve:    docente,
		IntensidadHoraria: intensidadHoraria,
	}

	return certificacion, nil
}
