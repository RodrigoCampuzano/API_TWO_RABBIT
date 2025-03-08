package repositories

import (
	"log"

	"API_TWO/src/core/db"
	"API_TWO/src/esp32/domain/entities"
)

type VentaRepositoryPostgres struct{}

func NewVentaRepositoryPostgres() *VentaRepositoryPostgres {
	return &VentaRepositoryPostgres{}
}

func (r *VentaRepositoryPostgres) RegistrarVenta(venta entities.Venta) error {
	var productoID int
	err := db.DB.QueryRow("SELECT id FROM productos WHERE codigo = $1", venta.Producto).Scan(&productoID)
	if err != nil {
		log.Println("Error al obtener ID del producto:", err)
		return err
	}

	_, err = db.DB.Exec("INSERT INTO ventas (producto_id, cantidad) VALUES ($1, $2)", productoID, venta.Cantidad)
	if err != nil {
		log.Println("Error al registrar la venta:", err)
		return err
	}

	return nil
}