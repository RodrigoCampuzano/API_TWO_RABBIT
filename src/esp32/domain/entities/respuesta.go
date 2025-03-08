package entities

type Respuesta struct {
	Estado   string `json:"estado"`
	Mensaje  string `json:"mensaje"`
	Producto string `json:"producto"`
	Cantidad int    `json:"cantidad"`
}