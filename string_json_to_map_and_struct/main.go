package main

import (
	"encoding/json"
	"fmt"
)

type Corona struct {
	Name    string
	Country string
	City    string
	Reason  string
}

func main() {
	coronaVirusJSON := `{
			"name" : "covid-11",
			"country" : "China",
			"city" : "Wuhan",
			"reason" : "Non vedge Food"
		}`

	var coronaMap map[string]interface{}
	var corona Corona

	json.Unmarshal([]byte(coronaVirusJSON), &corona)
	json.Unmarshal([]byte(coronaVirusJSON), &coronaMap)

	fmt.Println(corona)

	fmt.Println("Name :", coronaMap["name"],
		"\nCountry :", coronaMap["country"],
		"\nCity :", coronaMap["city"],
		"\nReason :", coronaMap["reason"])
}
