package main

import "fmt"

type WinTextbox struct{}

func (t WinTextbox) renderT() {
	fmt.Println("Rendering Windows textbox")
}
