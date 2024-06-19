package models

type InformacionCertificacionDve struct {
	InformacionDve    InformacionDVE         `json:"informacion_dve"`
	IntensidadHoraria []IntensidadHorariaDVE `json:"intensidad_horaria"`
	JefeTalentoHumano JefeTalentoHumano      `json:"informacion_certificacion_dve"`
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
}

type IntensidadHorariaDVE struct {
	Ano              string `json:"ano"`
	Periodo          string `json:"periodo"`
	NombreAsignatura string `json:"nombre_asignatura"`
	HorasSemana      string `json:"horas_semanales"`
	NumeroSemanas    string `json:"numero_semanas"`
	HorasSemestrales string `json:"horas_semestrales"`
	SalarioDocente   string `json:"salario_docente"`
}
type JefeTalentoHumano struct {
	Nombre string `json:"Nombre"`
	Cargo  string `json:"Cargo"`
}
