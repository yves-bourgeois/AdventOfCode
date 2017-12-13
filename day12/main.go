package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



var connections map[int][]int

func main() {
	connections = make(map[int][]int)
	parseData()

	traversePoints()

	cnt := 0
	fmt.Println(points)
	for _, pnt := range points {
		if pnt.nextPoint == nil {
			cnt++
		}
	}

	fmt.Println(len(connections))
	fmt.Println(cnt)
}

func traversePoints() {
	var uniquePaths [][]int
	var pathIndex map[int]int

	pathIndex = make(map[int]int)

	for program
}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		splitOnArrow := strings.Split(fileScanner.Text(), " <-> ")
		program, _ := strconv.Atoi(splitOnArrow[0])
		conns := strings.Split(splitOnArrow[1], ", ")
		connectedPrograms := make([]int, 0)
		for _, con := range conns {
			c, _ := strconv.Atoi(con)
			connectedPrograms = append(connectedPrograms, c)
		}
		connections[program] = connectedPrograms
	}
}
