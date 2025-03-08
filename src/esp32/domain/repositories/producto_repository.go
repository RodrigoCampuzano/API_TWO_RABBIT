package repositories

type ProductoRepository interface {
	ObtenerInventario(codigo string) (int, error)
	ActualizarInventario(codigo string, cantidad int) error
}