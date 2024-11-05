package models

type BodyAutenticacion struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Version  string `json:"version"`
}

type AutenticacionResponse struct {
	Token string `json:"token"`
}

type BodyCargaAcademica struct {
	Parametros struct {
		Rol            string `json:"rol"`
		Identificacion string `json:"identificacion"`
		Facultad       string `json:"facultad"`
		Proyecto       string `json:"proyecto"`
	} `json:"parametros"`
}

type CargaAcademica struct {
	Anio                  int    `json:"ANIO"`
	Periodo               int    `json:"PERIODO"`
	CodProyecto           int    `json:"COD_PROYECTO"`
	Proyecto              string `json:"PROYECTO"`
	DocDocente            int    `json:"DOC_DOCENTE"`
	Docente               string `json:"DOCENTE"`
	CodVinculacion        int    `json:"COD_VINCULACION"`
	TipoVinculacion       string `json:"TIPO_VINCULACION"`
	CodEspacio            int    `json:"COD_ESPACIO"`
	Espacio               string `json:"ESPACIO"`
	Grupo                 string `json:"GRUPO"`
	IdGrupo               int    `json:"ID_GRUPO"`
	IdDia                 int    `json:"ID_DIA"`
	Dia                   string `json:"DIA"`
	Hora                  int    `json:"HORA"`
	HoraLarga             string `json:"HORA_LARGA"`
	CodProyectoEstudiante int    `json:"COD_PROYECTO_ESTUDIANTE"`
	CodEstudiante         int    `json:"COD_ESTUDIANTE"`
	Estudiante            string `json:"ESTUDIANTE"`
}
