package web

type CreateTransactionPayload struct {
	ProductID    uint `json:"product_id"`
	Quantity 	 int  `json:"quantity"`
}