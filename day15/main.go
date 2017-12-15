package main

import (
	"fmt"
	"strconv"
)

const divide int64 = 2147483647

type generator struct {
	factor     int64
	value      int64
	binary16   string
	multipleOf int64
}

func (g *generator) produce() {
	for {
		g.value = (g.value * g.factor) % divide
		if g.value%g.multipleOf == 0 {
			break
		}
	}
	tempString := strconv.FormatInt(g.value, 2)
	tempString = "00000000000000" + tempString

	g.binary16 = tempString[len(tempString)-16 : len(tempString)]
}

func main() {
	var genA, genB generator
	var match int64

	genA = generator{factor: 16807, value: 722, multipleOf: 4}
	genB = generator{factor: 48271, value: 354, multipleOf: 8}

	match = 0

	for i := 0; i < 5000000; i++ {
		genA.produce()
		genB.produce()

		if genA.binary16 == genB.binary16 {
			match++
		}

	}
	fmt.Println(match)
}
