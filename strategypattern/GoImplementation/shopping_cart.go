package main

import (
    "errors"
    "fmt"
)

type ShoppingCart struct {
    strategy PaymentStrategy
}

func NewShoppingCart(strategy PaymentStrategy) *ShoppingCart {
    return &ShoppingCart{
        strategy: strategy,
    }
}

func (s *ShoppingCart) SetPaymentStrategy(
    strategy PaymentStrategy,
) {
    s.strategy = strategy
}

func (s *ShoppingCart) Checkout(amount float64) error {
    if s.strategy == nil {
        return errors.New("payment strategy not configured")
    }

    fmt.Println("Starting checkout...")

    if err := s.strategy.Pay(amount); err != nil {
        return fmt.Errorf("checkout failed: %w", err)
    }

    fmt.Println("Checkout successful!")
    return nil
}