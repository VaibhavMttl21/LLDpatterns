package main

type ChocolateDecorator struct{
	CoffeeDecorator
}

func NewChocolateDecorator(coffee Coffee) Coffee{
	return ChocolateDecorator{
		CoffeeDecorator: CoffeeDecorator{
			coffee:coffee,
		},
	}
}

func (c ChocolateDecorator) Cost() float64{
	return c.coffee.Cost()+50
}
func (c ChocolateDecorator) Description() string{
	return c.coffee.Description()+"Added chocolate"
}