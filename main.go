package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new build name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name : ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
	switch opt {
	case "a":
		name, _ := getInput("Enter name as ", reader)
		for {
			price, _ := getInput("Enter price as ", reader)
			fmt.Println(name, price)
			fmt.Println("You choose a")
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("Price must be a numberic!")
			}
			b.addItem(name, p)
			if err == nil {
				break
			}
		}

	case "s":
		fmt.Println("You choose s")
	case "t":
		fmt.Println("You choose t")
		tip, _ := getInput("Enter tips amount ($): ", reader)
		fmt.Println(tip)

	default:
		fmt.Println("Not valie option")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

	// fmt.Println(mybill)
	// mybill := newBill("Terng")
	// mybill.addItem("Onion soup", 4.50)
	// mybill.addItem("Chinken fried", 3.33)
	// mybill.addItem("Pad thai", 3.99)
	// mybill.updTip(10)

	// fmt.Println(mybill.format())

	// fmt.Println(mybill)
	// oriPath := "*.exe"
	// newPath := "/exe/*.exe"
	// err := os.Rename(oriPath, newPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
