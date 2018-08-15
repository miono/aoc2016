package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decompress(s string, p int) int {
	var instructions []string
	if strings.Index(s, "(") == -1 {
		return len(s)
	}
	ret := 0
	for strings.Index(s, "(") != -1 {
		boundStart := strings.Index(s, "(")
		ret += boundStart
		s = s[boundStart:]
		boundEnd := strings.Index(s, ")")
		instructions = strings.Split(s[1:boundEnd], "x")
		s = s[boundEnd+1:]
		numChars, _ := strconv.Atoi(instructions[0])
		numRepeat, _ := strconv.Atoi(instructions[1])
		if p == 1 {
			ret += len(s[:numChars]) * numRepeat
		} else if p == 2 {
			ret += decompress(s[:numChars], 2) * numRepeat
		}
		s = s[numChars:]
	}
	ret += len(s)
	return ret
}

func main() {
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fullString := scanner.Text()
	fmt.Println("Solution part 1:", decompress(fullString, 1))
	fmt.Println("Solution part 2:", decompress(fullString, 2))
}
