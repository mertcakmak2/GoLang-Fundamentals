package main

import "fmt"

func main() {

	m := make(map[string]int)
	//or
	// m := map[string]int{}
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// val := map[string]interface{}{}
	// or
	val := make(map[string]interface{})
	val["name"] = "mert"
	val["lastname"] = "cakmak"
	val["age"] = 24

	fmt.Printf("val: %v\n", val)
	fmt.Printf("val[\"name\"]: %v\n", val["name"])

	// or
	val2 := map[string]interface{}{
		"name":     "mert",
		"lastname": "cakmak",
		"age":      24,
	}
	fmt.Printf("val2: %v\n", val2)
}
