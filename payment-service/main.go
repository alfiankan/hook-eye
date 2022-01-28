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

func sendIfPaymentProcessed(data *kafka.Message) {
	time.Sleep(time.Second * 10)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "PAYMENT.PROCESSED"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data.Value,
	}, nil)

	p.Flush(15 * 1000)
}

func savePayment(data *kafka.Message) {

	var paymentData PaymentData
	err := json.Unmarshal(data.Value, &paymentData)
	if err != nil {
		log.Println(err)
	}
	log.Println("Payment With Reference ID ", paymentData.ReferenceID)
	sendIfPaymentProcessed(data)
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"auto.offset.reset": "earliest",
		"group.id":          "payment-service",
	})

	if err != nil {
		panic(err)
	}
	log.Println("Payment Service Running...")
	// consume if payment capture by front API
	c.SubscribeTopics([]string{"PAYMENT.CAPTURE"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// save payment
			savePayment(msg)
		} else {
			// if error
			log.Println(err)
		}
	}

	c.Close()
}
