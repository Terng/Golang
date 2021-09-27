package main

import (
	"fmt"
)

type bill struct {
	name   string
	iteams map[string]float64
	tip    float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:   name,
		iteams: map[string]float64{},
		tip:    0,
	}
	return b
}

// format bill
func (b *bill) format() string {
	fs := "Biill Breakdown: \n"
	var total float64 = 0

	for k, v := range b.iteams {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// addtip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

func (b *bill) updTip(tip float64) {
	b.tip = tip
}
func (b *bill) addItem(name string, price float64) {
	b.iteams[name] = price
}
