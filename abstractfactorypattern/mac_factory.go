package main

type Macfactory struct{}

func (m Macfactory) CreateButton() Button{
	return MacButton{}
}

func (m Macfactory) CreateTextBox() Textbox{
	return MacTextbox{}
}