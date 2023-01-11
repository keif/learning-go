package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var result float64

	value1 := getInput("first")
	value2 := getInput("second")
	operation := getOperation()

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}
	fmt.Printf("The result of of %v %v %v is %v\n", value1, operation, value2, result)

}

func getInput(prompt string) float64 {
	fmt.Printf("Enter your %v number: ", prompt)
	input, _ := reader.ReadString('\n')
	float, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Entry must be a number")
	}

	return float
}

func getOperation() string {
	fmt.Print("Select an operation (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValues(values ...float64) float64 {
	var total float64 = 0

	for _, v := range values {
		total += v
	}
	return total
}

func subtractValues(values ...float64) float64 {
	var total float64 = values[0]

	for i, v := range values {
		if i != 0 {
			total -= v
		}
	}
	return total
}

func multiplyValues(values ...float64) float64 {
	var total float64 = values[0]

	for i, v := range values {
		if i != 0 {
			total *= v
		}
	}
	return total
}

func divideValues(values ...float64) float64 {
	var total float64 = values[0]

	for i, v := range values {
		if i != 0 {
			total /= v
		}
	}
	return total
}
