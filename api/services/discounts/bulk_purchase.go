package discounts

import (
	"lanago/api/models"
)

// BulkPurchase has the logic to
func BulkPurchase(b *models.Basket, price float64, c Condition) {
	for _, item := range b.Items {
		if c(item) {
			qty := float64(item.Qty)

			free := qty * item.Price
			add := qty * price

			b.Total = b.Total - free + add
		}
	}
}
