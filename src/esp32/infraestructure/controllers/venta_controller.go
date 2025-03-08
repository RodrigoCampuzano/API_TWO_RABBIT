package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"API_TWO/src/esp32/application"
	"API_TWO/src/esp32/domain/entities"
)

type VentaController struct {
	ventaService *application.VentaService
}

func NewVentaController(ventaService *application.VentaService) *VentaController {
	return &VentaController{ventaService: ventaService}
}

func (c *VentaController) ProcesarVenta(w http.ResponseWriter, r *http.Request) {
	var venta entities.Venta
	json.NewDecoder(r.Body).Decode(&venta)

	err := c.ventaService.ProcesarVenta(venta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"estado": "error", "mensaje": err.Error()})
		return
	}

	respuesta := map[string]string{
		"estado":   "éxito",
		"mensaje":  "Venta registrada correctamente",
		"producto": venta.Producto,
		"cantidad": fmt.Sprintf("%d", venta.Cantidad),
	}
	json.NewEncoder(w).Encode(respuesta)
}