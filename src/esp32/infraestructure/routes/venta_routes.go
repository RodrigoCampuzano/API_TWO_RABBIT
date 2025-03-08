package routes

import (
	"net/http"

	"API_TWO/src/esp32/infraestructure/controllers"
	"API_TWO/src/esp32/infraestructure/repositories"
	"API_TWO/src/esp32/application"
)

func NewVentaRouter() http.Handler {
	mux := http.NewServeMux()

	// Inicializar dependencias
	ventaRepo := repositories.NewVentaRepositoryPostgres()
	productoRepo := repositories.NewProductoRepositoryPostgres()
	ventaService := application.NewVentaService(ventaRepo, productoRepo)
	ventaController := controllers.NewVentaController(ventaService)

	// Configurar rutas
	mux.HandleFunc("/procesar-venta", ventaController.ProcesarVenta)

	return mux
}