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