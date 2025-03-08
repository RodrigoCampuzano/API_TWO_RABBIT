package main

import (
	"log"
	"net/http"

	"API_TWO/src/core/db"
	"API_TWO/src/core/middleware"
	"API_TWO/src/esp32/infraestructure/routes"
)

func main() {
	// Inicializar la base de datos
	db.Init()

	// Configurar rutas
	router := routes.NewVentaRouter()

	// Aplicar middleware CORS
	handler := middleware.CORS(router)

	// Iniciar servidor
	log.Println("API escuchando en :8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}