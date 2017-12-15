package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var connections [][]int

func main() {
	parseData()

	connectionsJoined := true
	fmt.Println(connections)
	for connectionsJoined == true {

		connectionsJoined = false

		for i := 0; i < len(connections)-1; i++ {
			for j := 0; j < len(connections); j++ {
				if i != j {
					if elInCommon(connections[i], connections[j]) {
						//add them together
						connections[i] = append(connections[i], connections[j]...)
						//fmt.Println(connections[i])
						//fmt.Println(connections[j])
						connections = append(connections[0:j], connections[j+1:]...)
						//fmt.Println(connections)
						connectionsJoined = true
						break
					}
				}
			}
			if connectionsJoined == true {
				break
			}
		}

	}

	fmt.Println(len(connections))
}

func elInCommon(slice1 []int, slice2 []int) bool {
	found := false

	for _, el1 := range slice1 {
		for _, el2 := range slice2 {

			if el1 == el2 {
				found = true
				break
			}
		}
		if found == true {
			break
		}
	}

	return found
}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		splitOnArrow := strings.Split(fileScanner.Text(), " <-> ")
		program, _ := strconv.Atoi(splitOnArrow[0])
		conns := strings.Split(splitOnArrow[1], ", ")

		connectedPrograms := make([]int, 0)
		connectedPrograms = append(connectedPrograms, program)
		for _, con := range conns {
			c, _ := strconv.Atoi(con)
			connectedPrograms = append(connectedPrograms, c)
		}
		connections = append(connections, connectedPrograms)
	}
}
