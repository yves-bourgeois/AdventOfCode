package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var maze [][]string

type puppet struct {
	x, y        int
	orientation string //n, s, w, e
	letters     bytes.Buffer
	steps       int
}

func (p *puppet) move() bool {
	p.steps++

	switch p.orientation {
	case "n":
		p.x--
	case "e":
		p.y++
	case "s":
		p.x++
	case "w":
		p.y--
	}

	if maze[p.x][p.y] != "|" && maze[p.x][p.y] != "-" && maze[p.x][p.y] != "+" {
		p.letters.WriteString(maze[p.x][p.y])
	}

	if maze[p.x][p.y] == " " {
		return true
	}

	if maze[p.x][p.y] == "+" {
		if p.orientation == "n" || p.orientation == "s" {
			if maze[p.x][p.y+1] != " " {
				p.orientation = "e"
			} else {
				p.orientation = "w"
			}
		} else {
			if maze[p.x-1][p.y] != " " {
				p.orientation = "n"
			} else {
				p.orientation = "s"
			}
		}
	}

	return false

}

func main() {
	parseData()

	puppet := puppet{x: 0, y: 63, orientation: "s"}
	//puppet := puppet{x: 0, y: 5, orientation: "s"}

	fmt.Println(maze[puppet.x][puppet.y])

	for {

		if puppet.move() == true {
			break
		}
	}

	fmt.Println(puppet)
	fmt.Println(puppet.letters.String())
}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	maze = make([][]string, 201)

	j := 0
	for fileScanner.Scan() {

		line := make([]string, 0)
		for _, char := range fileScanner.Text() {
			line = append(line, string(char))
		}
		maze[j] = make([]string, len(line))
		maze[j] = line
		j++
	}
}
