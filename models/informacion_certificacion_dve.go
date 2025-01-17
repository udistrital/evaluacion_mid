package models

import "time"

type InformacionCertificacionDve struct {
	InformacionDve    InformacionDVE        `json:"informacion_dve"`
	IntensidadHoraria []IntensidadesDocente `json:"intensidad_horaria"`
	JefeTalentoHumano JefeTalentoHumano     `json:"informacion_certificacion_dve"`
}

type InformacionDVE struct {
	Activo             bool   `json:"activo"`
	NombreDocente      string `json:"nombre_docente"`
	NumeroDocumento    string `json:"numero_documento"`
	NivelAcademico     string `json:"nivel_academico"`
	Facultad           string `json:"facultad"`
	ProyectoCurricular string `json:"proyecto_curricular"`
	Categoria          string `json:"categoria"`
	Dedicacion         string `json:"dedicacion"`
	UltimoPago         string `json:"ultimo_pago_dve"`
}

type IntensidadHorariaDVE struct {
	Anio               int       `json:"Año"`
	Periodo            int       `json:"Periodo"`
	ProyectoCurricular string    `json:"ProyectoCurricular"`
	Asignaturas        string    `json:"Asignaturas"`
	FechaInicio        time.Time `json:"FechaInicio"`
	FechaFin           time.Time `json:"FechaFin"`
	HorasSemana        int       `json:"HorasSemanales"`
	NumeroSemanas      int       `json:"NumeroSemanas"`
	HorasSemestre      int       `json:"HorasSemestre"`
	TipoVinculacion    string    `json:"TipoVinculacion"`
	Categoria          string    `json:"Categoria"`
	Valor              string    `json:"Valor"`
	NivelAcademico     string    `json:"NivelAcademico"`
	Facultad           string    `json:"Facultad"`
	Resolucion         string    `json:"Resolucion"`
}

type IntensidadesDocente struct {
	Anio            int                 `json:"Año"`
	Periodo         int                 `json:"Periodo"`
	TipoVinculacion string              `json:"TipoVinculacion"`
	NumeroSemanas   int                 `json:"NumeroSemanas"`
	HorasSemanales  int                 `json:"HorasSemanales"`
	Resoluciones    []ResolucionDocente `json:"Resoluciones"`
}

type ResolucionDocente struct {
	ProyectoCurricular string    `json:"ProyectoCurricular"`
	Asignaturas        string    `json:"Asignaturas"`
	FechaInicio        time.Time `json:"FechaInicio"`
	FechaFin           time.Time `json:"FechaFin"`
	HorasSemanales     int       `json:"HorasSemanales"`
	NumeroSemanas      int       `json:"NumeroSemanas"`
	HorasSemestre      int       `json:"HorasSemestre"`
	TipoVinculacion    string    `json:"TipoVinculacion"`
	Categoria          string    `json:"Categoria"`
	Valor              string    `json:"Valor"`
	NivelAcademico     string    `json:"NivelAcademico"`
	Facultad           string    `json:"Facultad"`
	Resolucion         string    `json:"Resolucion"`
}

type JefeTalentoHumano struct {
	Nombre string `json:"Nombre"`
	Cargo  string `json:"Cargo"`
}
