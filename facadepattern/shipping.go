package main

import "fmt"

type ShippingService struct{}

func (s *ShippingService) CreateShipment(orderID string, address string) string {
	shipmentID := "SHIP-" + orderID

	fmt.Printf("Creating shipment %s for address: %s\n",shipmentID,address)

	return shipmentID
}
