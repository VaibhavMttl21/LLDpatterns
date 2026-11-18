package main

import (
    "fmt"
    "errors"
)

type CardPayment struct {
    cardNumber string
}

func NewCardPayment(cardNumber string) *CardPayment {
    return &CardPayment{
        cardNumber: cardNumber,
    }
}

func (c *CardPayment) Pay(amount float64) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }

    fmt.Printf(
        "Paid ₹%.2f using card ending in %s\n",
        amount,
        c.cardNumber[len(c.cardNumber)-4:],
    )

    return nil
}