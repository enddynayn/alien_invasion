package main

import (
	"fmt"
	"os"
)

func writeToFile() {
	file, error := os.Create("myFile.txt")

	if error != nil {
		panic(error)
	}

	data := []byte("hello World\n")

	file.Write(data)
	file.WriteString("invasion of world \n")
	file.Close()
}

func readFile() {
	file, error := os.Open("myFile.txt")

	if error != nil {
		panic(error)
	}

	data := make([]byte, 100)

	file.Read(data)
	fmt.Printf("The file data %s\n", string(data))
	file.Close()
}

func main() {
	writeToFile()
	readFile()
}
