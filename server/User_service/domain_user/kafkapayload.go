package domain_user

type KafkaMessagePayload struct {
	CustomerId uint
	Product
	Operation string
}
