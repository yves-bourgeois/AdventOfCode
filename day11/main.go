package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var circle int
var dist int

type hexagon struct {
	x, y int
}

func (h *hexagon) move(direction string) {
	switch direction {
	case "n":
		{
			h.y = h.y + 1
		}
	case "ne":
		{
			h.x = h.x + 1
			h.y = h.y + 1
		}
	case "se":
		{
			h.x = h.x + 1
		}
	case "s":
		{
			h.y = h.y - 1
		}
	case "sw":
		{
			h.x = h.x - 1
			h.y = h.y - 1
		}
	case "nw":
		{
			h.x = h.x - 1
		}

	}
}

func main() {
	fmt.Println("Day 11 then..")

	circle = 0
	dist = -1
	h := hexagon{x: 0, y: 0}
	maxH := hexagon{x: 0, y: 0}
	instructions := parseData()
	steps := 0
	for i, direction := range instructions {
		h.move(direction)
		if dist < distance(h) {
			dist = distance(h)
			steps = i
			maxH = h
		}
	}

	fmt.Printf("x: %v, y: %v\n", h.x, h.y)
	fmt.Println(distance(h))
	fmt.Println(distance(maxH))
	fmt.Println(steps)
	fmt.Println(len(instructions))
}

func distance(h hexagon) int {
	return int(math.Max(math.Abs(float64(h.y)), math.Max(math.Abs(float64(h.x)), math.Abs(float64(h.x-h.y)*-1))))
}
func parseData() []string {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	instructions := make([]string, 0)

	for fileScanner.Scan() {
		splitOnComma := strings.Split(fileScanner.Text(), ",")
		for _, elem := range splitOnComma {
			instructions = append(instructions, elem)
		}
	}

	return instructions
}
