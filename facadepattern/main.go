// That's the central idea: one simple interface hides the complexity of coordinating multiple subsystem classes.
package main

import "fmt"

func main() {
    orderFacade := NewOrderFacade()

    success := orderFacade.PlaceOrder("USER-101","LAPTOP-01",2,80000.00,"Chandigarh, India",)

    fmt.Println("\nFinal result:", success)
}

// 1.
// Client calls the facade
// main.go invokes PlaceOrder(...).
// The client doesn't need to know about the individual services.


// 2.
// Facade delegates
// OrderFacade calls InventoryService, then PaymentService, ShippingService, and NotificationService.
// Each service executes its own logic.


// 3.
// Facade returns the result
// PlaceOrder() returns true when all simulated steps succeed, or false if a checked step fails.