package models

import "time"

type Dependencia struct {
	Id                         int                          `json:"Id"`
	Nombre                     string                       `json:"Nombre"`
	TelefonoDependencia        string                       `json:"TelefonoDependencia"`
	CorreoElectronico          string                       `json:"CorreoElectronico"`
	Activo                     bool                         `json:"Activo"`
	FechaCreacion              time.Time                    `json:"FechaCreacion"`
	FechaModificacion          time.Time                    `json:"FechaModificacion"`
	DependenciaTipoDependencia []DependenciaTipoDependencia `json:"DependenciaTipoDependencia"`
}

type DependenciaTipoDependencia struct {
	Id                int             `json:"Id"`
	TipoDependenciaId TipoDependencia `json:"TipoDependenciaId"`
	DependenciaId     Dependencia     `json:"DependenciaId"`
	Activo            bool            `json:"Activo"`
	FechaCreacion     time.Time       `json:"FechaCreacion"`
	FechaModificacion time.Time       `json:"FechaModificacion"`
}

type TipoDependencia struct {
	Id                int       `json:"Id"`
	Nombre            string    `json:"Nombre"`
	Descripcion       string    `json:"Descripcion"`
	CodigoAbreviacion string    `json:"CodigoAbreviacion"`
	Activo            bool      `json:"Activo"`
	FechaCreacion     time.Time `json:"FechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion"`
}
