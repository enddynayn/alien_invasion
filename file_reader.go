package main

import (
	"bufio"
	"os"
)

func FileReader(file_path string) []string {
	file, error := os.Open(file_path)
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
	return lines
}
