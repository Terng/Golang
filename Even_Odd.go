package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Pongdilok Ketngam"

	fmt.Println(strings.Contains(greeting, "Pongdilok"))
	fmt.Println(strings.ReplaceAll(greeting, "Pongdilok", "Terng"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ok"))
	fmt.Println(strings.Split(greeting, "k K"))

	fmt.Println("Orignal string value =", greeting)

	ages := []int{1, 2, 3, 4, 5}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Terng", "Pongdilok", "Ketngam", "PK", "T"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Pongdilok"))
}
