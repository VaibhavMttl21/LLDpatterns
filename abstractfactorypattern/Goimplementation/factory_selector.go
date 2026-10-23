package main

import "fmt"

func GetGuiFactory(os string) (GUIFactory, error) {
	if os == "mac" {
		return Macfactory{}, nil
	} else if os == "win" {
		return Winfactory{}, nil
	}
	return nil, fmt.Errorf("unsupported OS: %s", os)
}
