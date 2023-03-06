package domain_order

type KafkaMessagePayload struct {
	CustomerId uint
	Product
	Operation string
}

type ProductPaylod struct {
	AddedIn         uint
	ProductID       uint
	ProductImageSrc string
	Name            string
	Description     string
	Price           float32
	Unit            int
	Available       bool
	Category        string
}
