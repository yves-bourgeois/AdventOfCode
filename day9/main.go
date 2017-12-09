package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 9 2017...")

	input := readFile()

	parseInput(input)
}

func parseInput(input string) {
	totalLevel := 0
	currLevel := 0
	garbageCount := 0
	garbage := false
	skipNext := false

	for _, char := range input {
		currCharacter := string(char)

		fmt.Printf("curCharacter %v: garbage == %v ; skipNext == %v ; currLevel == %v ; totalLevel == %v\n", currCharacter, garbage, skipNext, currLevel, totalLevel)

		switch {
		case currCharacter == "!" && skipNext == false:
			{
				skipNext = true
				continue
			}
		case skipNext == true:
			{
				//ignore character
				skipNext = false
				continue
			}
		case currCharacter == "<" && garbage == false:
			{ //garbage starts
				garbage = true
				continue
			}

		case currCharacter == ">":
			{
				//garbage is closed now
				garbage = false
				continue
			}
		case garbage == false && currCharacter == "{":
			{
				currLevel++
				continue
			}
		case garbage == false && currCharacter == "}":
			{
				totalLevel = totalLevel + currLevel
				currLevel--
				continue
			}
		case garbage == true:
			{
				garbageCount++
			}
		}

	}

	fmt.Printf("total level %v, total number of garbage %v\n", totalLevel, garbageCount)

}

func readFile() string {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		return fileScanner.Text()
	}

	return ""
}
