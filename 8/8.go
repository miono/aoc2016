package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type display struct {
	grid [][]bool
	w    int
	h    int
}

func newDisplay(w, h int) *display {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &display{grid: s, w: w, h: h}
}

func (d display) rect(w int, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			d.grid[j][i] = true
		}
	}
}

func (d display) rotate(mode string, place int, amount int) {
	if mode == "row" {
		inRow := d.grid[place]
		outRow := make([]bool, d.w)
		for i, status := range inRow {
			outRow[(i+amount)%len(inRow)] = status
		}
		d.grid[place] = outRow

	} else if mode == "column" {
		var inColumn []bool
		outColumn := make([]bool, d.h)
		for i := 0; i < d.h; i++ {
			inColumn = append(inColumn, d.grid[i][place])
		}
		for i, status := range inColumn {
			outColumn[(i+amount)%d.h] = status
		}
		for i, status := range outColumn {
			d.grid[i][place] = status
		}
	}
}

func (d display) print() {
	for _, line := range d.grid {
		for _, cell := range line {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print("·")
			}
		}
		fmt.Print("\n")
	}
}

func (d display) countLit() int {
	i := 0
	for _, line := range d.grid {
		for _, cell := range line {
			if cell {
				i++
			}
		}
	}
	return i
}

func main() {
	myDisplay := newDisplay(50, 6)

	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instr := strings.Split(scanner.Text(), " ")
		switch instr[0] {
		case "rect":
			coords := strings.Split(instr[1], "x")
			w, _ := strconv.Atoi(coords[0])
			h, _ := strconv.Atoi(coords[1])
			myDisplay.rect(w, h)
		case "rotate":
			mode := instr[1]
			amount, _ := strconv.Atoi(instr[4])
			place, _ := strconv.Atoi(strings.Split(instr[2], "=")[1])
			myDisplay.rotate(mode, place, amount)

		}

	}
	fmt.Println("This is what the display looks after all instructions have been made:")
	myDisplay.print()
	fmt.Printf("Amount of lit LEDs: %d\n", myDisplay.countLit())
}
