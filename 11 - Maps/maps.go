package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// We can use only a type for keys and a type for values
	user := map[string]string{
		"first_name": "Leonardo",
		"last_name":  "Monteiro",
	}
	fmt.Println(user["first_name"])

	// Map type into a map
	user2 := map[string]map[string]string{
		"name": {
			"first_name": "Leonardo",
			"last_name":  "Monteiro",
		},
		"course": {
			"name":       "ADS",
			"university": "UDF",
		},
	}
	fmt.Println(user2)

	// Remove map keys
	delete(user2, "course")
	fmt.Println(user2)

	// Add map keys
	user2["course"] = map[string]string{
		"name":       "ADS",
		"university": "UDF",
	}
	fmt.Println(user2)
}
