package main

import (
	"errors"
	"fmt"
)

type UpiPayment struct {
	UpiId string
}

func NewUpiPayemnt(UpiId string) *UpiPayment {
	return &UpiPayment{
		UpiId: UpiId,
	}
}

func (u *UpiPayment) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("Payemnt should be positive")
	}
	    fmt.Printf("Paid ₹%.2f using UPI ID %s\n",amount,u.UpiId,)

    return nil
}
