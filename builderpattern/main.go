package main

import "fmt"

func main() {
	builder := &GamingDesktopBuilder{}

	director := Director{
		builder: builder,
	}

	director.construct()

	desktop := builder.GetResult()

	desktop.Display()

	fmt.Println("Desktop built successfully!")
}
