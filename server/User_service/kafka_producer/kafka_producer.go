package kafka_producer

import (
	"fmt"
	"log"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderPlacer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewOrderPlacer(p *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 100),
	}
}

func (op *OrderPlacer) placeOrder(user domain_user.User) error {

	payload := []byte(fmt.Sprintf("%v", user))

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny},
		Value: payload},
		op.deliverych,
	)
	if err != nil {
		log.Fatal(err)
	}

	<-op.deliverych
	return nil

}

func ProduceCart(user domain_user.User) error {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "usergroup",
		"acks":              "all",
	})

	if err != nil {
		return err
	}

	op := NewOrderPlacer(p, "product")
	if err := op.placeOrder(user); err != nil {
		return err
	}

	log.Println("product sent......")

	return nil
}
