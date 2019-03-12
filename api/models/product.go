package models

// Product is the structure representation for a simple product
type Product struct {
	Code  string  `json:"code,required"`
	Name  string  `json:"name"`
	Price float64 `json:"price,required"`
	Qty   int64   `json:"qty,omitempty"`
}
