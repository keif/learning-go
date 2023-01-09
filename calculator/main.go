package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	errorStr := "Entry must be a number"

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your first number: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic(errorStr)
	}

	fmt.Print("Enter your second number: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic(errorStr)
	}

	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n", float1, float2, sum)

}
