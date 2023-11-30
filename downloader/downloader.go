package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func CheckFile(day string, year string) {
	checkForInput(day, year)
}

func readKey() (key string) {
	content, err := os.ReadFile("../secret/key.txt")
	if err != nil {
		log.Fatal(err)
	}
	removeNL := strings.ReplaceAll(string(content), "\n", "")
	return removeNL
}

func readInput(day string, year string, fileName string) {
	// create a request
	key := readKey()
	url := "https://adventofcode.com"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/day/%s/input", url, year, day), nil)
	if err != nil {
		log.Fatal("Error, reading request", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", key))

	// make a request
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading request", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body", err)
	}

	// Write to a file all at once
	if err := os.WriteFile(fileName, body, 0666); err != nil {
		log.Fatal("File could not be written", err)
	}
	log.Printf("File written: %s", fileName)
}

// Checks if an input file already exists
func checkForInput(day string, year string) {
	fileName := fmt.Sprintf("../input/%s_%s.txt", day, year)
	isFileExist := checkFileExists(fileName)

	if isFileExist {
		log.Printf("File %s exists, skip download", fileName)
	} else {
		// initiate download
		fmt.Println("File does not exist")
		readInput(day, year, fileName)
	}
}

func checkFileExists(filePath string) bool {
	_, err := os.Open(filePath) // For read access.
	return err == nil
}
