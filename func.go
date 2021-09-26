package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func SayBye(n string) {
	fmt.Printf("GoodBye %v \n", n)
}

func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Terng")
	// sayGreeting("pongdilok")
	// SayBye("Terng")
	// SayBye("pongdilok")
	// cycleName([]string{"Terng", "pongdilok", "Google"}, sayGreeting)
	// cycleName([]string{"Terng", "pongdilok", "Google"}, SayBye)

	a1 := circleArea(10)
	a2 := circleArea(10)

	fmt.Println(a1, a2)
}
