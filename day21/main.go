package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var drawing [][]string

var instructions map[string]string

func main() {
	fmt.Println("day 21..")

	//drawing = make([][]string, 0)

	drawing = [][]string{
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	}

}

func iterate(size int) {
	

	for i := 0; i < len(drawing[0])/size; i++ {

	}

}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	instructions = make(map[string]string)

	for fileScanner.Scan() {
		splitOnArrow := strings.Split(fileScanner.Text(), " => ")

		instructions[sortString(splitOnArrow[0])] = splitOnArrow[1]
	}
}

func sortString(toSort string) string {
	s := strings.Split(toSort, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
