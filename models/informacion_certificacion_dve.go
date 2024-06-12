package models

type InformacionCertificacionDve struct {
	InformacionDve    InformacionDVE         `json:"informacion_dve"`
	IntensidadHoraria []IntensidadHorariaDVE `json:"intensidad_horaria"`
}

type InformacionDVE struct {
	Activo             string `json:"activo"`
	NombreDocente      string `json:"nombre_docente"`
	NumeroDocumento    string `json:"numero_documento"`
	NivelAcademico     string `json:"nivel_academico"`
	Facultad           string `json:"facultad"`
	ProyectoCurricular string `json:"proyecto_curricular"`
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
