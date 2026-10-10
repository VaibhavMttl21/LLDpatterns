package main

type MilkDecorator struct{
	CoffeeDecorator
}
// MilkDecorator
// ├── embedded CoffeeDecorator
// │   └── coffee Coffee
// ├── Description()  ← MilkDecorator's own method
// └── Cost()         ← MilkDecorator's own method

func NewMilkDecorator(coffee Coffee) Coffee{
	return MilkDecorator{
		CoffeeDecorator: CoffeeDecorator{
			coffee:coffee,
		},
	}
}

// after calling constructor the structure
// MilkDecorator
//      |
//      └── CoffeeDecorator
//               |
//               └── coffee: PlainCoffee{}

// The return type is the interface, not the concrete type.
// This means the client only needs to know that it received a Coffee. It doesn't need to know that the concrete value is MilkDecorator.

func (m MilkDecorator) Cost() float64{
	return m.coffee.Cost()+30
}
func (m MilkDecorator) Description() string{
	return m.coffee.Description()+"Milk added"
}