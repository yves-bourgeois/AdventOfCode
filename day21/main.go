package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var drawing [][]string

var instructions map[string]string

func main() {
	fmt.Println("day 21..")

	//drawing = make([][]string, 0)

	drawing = [][]string{
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	}

	inputArray, solutionArray := parseData()

	for i := 0; i < 18; i++ {
		fmt.Printf("Tick %v \n", i+1)
		drawing = tick(drawing, inputArray, solutionArray)

		printMatrix(drawing)
	}

	cnt := 0
	for i := 0; i < len(drawing[0]); i++ {
		for j := 0; j < len(drawing[0]); j++ {
			if drawing[i][j] == "#" {
				cnt++
			}
		}
	}

	println(cnt)

}

func tick(drawing [][]string, inputArray [][][]string, solutionArray [][][]string) [][]string {
	var cardinality int
	var newDrawing [][]string
	dimension := len(drawing[0])

	if math.Mod(float64(dimension), 2) == 0 {
		cardinality = 2
	} else {
		cardinality = 3
	}
	var part [][]string

	part = make([][]string, cardinality)
	for i := 0; i < cardinality; i++ {
		part[i] = make([]string, cardinality)
	}

	newDimension := (cardinality + 1) * (dimension / cardinality)
	fmt.Println("New dimension: ", newDimension)

	newDrawing = make([][]string, newDimension)
	for i := 0; i < newDimension; i++ {
		newDrawing[i] = make([]string, newDimension)
	}

	for blockX := 0; blockX < dimension/cardinality; blockX++ {
		for blockY := 0; blockY < dimension/cardinality; blockY++ {

			for x := 0; x < cardinality; x++ {
				for y := 0; y < cardinality; y++ {
					part[x][y] = drawing[(blockX*cardinality)+x][(blockY*cardinality)+y]
				}
			}

			solution := findSolution(part, inputArray, solutionArray)
			if solution != nil {
				for x := 0; x < cardinality+1; x++ {
					for y := 0; y < cardinality+1; y++ {
						newDrawing[x+(blockX*(cardinality+1))][y+(blockY*(cardinality+1))] = solution[x][y]
					}
				}
			} else {
				fmt.Println("No solution found")
				printMatrix(part)
				os.Exit(-1)
			}

		}
	}

	return newDrawing
}

func equalArrays(arr1 [][]string, arr2 [][]string) bool {
	if len(arr1[0]) != len(arr2[0]) {
		return false
	}

	for i := 0; i < len(arr1[0]); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}
	return true
}

func findSolution(drawing [][]string, inputArray [][][]string, solutionArray [][][]string) [][]string {
	for i, input := range inputArray {
		if equalArrays(drawing, input) {
			return solutionArray[i]
		}

		flippedDrawing := flip(drawing)

		if equalArrays(flippedDrawing, input) {
			return solutionArray[i]
		}

		drawing = rotate(drawing) //90 degrees
		if equalArrays(drawing, input) {
			return solutionArray[i]
		}

		flippedDrawing = flip(drawing)
		if equalArrays(flippedDrawing, input) {
			return solutionArray[i]
		}

		drawing = rotate(drawing) //180 degrees
		if equalArrays(drawing, input) {
			return solutionArray[i]
		}

		flippedDrawing = flip(drawing)
		if equalArrays(flippedDrawing, input) {
			return solutionArray[i]
		}

		drawing = rotate(drawing) //270 degrees
		if equalArrays(drawing, input) {
			return solutionArray[i]
		}

		flippedDrawing = flip(drawing)
		if equalArrays(flippedDrawing, input) {
			return solutionArray[i]
		}
	}

	return nil

}

func parseData() (iArray [][][]string, oArray [][][]string) {
	var inputArray [][]string
	var outputArray [][]string

	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	instructions = make(map[string]string)

	for fileScanner.Scan() {
		splitOnArrow := strings.Split(fileScanner.Text(), " => ")

		if len(splitOnArrow[0]) == 5 {
			//array with 2 elements
			inputArray = make([][]string, 2)
			outputArray = make([][]string, 3)
		} else {
			inputArray = make([][]string, 3)
			outputArray = make([][]string, 4)
		}

		inputLines := strings.Split(splitOnArrow[0], "/")

		for i, line := range inputLines {
			lineElements := strings.Split(line, "")
			if len(splitOnArrow[0]) == 5 {
				inputArray[i] = make([]string, 2)
			} else {
				inputArray[i] = make([]string, 3)
			}
			for j, el := range lineElements {
				inputArray[i][j] = el
			}
		}

		outputLines := strings.Split(splitOnArrow[1], "/")

		for i, line := range outputLines {
			lineElements := strings.Split(line, "")
			if len(splitOnArrow[1]) == 11 {
				outputArray[i] = make([]string, 3)
			} else {
				outputArray[i] = make([]string, 4)
			}
			for j, el := range lineElements {
				outputArray[i][j] = el
			}
		}

		iArray = append(iArray, inputArray)
		oArray = append(oArray, outputArray)
	}

	return iArray, oArray
}

func rotate(m [][]string) [][]string {
	//Initialize return matrix
	d := len(m[0])
	retM := make([][]string, d)
	for i := 0; i < d; i++ {
		retM[i] = make([]string, d)
	}

	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			retM[i][j] = m[d-j-1][i]
			//retM[i][j] = m[j][i]
			//retM[i][j], retM[j][i] = m[j][i], m[i][j]
		}
	}

	return retM
}

func flip(m [][]string) [][]string {
	//Initialize return matrix
	d := len(m[0])
	retM := make([][]string, d)
	for i := 0; i < d; i++ {
		retM[i] = make([]string, d)
	}

	for i := 0; i < d; i++ {
		if d == 2 {
			retM[i][0] = m[i][1]
			retM[i][1] = m[i][0]
		}
		if d == 3 {
			retM[i][0] = m[i][2]
			retM[i][1] = m[i][1]
			retM[i][2] = m[i][0]
		}

	}

	return retM
}

func printMatrix(m [][]string) {
	for _, arr := range m {
		fmt.Println(arr)
	}
}
