package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid [][]string
var guards []guard

type guard struct {
	id           int
	position     int
	limit        int
	prevPosition int
}

type program struct {
	position int
	damage   int
}

func (g *guard) move() {
	if g.prevPosition < g.position {
		//going up
		if g.position+1 <= g.limit {
			g.prevPosition = g.position
			g.position++
		} else {
			//guard needs to return
			g.prevPosition = g.position
			g.position--
		}
	} else {
		//going down
		if g.position-1 > 0 {
			g.prevPosition = g.position
			g.position--
		} else {
			//guard needs to up again
			g.prevPosition = g.position
			g.position++
		}
	}

	g.mark()
}

func (g *guard) mark() {
	if g.prevPosition > 0 {
		grid[g.prevPosition][g.id] = ""
	}
	grid[g.position][g.id] = strconv.Itoa(g.limit)
}

func main() {
	fmt.Println("So it begins again.. day13")

	delay := 1
	parseData()

	caught := false

	for delay = 1; ; delay++ {
		caught = false
		for _, guard := range guards {
			if (guard.id+delay)%((guard.limit*2)-2) == 0 {
				caught = true
				break
			}
		}
		if caught == false {
			break
		}
	}

	fmt.Println(delay)
}

// func iterate(delay int) int {
// 	initGrid()

// 	parseData()

// 	//Init the guards
// 	for i := range guards {
// 		guards[i].move()
// 	}
// 	//printGrid()

// 	program := program{position: -1, damage: 0}

// 	for i := 0; i < 89+delay; i++ {
// 		if i > delay {
// 			if program.position > -1 && grid[1][program.position] != "X" {
// 				grid[1][program.position] = ""
// 			}

// 			program.position++ // = program.position + 1

// 			if grid[1][program.position] != "" {
// 				//Jkl, _ := strconv.Atoi(grid[1][program.position])
// 				//fmt.Printf("limit: %v for i = %v\n", grid[1][program.position], program.position)
// 				//program.damage = program.damage + (program.position * Jkl)
// 				return program.damage + 1
// 			}
// 			//grid[1][program.position] = "B"
// 		}
// 		for j := range guards {
// 			guards[j].move()
// 		}
// 		//printGrid()

// 		//time.Sleep(500 * time.Millisecond)
// 	}

// 	return program.damage

// }

// func initGrid() {
// 	grid = nil
// 	grid = make([][]string, 31)
// 	for i := 0; i < 31; i++ {
// 		grid[i] = make([]string, 89)
// 	}
// }

// func printGrid() {
// 	fmt.Println("------------------------------------")
// 	for i := 29; i > 0; i-- {
// 		for j := 0; j < 89; j++ {
// 			if grid[i][j] == "" {
// 				fmt.Printf(" ")
// 			} else {
// 				fmt.Printf("%v", grid[i][j])
// 			}
// 		}
// 		fmt.Printf("\n")
// 		// = make([]int, 30)
// 	}
// }

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	guards = nil
	guards = make([]guard, 0)

	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), ": ")

		id, _ := strconv.Atoi(split[0])
		limit, _ := strconv.Atoi(split[1])

		guard := guard{id: id, position: 0, limit: limit, prevPosition: -1}

		guards = append(guards, guard)
	}
}
