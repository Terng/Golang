package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	name := []string{"value", "value2", "value3", "value4"}

	// for i := 0; i < len(name); i++ {
	// 	fmt.Println(name[i])
	// }

	for _, value := range name {
		fmt.Printf("The value at index is: %v \n", value)
	}
}
