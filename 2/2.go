package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type keyPad struct {
	buttons  [][]string
	position keyPos
}

type keyPos struct {
	yPos int
	xPos int
}

func (k *keyPad) findButton(moves string) string {
	for _, x := range strings.Split(moves, "") {
		switch x {
		case "U":
			if k.buttons[k.position.yPos-1][k.position.xPos] != "." {
				k.position.yPos--
			}
		case "D":
			if k.buttons[k.position.yPos+1][k.position.xPos] != "." {
				k.position.yPos++
			}
		case "L":
			if k.buttons[k.position.yPos][k.position.xPos-1] != "." {
				k.position.xPos--
			}
		case "R":
			if k.buttons[k.position.yPos][k.position.xPos+1] != "." {
				k.position.xPos++
			}
		}
	}
	return k.buttons[k.position.yPos][k.position.xPos]
}

func main() {
	aKeyPad := keyPad{[][]string{
		{".", ".", ".", ".", "."},
		{".", "1", "2", "3", "."},
		{".", "4", "5", "6", "."},
		{".", "7", "8", "9", "."},
		{".", ".", ".", ".", "."}},
		keyPos{2, 2}}

	bKeyPad := keyPad{[][]string{
		{".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "1", ".", ".", "."},
		{".", ".", "2", "3", "4", ".", "."},
		{".", "5", "6", "7", "8", "9", "."},
		{".", ".", "A", "B", "C", ".", "."},
		{".", ".", ".", "D", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "."}},
		keyPos{3, 1}}

	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var aOut string
	var bOut string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		aOut = aOut + aKeyPad.findButton(scanner.Text())
		bOut = bOut + bKeyPad.findButton(scanner.Text())
	}
	fmt.Printf("The first code is: %s\n", aOut)
	fmt.Printf("The second code is: %s\n", bOut)

}
