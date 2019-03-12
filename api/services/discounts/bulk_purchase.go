package discounts

import "github.com/danteay/lanago/api/models"

// BulkPurchase has the logic to manage a bulk discount. It receives a
// model.Basket pointer to the current basket that you want to apply the
// discount, the price per unit that you want to apply and a condition
// to apply it.
//
// This function pass to the provided condition a models.Product structure
// to be validated, so you need to convert the interface by this way:
//
// e := elem.(models.Product)
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
