package main

import (
	"bufio"
	"downloader"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	downloader.CheckFile("1", "2023")
	parseDay1()
}

func findNumbersAndText(input string) int {
	r := regexp.MustCompile(`(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)`)
	matches := r.FindAllString(input, -1)
	// firstLast := matches[0] + matches[len(matches)-1]

	var parsedNum []string
	for _, parsed := range matches {
		parsedNum = append(parsedNum, wordToNumber(parsed))
	}

	// fmt.Println(parsedNum)
	firstLast := parsedNum[0] + parsedNum[len(parsedNum)-1]
	num, err := strconv.Atoi(firstLast)
	if err != nil {
		fmt.Println(err)
	}

	return num
}

func findNumbers(input string) int {
	r := regexp.MustCompile(`[0-9]`)
	matches := r.FindAllString(input, -1)

	// convert text to int

	firstLast := matches[0] + matches[len(matches)-1]
	// concatStr := strings.Join(matches, "")
	// fmt.Println(firstLast)
	num, err := strconv.Atoi(firstLast)
	if err != nil {
		fmt.Println(err)
	}

	return num
}

func replacemalignedWods(input string) string {
	input = strings.ReplaceAll(input, "oneight", "one8")
	input = strings.ReplaceAll(input, "twone", "two1")
	input = strings.ReplaceAll(input, "threight", "three8")
	input = strings.ReplaceAll(input, "fiveight", "five8")
	input = strings.ReplaceAll(input, "sevenine", "seven9")
	input = strings.ReplaceAll(input, "eightree", "eight3")
	input = strings.ReplaceAll(input, "eightwo", "eight2")
	input = strings.ReplaceAll(input, "nineight", "nine8")
	return input
}

func wordToNumber(word string) string {
	switch strings.ToLower(word) {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "1":
		return "1"
	case "2":
		return "2"
	case "3":
		return "3"
	case "4":
		return "4"
	case "5":
		return "5"
	case "6":
		return "6"
	case "7":
		return "7"
	case "8":
		return "8"
	default:
		return "9"
	}
}

func parseDay1() {
	file, err := os.Open("../input/1_2023.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string

	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}

	file.Close()

	// part1
	part1(textLines)
	// part2
	part2(textLines)
}

func part1(textLines []string) {
	var intArr []int
	for _, line := range textLines {
		intArr = append(intArr, findNumbers(line))
	}
	// fmt.Print(intArr)

	resSum := 0
	for _, int := range intArr {
		resSum += int
		// fmt.Println(resSum)
	}

	// fmt.Print(resSum)
}

func part2(textLines []string) {
	var intArr []int
	for _, line := range textLines {
		replacedLine := replacemalignedWods(line)
		intArr = append(intArr, findNumbersAndText(replacedLine))
	}
	// fmt.Print(intArr)

	resSum := 0
	for _, int := range intArr {
		resSum += int
		// fmt.Println(resSum)
	}

	fmt.Print(resSum)
}
