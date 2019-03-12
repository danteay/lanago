package discounts

import "github.com/danteay/lanago/api/models"

// TwoXOne apply 2 x 1 promotion to the current basket. It receives the a
// pointer to te current basket as first parameter and as a second parameter
// a condition to apply the discount.
//
// This function pass to the provided condition a models.Product structure
// to be validated, so you need to convert the interface by this way:
//
// e := elem.(models.Product)
func TwoXOne(b *models.Basket, c Condition) {
	for _, item := range b.Items {
		if c(item) {
			free := item.Qty / 2

			b.Total = b.Total - (float64(free) * item.Price)
		}
	}
}
