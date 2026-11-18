package main

type PlainCoffeee struct{}

func (p PlainCoffeee) Cost() float64{
	return 100
}
func (p PlainCoffeee) Description() string{
	return "Plain Coffee"
}