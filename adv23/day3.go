package main

import (
	"bufio"
	"downloader"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	downloader.CheckFile("3", "2023")
	input, lineCount := readerD3()
	day3Part1(input, lineCount)
}

func day3Part1(input []string, lineCount int) {
	re := regexp.MustCompile(`\b\d+\b`)
	reNeg := regexp.MustCompile(`[^.\d]`)

	var sum int
	for idx, line := range input {
		// get indexes of al the matches in the line
		matches := re.FindAllStringSubmatchIndex(line, -1)
		// loop over indexes
		for _, match := range matches {
			// set boundaries
			var inner, outer int
			if match[0] > 0 {
				inner = match[0] - 1
			}
			if match[1] < len(line)-1 {
				outer = match[1] + 1
			} else {
				outer = len(line) - 1
			}

			if idx == 0 {
				// first line, only look below
				if reNeg.MatchString(line[inner:outer]) || reNeg.MatchString(input[idx+1][inner:outer]) {
					// fmt.Printf("got it at %d, %d", inner, outer)
					num, _ := strconv.Atoi(line[match[0]:match[1]])
					sum += num
				}
			} else if idx == (lineCount - 1) {
				// last line, look only line up
				if reNeg.MatchString(line[inner:outer]) || reNeg.MatchString(input[idx-1][inner:outer]) {
					// fmt.Printf("got it at %d, %d", inner, outer)
					num, _ := strconv.Atoi(line[match[0]:match[1]])
					sum += num
				}
			} else {
				// middle
				if reNeg.MatchString(line[inner:outer]) || reNeg.MatchString(input[idx-1][inner:outer]) || reNeg.MatchString(input[idx+1][inner:outer]) {
					// fmt.Printf("got it at %d, %d", inner, outer)
					num, _ := strconv.Atoi(line[match[0]:match[1]])
					sum += num
				}
			}
		}
	}
	fmt.Print(sum)
}

func readerD3() ([]string, int) {
	file, err := os.Open("../input/3_2023.txt")
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
