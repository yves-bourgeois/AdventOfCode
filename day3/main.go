package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	x int64
	y int64
}

//anonymous function: up / down / left / right
type action func(c coordinate) coordinate

//Keeps all calculated coordinates
var coordinates map[coordinate]int64

func traverse(target int64, dimension int, c coordinate) coordinate {
	//Dimension:
	//	0: starter square
	//	1: first circle     8 squares
	//	2: second circle	16 squares
	//	3: third circle		24 squares
	fmt.Printf("Number of elements in this circle: %v, currently at (%v, %v)\n", (dimension * 8), c.x, c.y)

	targetReached := false

	//until we break..
	for {
		stepsPerDirection := (dimension * 8) / 4

		//make one step to the right to start the new cycle
		c, targetReached = step(right, c, 1, target)

		//step up -- left -- down -- right until the target is reached
		if targetReached == false {
			c, targetReached = step(up, c, stepsPerDirection-1, target)
		}

		if targetReached == false {
			c, targetReached = step(left, c, stepsPerDirection, target)
		}

		if targetReached == false {
			c, targetReached = step(down, c, stepsPerDirection, target)
		}

		if targetReached == false {
			c, targetReached = step(right, c, stepsPerDirection, target)
		}

		if targetReached == false {
			c = traverse(target, dimension+1, c)
			break
		} else {
			break
		}

	}

	return c
}

//Step
// in a certain direction (defined by anonymous function) act
// nrSteps times
// until a target is reached
func step(act action, c coordinate, nrSteps int, target int64) (coordinate, bool) {
	targetReached := false

	for i := 0; i < nrSteps && targetReached == false; i++ {
		c = act(c)
		val := c.calcValue()
		if val > target {
			fmt.Printf("Target reached: %v\n", val)
			targetReached = true
		}
	}

	return c, targetReached
}

//Add up all values that are surrounding this coordinate
//If no coordinate is found: it will return 0
func (c *coordinate) calcValue() int64 {
	val1 := coordinates[coordinate{c.x - 1, c.y - 1}]
	val2 := coordinates[coordinate{c.x - 1, c.y}]
	val3 := coordinates[coordinate{c.x - 1, c.y + 1}]

	val4 := coordinates[coordinate{c.x, c.y - 1}]
	val6 := coordinates[coordinate{c.x, c.y + 1}]

	val7 := coordinates[coordinate{c.x + 1, c.y - 1}]
	val8 := coordinates[coordinate{c.x + 1, c.y}]
	val9 := coordinates[coordinate{c.x + 1, c.y + 1}]

	coordinates[coordinate{x: c.x, y: c.y}] = val1 + val2 + val3 + val4 + val6 + val7 + val8 + val9

	return val1 + val2 + val3 + val4 + val6 + val7 + val8 + val9
}

func main() {
	fmt.Println("A good start..")

	target := 347991

	//	37	36	35	34	33	32	31
	//	38	17  16  15  14  13	30
	//	39	18   5   4   3  12	29
	//	40	19   6   1   2  11	28
	//	41	20   7   8   9  10	27
	//	42	21  22  23	24	25	26
	//	43	44	45	46	47	48	49

	coordinates = make(map[coordinate]int64)

	//Initialize map: 1 goes in the middle
	c := coordinate{0, 0}

	coordinates[c] = 1

	c = traverse(int64(target), 1, c)

	fmt.Printf("X = %v \n", c.x)
	fmt.Printf("Y = %v \n", c.y)

	//taxi distance
	distance := math.Abs(float64(c.y)) + math.Abs(float64(c.x))

	fmt.Printf("Distance to value %v is %v", coordinates[c], distance)
}

func left(c coordinate) coordinate {
	return coordinate{x: c.x - 1, y: c.y}
}

func right(c coordinate) coordinate {
	return coordinate{x: c.x + 1, y: c.y}
}

func up(c coordinate) coordinate {
	return coordinate{x: c.x, y: c.y + 1}
}

func down(c coordinate) coordinate {
	return coordinate{x: c.x, y: c.y - 1}
}
