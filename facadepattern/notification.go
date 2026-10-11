package main

import "fmt"

type NotificationService struct{}

func (n *NotificationService) SendOrderConfirmation(userID string,orderID string,) {
    fmt.Printf("Order confirmation sent to user %s for order %s\n", userID,orderID)
}