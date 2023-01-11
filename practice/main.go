package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Evaluate expressions with switch statements")

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "It's Sunday"
		// fallthrough
	case 2:
		result = "It's Monday"
		// fallthrough
	case 3:
		result = "It's Tuesday"
		// fallthrough
	case 4:
		result = "It's Wednesday"
		// fallthrough
	case 5:
		result = "It's Thursday"
		// fallthrough
	case 6:
		result = "It's Friday"
		// fallthrough
	case 7:
		result = "It's Saturday"
		// fallthrough
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)

}
