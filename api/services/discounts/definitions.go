package discounts

// Condition is the definition of a discount condition to validate
// if its applicable. It receives an interface element in what the
// discount try to be applied. if the conditions are satisfied,
// this function returns true, and false in other cases.
type Condition func(elem interface{}) bool
