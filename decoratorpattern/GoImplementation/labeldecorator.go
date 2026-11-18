package main

type LabelDecorator struct {
    CoffeeDecorator
}

func NewLabelDecorator(coffee Coffee) Coffee{
	return LabelDecorator{
		CoffeeDecorator:CoffeeDecorator{
			coffee:coffee,
		},
	}
}

func (d LabelDecorator) Description() string {
    return "Special: " + d.coffee.Description()
}