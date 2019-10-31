package services

import (
	"bufio"
	"os"
)

// FileReader reads lines from a file.
// It returns a splice of lines.
func FileReader(file_path string) []string {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
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
