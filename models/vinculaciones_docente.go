package models

import "time"

type ResolucionVinculacionDocente struct {
	Id                int    `json:"Id"`
	FacultadId        int    `json:"FacultadId"`
	Dedicacion        string `json:"Dedicacion"`
	NivelAcademico    string `json:"NivelAcademico"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}

type VinculacionesDocenteResolucion struct {
	Id                           int                          `json:"Id"`
	NumeroContrato               string                       `json:"NumeroContrato"`
	Vigencia                     int                          `json:"Vigencia"`
	PersonaId                    int64                        `json:"PersonaId"`
	NumeroHorasSemanales         int                          `json:"NumeroHorasSemanales"`
	NumeroSemanas                int                          `json:"NumeroSemanas"`
	PuntoSalarialId              int                          `json:"PuntoSalarialId"`
	SalarioMinimoId              int                          `json:"SalarioMinimoId"`
	ResolucionVinculacionDocente ResolucionVinculacionDocente `json:"ResolucionVinculacionDocenteId"`
	DedicacionId                 int                          `json:"DedicacionId"`
	ProyectoCurricularId         int                          `json:"ProyectoCurricularId"`
	ValorContrato                float64                      `json:"ValorContrato"`
	Categoria                    string                       `json:"Categoria"`
	DependenciaAcademica         int                          `json:"DependenciaAcademica"`
	NumeroRp                     int                          `json:"NumeroRp"`
	VigenciaRp                   int                          `json:"VigenciaRp"`
	FechaInicio                  time.Time                    `json:"FechaInicio"`
	Activo                       bool                         `json:"Activo"`
	FechaCreacion                string                       `json:"FechaCreacion"`
	FechaModificacion            string                       `json:"FechaModificacion"`
	NumeroHorasTrabajadas        int                          `json:"NumeroHorasTrabajadas"`
}

type VinculacionesDocente struct {
	NumeroContrato         string    `json:"NumeroContrato"`
	Vigencia               int       `json:"Vigencia"`
	Periodo                int       `json:"Periodo"`
	FechaInicio            time.Time `json:"FechaInicio"`
	FechaFin               time.Time `json:"FechaFin"`
	NumeroHorasSemanales   int       `json:"NumeroHorasSemanales"`
	NumeroSemanas          int       `json:"NumeroSemanas"`
	NumeroHorasSemestrales int       `json:"NumeroHorasSemestrales"`
	Dedicacion             string    `json:"Dedicacion"`
	ProyectoCurricular     string    `json:"ProyectoCurricular"`
	DependenciaAcademica   int       `json:"DependenciaAcademica"`
}

type VinculacionesDocenteOld struct {
	Docente struct {
		VinculacionesOld []VinculacionOld `json:"vinculaciones"`
	} `json:"docente"`
}

type VinculacionOld struct {
	Vigencia             string    `json:"vigencia"`
	Descripcion          string    `json:"descripcion"`
	NumeroContrato       string    `json:"numero_contrato"`
	FechaInicio          time.Time `json:"fecha_inicio"`
	NumeroHorasSemanales string    `json:"numero_horas_semanales"`
	Periodo              string    `json:"periodo"`
	FechaFin             time.Time `json:"fecha_fin"`
	ProyectoCurricularId string    `json:"proyecto_curricular_id"`
	NumeroSemanas        string    `json:"numero_semanas"`
	DedicacionId         string    `json:"dedicacion_id"`
}
