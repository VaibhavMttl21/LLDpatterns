package main

import "fmt"

type OrderFacade struct {
	inventory    *InventoryService
	payment      *PaymentService
	shipping     *ShippingService
	notification *NotificationService
}

//NewOrderFacade() creates the facade and its four subsystem objects.
// The fields are pointers, such as *InventoryService, so the facade stores references to these objects. The client doesn't need to create or manage each service itself.

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory:    &InventoryService{},
		payment:      &PaymentService{},
		shipping:     &ShippingService{},
		notification: &NotificationService{},
	}
}

func (o OrderFacade) PlaceOrder(UserId string, itemID string, qty int, amount float64, address string) bool {
	fmt.Println("\n starting order process...")

	//1. check inventory
	if !o.inventory.CheckStock(itemID, qty) {
		fmt.Println("Order Failed :insufficeint stock or invalid quantity")
		return false
	}

	//2. process payment
	if !o.payment.ProcessPayment(UserId, amount) {
		fmt.Println("Order failed:payment failed")
		return false
	}

	//3. create shipment
	orderID := "ORD-1001"
	shipmentID := o.shipping.CreateShipment(orderID, address)

	if shipmentID == "" {
		fmt.Println("Order failed: shipment creation failed")
		return false
	}

	//4. Notification
	o.notification.SendOrderConfirmation(UserId, orderID)

	fmt.Println("Order placed successfully!")
	return true
}
