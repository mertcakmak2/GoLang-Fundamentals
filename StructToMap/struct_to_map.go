package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	convertStructToMap()
	convertMapToStruct()
}

func convertStructToMap() {

	var person Person

	personMap := map[string]interface{}{}
	personMap["Name"] = "mert"
	personMap["Age"] = 24

	marshalPersonMap, _ := json.Marshal(personMap)
	json.Unmarshal(marshalPersonMap, &person)

	fmt.Println(person)
}

func convertMapToStruct() {

	person := Person{Name: "John", Age: 34}
	var personMap map[string]interface{}

	marshalPersonStruct, _ := json.Marshal(person)
	json.Unmarshal(marshalPersonStruct, &personMap)

	fmt.Println(personMap)
}
