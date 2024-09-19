package model

type Order struct {
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Amount     float64 `json:"amount"`
}
