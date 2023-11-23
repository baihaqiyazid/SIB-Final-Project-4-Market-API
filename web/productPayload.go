package web

type CreateProductPayload struct {
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryId uint   `json:"category_id"`
}
