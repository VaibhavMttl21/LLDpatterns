package main

import "fmt"

type WinButton struct{}

func (b WinButton) render() {
	fmt.Println("Rendering Windows button")
}
