package main

import "fmt"

type XMLToJSONAdapter struct {
	xmlParser XMLParser
}

func (a XMLToJSONAdapter) Analyze(jsonData string) {

	xmlData := a.xmlParser.GetXMLData()

	jsonData = convertXMLToJSON(xmlData)

	fmt.Println("Adapter converted XML to JSON:")
	fmt.Println(jsonData)

	fmt.Println("Analytics processing completed.")
}

func convertXMLToJSON(xmlData string) string {
	return `{"name":"Vaibhav","age":21}`
}