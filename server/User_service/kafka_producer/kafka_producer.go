package kafka_producer

import (
	"encoding/json"
	"log"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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

func (op *OrderPlacer) placeOrder(cartpayload domain_user.KafkaMessagePayload) error {

	payload, err := json.Marshal(cartpayload)

	if err != nil {
		log.Fatal(err)
	}

	err = op.producer.Produce(&kafka.Message{
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

func ProduceCart(cartpayload domain_user.KafkaMessagePayload) error {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "unique6",
		"acks":              "all",
	})

	if err != nil {
		return err
	}

	op := NewOrderPlacer(p, "cart")
	if err := op.placeOrder(cartpayload); err != nil {
		return err
	}

	log.Println("product sent......")

	return nil
}
