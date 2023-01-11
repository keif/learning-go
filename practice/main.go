package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Parse JSON-formatted text")

	resp, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	content := string(bytes)
	fmt.Print(content)

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name, " ", tour.Price)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)

		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name, Price string
}
