package repositories

import "API_TWO/src/esp32/domain/entities"

type VentaRepository interface {
	RegistrarVenta(venta entities.Venta) error
}