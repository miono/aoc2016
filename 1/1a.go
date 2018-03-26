package main

import (
	"fmt"
	"strings"
	"strconv"
)

func catch(e error) {
	fmt.Println(e)
}
type sleigh struct {
	xPos      int
	yPos      int
	direction Direction
}

type Direction int8

const (
	N Direction = iota
	E
	S
	W
)

func (s *sleigh) turnAndMove(direction string, distance int) {
	// TURN
	switch  direction {
	case "R":
		s.direction = s.direction + 1%4
	case "L":
		s.direction = s.direction - 1%4
	}
	switch s.direction {
	case N:
		s.yPos += distance
	case S:
		s.yPos -= distance
	case E:
		s.xPos += distance
	case W:
		s.xPos -= distance
	}
}

func getMoveSlice(moves string) []string {
	return strings.Split(moves, ", ")
}

func main() {
	newSleigh := sleigh{0, 0, N}
	inputString := "R4, R3, R5, L3, L5, R2, L2, R5, L2, R5, R5, R5, R1, R3, L2, L2, L1, R5, L3, R1, L2, R1, L3, L5, L1, R3, L4, R2, R4, L3, L1, R4, L4, R3, L5, L3, R188, R4, L1, R48, L5, R4, R71, R3, L2, R188, L3, R2, L3, R3, L5, L1, R1, L2, L4, L2, R5, L3, R3, R3, R4, L3, L4, R5, L4, L4, R3, R4, L4, R1, L3, L1, L1, R4, R1, L4, R1, L1, L3, R2, L2, R2, L1, R5, R3, R4, L5, R2, R5, L5, R1, R2, L1, L3, R3, R1, R3, L4, R4, L4, L1, R1, L2, L2, L4, R1, L3, R4, L2, R3, L1, L5, R4, R5, R2, R5, R1, R5, R1, R3, L3, L2, L2, L5, R2, L2, R5, R5, L2, R3, L5, R5, L2, R4, R2, L1, R3, L5, R3, R2, R5, L1, R3, L2, R2, R1"
	moves := getMoveSlice(inputString)
	for i := 0; i < len(moves); i++ {
		direction := moves[i][:1]
		distance, err := strconv.Atoi(moves[i][1:])
		if err != nil {
			catch(err)
		}
		newSleigh.turnAndMove(direction,distance)
	}
	fmt.Printf("Distance to the end of the moves is: %d\n", newSleigh.xPos + newSleigh.yPos)
}

