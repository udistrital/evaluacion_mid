package models

type DetallePreliquidacion struct {
	Id                       int
	ContratoPreliquidacionId *ContratoPreliquidacion
	ValorCalculado           float64
	DiasLiquidados           float64
	DiasEspecificos          string
	TipoPreliquidacionId     int
	ConceptoNominaId         *ConceptoNomina
	EstadoDisponibilidadId   int
	Activo                   bool
	FechaCreacion            string
	FechaModificacion        string
	NombreCompleto           string
	Documento                string
	NumeroContrato           string
	VigenciaContrato         int
	Persona                  int
}
