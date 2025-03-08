package repositories

import "API_TWO/src/esp32/domain/entities"

type RespuestaRepository interface {
	EnviarRespuesta(respuesta entities.Respuesta) error
}