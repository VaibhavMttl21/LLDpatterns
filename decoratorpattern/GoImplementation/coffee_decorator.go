// Responsibility 1: Store the wrapped object
// coffee Coffee
// This holds whichever coffee object we wrap: plain coffee, coffee with milk, or coffee with other ingredients.
// Responsibility 2: Forward method calls
// func (d CoffeeDecorator) Cost() float64 {
//     return d.coffee.Cost()
// }
// If a decorator does not need to change the price, it can use this default behavior instead of writing the forwarding code again.

package main

type CoffeeDecorator struct{
	coffee Coffee
}

func (d CoffeeDecorator) Description() string{
	return d.coffee.Description()
}
func (d CoffeeDecorator) Cost() float64{
	return d.coffee.Cost()
}