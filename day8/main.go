package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var instructions []instruction
var registers map[string]*register

type register struct {
	register string
	value    int
}

type instruction struct {
	target string
	action
	condition
}

type action struct {
	action string
	amount int
}

type condition struct {
	register  string
	condition string
	value     int
}

func (c *condition) test() bool {
	switch c.condition {
	case "<":
		if registers[c.register].value < c.value {
			return true
		} else {
			return false
		}
	case "<=":
		if registers[c.register].value <= c.value {
			return true
		} else {
			return false
		}
	case ">=":
		if registers[c.register].value >= c.value {
			return true
		} else {
			return false
		}
	case ">":
		if registers[c.register].value > c.value {
			return true
		} else {
			return false
		}
	case "==":
		if registers[c.register].value == c.value {
			return true
		} else {
			return false
		}
	case "!=":
		if registers[c.register].value != c.value {
			return true
		} else {
			return false
		}
	default:
		fmt.Println("Condition not found: " + c.condition)
	}

	return false
}

func (a *action) getNewValue(register string) int {
	switch a.action {
	case "inc":
		return registers[register].value + a.amount
	case "dec":
		return registers[register].value - a.amount
	default:
		fmt.Println("action not found: " + a.action)
	}
	return registers[register].value
}

func main() {
	registers = make(map[string]*register)
	//instructions = make([]instruction)

	parseData()

	max := -99999

	for _, instruction := range instructions {
		if instruction.test() == true {
			registers[instruction.target].value = instruction.getNewValue(instruction.target)

			for _, register := range registers {
				if max < register.value {
					max = register.value
				}
			}
		}

	}

	fmt.Println(max)
}

func parseData() {
	fileHandle, _ := os.Open("testdata.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		splitOnIf := strings.Split(fileScanner.Text(), "if")
		firstPart := strings.Split(splitOnIf[0], " ")
		secondPart := strings.Split(strings.Trim(splitOnIf[1], " "), " ")

		amount, _ := strconv.Atoi(firstPart[2])
		val, _ := strconv.Atoi(secondPart[2])

		instruction := instruction{target: firstPart[0],
			action: action{action: firstPart[1], amount: amount}, condition: condition{register: strings.Trim(secondPart[0], " "), condition: secondPart[1], value: val}}

		instructions = append(instructions, instruction)

		registers[firstPart[0]] = &register{register: firstPart[0], value: 0}

	}
}
