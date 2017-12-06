package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Day 4 ay...")

	fileContent, _ := ioutil.ReadFile("./testData")

	fileCSV := csv.NewReader(strings.NewReader(string(fileContent)))

	fileCSV.Comma = ' '

	cntUniquePass := 0

	for {
		line, err := fileCSV.Read()

		if err == io.EOF {
			fmt.Println("End of file reached..")
			break
		}

			passMap := make(map[string]string)

		unique := true
		for _, record := range line {
			sortedRecord := sortString(record)
			if _, ok := passMap[sortedRecord]; ok {
				fmt.Println("Not unique..")
				unique = false
				break
			} else {
				passMap[sortedRecord] = record
				//do nothing
			}
		}

		if unique == true {
			cntUniquePass++
		}

	}

	fmt.Println(cntUniquePass)
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")

}
