package domain_order

type CreateOrderRequest struct {
	CustomerId    uint `json:"customerid"`
	TransactionId uint `json:"transactionid"`
}

type FetchOrderRequest struct {
	CustomerId uint `json:"customerid"`
}
