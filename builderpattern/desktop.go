// product side 

package main

import "fmt"

type Desktop struct {
	CPU     string
	RAM     int
	Storage int
	GPU     string
}

func (d Desktop) Display() {
	fmt.Println("Desktop Configuration:")
	fmt.Println("CPU:", d.CPU)
	fmt.Println("RAM:", d.RAM, "GB")
	fmt.Println("Storage:", d.Storage, "GB")
	fmt.Println("GPU:", d.GPU)
}
