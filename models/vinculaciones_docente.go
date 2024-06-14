package models

type ResolucionVinculacionDocente struct {
	Id                int    `json:"Id"`
	FacultadId        int    `json:"FacultadId"`
	Dedicacion        string `json:"Dedicacion"`
	NivelAcademico    string `json:"NivelAcademico"`
	Activo            bool   `json:"Activo"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}

type VinculacionesDocente struct {
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
	FechaInicio                  string                       `json:"FechaInicio"`
	Activo                       bool                         `json:"Activo"`
	FechaCreacion                string                       `json:"FechaCreacion"`
	FechaModificacion            string                       `json:"FechaModificacion"`
	NumeroHorasTrabajadas        int                          `json:"NumeroHorasTrabajadas"`
}
