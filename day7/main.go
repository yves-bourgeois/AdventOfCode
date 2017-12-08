package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nodes map[string]*node

type node struct {
	name     string
	weight   int
	prevNode *node
	nextNode []*node
}

func main() {
	fmt.Println("day 7: And so it begins..")

	nodes = make(map[string]*node)

	parseData()

	var rootNode node
	for _, node := range nodes {
		if node.prevNode == nil {
			rootNode = *node
			break
		}
	}

	fmt.Println("Rootnode: " + rootNode.name)

	findBalance(rootNode.name)
}

func findBalance(rootNode string) {
	for _, node := range nodes[rootNode].nextNode {
		fmt.Println("Node " + node.name + " has a weight of " + strconv.Itoa(node.weight))
		fmt.Println("This program: " + node.name + " has a total weight of " + strconv.Itoa(node.calcTotalWeight()))
	}

	fmt.Println("-------")
	for _, node := range nodes["gtervu"].nextNode {
		fmt.Println("Node " + node.name + " has a weight of " + strconv.Itoa(node.weight))
		fmt.Println("This program: " + node.name + " has a total weight of " + strconv.Itoa(node.calcTotalWeight()))
	}

	fmt.Println(75514 - 75509)
	fmt.Println(2260 - 5)
}

func (n *node) calcTotalWeight() int {
	weight := n.weight

	for _, child := range n.nextNode {
		weight = weight + child.calcTotalWeight()
	}

	return weight

}

func parseData() {
	fileHandle, _ := os.Open("input")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		splitOnArrow := strings.Split(fileScanner.Text(), " -> ")

		nodeInfo := strings.Split(splitOnArrow[0], " ")

		nodeName := strings.Trim(nodeInfo[0], " ")
		nodeWeight, _ := strconv.Atoi(strings.Replace(strings.Replace(strings.Trim(nodeInfo[1], " "), "(", "", -1), ")", "", -1))

		if _, ok := nodes[nodeName]; ok {
			//Node already exists: only set its weight
			nodes[nodeName].weight = nodeWeight
		} else {
			node := node{name: nodeName, weight: nodeWeight, prevNode: nil, nextNode: nil}
			nodes[nodeName] = &node
		}

		if len(splitOnArrow) == 2 {
			//children found
			children := strings.Split(splitOnArrow[1], ", ")
			for _, child := range children {
				if _, ok := nodes[child]; ok == false {
					//preventifly create child node
					node := node{name: child, weight: -1, prevNode: nil, nextNode: nil}
					nodes[child] = &node
				}
				nodes[child].prevNode = nodes[nodeName]
				nodes[nodeName].nextNode = append(nodes[nodeName].nextNode, nodes[child])
			}
		}

	}
}
