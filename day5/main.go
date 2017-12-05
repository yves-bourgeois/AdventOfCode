package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Here we go again..")

	increment := 0
	index := 0

	input, _ := ioutil.ReadFile("./testdata")

	lines := strings.Split(string(input), "\n")

	//fmt.Printf("File contents: %s", input)

	for increment = 0; index >= 0 && index < len(lines); increment++ {
		newVal := 0

		currVal, _ := strconv.Atoi(lines[index])
		if currVal >= 3 {
			newVal = currVal - 1
		} else {
			newVal = currVal + 1
		}

		//fmt.Printf("at [%v]: currVal = %v;   newVal = %v\n", index, currVal, newVal)

		lines[index] = strconv.Itoa(newVal)

		index = index + currVal
	}

	fmt.Printf("did a break after.. %v", increment)

}
