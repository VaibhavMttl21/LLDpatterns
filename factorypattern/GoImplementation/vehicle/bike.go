package vehicle

import "fmt"

type Car struct{}

func (c Car) Create() {
    fmt.Println("Creating Car")
}