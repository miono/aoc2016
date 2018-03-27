package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func catch(e error) {
	fmt.Println(e)
}

type sleigh struct {
	curPos           sleighPos
	direction        int8
	allPos           []sleighPos
	firstDoubleVisit sleighPos
	doubleHit        bool
}

type sleighPos struct {
	xPos int
	yPos int
}

const (
	N int8 = iota
	E
	S
	W
)

func (s *sleigh) turnAndMove(direction string, distance int) {
	// TURN
	switch direction {
	case "R":
		s.direction = (s.direction + 1) % 4
	case "L":
		s.direction = (s.direction + 3) % 4
	}

	// MOVE
	switch s.direction {
	case N:
		for i := 0; i < distance; i++ {
			s.curPos.yPos++
			s.checkPositions(s.curPos)
		}
	case S:
		for i := 0; i < distance; i++ {
			s.curPos.yPos--
			s.checkPositions(s.curPos)
		}
	case E:
		for i := 0; i < distance; i++ {
			s.curPos.xPos++
			s.checkPositions(s.curPos)
		}
	case W:
		for i := 0; i < distance; i++ {
			s.curPos.xPos--
			s.checkPositions(s.curPos)
		}
	}
}

func (s *sleigh) checkPositions(newPos sleighPos) {
	if s.doubleHit == false {
		for i := range s.allPos {
			if newPos == s.allPos[i] {
				s.firstDoubleVisit = newPos
				s.doubleHit = true
			}
		}
	}
	s.allPos = append(s.allPos, newPos)
}

func main() {
	newSleigh := sleigh{sleighPos{0, 0}, N, make([]sleighPos, 0), sleighPos{}, false}
	newSleigh.allPos = append(newSleigh.allPos, sleighPos{0, 0})
	inputString := "R4, R3, R5, L3, L5, R2, L2, R5, L2, R5, R5, R5, R1, R3, L2, L2, L1, R5, L3, R1, L2, R1, L3, L5, L1, R3, L4, R2, R4, L3, L1, R4, L4, R3, L5, L3, R188, R4, L1, R48, L5, R4, R71, R3, L2, R188, L3, R2, L3, R3, L5, L1, R1, L2, L4, L2, R5, L3, R3, R3, R4, L3, L4, R5, L4, L4, R3, R4, L4, R1, L3, L1, L1, R4, R1, L4, R1, L1, L3, R2, L2, R2, L1, R5, R3, R4, L5, R2, R5, L5, R1, R2, L1, L3, R3, R1, R3, L4, R4, L4, L1, R1, L2, L2, L4, R1, L3, R4, L2, R3, L1, L5, R4, R5, R2, R5, R1, R5, R1, R3, L3, L2, L2, L5, R2, L2, R5, R5, L2, R3, L5, R5, L2, R4, R2, L1, R3, L5, R3, R2, R5, L1, R3, L2, R2, R1"
	for _, x := range strings.Split(inputString, ", ") {
		direction := x[:1]
		distance, err := strconv.Atoi(x[1:])
		if err != nil {
			catch(err)
		}
		newSleigh.turnAndMove(direction, distance)
	}
	fmt.Printf("Distance to the end of the moves is: %d\n", newSleigh.curPos.xPos+newSleigh.curPos.yPos)
	fmt.Printf("Distance to first double visit is: %v", intAbs(newSleigh.firstDoubleVisit.xPos)+intAbs(newSleigh.firstDoubleVisit.yPos))
	//	fmt.Println(newSleigh.allPos)

}
