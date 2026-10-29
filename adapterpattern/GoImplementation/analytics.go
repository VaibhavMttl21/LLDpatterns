package main

import "fmt"

type AnalyticsTool interface {
	Analyze(jsonData string)
}

type AnalyticsClient struct {
	tool AnalyticsTool
}

func (c AnalyticsClient) Run(jsonData string) {
	fmt.Println("Analytics Client:")
	c.tool.Analyze(jsonData)
}