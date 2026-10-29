package main

func main() {

	xmlParser := XMLParser{}

	adapter := XMLToJSONAdapter{
		xmlParser: xmlParser,
	}

	client := AnalyticsClient{
		tool: adapter,
	}

	client.Run("")
}


// AnalyticsClient
//       |
//       ↓
// AnalyticsTool
//       ↑
//       |
// XMLToJSONAdapter ───→ XMLParser


// Client
//   |
//   | Analyze()
//   ↓
// AnalyticsTool
//   ↑
//   |
// XMLToJSONAdapter
//   |
//   | GetXMLData()
//   ↓
// XMLParser