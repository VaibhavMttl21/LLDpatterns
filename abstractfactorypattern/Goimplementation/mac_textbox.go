package main

import "fmt"

type MacTextbox struct{}

func (t MacTextbox) renderT() {
	fmt.Println("Rendering Mac textbox")
}
