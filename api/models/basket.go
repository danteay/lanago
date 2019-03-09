package models

// Basket is the representation of a Checkout basket
type Basket struct {
	ID    string    `json:"id"`
	Total float64   `json:"total"`
	Items []Product `json:"items"`
}
