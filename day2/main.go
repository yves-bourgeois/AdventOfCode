package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Beginning to read file..")

	data, err := ioutil.ReadFile("./testdata")

	if err != nil {
		fmt.Println("Error whilst opening file.. " + err.Error())
	}

	csvReader := csv.NewReader(strings.NewReader(string(data)))
	csvReader.Comma = '\t'

	checkSum := 0

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			fmt.Println("End of file reached..")
			break
		}

		//Process line
		res := 0

		for _, el := range line {
			elInt, _ := strconv.Atoi(el)

			for _, el2 := range line {
				el2Int, _ := strconv.Atoi(el2)

				if elInt%el2Int == 0 && elInt != el2Int && el2Int != 0 {
					res = elInt / el2Int
					break
				}
			}

			if res != 0 {
				break
			}
		}
		checkSum = checkSum + res
	}

	fmt.Printf("Result: %d", checkSum)
}
