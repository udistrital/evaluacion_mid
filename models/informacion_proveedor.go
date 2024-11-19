package models

type InformacionProveedor struct {
	Id                      int                `orm:"column(id);pk"`
	Tipopersona             string             `orm:"column(tipopersona)"`
	NumDocumento            string             `orm:"column(numdocumento)"`
	IdCiudadContacto        float64            `orm:"column(id_ciudad_contacto)"`
	Direccion               string             `orm:"column(direccion)"`
	Correo                  string             `orm:"column(correo)"`
	Web                     string             `orm:"column(web)"`
	NomAsesor               string             `orm:"column(nom_asesor)"`
	TelAsesor               string             `orm:"column(tel_asesor)"`
	Descripcion             string             `orm:"column(descripcion)"`
	PuntajeEvaluacion       float64            `orm:"column(puntaje_evaluacion)"`
	ClasificacionEvaluacion string             `orm:"column(clasificacion_evaluacion)"`
	Estado                  *ParametroEstandar `orm:"column(estado);rel(fk)"`
	TipoCuentaBancaria      string             `orm:"column(tipo_cuenta_bancaria)"`
	NumCuentaBancaria       string             `orm:"column(num_cuenta_bancaria)"`
	IdEntidadBancaria       float64            `orm:"column(id_entidad_bancaria)"`
	FechaRegistro           string             `orm:"column(fecha_registro)"`
	FechaUltimaModificacion string             `orm:"column(fecha_ultima_modificacion)"`
	NomProveedor            string             `orm:"column(nom_proveedor)"`
	Anexorut                string             `orm:"column(anexorut)"`
	Anexorup                string             `orm:"column(anexorup)"`
	RegimenContributivo     string             `orm:"column(regimen_contributivo)"`
}

type ParametroEstandar struct {
	Id                   int    `orm:"column(id);pk"`
	ClaseParametro       string `orm:"column(clase_parametro)"`
	ValorParametro       string `orm:"column(valor_parametro)"`
	DescripcionParametro string `orm:"column(descripcion_parametro)"`
}
