package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type register struct {
	value int
}

type instruction struct {
	instr string
	arg1  string
	arg2  string
}

var instructions []instruction
var sound int

var snd int
var sent int

func execute(id string, registers map[string]*register, sndChan chan int, rcvChan chan int) {
	currPosition := 0
	//soundFound := -1
	for {
		i := instructions[currPosition]
		switch i.instr {
		case "snd":
			{

				iarg1, err := strconv.Atoi(i.arg1)
				if err != nil {
					iarg1 = registers[i.arg1].value
				}
				sent++
				snd++
				select {
				case sndChan <- iarg1:
					fmt.Println("sent message")
				default:
					fmt.Println("no message sent")
				}

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
				registers[i.arg1].value = <-rcvChan
				fmt.Println("Received a value")
			}
		case "jgz":
			iarg1, err := strconv.Atoi(i.arg1)
			if err != nil {
				iarg1 = registers[i.arg1].value
			}

			if iarg1 > 0 {
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
		if currPosition < 0 || currPosition >= len(instructions) {
			fmt.Printf("%v: Breaking.... %v.. currPos: %v\n", id, sent, currPosition)
			break
		}
	}
}

func main() {
	p1Chan := make(chan int, 10000000)
	p2Chan := make(chan int, 10000000)

	var registersP1 map[string]*register
	var registersP2 map[string]*register

	sound = 0

	registersP1 = parseData()
	registersP2 = parseData()

	registersP2["p"].value = 1

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		execute("p1", registersP1, p1Chan, p2Chan)
	}()

	func() {
		defer wg.Done()
		execute("p2", registersP2, p2Chan, p1Chan)
	}()

	wg.Wait()

	fmt.Println("done..")

}

func printRegisters(reg map[string]*register) {
	for reg, val := range reg {
		fmt.Printf("reg %v, val %v\n", reg, val.value)
	}
}

func parseData() map[string]*register {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	registers := make(map[string]*register)
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

	return registers

}
