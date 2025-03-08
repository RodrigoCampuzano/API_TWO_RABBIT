package application

import (
	"fmt"
	"API_TWO/src/esp32/domain/entities"
	"API_TWO/src/esp32/domain/repositories"
)

type VentaService struct {
	ventaRepo         repositories.VentaRepository
	productoRepo      repositories.ProductoRepository
	respuestaRepo     repositories.RespuestaRepository
}

func NewVentaService(
	ventaRepo repositories.VentaRepository,
	productoRepo repositories.ProductoRepository,
	respuestaRepo repositories.RespuestaRepository,
) *VentaService {
	return &VentaService{
		ventaRepo:     ventaRepo,
		productoRepo:  productoRepo,
		respuestaRepo: respuestaRepo,
	}
}

func (s *VentaService) ProcesarVenta(venta entities.Venta) error {
	// Verificar inventario
	inventario, err := s.productoRepo.ObtenerInventario(venta.Producto)
	if err != nil {
		return err
	}

	if inventario < venta.Cantidad {
		return fmt.Errorf("inventario insuficiente")
	}

	// Registrar la venta
	err = s.ventaRepo.RegistrarVenta(venta)
	if err != nil {
		return err
	}

	// Actualizar inventario
	err = s.productoRepo.ActualizarInventario(venta.Producto, venta.Cantidad)
	if err != nil {
		return err
	}

	// Crear la respuesta
	respuesta := entities.Respuesta{
		Estado:   "éxito",
		Mensaje:  "Venta registrada correctamente",
		Producto: venta.Producto,
		Cantidad: venta.Cantidad,
	}

	// Enviar la respuesta a la nueva cola
	err = s.respuestaRepo.EnviarRespuesta(respuesta)
	if err != nil {
		return err
	}

	return nil
}