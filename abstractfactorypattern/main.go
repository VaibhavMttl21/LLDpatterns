package main

import "fmt"

func main() {
	var os string

	fmt.Print("Enter OS (mac/win): ")
	fmt.Scanln(&os)

	factory,err := GetGuiFactory(os)
	if(err!=nil){
		return 
	}

	button := factory.CreateButton()
	textbox := factory.CreateTextBox()

	button.render()
	textbox.renderT()
}
