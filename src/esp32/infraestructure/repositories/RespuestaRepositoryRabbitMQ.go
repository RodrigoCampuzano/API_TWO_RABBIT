package repositories

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"API_TWO/src/esp32/domain/entities"
)

type RespuestaRepositoryRabbitMQ struct{}

func NewRespuestaRepositoryRabbitMQ() *RespuestaRepositoryRabbitMQ {
	return &RespuestaRepositoryRabbitMQ{}
}

func (r *RespuestaRepositoryRabbitMQ) EnviarRespuesta(respuesta entities.Respuesta) error {
	conn, err := amqp.Dial("amqp://rodrigo:123456789@52.0.68.153:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declarar la nueva cola para respuestas
	q, err := ch.QueueDeclare(
		"cola_respuestas", // nombre de la cola
		false,             // durable
		false,             // eliminar cuando no esté en uso
		false,             // exclusiva
		false,             // no esperar
		nil,               // argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Convertir la respuesta a JSON
	body, err := json.Marshal(respuesta)
	if err != nil {
		return err
	}

	// Publicar la respuesta en la cola
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}