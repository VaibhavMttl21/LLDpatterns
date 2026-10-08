package main

import "fmt"

type XMLParser struct{}

func (x XMLParser) GetXMLData() string {
	xmlData := `<user><name>Vaibhav</name><age>21</age></user>`

	fmt.Println("XML Parser produced:")
	fmt.Println(xmlData)

	return xmlData
}