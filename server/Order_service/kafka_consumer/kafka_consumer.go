package kafka_consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/domain_order"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

func ConsumeCart(db *gorm.DB) {

	topic := "cart"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "cart_group",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Fatal(err)
	}
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}
	//infinite loop for receiving object
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			//decode the byte message received from kafka into desired model
			var data domain_order.KafkaMessagePayload
			err := json.Unmarshal(e.Value, &data)
			if err != nil {
				log.Fatal(err)
			}
			//-------------------------------------------------------------//
			if data.Operation == "Add" {
				AddToCart(db, data.CustomerId, data.Product)
			} else {
				RemoveToCart(db, data.CustomerId, data.Product)
			}
		case *kafka.Error:
			fmt.Printf("%s\n", e)
		}
	}
}
