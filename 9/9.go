package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var outString string
	f, err := os.Open("./input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fullString := scanner.Text()
	reFind, err := regexp.Compile("\\(\\d+x\\d+\\)")
	if err != nil {
		panic(err)
	}
	reSplit, err := regexp.Compile("[)(x]")
	if err != nil {
		panic(err)
	}
	for {
		// fmt.Println(outString)
		bounds := reFind.FindStringIndex(fullString)
		if bounds == nil {
			outString = outString + fullString
			break
		}
		outString = outString + fullString[:bounds[0]]
		fmt.Println(bounds)
		instruction := fullString[bounds[0]:bounds[1]]
		fmt.Println(instruction)
		charAndRepeat := reSplit.Split(instruction, -1)
		numChar, _ := strconv.Atoi(charAndRepeat[1])
		numRepeat, _ := strconv.Atoi(charAndRepeat[2])
		charsToRepeat := fullString[bounds[1] : bounds[1]+numChar]
		stringToAdd := strings.Repeat(charsToRepeat, numRepeat)
		outString = outString + stringToAdd
		fullString = fullString[bounds[1]+numChar:]

	}
	fmt.Println(len(outString))

}
