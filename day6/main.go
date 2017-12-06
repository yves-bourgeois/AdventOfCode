package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	memoryKey := ""
	states := make(map[string]struct{})

	fmt.Println("Day 6 ay..")
	memoryBanks := getData()

	for i := 0; true; i++ {
		memoryBanks := distribute(memoryBanks)

		memoryKey = arrayToString(memoryBanks)

		if _, ok := states[memoryKey]; ok {
			fmt.Printf("Memory banks distributed after %v\n", i+1)
			break
		}

		states[memoryKey] = struct{}{}

		//time.Sleep(1 * time.Second)
	}

	for i := 0; true; i++ {
		memoryBanks := distribute(memoryBanks)

		key := arrayToString(memoryBanks)

		if key == memoryKey {
			fmt.Printf("Memory bank key found again after %v\n", i+1)
			break
		}

		states[memoryKey] = struct{}{}

		//time.Sleep(1 * time.Second)
	}
}

func getData() []int {
	var memoryBanks []int
	fileContent, _ := ioutil.ReadFile("testdata")
	memoryString := string(fileContent)

	for _, mem := range strings.Fields(memoryString) {
		memInt, _ := strconv.Atoi(mem)

		memoryBanks = append(memoryBanks, memInt)
	}

	return memoryBanks
}
func distribute(memoryBanks []int) []int {
	fmt.Printf("Ditributing %v", memoryBanks)

	var index int

	distValue := -9999

	//get maximum
	for i, val := range memoryBanks {
		if val > distValue {
			index = i
			distValue = val
		}
	}

	fmt.Printf("Value found: %v, position: %v\n", distValue, index)

	//set current block to zero
	memoryBanks[index] = 0

	for i := distValue; i > 0; i-- {
		index = (index + 1) % len(memoryBanks)
		memoryBanks[index] = memoryBanks[index] + 1
	}

	return memoryBanks
}

func arrayToString(memoryBanks []int) string {
	var buff bytes.Buffer

	for _, val := range memoryBanks {
		buff.WriteString(strconv.Itoa(val))
		buff.WriteString("-")
	}

	return buff.String()
}
