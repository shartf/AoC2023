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

func day2() {
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
		resGame := parserPart1(line)
		// fmt.Printf("Game No. %d has %d red, %d green, %d blue cubes\n", resGame.GameID, resGame.Red, resGame.Green, resGame.Blue)
		// if resGame.Blue <= 14 && resGame.Red <= 12 && resGame.Green <= 13 {
		// 	fmt.Printf("Game No. %d has %d red, %d green, %d blue cubes\n", resGame.GameID, resGame.Red, resGame.Green, resGame.Blue)
		// 	gameSum += resGame.GameID
		gameSum += resGame
	}
	fmt.Printf("Answer to part 1 is: %d \n", gameSum)

	// part22
	var towerOfPower int
	for _, line := range input {
		resPart2 := parserPart2(line)
		towerOfPower += resPart2
	}

	fmt.Printf("Answer to part 2 is: %d \n", towerOfPower)
}

func parserPart1(line string) int {
	reGame := regexp.MustCompile(`Game (\d+)`)
	reblue := regexp.MustCompile(`(\d+) blue`)
	regreen := regexp.MustCompile(`(\d+) green`)
	rered := regexp.MustCompile(`(\d+) red`)

	// split in game number and games
	gameNoAndGames := strings.Split(line, ":")
	matchesGame := reGame.FindAllStringSubmatch(gameNoAndGames[0], -1)
	gameCount := counter(matchesGame)

	// if any of the values is bigger than allowed, value is switched to false
	validGame := true
	// check single games
	singleGames := strings.Split(gameNoAndGames[1], ";")
	for _, game := range singleGames {

		matchesBlue := reblue.FindAllStringSubmatch(game, -1)
		matchesGreen := regreen.FindAllStringSubmatch(game, -1)
		matchesRed := rered.FindAllStringSubmatch(game, -1)

		blueCount := counter(matchesBlue)
		greenCount := counter(matchesGreen)
		redCount := counter(matchesRed)

		// check for conditions
		// the variable is set to true initially, but even if one game is not valid, the whole set is invalid, so we look for one game only
		if redCount > 12 || greenCount > 13 || blueCount > 14 {
			validGame = false
		}
	}

	if !validGame {
		return 0
	} else {
		return gameCount
	}
}

func parserPart2(line string) int {
	reblue := regexp.MustCompile(`(\d+) blue`)
	regreen := regexp.MustCompile(`(\d+) green`)
	rered := regexp.MustCompile(`(\d+) red`)

	// split in game number and games
	gameNoAndGames := strings.Split(line, ":")

	blueMax := 0
	greenMax := 0
	redMax := 0

	singleGames := strings.Split(gameNoAndGames[1], ";")
	for _, game := range singleGames {

		matchesBlue := reblue.FindAllStringSubmatch(game, -1)
		matchesGreen := regreen.FindAllStringSubmatch(game, -1)
		matchesRed := rered.FindAllStringSubmatch(game, -1)

		blueCount := counter(matchesBlue)
		greenCount := counter(matchesGreen)
		redCount := counter(matchesRed)

		if blueCount > blueMax {
			blueMax = blueCount
		}
		if greenCount > greenMax {
			greenMax = greenCount
		}
		if redCount > redMax {
			redMax = redCount
		}
	}

	return (blueMax * greenMax * redMax)
}

// Those, who can read, are in a clear advantage.
// Sadly, it is not me, so I leave it as a sad statement of me not being able to understand the requirements.
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
