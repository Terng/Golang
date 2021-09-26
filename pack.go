package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		"Test1": 50.0,
		"Test2": 51.0,
		"Test3": 512.0,
		"Test4": 5133.0,
	}

	fmt.Println(menu)
	fmt.Println(menu["Test2"])
	// fmt.Println(" -- ")
	for k, v := range menu {
		fmt.Println("The key", k, "is", v)
	}
}
