package main

type GUIFactory interface{
	CreateButton() Button
	CreateTextBox() Textbox
}

