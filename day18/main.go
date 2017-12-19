package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type register struct {
	value int
}

type instruction struct {
	instr string
	arg1  string
	arg2  string
}

var registers map[string]*register
var instructions []instruction
var sound int
var currPosition int

func execute() {
	soundFound := -1
	for {
		i := instructions[currPosition]
		fmt.Println(i)
		switch i.instr {
		case "snd":
			{
				sound = registers[i.arg1].value
			}
		case "set":
			{
				iarg2, err := strconv.Atoi(i.arg2)
				if err != nil {
					iarg2 = registers[i.arg2].value
				}
				registers[i.arg1].value = iarg2
			}
		case "add":
			{
				iarg2, err := strconv.Atoi(i.arg2)
				if err != nil {
					iarg2 = registers[i.arg2].value
				}
				registers[i.arg1].value += iarg2
			}
		case "mul":
			iarg2, err := strconv.Atoi(i.arg2)
			if err != nil {
				iarg2 = registers[i.arg2].value
			}
			registers[i.arg1].value *= iarg2
		case "mod":
			iarg2, err := strconv.Atoi(i.arg2)
			if err != nil {
				iarg2 = registers[i.arg2].value
			}
			registers[i.arg1].value = registers[i.arg1].value % iarg2
		case "rcv":
			if registers[i.arg1].value != 0 {
				registers[i.arg1].value = sound
				fmt.Printf("Lastly played: %v ]\n", sound)
				soundFound = sound
			}
		case "jgz":
			if registers[i.arg1].value > 0 {
				iarg2, err := strconv.Atoi(i.arg2)
				if err != nil {
					iarg2 = registers[i.arg2].value
				}
				currPosition += iarg2
			} else {
				currPosition++
			}
		}

		if i.instr != "jgz" {
			currPosition++
		}

		//Quit loop if it jumps outside the instruction set
		if currPosition < 0 || currPosition >= len(instructions) || soundFound != -1 {
			break
		}
	}
}

func main() {
	sound = 0
	currPosition = 0

	parseData()

	execute()

	printRegisters()
}

func printRegisters() {
	for reg, val := range registers {
		fmt.Printf("reg %v, val %v\n", reg, val.value)
	}
}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	registers = make(map[string]*register)
	instructions = make([]instruction, 0)

	for fileScanner.Scan() {
		splitOnSpace := strings.Split(fileScanner.Text(), " ")
		if len(splitOnSpace) == 2 {
			instruction := instruction{instr: splitOnSpace[0], arg1: splitOnSpace[1]}
			instructions = append(instructions, instruction)
		} else {
			instruction := instruction{instr: splitOnSpace[0], arg1: splitOnSpace[1], arg2: splitOnSpace[2]}
			instructions = append(instructions, instruction)
		}
		registers[splitOnSpace[1]] = &register{value: 0}
	}

}
