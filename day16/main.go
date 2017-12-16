package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var programs string

var log map[string]struct{}

func spin(position int) {
	out := []rune(programs)

	temp := make([]rune, 0)

	temp = append(temp, out[len(programs)-position:len(out)]...)
	temp = append(temp, out[0:len(out)-position]...)

	programs = string(temp)

}

func exchange(posA int, posB int) {
	out := []rune(programs)
	temp := out[posA]
	out[posA] = out[posB]
	out[posB] = temp
	programs = string(out)

}

func partner(a string, b string) {
	exchange(strings.Index(programs, a), strings.Index(programs, b))
}

func dance(instruction string) {
	if string(instruction[0]) == "s" {
		steps, _ := strconv.Atoi(string(instruction[1:len(instruction)]))
		spin(steps)
	}

	if string(instruction[0]) == "x" {
		splitOnd := strings.Split(string(instruction[1:len(instruction)]), "/")

		posA, _ := strconv.Atoi(string(splitOnd[0]))
		posB, _ := strconv.Atoi(string(splitOnd[1]))

		exchange(posA, posB)
	}

	if string(instruction[0]) == "p" {
		splitOnd := strings.Split(string(instruction[1:len(instruction)]), "/")

		partner(splitOnd[0], splitOnd[1])
	}
}

func main() {
	programs = "abcdefghijklmnop"

	log = make(map[string]struct{})
	log[programs] = struct{}{}

	instructions := parseData()
	//fmt.Println(instructions)

	fmt.Printf("Every 60 times the loop is the same.  After 1 billion loops, we would be on %v\n", 1000000000%60)

	for i := 0; i < 40; i++ {
		for _, instruction := range instructions {
			dance(instruction)
		}
		if _, ok := log[programs]; ok {
			fmt.Printf("Found in map after %v iterations\n", i)
			fmt.Println(programs)
		}
		log[programs] = struct{}{}
	}
	fmt.Println(programs)
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
