package entities

type Venta struct {
	Producto string `json:"producto"`
	Cantidad int    `json:"cantidad"`
}