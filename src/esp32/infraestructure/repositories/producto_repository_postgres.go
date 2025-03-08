package repositories

import (
	"log"

	"API_TWO/src/core/db"
)

type ProductoRepositoryPostgres struct{}

func NewProductoRepositoryPostgres() *ProductoRepositoryPostgres {
	return &ProductoRepositoryPostgres{}
}

func (r *ProductoRepositoryPostgres) ObtenerInventario(codigo string) (int, error) {
	var inventario int
	err := db.DB.QueryRow("SELECT inventario FROM productos WHERE codigo = $1", codigo).Scan(&inventario)
	if err != nil {
		log.Println("Error al verificar inventario:", err)
		return 0, err
	}
	return inventario, nil
}

func (r *ProductoRepositoryPostgres) ActualizarInventario(codigo string, cantidad int) error {
	_, err := db.DB.Exec("UPDATE productos SET inventario = inventario - $1 WHERE codigo = $2", cantidad, codigo)
	if err != nil {
		log.Println("Error al actualizar inventario:", err)
		return err
	}
	return nil
}