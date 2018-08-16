package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bots := make(map[string]*Bot)
	outputs := make(map[string]*Output)
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// First step is to create all the bots with their configurations, also create the outputs.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() != false {
		instr := scanner.Text()
		if strings.Index(instr, "bot") == 0 {
			var botID, lowID, highID int
			var lowType, highType string
			_, err = fmt.Sscanf(instr, "bot %d gives low to %s %d and high to %s %d", &botID, &lowType, &lowID, &highType, &highID)
			if err != nil {
				panic(err)
			}
			if _, found := bots["bot"+strconv.Itoa(botID)]; !found {
				lowBot := lowType + strconv.Itoa(lowID)
				highBot := highType + strconv.Itoa(highID)
				bots["bot"+strconv.Itoa(botID)] = newBot(botID, lowBot, highBot)
				if lowType == "output" {
					if _, found := outputs["output"+strconv.Itoa(lowID)]; !found {
						outputs["output"+strconv.Itoa(lowID)] = newOutput(lowID)
					}
				}
				if highType == "output" {
					if _, found := outputs["output"+strconv.Itoa(highID)]; !found {
						outputs["output"+strconv.Itoa(highID)] = newOutput(highID)
					}
				}
			}

		}
	}
	// Rescan the input for value-lines to get some values into our bots
	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	for scanner.Scan() != false {
		instr := scanner.Text()
		if strings.Index(instr, "value") == 0 {
			var botID, value int
			_, err = fmt.Sscanf(instr, "value %d goes to bot %d", &value, &botID)
			if err != nil {
				panic(err)
			}
			bots["bot"+strconv.Itoa(botID)].receive(bots, outputs, value)
		}
	}

	fmt.Println("Solution for part 2:", outputs["output0"].received[0]*outputs["output1"].received[0]*outputs["output2"].received[0])
}

// Bot is a struct that contains the configuration for a bot
type Bot struct {
	id      int
	lowOut  string
	highOut string
	slots   []int
}

func newBot(id int, lowOut string, highOut string) *Bot {
	return &Bot{
		id:      id,
		lowOut:  lowOut,
		highOut: highOut,
	}
}

func (b *Bot) receive(bots map[string]*Bot, outputs map[string]*Output, value int) {
	b.slots = append(b.slots, value)
	if len(b.slots) == 2 {
		if b.slots[0] == 17 || b.slots[1] == 17 {
			if b.slots[0] == 61 || b.slots[1] == 61 {
				fmt.Println("Solution for part 1:", b.id)
			}
		}
		if b.slots[0] > b.slots[1] {
			if strings.Index(b.lowOut, "bot") == 0 {
				bots[b.lowOut].receive(bots, outputs, b.slots[1])
			} else if strings.Index(b.lowOut, "output") == 0 {
				outputs[b.lowOut].receive(b.slots[1])
			}
			if strings.Index(b.highOut, "bot") == 0 {
				bots[b.highOut].receive(bots, outputs, b.slots[0])
			} else if strings.Index(b.highOut, "output") == 0 {
				outputs[b.highOut].receive(b.slots[1])
			}
		} else if b.slots[0] < b.slots[1] {
			if strings.Index(b.lowOut, "bot") == 0 {
				bots[b.lowOut].receive(bots, outputs, b.slots[0])
			} else if strings.Index(b.lowOut, "output") == 0 {
				outputs[b.lowOut].receive(b.slots[0])
			}
			if strings.Index(b.highOut, "bot") == 0 {
				bots[b.highOut].receive(bots, outputs, b.slots[1])
			} else if strings.Index(b.highOut, "output") == 0 {
				outputs[b.highOut].receive(b.slots[0])
			}
		}
		b.slots = nil
	}
}

// Output is a struct that contains an output slot
type Output struct {
	id            int
	received      []int
	totalReceived int
}

func newOutput(id int) *Output {
	return &Output{
		id:            id,
		totalReceived: 0,
	}
}

func (o *Output) receive(value int) {
	o.totalReceived += value
	o.received = append(o.received, value)
}
