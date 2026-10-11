package main

import "fmt"

type PaymentService struct{}

func (p *PaymentService) ProcessPayment(userId string, amount float64) bool {
	fmt.Printf("Processing Payment of ₹%.2f for user %s\n",amount,userId)
	return amount > 0
}
