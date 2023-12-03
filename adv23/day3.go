package main

import (
	"bufio"
	"downloader"
	"fmt"
	"log"
	"os"
)

func main() {
	downloader.CheckFile("3", "2023")
	input, lineCount := readerD3()
	day3Part1(input, lineCount)
	fmt.Println(input)
	fmt.Println(lineCount)
}

func day3Part1(input []string, lineCount int) {
	for idx := range input {
		if idx, line == 0 {
			// first line, look only line below
		} else if idx == (lineCount - 1) {
			// last line, look only line up
		} else {
			// middle
		}
	}
}

func readerD3() ([]string, int) {
	file, err := os.Open("../input/3_2023_test.txt")
	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLine []string
	lineCount := 0
	for scanner.Scan() {
		textLine = append(textLine, scanner.Text())
		lineCount++
	}
	file.Close()
	return textLine, lineCount
}
