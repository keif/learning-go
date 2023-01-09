package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[0])
	fmt.Println(colors[1])
	fmt.Println(colors[2])

	var numbers = [5]int{5, 4, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))

}
