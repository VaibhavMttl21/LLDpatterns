package main

type Winfactory struct{}

func (m Winfactory) CreateButton() Button{
	return WinButton{}
}

func (m Winfactory) CreateTextBox() Textbox{
	return WinTextbox{}
}