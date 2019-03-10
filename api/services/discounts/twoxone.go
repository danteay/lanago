package discounts

import "github.com/danteay/lanago/api/models"

// TwoXOne apply 2 x 1 promotion to the current basquet.
// It receives the a pointer to te current basket as first parameter
// and as a secondparameter a clousere with the conditions to apply
// the discount.
func TwoXOne(b *models.Basket, c Condition) {
	for _, item := range b.Items {
		if c(item) {
			free := item.Qty / 2

			b.Total = b.Total - (float64(free) * item.Price)
		}
	}
}
