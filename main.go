package main

import (
	"bufio"
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

// https://kgrz.io/reading-files-in-go-an-overview.html
func readWithBufio() {
	file, error := os.Open("myFile.txt")
	if error != nil {
		panic(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("read lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
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
	readWithBufio()
}
