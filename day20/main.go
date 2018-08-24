package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var particles []*particle
var minDistID int
var minDist int
var disabled int

type particle struct {
	id               int
	x, y, z          int
	xacc, yacc, zacc int
	xvel, yvel, zvel int
	dist             int
	prevDist         int
	enabled          bool
}

func (p *particle) tick() {
	if p.enabled == true {
		p.xvel += p.xacc
		p.yvel += p.yacc
		p.zvel += p.zacc

		p.x += p.xvel
		p.y += p.yvel
		p.z += p.zvel

		dist := int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z)))

		if dist < minDist {
			minDistID = p.id
			minDist = dist
		}
		p.prevDist = int(dist)

	}
}
func main() {
	fmt.Println("Day 20 ay...")
	parseData()
	fmt.Println(particles)
	minDist = 999999999
	disabled = 0

	for j := 0; j < 1000000; j++ {
		for i := 0; i < 1000; i++ {
			minDist = 999999999
			fmt.Println(len(particles))
			for _, p := range particles {
				p.tick()
			}

			particles = removeCollisions(particles)

			fmt.Printf("in the mean time: loop %v, disabled %v, minDist %v, minDistID %v.. Length: %v\n", j*10000, disabled, minDist, minDistID, len(particles))
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("------------")
	fmt.Println(minDist)
	fmt.Println(minDistID)

}

func removeCollisions(particles []*particle) []*particle {
	retParticles := make([]*particle, 0)

	for i, p := range particles {
		moreThanOne := false

		for i2, p2 := range particles {
			if p.x == p2.x && p.y == p2.y && p.z == p2.z && i != i2 {
				moreThanOne = true
				break
			}

		}

		if !moreThanOne {
			retParticles = append(retParticles, p)
		}

	}

	return retParticles
}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	particles = make([]*particle, 0)
	i := 0

	for fileScanner.Scan() {
		particle := particle{}

		particle.id = i
		particle.enabled = true
		particle.prevDist = 9999999999
		i++
		//p=<4459,-902,279>, v=<-48,66,-14>, a=<-15,-2,0>
		splitOnComma := strings.Split(fileScanner.Text(), ", ")

		positionPart := strings.Split(splitOnComma[0][3:len(splitOnComma[0])-1], ",")
		fmt.Println(positionPart)
		particle.x, _ = strconv.Atoi(string(positionPart[0]))
		particle.y, _ = strconv.Atoi(string(positionPart[1]))
		particle.z, _ = strconv.Atoi(string(positionPart[2]))

		velocityPart := strings.Split(splitOnComma[1][3:len(splitOnComma[1])-1], ",")
		particle.xvel, _ = strconv.Atoi(string(velocityPart[0]))
		particle.yvel, _ = strconv.Atoi(string(velocityPart[1]))
		particle.zvel, _ = strconv.Atoi(string(velocityPart[2]))

		acceleratorPart := strings.Split(splitOnComma[2][3:len(splitOnComma[2])-1], ",")
		particle.xacc, _ = strconv.Atoi(string(acceleratorPart[0]))
		particle.yacc, _ = strconv.Atoi(string(acceleratorPart[1]))
		particle.zacc, _ = strconv.Atoi(string(acceleratorPart[2]))

		particles = append(particles, &particle)
	}
}
