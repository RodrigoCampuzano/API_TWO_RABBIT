package entities

type Producto struct {
	ID         int    `json:"id"`
	Codigo     string `json:"codigo"`
	Nombre     string `json:"nombre"`
	Inventario int    `json:"inventario"`
}