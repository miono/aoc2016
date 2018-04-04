package main

import (
	"bufio"
	"fmt"
	"os"
)

func getRealMsg(straangs []string, mode byte) string {

	var outString string
	for i := range straangs[0] {
		var charList []string
		for _, x := range straangs {
			charList = append(charList, string(x[i]))
		}
		charMap := make(map[string]int)
		for _, char := range charList {
			if _, ok := charMap[char]; ok {
				charMap[char]++
			} else {
				charMap[char] = 1
			}
		}
		// Get most common character
		if mode == 'a' {
			var highestChar string
			var highestCount int
			for k, v := range charMap {
				if v > highestCount {
					highestChar = k
					highestCount = v
				}
			}
			outString = outString + highestChar
		}
		if mode == 'b' {
			var highestChar string
			var highestCount = 1000
			for k, v := range charMap {
				if v < highestCount {
					highestChar = k
					highestCount = v
				}
			}
			outString = outString + highestChar
		}

	}

	return outString
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	var msgs []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		msgs = append(msgs, scanner.Text())
	}
	fmt.Printf("The first password is %v\n", getRealMsg(msgs, 'a'))
	fmt.Printf("The second password is %v\n", getRealMsg(msgs, 'b'))
}
