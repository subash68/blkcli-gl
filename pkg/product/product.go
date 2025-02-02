package product

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"product_name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
