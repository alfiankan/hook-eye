package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type PaymentData struct {
	Event	string `json:"event"`
	ReferenceID string `json:"reference_id"`
	Total int `json:"total"`
}

func savePayment(data []byte)  {

	var paymentData PaymentData
	err := json.Unmarshal(data, &paymentData)
	if err != nil {
		log.Println(err)
	}
	log.Println("Payment With Reference ID ", paymentData.ReferenceID)
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"PAYMENT.CAPTURE"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// save payment
			savePayment(msg.Value)
		} else {
			// if error
			log.Println(msg.TopicPartition, string(msg.Value))
		}
	}

	c.Close()
}