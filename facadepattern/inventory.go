package main

import "fmt"

type InventoryService struct{}

func (i *InventoryService) CheckStock(itemID string, qty int) bool {
	fmt.Println("checking stock for", itemID)

	if qty <= 0{
		return false
	}

	return qty<=10
}
