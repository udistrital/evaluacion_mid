package helpers

import (
	"encoding/json"
	"strconv"
	"time"

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

	// Solicitud para obtener los datos del proveedor
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numeroDocumento, &respuesta_peticion); err == nil && response == 200 {
		if len(respuesta_peticion) > 0 {
			proveedorJson, err := json.Marshal(respuesta_peticion[0])
			if err != nil {
				return docente, outputError
			} else {
				// Convertir los datos al modelo de información proveedor
				if err = json.Unmarshal(proveedorJson, &informacion_proveedor); err != nil {
					return docente, outputError
				} else {
					var respuesta_contrato []map[string]interface{}
					var informacion_contrato models.ContratoGeneral
					// Validar que el docente sea de vinculación especial y tenga un contrato de prestación de servicios
					if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_general/?query=Contratista:"+strconv.Itoa(informacion_proveedor.Id), &respuesta_contrato); err == nil && response == 200 {
						if len(respuesta_contrato) > 0 {
							contratoJson, err := json.Marshal(respuesta_contrato[0])
							if err != nil {
								return docente, outputError
							} else {
								if err = json.Unmarshal(contratoJson, &informacion_contrato); err == nil && informacion_contrato.TipoContrato.Id == 18 {
									docente.NombreDocente = informacion_proveedor.NomProveedor
									docente.NumeroDocumento = informacion_proveedor.NumDocumento

									// Traer información de la vinculación del docente
									var respuesta_vinculacion map[string]interface{}
									order := "&order=desc"
									sortby := "&sortby=FechaCreacion"

									if response, err := getJsonTest(beego.AppConfig.String("UrlCrudResoluciones")+"/vinculacion_docente/?query=PersonaId:"+numeroDocumento+sortby+order, &respuesta_vinculacion); err == nil && response == 200 {
										if len(respuesta_vinculacion) > 0 {
											var vinculacionDocente models.VinculacionesDocente
											vinculacionJson, err := json.Marshal(respuesta_vinculacion["Data"].([]interface{})[0])
											if err != nil {
												return docente, outputError
											} else {
												if err = json.Unmarshal(vinculacionJson, &vinculacionDocente); err != nil {
													return docente, outputError
												} else {
													docente.Facultad, _ = NombreFacultadProyecto(vinculacionDocente.ResolucionVinculacionDocente.FacultadId)
													docente.ProyectoCurricular, _ = NombreFacultadProyecto(vinculacionDocente.ProyectoCurricularId)
													docente.NivelAcademico = vinculacionDocente.ResolucionVinculacionDocente.NivelAcademico
													docente.Categoria = vinculacionDocente.Categoria
													docente.Dedicacion = vinculacionDocente.ResolucionVinculacionDocente.Dedicacion

													var respuesta_resolucion map[string]interface{}
													var resolucion models.Resolucion
													if response, err := getJsonTest(beego.AppConfig.String("UrlCrudResoluciones")+"/resolucion/?query=Id:"+strconv.Itoa(vinculacionDocente.ResolucionVinculacionDocente.Id), &respuesta_resolucion); err == nil && response == 200 {
														if len(respuesta_resolucion) > 0 {
															resolucionJson, err := json.Marshal(respuesta_resolucion["Data"].([]interface{})[0])
															if err != nil {
																return docente, outputError
															} else {
																if err = json.Unmarshal(resolucionJson, &resolucion); err != nil {
																	return docente, outputError
																} else {
																	if resolucion.FechaInicio == nil || resolucion.FechaFin == nil {
																		docente.Activo = false

																	}
																	//Verificacr si la fecha actual es mas reciente que la fecha de inicio y la fecha de fin del contrato
																	if resolucion.FechaInicio != nil && resolucion.FechaFin != nil {
																		if (*resolucion.FechaInicio).Before(*resolucion.FechaFin) {
																			if (*resolucion.FechaInicio).Before(time.Now()) && (*resolucion.FechaFin).After(time.Now()) {
																				docente.Activo = true
																			} else {
																				docente.Activo = false
																			}
																		} else {
																			docente.Activo = false
																		}
																	} else {
																		docente.Activo = false
																	}
																}
															}
														}
													}

												}
											}
										}
									}
								} else {
									return docente, map[string]interface{}{
										"Succes":  false,
										"Status":  404,
										"Message": "El docente no es de vinculación especial y no tiene un contrato de prestación de servicios",
										"Funcion": "InformacionDve",
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return docente, nil
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
	docente, errorDocente := InformacionDve(numeroDocumento[0])
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

//helper que dado un codigo de facultad retorna el nombre de la facultad o del proyecto curricular
func NombreFacultadProyecto(codigoFacultadProyecto int) (nombreFacultadProyecto string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener el nombre de la facultad",
				"Funcion": "NombreFacultad",
			}
			panic(outputError)
		}
	}()

	var respuesta_peticion []map[string]interface{}
	var dependencia models.Dependencia

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudOikos")+"/dependencia/?query=Id:"+strconv.Itoa(codigoFacultadProyecto), &respuesta_peticion); err == nil && response == 200 {
		if len(respuesta_peticion) > 0 {
			dependenciaJson, err := json.Marshal(respuesta_peticion[0])
			if err != nil {
				outputError = map[string]interface{}{
					"Succes":  false,
					"Status":  404,
					"Message": "Error al obtener el nombre de la facultad o el proyecto curricular",
					"Funcion": "NombreFacultad",
				}
				return nombreFacultadProyecto, outputError
			} else {
				if err = json.Unmarshal(dependenciaJson, &dependencia); err != nil {
					outputError = map[string]interface{}{
						"Succes":  false,
						"Status":  404,
						"Message": "Error al obtener el nombre de la facultad o el proyecto curricular",
						"Funcion": "NombreFacultad",
					}
					return nombreFacultadProyecto, outputError
				} else {
					nombreFacultadProyecto = dependencia.Nombre
					return nombreFacultadProyecto, nil
				}
			}
		} else {
			return nombreFacultadProyecto, map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "No se encontro la facultad o el proyecto curricular",
				"Funcion": "NombreFacultad",
			}
		}
	}
	return
}
