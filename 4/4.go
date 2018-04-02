package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	roomName string
	sectorID int
	checksum string
}

func (r room) checkRoom() int {
	var realChecksum string
	type kv struct {
		key   string
		value int
	}
	charMap := make(map[string]int)
	for _, char := range strings.Split(r.roomName, "") {
		if _, ok := charMap[char]; ok {
			charMap[char]++
		} else {
			charMap[char] = 1
		}
	}
	var sortedMap []kv
	for k, v := range charMap {
		sortedMap = append(sortedMap, kv{k, v})
	}
	sort.Slice(sortedMap, func(i, j int) bool {
		return sortedMap[i].value > sortedMap[j].value
	})
	for i := sortedMap[0].value; len(realChecksum) <= 5; i-- {
		var keyz []string
		for _, v := range sortedMap {
			if v.value == i {
				keyz = append(keyz, v.key)
			}
		}
		sort.Strings(keyz)
		realChecksum = realChecksum + strings.Join(keyz, "")

	}
	if realChecksum[:5] == r.checksum {
		return r.sectorID
	} else {
		return 0
	}
}

func main() {
	// var rooms []room
	var sectorSum int
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		roomData := scanner.Text()
		re, _ := regexp.Compile("[-\\[\\]]")
		var roomName string
		dashSplitted := re.Split(roomData, -1)
		for _, x := range dashSplitted[:len(dashSplitted)-3] {
			roomName = roomName + x
		}
		sectorID, _ := strconv.Atoi(dashSplitted[len(dashSplitted)-3])
		checksum := dashSplitted[len(dashSplitted)-2]
		newRoom := room{roomName, sectorID, checksum}
		sectorSum += newRoom.checkRoom()
	}
	fmt.Printf("The sum of sectorIDs with real rooms is %v\n", sectorSum)
}
