package models

import "time"

type ParametrosData struct {
	Data []Parametro `json:"Data"`
}

type Parametro struct {
	Id                int             `json:"Id"`
	Nombre            string          `json:"Nombre"`
	Descripcion       string          `json:"Descripcion"`
	CodigoAbreviacion string          `json:"CodigoAbreviacion"`
	Activo            bool            `json:"Activo"`
	NumeroOrden       int             `json:"NumeroOrden"`
	FechaCreacion     time.Time       `json:"FechaCreacion"`
	FechaModificacion time.Time       `json:"FechaModificacion"`
	TipoParametroId   TipoParametro   `json:"TipoParametroId"`
	ParametroPadreId  *ParametroPadre `json:"ParametroPadreId"`
}

type TipoParametro struct {
	Id                int        `json:"Id"`
	Nombre            string     `json:"Nombre"`
	Descripcion       string     `json:"Descripcion"`
	CodigoAbreviacion string     `json:"CodigoAbreviacion"`
	Activo            bool       `json:"Activo"`
	NumeroOrden       int        `json:"NumeroOrden"`
	FechaCreacion     time.Time  `json:"FechaCreacion"`
	FechaModificacion time.Time  `json:"FechaModificacion"`
	AreaTipoId        AreaTipoId `json:"AreaTipoId"`
}

type AreaTipoId struct {
	Id                int       `json:"Id"`
	Nombre            string    `json:"Nombre"`
	Descripcion       string    `json:"Descripcion"`
	CodigoAbreviacion string    `json:"CodigoAbreviacion"`
	Activo            bool      `json:"Activo"`
	NumeroOrden       int       `json:"NumeroOrden"`
	FechaCreacion     time.Time `json:"FechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion"`
}

type ParametroPadre struct {
	Id                int             `json:"Id"`
	Nombre            string          `json:"Nombre"`
	Descripcion       string          `json:"Descripcion"`
	CodigoAbreviacion string          `json:"CodigoAbreviacion"`
	Activo            bool            `json:"Activo"`
	NumeroOrden       int             `json:"NumeroOrden"`
	FechaCreacion     time.Time       `json:"FechaCreacion"`
	FechaModificacion time.Time       `json:"FechaModificacion"`
	TipoParametroId   TipoParametro   `json:"TipoParametroId"`
	ParametroPadreId  *ParametroPadre `json:"ParametroPadreId"`
}
