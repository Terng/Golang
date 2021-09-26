package main

import (
	"fmt"
)

func main() {
	for i := 100; i < 200; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Is even number")
		} else {
			fmt.Println(i, "Is odd number")
		}
	}
}
