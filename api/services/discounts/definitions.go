package discounts

// Condition is the definition of a condition closure to valite if
// a discount can be appliend to a basket
type Condition func(elem interface{}) bool