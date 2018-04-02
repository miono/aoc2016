package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkTriangle(t []int) bool {
	sort.Ints(t)
	if t[0]+t[1] > t[2] {
		return true
	}
	return false
}

func a(f *os.File) int {
	var validTriangles int
	aScanner := bufio.NewScanner(f)
	for aScanner.Scan() {
		stringsides := strings.Fields(aScanner.Text())
		var sides []int
		for _, x := range stringsides {
			side, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			sides = append(sides, side)
		}
		if checkTriangle(sides) {
			validTriangles++
		}

	}
	return validTriangles
}

func b(f *os.File) int {
	var validTriangles int
	bScanner := bufio.NewScanner(f)
	for bScanner.Scan() {
		line1 := strings.Fields(bScanner.Text())
		bScanner.Scan()
		line2 := strings.Fields(bScanner.Text())
		bScanner.Scan()
		line3 := strings.Fields(bScanner.Text())
		for i := 0; i < 3; i++ {
			side1, _ := strconv.Atoi(line1[i])
			side2, _ := strconv.Atoi(line2[i])
			side3, _ := strconv.Atoi(line3[i])
			if checkTriangle([]int{side1, side2, side3}){
				validTriangles++
			}
		}
	}
	return validTriangles
}

func main() {
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The number of valid trinagles in A: %v\n", a(f))
	f.Seek(0, 0)
	fmt.Printf("The number of valid trinagles in B: %v\n", b(f))
}
