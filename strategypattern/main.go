package main

import "fmt"

func main() {
    // Start with UPI payment.
    upi := NewUpiPayemnt("user@upi")
    cart := NewShoppingCart(upi)

    if err := cart.Checkout(1500); err != nil {
        fmt.Println(err)
    }

    fmt.Println()

    // Switch to card payment at runtime.
    card := NewCardPayment("1234567812345678")
    cart.SetPaymentStrategy(card)

    if err := cart.Checkout(2500); err != nil {
        fmt.Println(err)
    }
}


// NewShoppingCart(upi) stores the UPI strategy.
// Checkout(1500) calls s.strategy.Pay(1500).
// The UPI implementation handles the call.
// SetPaymentStrategy(card) replaces the strategy.
// The next checkout invokes the card implementation instead.