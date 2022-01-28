package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type PaymentData struct {
	Event       string `json:"event"`
	ReferenceID string `json:"reference_id"`
	Total       string `json:"total"`
}

func hitWebhook(data *kafka.Message) {

	var paymentData PaymentData
	err := json.Unmarshal(data.Value, &paymentData)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second * 1)
	log.Println("Webhhok Called With Data", string(data.Value))
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"auto.offset.reset": "earliest",
		"group.id":          "webhook-service",
	})

	if err != nil {
		panic(err)
	}
	log.Println("Webhook Running...")
	// consume if payment processed run webhook
	c.SubscribeTopics([]string{"PAYMENT.PROCESSED"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// save payment
			hitWebhook(msg)
		} else {
			// if error
			log.Println(err)
		}
	}

	c.Close()
}
