package main

import (
	"fmt"
	"strings"
)

func getInit(n string) (string, string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var inti []string
	for _, v := range names {
		inti = append(inti, v[:1])
	}
	if len(inti) > 1 {
		return inti[0], inti[1], inti[2]
	}
	return inti[0], "_", "_"
}

func main() {
	fn1, sn1, ls1 := getInit("Pongdilok Ketngam Test")
	fmt.Println(fn1, sn1, ls1)

	fn2, sn2, ls2 := getInit("Terng Pongdilok sdsd")
	fmt.Println(fn2, sn2, ls2)

	fn3, sn3, ls3 := getInit("PK")
	fmt.Println(fn3, sn3, ls3)

}
