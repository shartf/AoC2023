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
	downloader.CheckFile("2", "2023")
	input := reader()
	part21(input)
}

type Game struct {
	GameID int
	Blue   int
	Red    int
	Green  int
}

func part21(input []string) {
	var gameSum int
	for _, line := range input {
		resGame := parser(line)
		// fmt.Printf("Game No. %d has %d red, %d green, %d blue cubes\n", resGame.GameID, resGame.Red, resGame.Green, resGame.Blue)
		if resGame.Blue <= 14 && resGame.Red <= 12 && resGame.Green <= 13 {
			fmt.Printf("Game No. %d has %d red, %d green, %d blue cubes\n", resGame.GameID, resGame.Red, resGame.Green, resGame.Blue)
			gameSum += resGame.GameID
		}
	}
	fmt.Print(gameSum)
}

func parser(line string) Game {
	// split into game number and draws
	// gameResults := strings.Split(line, ":")
	// fmt.Println(gameResults)
	reGame := regexp.MustCompile(`Game (\d+)`)
	reblue := regexp.MustCompile(`(\d+) blue`)
	regreen := regexp.MustCompile(`(\d+) green`)
	rered := regexp.MustCompile(`(\d+) red`)

	matchesGame := reGame.FindAllStringSubmatch(line, -1)
	matchesBlue := reblue.FindAllStringSubmatch(line, -1)
	matchesGreen := regreen.FindAllStringSubmatch(line, -1)
	matchesRed := rered.FindAllStringSubmatch(line, -1)

	gameCount := counter(matchesGame)
	blueCount := counter(matchesBlue)
	greenCount := counter(matchesGreen)
	redCount := counter(matchesRed)

	return Game{GameID: gameCount, Blue: blueCount, Red: redCount, Green: greenCount}
}

// return a final number of items in an array
func counter(gamesNum [][]string) int {
	var count int
	for _, game := range gamesNum {
		res, _ := strconv.Atoi(game[1])
		count += res
	}
	return count
}

func reader() []string {
	file, err := os.Open("../input/2_2023.txt")
	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLine []string
	for scanner.Scan() {
		textLine = append(textLine, scanner.Text())
	}
	file.Close()
	return textLine
}
