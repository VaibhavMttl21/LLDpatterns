package vehicle

import "fmt"

type Bike struct{}

func (b Bike) Create() {
    fmt.Println("Creating Bike")
}