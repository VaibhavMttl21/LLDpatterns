package main

import "fmt"

type MacButton struct{}

func (b MacButton) render() {
	fmt.Println("Rendering mac button")
}
