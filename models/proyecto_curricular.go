package models

import "time"

type ProyectoCurricular struct {
	Id                         int       `json:"Id"`
	Nombre                     string    `json:"Nombre"`
	TelefonoDependencia        string    `json:"TelefonoDependencia"`
	CorreoElectronico          string    `json:"CorreoElectronico"`
	Activo                     bool      `json:"Activo"`
	FechaCreacion              time.Time `json:"FechaCreacion"`
	FechaModificacion          time.Time `json:"FechaModificacion"`
	DependenciaTipoDependencia *int      `json:"DependenciaTipoDependencia"`
}
