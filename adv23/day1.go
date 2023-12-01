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
	downloader.CheckFile("1", "2023")
	parseDay1()
}

func findNumbers(input string) int {
	r := regexp.MustCompile(`[0-9]`)
	matches := r.FindAllString(input, -1)
	firstLast := matches[0] + matches[len(matches)-1]
	// concatStr := strings.Join(matches, "")
	// fmt.Println(firstLast)
	num, err := strconv.Atoi(firstLast)
	if err != nil {
		fmt.Println(err)
	}

	return num
}

func parseDay1() {
	file, err := os.Open("../input/1_2023.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	var intArr []int

	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}

	file.Close()

	for _, line := range textLines {
		intArr = append(intArr, findNumbers(line))
	}
	// fmt.Print(intArr)

	resSum := 0
	for _, int := range intArr {
		resSum += int
		// fmt.Println(resSum)
	}

	fmt.Print(resSum)
}
