package main

import (
	"fmt"
)

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	names := []string{"name1", "name2", "name3", "name4", "name5"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at POS", index)
			continue
		}
		if index > 2 {
			fmt.Println("Bracking at POS", index)
			break
		}

		fmt.Printf("The value at POS %v is %v \n", index, value)
	}
}
