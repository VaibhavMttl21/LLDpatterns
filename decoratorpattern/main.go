package main

import "fmt"

func main() {
	var coffee Coffee = PlainCoffeee{}
	fmt.Println("Cost=", coffee.Cost(), "Description=", coffee.Description())

	coffee = NewMilkDecorator(coffee)
	fmt.Println("Cost=", coffee.Cost(), "Description=", coffee.Description())

	coffee = NewChocolateDecorator(coffee)
	fmt.Println("Cost=", coffee.Cost(), "Description=", coffee.Description())

	coffee = NewLabelDecorator(coffee)
	fmt.Println("Cost=", coffee.Cost(), "Description=", coffee.Description())
}
