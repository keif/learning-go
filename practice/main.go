package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Files")

	content := "Hello from Go!"
	fileName := "./fromString.txt"

	file, err := os.Create(fileName)
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close() // always close files
	defer readFile(fileName)
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
