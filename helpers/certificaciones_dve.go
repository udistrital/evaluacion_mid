package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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
						if respuesta_contrato != nil {
							contratoJson, err := json.Marshal(respuesta_contrato[0])
							if err != nil {
								return docente, outputError
							} else {
								if err = json.Unmarshal(contratoJson, &informacion_contrato); err == nil && informacion_contrato.TipoContrato.Id == 18 || informacion_contrato.TipoContrato.Id == 3 {
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

													pago, err := ObtenerUltimoSalarioDve(numeroDocumento)

													if err != nil {
														docente.UltimoPago = "Este docente no cuenta con un contrato por lo tanto no es posible obtener el ultimo salario"
													} else {
														docente.UltimoPago = pago
													}

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

	fmt.Println("Periodo inicio: ", periodoInicial)
	fmt.Println("Periodo fin: ", periodoFinal)

	// Construir el json para enviar los datos de la autenticación
	var bodyAutenticacion models.BodyAutenticacion
	bodyAutenticacion.Username = beego.AppConfig.String("UsuarioAutenticacion")
	bodyAutenticacion.Password = beego.AppConfig.String("ContrasenaAutenticacion")
	bodyAutenticacion.Version = beego.AppConfig.String("VersionAutenticacion")

	var respuesta_peticion map[string]interface{}
	var response_autenticacion models.AutenticacionResponse

	// Solicitud de autenticación
	if err := sendJsonAutenticacion(beego.AppConfig.String("UrlAutenticacion"), "POST", &respuesta_peticion, bodyAutenticacion); err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al autenticar el usuario",
			"Funcion": "IntensidadHorariaDve",
		}
		return intensidadHoraria, outputError
	}

	// Convertir la respuesta de la autenticación al modelo de autenticación
	responseJson, err := json.Marshal(respuesta_peticion)
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener la respuesta de la autenticación",
			"Funcion": "IntensidadHorariaDve",
		}
		return intensidadHoraria, outputError
	}

	json.Unmarshal(responseJson, &response_autenticacion)

	// Creacion del json para obtener la carga academica
	var bodyCargaAcademica models.BodyCargaAcademica
	var cargas_academicas []models.CargaAcademica
	//var respuesta_carga_academica map[string]interface{}

	bodyCargaAcademica.Parametros.Identificacion = numeroDocumento[0]

	if err := sendJsonWithToken(beego.AppConfig.String("UrlCargaAcademica"), "POST", &cargas_academicas, bodyCargaAcademica, response_autenticacion.Token); err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener la carga academica",
			"Funcion": "IntensidadHorariaDve",
		}
		return intensidadHoraria, outputError
	}

	cargasFiltradas := []models.CargaAcademica{}
	for _, carga := range cargas_academicas {
		incluir := true
		// Filtrado por periodoInicial y periodoFinal si están presentes
		if len(periodoInicial) > 0 {
			anioInicial, periodoInicio, _ := ObtenerPeriodo(periodoInicial[0])
			if carga.Anio < anioInicial || (carga.Anio == anioInicial && carga.Periodo < periodoInicio) {
				incluir = false
			}
		}
		if len(periodoFinal) > 0 {
			anioFinal, periodoFin, _ := ObtenerPeriodo(periodoFinal[0])
			if carga.Anio > anioFinal || (carga.Anio == anioFinal && carga.Periodo > periodoFin) {
				incluir = false
			}
		}
		if incluir {
			cargasFiltradas = append(cargasFiltradas, carga)
		}
	}

	for _, carga := range cargasFiltradas {

		// Filtrar por vinculaciones si están presentes
		if len(vinculaciones) > 0 {
			encontrado := false
			for _, vinculacion := range vinculaciones {
				if carga.TipoVinculacion == vinculacion {
					encontrado = true
					break
				}
			}
			if !encontrado {
				continue
			}
		}

		// Crear el objeto de intensidad horaria
		intensidad := models.IntensidadHorariaDVE{
			Ano:              strconv.Itoa(carga.Anio),
			Periodo:          fmt.Sprintf("%d", carga.Periodo),
			NombreAsignatura: carga.Espacio,
			HorasSemana:      fmt.Sprintf("%d", carga.Hora),
			NumeroSemanas:    "16",
			HorasSemestrales: fmt.Sprintf("%d", carga.Hora*16),
		}

		intensidadHoraria = append(intensidadHoraria, intensidad)
	}

	return intensidadHoraria, nil

}

// helper que retorna la informacion de un docente y el listado de intensidades horarias en un solo objeto
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
	fmt.Println("Intensidad horaria: ", intensidadHoraria)
	if errorIntensidadHoraria != nil {
		return certificacion, errorIntensidadHoraria
	}

	certificacion = models.InformacionCertificacionDve{
		InformacionDve:    docente,
		IntensidadHoraria: intensidadHoraria,
	}

	fmt.Println("Certificacion: ", certificacion)
	infoTalentoHumano, errInfoTelento := InformacionJefeTalentoHumano()
	if errInfoTelento == nil {
		certificacion.JefeTalentoHumano = *infoTalentoHumano
	}

	return certificacion, nil
}

// helper que dado un codigo de facultad retorna el nombre de la facultad o del proyecto curricular
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
func InformacionJefeTalentoHumano() (JefeTalentoHumano *models.JefeTalentoHumano, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener  Informacion del Jefe de TalentoHumano",
				"Funcion": "InformacionJefeTalentoHumano",
			}
			panic(outputError)
		}
	}()

	var respuesta []models.SupervisorContrato

	var query = "/supervisor_contrato?query=DependenciaSupervisor:DEP633&sortby=FechaInicio&order=desc"
	fmt.Println("Url talento humano: ", beego.AppConfig.String("UrlcrudAgoraProduccion")+query)
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgoraProduccion")+query, &respuesta); (err == nil) && (response == 200) {
		if respuesta != nil {

			if respuesta[0].FechaFin.After(time.Now()) {

				JefeTalentoHumano = &models.JefeTalentoHumano{
					Nombre: respuesta[0].Nombre,
					Cargo:  respuesta[0].Cargo,
				}
			} else {
				outputError = map[string]interface{}{
					"Succes":  true,
					"Status":  200,
					"Message": "No hay JefeTalentoHumano activo para la dependencia",
					"Funcion": "InformacionJefeTalentoHumano",
				}
				panic(outputError)
			}

		} else {

			return nil, outputError
		}
	}

	return JefeTalentoHumano, nil
}

// funcion para que dado una fecha del estilo "2022-09-01T12:00:00Z" retorne el año y el mes, es decir 2022 y 9
func ObtenerAnoMes(fecha time.Time) (ano string, mes string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener el año y el mes",
				"Funcion": "ObtenerAnoMes",
			}
			panic(outputError)
		}
	}()

	ano = strconv.Itoa(fecha.Year())
	mes = strconv.Itoa(int(fecha.Month()))

	return ano, mes, nil
}

func ObtenerUltimoSalarioDve(numeroDocumentoDve string) (ultimoSalario string, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener el ultimo salario del docente, el docente no tiene salarios registrados",
				"Funcion": "ObtenerUltimoSalaraDve",
			}
		}
	}()
	var respuesta_peticion map[string]interface{}
	var detalle_preliquidacion []models.DetallePreliquidacion
	var fecha time.Time
	var nomina string

	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudTitan")+"/detalle_preliquidacion/?query=ContratoPreliquidacionId.ContratoId.Documento:"+numeroDocumentoDve, &respuesta_peticion); err == nil && response == 200 {
		if len(respuesta_peticion) > 0 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &detalle_preliquidacion)
			for _, detalle := range detalle_preliquidacion {
				if fecha.IsZero() {
					fecha = detalle.ContratoPreliquidacionId.ContratoId.FechaInicio
					nomina = strconv.Itoa(detalle.ContratoPreliquidacionId.PreliquidacionId.NominaId)
				} else {
					if fecha.Before(detalle.ContratoPreliquidacionId.ContratoId.FechaInicio) {
						fecha = detalle.ContratoPreliquidacionId.ContratoId.FechaInicio
						nomina = strconv.Itoa(detalle.ContratoPreliquidacionId.PreliquidacionId.NominaId)
					} else {
						fecha = fecha
						nomina = nomina
					}
				}
			}
			ano, mes, errorAnoMes := ObtenerAnoMes(fecha)
			if errorAnoMes != nil {
				return ultimoSalario, errorAnoMes
			} else {
				var respuesta_detalle map[string]interface{}
				var detalles []models.Detalle
				if response, err := getJsonTest(beego.AppConfig.String("UrlMidTitan")+"/detalle_preliquidacion/obtener_detalle_DVE/"+ano+"/"+mes+"/"+numeroDocumentoDve+"/"+nomina, &respuesta_detalle); err == nil && response == 200 {
					if len(respuesta_detalle) > 0 {
						LimpiezaRespuestaRefactor(respuesta_detalle, &detalles)
						for _, detalle := range detalles {
							for _, det := range detalle.Detalle {
								if det.ConceptoNominaId.Id == 152 {
									ultimoSalario = FormatMoney(int(det.ValorCalculado), 0)
									return ultimoSalario, nil
								} else {
									continue
								}
							}
						}
					}
				}
			}
		} else {
			return "ultimoSalario", map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "No se encontro el salario del docente",
				"Funcion": "ObtenerUltimoSalaraDve",
			}
		}
	}
	return ultimoSalario, nil
}

func ObtenerPeriodo(periodo string) (int, int, error) {
	partes := strings.Split(periodo, "-")
	if len(partes) != 2 {
		return 0, 0, errors.New("Formato de periodo inválido")
	}

	// Convertir el año a entero
	anio, err := strconv.Atoi(partes[0])
	if err != nil {
		return 0, 0, errors.New("Error al convertir el año a número")
	}

	// Determinar el periodo (I = 1, II = 2, III = 3, etc.)
	var periodoNum int
	switch partes[1] {
	case "I":
		periodoNum = 1
	case "II":
		periodoNum = 2
	case "III":
		periodoNum = 3
	default:
		return 0, 0, errors.New("Periodo no reconocido")
	}

	return anio, periodoNum, nil
}

// Calcula las horas entre dos tiempos en formato "5PM-6PM".
func calcularHoras(horaLarga string) (int, error) {
	partes := strings.Split(horaLarga, "-")
	if len(partes) != 2 {
		return 0, fmt.Errorf("Formato de HORA_LARGA inválido")
	}

	inicio, err := parseHora(partes[0])
	if err != nil {
		return 0, err
	}

	fin, err := parseHora(partes[1])
	if err != nil {
		return 0, err
	}

	if fin < inicio {
		fin += 24
	}

	return fin - inicio, nil
}

// Convierte una hora en formato "5PM" o "10AM" a un valor en horas (24 horas).
func parseHora(hora string) (int, error) {
	hora = strings.TrimSpace(strings.ToUpper(hora))
	if strings.HasSuffix(hora, "PM") || strings.HasSuffix(hora, "AM") {
		t := hora[:len(hora)-2]
		h, err := strconv.Atoi(t)
		if err != nil {
			return 0, fmt.Errorf("Error al parsear hora: %v", err)
		}
		if strings.HasSuffix(hora, "PM") && h != 12 {
			h += 12
		} else if strings.HasSuffix(hora, "AM") && h == 12 {
			h = 0
		}
		return h, nil
	}
	return 0, fmt.Errorf("Formato de hora inválido: %s", hora)
}
