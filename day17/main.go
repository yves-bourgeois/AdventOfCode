package main

import "fmt"

type spinlock struct {
	buffer       []int64
	bufferSize   int64
	currPosition int64
	value        int64
}

var maxValue int64

func (s *spinlock) algorithm(nrSteps int) {
	s.value++

	// s.currPosition = (s.currPosition + nrSteps) % len(s.buffer)

	// if len(s.buffer) == 1 {
	// 	s.buffer = append(s.buffer, s.value)
	// } else {
	// 	s.buffer = append(s.buffer, 0)
	// 	copy(s.buffer[s.currPosition+1:], s.buffer[s.currPosition:])
	// 	s.buffer[s.currPosition] = s.value
	// 	s.currPosition = (s.currPosition + 1) % len(s.buffer)

	// }
	s.currPosition = (s.currPosition + int64(nrSteps)) % s.bufferSize
	s.bufferSize++

	s.currPosition = (s.currPosition + 1) % s.bufferSize
	//	}

	if s.currPosition == 1 {
		maxValue = s.value
	}

	//	fmt.Printf("CurrPos: %v, value: %v, buffSize: %v\n", s.currPosition, s.value, s.bufferSize)
	// if s.value == 1 {
	// 	s.bufferSize++
	// 	//s.buffer = append(s.buffer, s.value)
	// } else {
	// 	s.currPosition = int(int64((s.currPosition + nrSteps)) % s.bufferSize)

	// 	// s.buffer = append(s.buffer, 0)
	// 	// copy(s.buffer[s.currPosition+1:], s.buffer[s.currPosition:])
	// 	// s.buffer[s.currPosition] = s.value

	// 	s.currPosition = int(int64((s.currPosition + 1)) % s.bufferSize)
	// 	s.bufferSize++
	// }

	// if s.currPosition == 1 {
	// 	maxValue = s.value
	// }
}

func main() {
	maxValue = 0
	spinlock := spinlock{buffer: []int64{0}, bufferSize: 1, currPosition: 0, value: 0}

	for i := 0; i < 50000000; i++ {
		spinlock.algorithm(363)
		//	fmt.Printf("Buffer: %v, currPosition: %v\n", spinlock.buffer, spinlock.currPosition)
	}
	//265
	//fmt.Printf("And for the answer.... second item: %v, %v, %v, %v", spinlock.buffer[1], spinlock.buffer[spinlock.currPosition-1], spinlock.buffer[spinlock.currPosition], spinlock.buffer[spinlock.currPosition+1])
	//fmt.Println(spinlock.buffer)
	fmt.Println(maxValue)
}

//1449 265 1457 532 1531 1453 1215 581
