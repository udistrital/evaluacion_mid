package models

type Detalle struct {
	Contrato        string
	Vigencia        int
	TotalDevengado  float64
	TotalDescuentos float64
	TotalPago       float64
	Detalle         []DetallePreliquidacion
}
