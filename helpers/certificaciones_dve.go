package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
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

	// Recuperar la información del proveedor

	var respuesta_proveedor []map[string]interface{}
	var informacion_proveedor models.InformacionProveedor

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numeroDocumento, &respuesta_proveedor); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener la informacion del proveedor",
			"Funcion": "InformacionDve",
		}
		return docente, outputError
	}

	if len(respuesta_proveedor) == 0 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "No se encontro ningun proveedor registrado con el numero de documento " + numeroDocumento,
			"Funcion": "InformacionDve",
		}
		return docente, outputError
	}

	proveedorJson, err := json.Marshal(respuesta_proveedor[0])
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al convertir la informacion del proveedor",
			"Funcion": "InformacionDve",
		}
		return docente, outputError
	}

	json.Unmarshal(proveedorJson, &informacion_proveedor)

	// Traer la informacion de la vinculacion del docente
	//var respuesta_vinculacion map[string]interface{}
	// var vinculacionDocente []models.VinculacionesDocenteResolucion
	//var ultimaVinculacion models.VinculacionesDocenteResolucion

	//fmt.Println("Url vinculacion: ", beego.AppConfig.String("UrlCrudResoluciones")+"/vinculacion_docente/?query=PersonaId:"+numeroDocumento+"&sortby=FechaCreacion&order=desc")
	// if response, err := getJsonTest(beego.AppConfig.String("UrlCrudResoluciones")+"/vinculacion_docente/?query=PersonaId:"+numeroDocumento+"&sortby=FechaCreacion&order=desc", &respuesta_vinculacion); err != nil && response != 200 {
	// 	outputError = map[string]interface{}{
	// 		"Succes":  false,
	// 		"Status":  404,
	// 		"Message": "Error al obtener la informacion de la vinculacion del docente",
	// 		"Funcion": "InformacionDve",
	// 	}
	// 	return docente, outputError
	// }

	// if len(respuesta_vinculacion) == 0 {
	// 	outputError = map[string]interface{}{
	// 		"Succes":  false,
	// 		"Status":  404,
	// 		"Message": "No se encontro la informacion de la vinculacion del docente",
	// 		"Funcion": "InformacionDve",
	// 	}
	// 	return docente, outputError
	// }

	// LimpiezaRespuestaRefactor(respuesta_vinculacion, &vinculacionDocente)

	//var fechaMayor time.Time
	// for _, vinculacion := range vinculacionDocente {
	// 	if vinculacion.NumeroContrato == "" {
	// 		continue
	// 	}
	// 	if fechaMayor.Before(vinculacion.FechaInicio.AddDate(0, 0, 7*vinculacion.NumeroSemanas)) {
	// 		fechaMayor = vinculacion.FechaInicio.AddDate(0, 0, 7*vinculacion.NumeroSemanas)
	// 		ultimaVinculacion = vinculacion
	// 	}
	// }

	//var respuesta_resolucion map[string]interface{}
	//var resolucion models.Resolucion

	//fmt.Println("Url resolucion: ", beego.AppConfig.String("UrlCrudResoluciones")+"/resolucion/?query=Id:"+strconv.Itoa(vinculacionDocente[0].ResolucionVinculacionDocente.Id))
	// if response, err := getJsonTest(beego.AppConfig.String("UrlCrudResoluciones")+"/resolucion/?query=Id:"+strconv.Itoa(vinculacionDocente[0].ResolucionVinculacionDocente.Id), &respuesta_resolucion); err != nil && response != 200 {
	// 	outputError = map[string]interface{}{
	// 		"Succes":  false,
	// 		"Status":  404,
	// 		"Message": "Error al obtener la informacion de la resolucion",
	// 		"Funcion": "InformacionDve",
	// 	}
	// 	return docente, outputError
	// }

	// if len(respuesta_resolucion) == 0 {
	// 	outputError = map[string]interface{}{
	// 		"Succes":  false,
	// 		"Status":  404,
	// 		"Message": "No se encontro la informacion de la resolucion",
	// 		"Funcion": "InformacionDve",
	// 	}
	// 	return docente, outputError
	// }

	// resolucionJson, err := json.Marshal(respuesta_resolucion["Data"].([]interface{})[0])
	// if err != nil {
	// 	outputError = map[string]interface{}{
	// 		"Succes":  false,
	// 		"Status":  404,
	// 		"Message": "Error al convertir la informacion de la resolucion",
	// 		"Funcion": "InformacionDve",
	// 	}
	// 	return docente, outputError
	// }

	// json.Unmarshal(resolucionJson, &resolucion)

	// if (resolucion.FechaInicio.IsZero() || resolucion.FechaInicio.Year() == 1) || (resolucion.FechaFin.IsZero() || resolucion.FechaFin.Year() == 1) {
	// 	docente.Activo = false

	// }

	// if !resolucion.FechaInicio.IsZero() && !resolucion.FechaFin.IsZero() {
	// 	if (resolucion.FechaInicio).Before(resolucion.FechaFin) {
	// 		if (resolucion.FechaInicio).Before(time.Now()) && (resolucion.FechaFin).After(time.Now()) {
	// 			docente.Activo = true
	// 		} else {
	// 			docente.Activo = false
	// 		}
	// 	} else {
	// 		docente.Activo = false
	// 	}
	// } else {
	// 	docente.Activo = false
	// }

	// if incluirSalario {
	// 	pago, err := ObtenerUltimoSalarioDve(numeroDocumento)

	// 	if err != nil {
	// 		docente.UltimoPago = "No se encontro un salario para el docente"
	// 	} else {
	// 		docente.UltimoPago = pago
	// 	}
	// }

	docente.NombreDocente = informacion_proveedor.NomProveedor
	docente.NumeroDocumento = informacion_proveedor.NumDocumento
	//docente.Facultad, _ = NombreFacultadProyecto(ultimaVinculacion.ResolucionVinculacionDocente.FacultadId)
	//docente.ProyectoCurricular, _ = NombreFacultadProyecto(ultimaVinculacion.ProyectoCurricularId)
	//docente.NivelAcademico = ultimaVinculacion.ResolucionVinculacionDocente.NivelAcademico
	//docente.Categoria = ultimaVinculacion.Categoria
	//docente.Dedicacion = ultimaVinculacion.ResolucionVinculacionDocente.Dedicacion

	return docente, nil

}

func VinculacionesDocenteDve(docDocente string) (vinculacionesDocente []models.VinculacionesDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener las vinculaciones del docente",
				"Funcion": "vinculacionesDocenteDve",
			}
		}
	}()

	var vinculaciones_old bool
	var vinculaciones_new bool
	var respuesta_proveedor []map[string]interface{}
	var informacion_proveedor models.InformacionProveedor

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+docDocente, &respuesta_proveedor); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener la informacion del proveedor",
			"Funcion": "vinculacionesDocenteDve",
		}
		return vinculacionesDocente, outputError
	}

	if len(respuesta_proveedor) > 0 {
		proveedorJson, err := json.Marshal(respuesta_proveedor[0])
		if err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la informacion del proveedor",
				"Funcion": "vinculacionesDocenteDve",
			}
			return vinculacionesDocente, outputError
		} else {
			if err = json.Unmarshal(proveedorJson, &informacion_proveedor); err != nil {
				outputError = map[string]interface{}{
					"Succes":  false,
					"Status":  404,
					"Message": "Error al obtener la informacion del proveedor",
					"Funcion": "vinculacionesDocenteDve",
				}
				return vinculacionesDocente, outputError
			}
		}
	}

	var respuesta_contrato []map[string]interface{}
	var contratos []models.ContratoGeneral
	//fmt.Println("Url contrato: ", beego.AppConfig.String("UrlcrudAgora")+"/contrato_general/?query=TipoContrato.Id__in:2|3|18,Contratista:"+strconv.Itoa(informacion_proveedor.Id)+"&limit=-1")
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_general/?query=TipoContrato.Id__in:2|3|18,Contratista:"+strconv.Itoa(informacion_proveedor.Id)+"&limit=-1", &respuesta_contrato); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener los contratos del docente",
			"Funcion": "vinculacionesDocenteDve",
		}
	}

	if len(respuesta_contrato) == 0 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "El docente no tiene contratos de vinculacion especial",
			"Funcion": "vinculacionesDocenteDve",
		}
		return vinculacionesDocente, outputError
	}

	contratoJson, err := json.Marshal(respuesta_contrato)
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener los contratos del docente",
			"Funcion": "vinculacionesDocenteDve",
		}
		return vinculacionesDocente, outputError
	}

	json.Unmarshal(contratoJson, &contratos)

	for _, contrato := range contratos {
		if contrato.VigenciaContrato < 2023 {
			vinculaciones_old = true
			break
		}
	}

	for _, contrato := range contratos {
		if contrato.VigenciaContrato >= 2023 {
			vinculaciones_new = true
			break
		}
	}

	if vinculaciones_new {
		vinculacionesNew, error := ObtenerVinculacionesNew(docDocente)
		if error != nil {
			return vinculacionesDocente, error
		}

		if len(vinculacionesNew) > 0 {
			vinculacionesDocente = append(vinculacionesDocente, vinculacionesNew...)
		}
	}

	// aqui se debe implementar el servicio para obtener las vinculaciones anteriores al 2023
	if vinculaciones_old {
		vinculacionesOld, error := ObtenerVinculacionesOld(docDocente)
		if error != nil {
			return vinculacionesDocente, error
		}

		if len(vinculacionesOld) > 0 {
			vinculacionesDocente = append(vinculacionesDocente, vinculacionesOld...)
		}
	}

	OrdenarVinculaciones(vinculacionesDocente)

	return vinculacionesDocente, nil

}

func ObtenerVinculacionesNew(docDocente string) (vinculacionesDocente []models.VinculacionesDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener las vinculaciones a partir del 2023 del docente",
				"Funcion": "obtenerVinculacionesNew",
			}
		}
	}()

	var respuesta_vinculacion map[string]interface{}
	var vinculaciones_docente []models.VinculacionesDocenteResolucion

	//fmt.Println("Url vinculacion: ", beego.AppConfig.String("UrlCrudResoluciones")+"/vinculacion_docente/?query=PersonaId:"+docDocente+"&limit=-1&sortby=VigenciaContrato&order=desc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudResoluciones")+"/vinculacion_docente/?query=PersonaId:"+docDocente+"&limit=-1&sortby=Vigencia&order=desc", &respuesta_vinculacion); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener las vinculaciones del docente",
			"Funcion": "obtenerVinculacionesNew",
		}
		return vinculacionesDocente, outputError
	}

	if len(respuesta_vinculacion) == 0 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "El docente no tiene contratos de vinculacion especial",
			"Funcion": "obtenerVinculacionesNew",
		}
		return vinculacionesDocente, outputError
	}

	LimpiezaRespuestaRefactor(respuesta_vinculacion, &vinculaciones_docente)

	for _, vinculacion := range vinculaciones_docente {

		if vinculacion.NumeroContrato == "" {
			continue
		}
		var respuesta_dependencia map[string]interface{}
		var proyecto_curricular models.ProyectoCurricular
		var parametros models.ParametrosData
		var respuesta_resolucion map[string]interface{}
		var resolucion models.Resolucion

		// Peticion para traer el proyecto curricular
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.ProyectoCurricularId), &respuesta_dependencia); err != nil && response != 200 {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la dependencia",
				"Funcion": "obtenerVinculacionesNew",
			}
			return vinculacionesDocente, outputError
		}

		jsonProyectoCurricular, err := json.Marshal(respuesta_dependencia)
		if err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la dependencia",
				"Funcion": "obtenerVinculacionesNew",
			}
			return vinculacionesDocente, outputError
		}
		json.Unmarshal(jsonProyectoCurricular, &proyecto_curricular)

		// Peticion para traer la dedicacion
		//fmt.Println("Url parametros: ", beego.AppConfig.String("UrlParametrosCrud")+"/parametro/?query=Id:"+strconv.Itoa(vinculacion.DedicacionId))
		if response, err := getJsonTest(beego.AppConfig.String("UrlParametrosCrud")+"/parametro/?query=Id:"+strconv.Itoa(vinculacion.DedicacionId), &parametros); err != nil && response != 200 {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la dedicacion",
				"Funcion": "obtenerVinculacionesNew",
			}
			return vinculacionesDocente, outputError
		}

		// Peticion para traer el periodo academico

		//fmt.Println("Url resolucion: ", beego.AppConfig.String("UrlCrudResoluciones")+"/resolucion/"+strconv.Itoa(vinculacion.ResolucionVinculacionDocente.Id))
		if response, err := getJsonTest(beego.AppConfig.String("UrlCrudResoluciones")+"/resolucion/"+strconv.Itoa(vinculacion.ResolucionVinculacionDocente.Id), &respuesta_resolucion); err != nil && response != 200 {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la resolucion",
				"Funcion": "obtenerVinculacionesNew",
			}
			return vinculacionesDocente, outputError
		}

		LimpiezaRespuestaRefactor(respuesta_resolucion, &resolucion)

		vinculacionDocente := models.VinculacionesDocente{
			NumeroContrato: vinculacion.NumeroContrato,
			Vigencia:       vinculacion.Vigencia,
			FechaInicio:    vinculacion.FechaInicio,
			//FechaFin:               vinculacion.FechaInicio.AddDate(0, 0, 7*vinculacion.NumeroSemanas),
			NumeroHorasSemanales:   vinculacion.NumeroHorasSemanales,
			NumeroSemanas:          vinculacion.NumeroSemanas,
			NumeroHorasSemestrales: vinculacion.NumeroHorasSemanales * vinculacion.NumeroSemanas,
			ProyectoCurricular:     proyecto_curricular.Nombre,
			DependenciaAcademica:   vinculacion.DependenciaAcademica,
		}

		if resolucion.FechaFin.IsZero() || resolucion.FechaFin.Year() == 1 {
			vinculacionDocente.FechaFin = vinculacion.FechaInicio.AddDate(0, 0, 7*vinculacion.NumeroSemanas)
		} else {
			vinculacionDocente.FechaFin = resolucion.FechaFin
		}

		if (resolucion != models.Resolucion{}) {
			vinculacionDocente.Periodo = resolucion.Periodo
		}

		if len(parametros.Data) == 0 {
			vinculacionDocente.Dedicacion = "No se encontro la dedicacion del docente"
		} else {
			vinculacionDocente.Dedicacion = parametros.Data[0].Nombre
		}

		vinculacionesDocente = append(vinculacionesDocente, vinculacionDocente)
	}

	return vinculacionesDocente, nil

}

func ObtenerVinculacionesOld(docDocente string) (vinculacionesDocente []models.VinculacionesDocente, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener las vinculaciones antes del 2023 del docente",
				"Funcion": "obtenerVinculacionesNew",
			}
		}
	}()

	var vinculacionesOld models.VinculacionesDocenteOld
	var respuesta_vinculacion_old map[string]interface{}
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJbpm")+"/vinculacion_docente/"+docDocente, &respuesta_vinculacion_old); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener las vinculaciones anteriores al 2023 del docente",
			"Funcion": "vinculacionesDocenteDve",
		}
		return vinculacionesDocente, outputError
	}

	vinculacionesOldJson, err := json.Marshal(respuesta_vinculacion_old)
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al deserializar las vinculaciones anteriores al 2023 del docente",
			"Funcion": "vinculacionesDocenteDve",
		}
		return vinculacionesDocente, outputError
	}

	json.Unmarshal(vinculacionesOldJson, &vinculacionesOld)

	for _, vinculacion := range vinculacionesOld.Docente.VinculacionesOld {
		var proyecto_curricular models.ProyectoCurricular
		var respuesta_dependencia map[string]interface{}
		// Peticion para traer el proyecto curricular
		//fmt.Println("Url dependencia: ", beego.AppConfig.String("UrlcrudOikos")+"/dependencia/"+vinculacion.ProyectoCurricularId)
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudOikos")+"/dependencia/"+vinculacion.ProyectoCurricularId, &respuesta_dependencia); err != nil && response != 200 {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la dependencia",
				"Funcion": "obtenerVinculacionesNew",
			}
			return vinculacionesDocente, outputError
		}

		jsonProyectoCurricular, err := json.Marshal(respuesta_dependencia)
		if err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  404,
				"Message": "Error al obtener la dependencia",
				"Funcion": "obtenerVinculacionesNew",
			}
			return vinculacionesDocente, outputError
		}
		json.Unmarshal(jsonProyectoCurricular, &proyecto_curricular)

		vigencia, _ := strconv.Atoi(vinculacion.Vigencia)
		periodo, _ := strconv.Atoi(vinculacion.Periodo)
		numero_horas_semanales, _ := strconv.Atoi(vinculacion.NumeroHorasSemanales)
		numero_semanas, _ := strconv.Atoi(vinculacion.NumeroSemanas)
		dependencia_academica, _ := strconv.Atoi(vinculacion.ProyectoCurricularId)
		vinculacionDocente := models.VinculacionesDocente{
			NumeroContrato:         vinculacion.NumeroContrato,
			Vigencia:               vigencia,
			Periodo:                periodo,
			FechaInicio:            vinculacion.FechaInicio,
			FechaFin:               vinculacion.FechaFin,
			NumeroHorasSemanales:   numero_horas_semanales,
			NumeroSemanas:          numero_semanas,
			NumeroHorasSemestrales: numero_horas_semanales * numero_semanas,
			Dedicacion:             vinculacion.Descripcion,
			ProyectoCurricular:     proyecto_curricular.Nombre,
			DependenciaAcademica:   dependencia_academica,
		}

		vinculacionesDocente = append(vinculacionesDocente, vinculacionDocente)
	}

	return vinculacionesDocente, nil
}

func IntensidadHorariaDve(numeroDocumento []string, periodoInicial []string, periodoFinal []string, vinculaciones []string) (intensidadHoraria []models.IntensidadHorariaDVE, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": fmt.Sprintf(`Error al obtener la intensidad horaria del docente %s`, numeroDocumento),
				"Funcion": "InformacionCertificacionDve",
			}
			panic(outputError)
		}
	}()

	//Obtener los datos de las vinculaciones del docente
	var vinculaciones_docente models.Docentes
	var respuesta_vinculacion map[string]interface{}
	var vinculacionesFiltradas []models.Certificado
	//fmt.Println("Url bodega:", beego.AppConfig.String("UrlBodegaProduccion")+"/certificado_dve/"+numeroDocumento[0])
	if err := getJsonWSO2(beego.AppConfig.String("UrlBodegaProduccion")+"/certificado_dve/"+numeroDocumento[0], &respuesta_vinculacion); err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": fmt.Sprintf(`Error al obtener las vinculaciones del docente %s`, numeroDocumento),
			"Funcion": "IntensidadHorariaDve",
		}
		return intensidadHoraria, outputError
	}

	vinculacionesJson, err := json.Marshal(respuesta_vinculacion)
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": fmt.Sprintf(`Error al obtener las vinculaciones del docente %s`, numeroDocumento),
			"Funcion": "IntensidadHorariaDve",
		}
		return intensidadHoraria, outputError
	}

	json.Unmarshal(vinculacionesJson, &vinculaciones_docente)

	for _, vinculacion := range vinculaciones_docente.Docente.Certificados {

		vigencia, _ := strconv.Atoi(strings.Split(vinculacion.Periodo, "-")[0])
		periodo, _ := strconv.Atoi(strings.Split(vinculacion.Periodo, "-")[1])
		incluir := true

		if len(periodoInicial) > 0 {
			anioInicial, periodoInicialNum, err := ObtenerPeriodo(periodoInicial[0])
			if err != nil {
				return intensidadHoraria, map[string]interface{}{
					"Succes":  false,
					"Status":  404,
					"Message": err,
					"Funcion": "IntensidadHorariaDve",
				}
			}

			if vigencia < anioInicial || (vigencia == anioInicial && periodo < periodoInicialNum) {
				incluir = false
			}
		}

		if len(periodoFinal) > 0 {
			anioFinal, periodoFinalNum, err := ObtenerPeriodo(periodoFinal[0])
			if err != nil {
				return intensidadHoraria, map[string]interface{}{
					"Succes":  false,
					"Status":  404,
					"Message": err,
					"Funcion": "IntensidadHorariaDve",
				}
			}

			if (incluir && vigencia > anioFinal) || (incluir && vigencia == anioFinal && periodo > periodoFinalNum) {
				incluir = false
			}
		}

		if len(vinculaciones) > 0 {

			encontrado := false

			for _, v := range vinculaciones {
				if incluir && v == vinculacion.Periodo {
					encontrado = true
					break
				}
			}

			if !encontrado {
				incluir = false
			}
		}

		if incluir {
			vinculacionesFiltradas = append(vinculacionesFiltradas, vinculacion)
		}
	}

	for _, vinculacion := range vinculacionesFiltradas {
		var intensidad models.IntensidadHorariaDVE
		intensidad.Anio, _ = strconv.Atoi(strings.Split(vinculacion.Periodo, "-")[0])
		intensidad.Periodo, _ = strconv.Atoi(strings.Split(vinculacion.Periodo, "-")[1])
		intensidad.ProyectoCurricular = vinculacion.Proyecto
		intensidad.Asignaturas = strings.TrimSuffix(vinculacion.Asignatura, ",")
		intensidad.FechaInicio = vinculacion.FechaInicio
		intensidad.FechaFin = vinculacion.FechaFin
		intensidad.HorasSemana, _ = strconv.Atoi(vinculacion.NumeroHorasSemanales)
		intensidad.NumeroSemanas, _ = strconv.Atoi(vinculacion.NumeroSemanas)
		intensidad.HorasSemestre = intensidad.HorasSemana * intensidad.NumeroSemanas
		intensidad.TipoVinculacion = vinculacion.Dedicacion
		intensidad.Categoria = vinculacion.Categoria
		intensidad.Valor = vinculacion.Valor
		intensidad.NivelAcademico = vinculacion.NivelAcademico
		intensidad.Facultad = vinculacion.Facultad
		intensidad.Resolucion = vinculacion.NumeroResolucion
		intensidadHoraria = append(intensidadHoraria, intensidad)
	}

	OrdenarIntensidadHoraria(intensidadHoraria)

	return intensidadHoraria, nil
}

// func IntensidadHorariaDve(numeroDocumento []string, periodoInicial []string, periodoFinal []string, vinculaciones []string) (intensidadHoraria []models.IntensidadHorariaDVE, outputError map[string]interface{}) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			outputError = map[string]interface{}{
// 				"Succes":  false,
// 				"Status":  502,
// 				"Message": "Error al obtener la intensidad horaria del docente",
// 				"Funcion": "IntensidadHorariaDve",
// 			}
// 			panic(outputError)
// 		}
// 	}()

// 	// Construir el json para enviar los datos de la autenticación
// 	var bodyAutenticacion models.BodyAutenticacion
// 	bodyAutenticacion.Username = beego.AppConfig.String("UsuarioAutenticacion")
// 	bodyAutenticacion.Password = beego.AppConfig.String("ContrasenaAutenticacion")
// 	bodyAutenticacion.Version = beego.AppConfig.String("VersionAutenticacion")

// 	var respuesta_peticion map[string]interface{}
// 	var response_autenticacion models.AutenticacionResponse

// 	// Solicitud de autenticación
// 	if err := sendJsonAutenticacion(beego.AppConfig.String("UrlAutenticacion"), "POST", &respuesta_peticion, bodyAutenticacion); err != nil {
// 		outputError = map[string]interface{}{
// 			"Succes":  false,
// 			"Status":  404,
// 			"Message": "Error al autenticar el usuario",
// 			"Funcion": "IntensidadHorariaDve",
// 		}
// 		return intensidadHoraria, outputError
// 	}

// 	// Convertir la respuesta de la autenticación al modelo de autenticación
// 	responseJson, err := json.Marshal(respuesta_peticion)
// 	if err != nil {
// 		outputError = map[string]interface{}{
// 			"Succes":  false,
// 			"Status":  404,
// 			"Message": "Error al obtener la respuesta de la autenticación",
// 			"Funcion": "IntensidadHorariaDve",
// 		}
// 		return intensidadHoraria, outputError
// 	}

// 	json.Unmarshal(responseJson, &response_autenticacion)

// 	// Creacion del json para obtener la carga academica
// 	var bodyCargaAcademica models.BodyCargaAcademica
// 	var cargas_academicas []models.CargaAcademica
// 	//var respuesta_carga_academica map[string]interface{}

// 	bodyCargaAcademica.Parametros.Identificacion = numeroDocumento[0]

// 	if err := sendJsonWithToken(beego.AppConfig.String("UrlCargaAcademica"), "POST", &cargas_academicas, bodyCargaAcademica, response_autenticacion.Token); err != nil {
// 		outputError = map[string]interface{}{
// 			"Succes":  true,
// 			"Status":  200,
// 			"Message": "No se encontraron cargas academicas para el docente",
// 			"Funcion": "IntensidadHorariaDve",
// 		}
// 	}

// 	materiasUnicas := make(map[string]models.MateriaUnica)
// 	listaMateriasUnicas := []models.MateriaUnica{}

// 	for _, carga := range cargas_academicas {

// 		clave := fmt.Sprintf("%d-%d-%d", carga.Anio, carga.Periodo, carga.CodProyecto)

// 		if _, existe := materiasUnicas[clave]; !existe {
// 			materiasUnicas[clave] = models.MateriaUnica{
// 				Anio:            carga.Anio,
// 				Periodo:         carga.Periodo,
// 				Proyecto:        carga.Proyecto,
// 				Docente:         carga.Docente,
// 				CodProyecto:     carga.CodProyectoEstudiante,
// 				TipoVinculacion: carga.TipoVinculacion,
// 				Espacio:         carga.Espacio,
// 				IDGrupo:         carga.IdGrupo,
// 			}
// 		}
// 	}

// 	for _materia := range materiasUnicas {
// 		listaMateriasUnicas = append(listaMateriasUnicas, materiasUnicas[_materia])
// 	}

// 	// Traer las vinculaciones del docente
// 	vinculacionesDocente, error := VinculacionesDocenteDve(numeroDocumento[0])
// 	if error != nil {
// 		return intensidadHoraria, error
// 	}

// 	vinculacionesFiltradas := []models.VinculacionesDocente{}

// 	// Filtrar las vinculaciones por periodo
// 	for _, vinculacion := range vinculacionesDocente {
// 		incluir := true

// 		if len(periodoInicial) > 0 {
// 			anioInicial, periodoInicialNum, err := ObtenerPeriodo(periodoInicial[0])
// 			if err != nil {
// 				return intensidadHoraria, map[string]interface{}{
// 					"Succes":  false,
// 					"Status":  404,
// 					"Message": err,
// 					"Funcion": "IntensidadHorariaDve",
// 				}
// 			}

// 			if vinculacion.Vigencia < anioInicial || (vinculacion.Vigencia == anioInicial && vinculacion.Periodo < periodoInicialNum) {
// 				incluir = false
// 			}
// 		}

// 		if len(periodoFinal) > 0 {
// 			anioFinal, periodoFinalNum, err := ObtenerPeriodo(periodoFinal[0])
// 			if err != nil {
// 				return intensidadHoraria, map[string]interface{}{
// 					"Succes":  false,
// 					"Status":  404,
// 					"Message": err,
// 					"Funcion": "IntensidadHorariaDve",
// 				}
// 			}

// 			if vinculacion.Vigencia > anioFinal || (vinculacion.Vigencia == anioFinal && vinculacion.Periodo > periodoFinalNum) {
// 				incluir = false
// 			}
// 		}

// 		if len(vinculaciones) > 0 {
// 			//clave := fmt.Sprintf("%d-%d", vinculacion.Vigencia, vinculacion.Periodo, vinculacion.DependenciaAcademica)
// 			clave := fmt.Sprintf("%d-%d", vinculacion.Vigencia, vinculacion.Periodo)
// 			encontrado := false

// 			for _, v := range vinculaciones {
// 				if v == clave {
// 					encontrado = true
// 					break
// 				}
// 			}

// 			if !encontrado {
// 				incluir = false
// 			}
// 		}

// 		if incluir {
// 			vinculacionesFiltradas = append(vinculacionesFiltradas, vinculacion)
// 		}
// 	}

// 	// Retornar las coincidencias de las materias unicas con las vinculaciones del docente
// 	for _, vinculacion := range vinculacionesFiltradas {
// 		var intensidad models.IntensidadHorariaDVE
// 		intensidad.Anio = vinculacion.Vigencia
// 		intensidad.Periodo = vinculacion.Periodo
// 		intensidad.ProyectoCurricular = vinculacion.ProyectoCurricular
// 		intensidad.HorasSemana = vinculacion.NumeroHorasSemanales
// 		intensidad.NumeroSemanas = vinculacion.NumeroSemanas
// 		intensidad.HorasSemestre = vinculacion.NumeroHorasSemestrales
// 		intensidad.TipoVinculacion = vinculacion.Dedicacion
// 		intensidad.FechaInicio = vinculacion.FechaInicio
// 		intensidad.FechaFin = vinculacion.FechaFin
// 		for _, materia := range listaMateriasUnicas {
// 			if vinculacion.Vigencia == materia.Anio && vinculacion.Periodo == materia.Periodo {
// 				intensidad.Asignaturas = append(intensidad.Asignaturas, materia.Espacio)
// 			}
// 		}
// 		intensidadHoraria = append(intensidadHoraria, intensidad)
// 	}

// 	for _, materia := range listaMateriasUnicas {
// 		fmt.Println("{")
// 		fmt.Println("Año: ", materia.Anio)
// 		fmt.Println("Periodo: ", materia.Periodo)
// 		fmt.Println("Proyecto: ", materia.Proyecto)
// 		fmt.Println("Docente: ", materia.Docente)
// 		fmt.Println("Tipo vinculacion: ", materia.TipoVinculacion)
// 		fmt.Println("Espacio: ", materia.Espacio)
// 		fmt.Println("ID Grupo: ", materia.IDGrupo)
// 		fmt.Println("}")
// 	}

// 	OrdenarIntensidadHoraria(intensidadHoraria)

// 	return intensidadHoraria, nil

// }

func OrdenarIntensidadHoraria(vinculaciones []models.IntensidadHorariaDVE) {
	sort.Slice(vinculaciones, func(i, j int) bool {

		if vinculaciones[i].Anio == vinculaciones[j].Anio && vinculaciones[i].Periodo == vinculaciones[j].Periodo {
			if vinculaciones[i].FechaInicio.Equal(vinculaciones[j].FechaInicio) {
				return vinculaciones[i].Valor < vinculaciones[j].Valor
			}
			return vinculaciones[i].FechaFin.After(vinculaciones[j].FechaFin)
		}
		// Comparar Vigencia (años)
		if vinculaciones[i].Anio != vinculaciones[j].Anio {
			return vinculaciones[i].Anio > vinculaciones[j].Anio
		}
		// Comparar Período si los años son iguales
		return vinculaciones[i].Periodo > vinculaciones[j].Periodo
	})
}

func OrdenarVinculaciones(vinculaciones []models.VinculacionesDocente) {
	sort.Slice(vinculaciones, func(i, j int) bool {
		// Comparar Vigencia (años)
		if vinculaciones[i].Vigencia != vinculaciones[j].Vigencia {
			return vinculaciones[i].Vigencia > vinculaciones[j].Vigencia
		}
		// Comparar Período si los años son iguales
		return vinculaciones[i].Periodo > vinculaciones[j].Periodo
	})
}

// helper que retorna la informacion de un docente y el listado de intensidades horarias en un solo objeto
func InformacionCertificacionDve(numeroDocumento []string, periodoInicial []string, periodoFinal []string, vinculaciones []string, incluirSalario bool) (certificacion models.InformacionCertificacionDve, outputError map[string]interface{}) {
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

	fechaActual := time.Now()
	intensidadHoraria, errorIntensidadHoraria := IntensidadHorariaDve(numeroDocumento, periodoInicial, periodoFinal, vinculaciones)
	if errorIntensidadHoraria != nil {
		return certificacion, errorIntensidadHoraria
	}

	docente, errorDocente := InformacionDve(numeroDocumento[0])
	if errorDocente != nil {
		return certificacion, errorDocente
	}

	if incluirSalario {
		docente.UltimoPago = intensidadHoraria[0].Valor
	}

	if !intensidadHoraria[0].FechaInicio.IsZero() && !fechaActual.Before(intensidadHoraria[0].FechaInicio) && !intensidadHoraria[0].FechaFin.IsZero() && !fechaActual.After(intensidadHoraria[0].FechaFin) {
		docente.Activo = true
	} else {
		docente.Activo = false
	}

	certificacion = models.InformacionCertificacionDve{
		InformacionDve:    docente,
		IntensidadHoraria: intensidadHoraria,
	}

	infoTalentoHumano, errInfoTalento := InformacionJefeTalentoHumano()
	if errInfoTalento == nil {
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

	var respuesta_preliquidacion map[string]interface{}
	var detalle_preliquidacion []models.DetallePreliquidacion
	var fecha time.Time
	var nomina string

	//fmt.Println("Url detalle preliquidacion: ", beego.AppConfig.String("UrlCrudTitan")+"/detalle_preliquidacion/?query=ContratoPreliquidacionId.ContratoId.Documento:"+numeroDocumentoDve+"&sortby=FechaCreacion&order=desc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudTitan")+"/detalle_preliquidacion/?query=ContratoPreliquidacionId.ContratoId.Documento:"+numeroDocumentoDve+"&sortby=FechaCreacion&order=desc", &respuesta_preliquidacion); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener el ultimo salario del docente",
			"Funcion": "ObtenerUltimoSalaraDve",
		}
		return ultimoSalario, outputError
	}

	if len(respuesta_preliquidacion) == 0 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "El docente no tiene salarios registrados",
			"Funcion": "ObtenerUltimoSalaraDve",
		}
		return ultimoSalario, outputError
	}

	LimpiezaRespuestaRefactor(respuesta_preliquidacion, &detalle_preliquidacion)

	for _, detalle := range detalle_preliquidacion {

		if fecha.IsZero() {
			fecha = detalle.ContratoPreliquidacionId.ContratoId.FechaInicio
			nomina = strconv.Itoa(detalle.ContratoPreliquidacionId.PreliquidacionId.NominaId)
		}

		if detalle.ContratoPreliquidacionId.ContratoId.FechaInicio.After(fecha) {
			fecha = detalle.ContratoPreliquidacionId.ContratoId.FechaInicio
			nomina = strconv.Itoa(detalle.ContratoPreliquidacionId.PreliquidacionId.NominaId)
		}
	}

	anio, mes, error := ObtenerAnoMes(fecha)
	if error != nil {
		return ultimoSalario, error
	}

	var respuesta_detalle map[string]interface{}
	var detalles []models.Detalle

	//fmt.Println("Url detalle preliquidacion: ", beego.AppConfig.String("UrlMidTitan")+"/detalle_preliquidacion/obtener_detalle_DVE/"+anio+"/"+mes+"/"+numeroDocumentoDve+"/"+nomina)
	if response, err := getJsonTest(beego.AppConfig.String("UrlMidTitan")+"/detalle_preliquidacion/obtener_detalle_DVE/"+anio+"/"+mes+"/"+numeroDocumentoDve+"/"+nomina, &respuesta_detalle); err != nil && response != 200 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "Error al obtener el ultimo salario del docente",
			"Funcion": "ObtenerUltimoSalaraDve",
		}
		return ultimoSalario, outputError
	}

	if len(respuesta_detalle) == 0 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "El docente no tiene salarios registrados",
			"Funcion": "ObtenerUltimoSalaraDve",
		}
		return ultimoSalario, outputError
	}

	LimpiezaRespuestaRefactor(respuesta_detalle, &detalles)

	// for _, detalle := range detalles {
	// 	for _, det := range detalle.Detalle {
	// 		if det.ConceptoNominaId.Id == 152 {
	// 			ultimoSalario = FormatMoney(int(det.ValorCalculado), 0)
	// 			return ultimoSalario, nil
	// 		} else {
	// 			continue
	// 		}
	// 	}
	// }

	if len(detalles) == 0 {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  404,
			"Message": "El docente no tiene salarios registrados",
			"Funcion": "ObtenerUltimoSalaraDve",
		}
		return ultimoSalario, outputError
	}

	ultimoSalario = FormatMoney(int(detalles[0].TotalDevengado), 0)

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
