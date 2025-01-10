package models

import "time"

type Docente struct {
	Certificados []Certificado `json:"certificados"`
}

type Docentes struct {
	Docente Docente `json:"docente"`
}

type Certificado struct {
	Vigencia             string    `json:"vigencia"`
	Dedicacion           string    `json:"dedicacion"`
	Asignatura           string    `json:"asignatura"`
	FechaInicio          time.Time `json:"fecha_inicio"`
	CodigoProyectoSnies  string    `json:"codigo_proyecto_snies"`
	Periodo              string    `json:"periodo"`
	CodigoDocente        string    `json:"codigo_docente"`
	CodigoFacultad       string    `json:"codigo_facultad"`
	NumeroSemanas        string    `json:"numero_semanas"`
	Categoria            string    `json:"categoria"`
	Valor                string    `json:"valor"`
	Docente              string    `json:"docente"`
	Proyecto             string    `json:"proyecto"`
	NumeroResolucion     string    `json:"numero_resolucion"`
	NumeroHorasSemanales string    `json:"numero_horas_semanales"`
	CodigoProyecto       string    `json:"codigo_proyecto"`
	FechaFin             time.Time `json:"fecha_fin"`
	NivelAcademico       string    `json:"nivel_academico"`
	Facultad             string    `json:"facultad"`
}
