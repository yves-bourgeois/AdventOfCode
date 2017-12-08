package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var nodes map[string]node

type node struct {
	name     string
	prevNode *node
	nextNode *[]node
	weight   int
}

func main() {
	fmt.Println("day 7: And so it begins..")

	//	found := false
	foundElem := ""

	getData()

	// for _, el := range parent {
	// 	found = false
	// 	for _, el2 := range child {
	// 		if el == el2 {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// 	if found == false {
	// 		foundElem = el
	// 		break
	// 	}
	// }
	fmt.Println("jkl;;;")
	fmt.Println(foundElem)
}

func getData() {
	nodes := make(map[string]node)

	data, err := ioutil.ReadFile("./input")

	if err != nil {
		fmt.Println("Error whilst opening file.. " + err.Error())
	}

	csvReader := csv.NewReader(strings.NewReader(string(data)))
	csvReader.Comma = ' '

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			fmt.Println("End of file reached..")
			break
		}

		for i, el := range line {
			var nd node
			if i == 0 {
				nd.name = el
			}
			if i == 1 {
				el = strings.Replace(el, "(", "", -1)
				el = strings.Replace(el, ")", "", -1)
				nd.weight, _ = strconv.Atoi(el)
			}
			nodes[nd.name] = nd
		}
	}

	data, _ = ioutil.ReadFile("input")
	//memoryString := string(fileContent)

	stream := bufio.NewReader(strings.NewReader(string(data)))

	for {
		line, _, err := stream.ReadLine()
		if err != nil {
			break
		}

		strLine := string(line)

		splitLine := strings.Split(line, "->")

	}

}
