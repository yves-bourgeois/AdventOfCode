package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 10...")

	list := initList()

	instructions := parseData()

	skipSize := 0
	curPosition := 0
	for i := 0; i < 64; i++ {
		for _, instruction := range instructions {
			curPosition, skipSize = knot(list, curPosition, instruction, skipSize)
		}
	}
	//fmt.Println(list[0] * list[1])

	list2 := make([]byte, 0)
	for i := 0; i < 16; i++ {
		a := (list[(i*16)+0] ^ list[(i*16)+1] ^ list[(i*16)+2] ^ list[(i*16)+3] ^ list[(i*16)+4] ^ list[(i*16)+5] ^ list[(i*16)+6] ^ list[(i*16)+7] ^ list[(i*16)+8] ^ list[(i*16)+9] ^ list[(i*16)+10] ^ list[(i*16)+11] ^ list[(i*16)+12] ^ list[(i*16)+13] ^ list[(i*16)+14] ^ list[(i*16)+15])
		list2 = append(list2, byte(a))
	}

	fmt.Println(list2)
	fmt.Println(hex.EncodeToString(list2))
}

func knot(list []int, currPosition int, length int, skipSize int) (int, int) {
	fmt.Printf("currPosition: %v, length: %v, skipSize: %v\n", currPosition, length, skipSize)

	if currPosition+length < len(list) {
		subList := list[currPosition : currPosition+length]

		subList = reverse(subList)

		currPosition = (currPosition + length + skipSize) % len(list)
	} else {
		toEnd := len(list) - currPosition

		subList := append(list[currPosition:len(list)], list[0:length-toEnd]...)
		subList = reverse(subList)

		for i := currPosition; i < len(list); i++ {
			list[i] = subList[i-currPosition]
		}

		for i := 0; i < length-toEnd; i++ {
			list[i] = subList[i+toEnd]
		}

		currPosition = (length - toEnd + skipSize) % len(list)

	}

	skipSize++

	return currPosition, skipSize
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func initList() []int {
	var list []int
	list = make([]int, 0)

	for i := 0; i < 256; i++ {
		list = append(list, i)
	}
	return list
}

func parseData() []int {
	var instructions []int

	instructions = make([]int, 0)

	fileHandle, _ := os.Open("testdata")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {

		for _, el := range fileScanner.Text() {
			instructions = append(instructions, int(el))
		}

	}

	instructions = append(instructions, 17)
	instructions = append(instructions, 31)
	instructions = append(instructions, 73)
	instructions = append(instructions, 47)
	instructions = append(instructions, 23)

	fmt.Println(instructions)
	return instructions
}
