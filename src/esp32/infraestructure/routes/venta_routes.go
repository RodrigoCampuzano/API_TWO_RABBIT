package routes

import (
	"net/http"

	"API_TWO/src/esp32/application"
	"API_TWO/src/esp32/infraestructure/controllers"
	"API_TWO/src/esp32/infraestructure/repositories"
)

func NewVentaRouter() http.Handler {
	mux := http.NewServeMux()

	// Inicializar dependencias
	ventaRepo := repositories.NewVentaRepositoryPostgres()
	productoRepo := repositories.NewProductoRepositoryPostgres()
	respuestaRepo := repositories.NewRespuestaRepositoryRabbitMQ()
	ventaService := application.NewVentaService(ventaRepo, productoRepo, respuestaRepo)
	ventaController := controllers.NewVentaController(ventaService)

	// Configurar rutas
	mux.HandleFunc("/procesar-venta", ventaController.ProcesarVenta)

	return mux
}