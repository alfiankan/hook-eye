package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type PaymentData struct {
	Event       string `json:"event"`
	ReferenceID string `json:"reference_id"`
	Total       string `json:"total"`
}

func main() {
	app := fiber.New()

	kafka, errKafka := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"auto.offset.reset": "earliest",
		"group.id":          "websocket-service",
	})

	if errKafka != nil {
		panic(errKafka)
	}
	log.Println("Web Socket Running...")
	// consume if payment processed run webhook
	kafka.SubscribeTopics([]string{"PAYMENT.PROCESSED"}, nil)

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/notification/:client_id", websocket.New(func(c *websocket.Conn) {
		clientID := c.Params("client_id")
		for {
			msg, err := kafka.ReadMessage(-1)
			if err == nil {
				// get reference client id
				var paymentData PaymentData
				err := json.Unmarshal(msg.Value, &paymentData)
				if err != nil {
					log.Println(err)
				}
				log.Println(paymentData.ReferenceID, clientID)
				if paymentData.ReferenceID == clientID {
					log.Println("Notif sent", string(msg.Value), " to ")
					if err := c.WriteMessage(1, msg.Value); err != nil {
						log.Println("write:", err)
					}
				}

			} else {
				// if error
				log.Println(err)
			}
		}

	}))

	log.Fatal(app.Listen(":3000"))
}
