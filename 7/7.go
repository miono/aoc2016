package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ip struct {
	addresses []string
	hypernets []string
}

func (i ip) verifyTLS() bool {
	addrsVerify := false
	netsVerify := true
	for _, x := range i.addresses {
		if checkForAbba(x) {
			addrsVerify = true
		}
	}
	for _, x := range i.hypernets {
		if checkForAbba(x) {
			netsVerify = false
		}
	}
	if netsVerify && addrsVerify {
		return true
	}
	return false
}

func (i ip) verifySSL() bool {
	var allAbas []string
	for _, x := range i.addresses {
		allAbas = append(allAbas, getAbas(x)...)
	}
	if len(allAbas) == 0 {
		return false
	}
	for _, x := range allAbas {
		y := string(x[1]) + string(x[0]) + string(x[1])
		fmt.Println(y)
		for _, z := range i.hypernets {
			if strings.Contains(z, y) {
				return true
			}
		}
	}
	return false
}

func checkForAbba(s string) bool {
	stringSlice := strings.Split(s, "")
	for i := 0; i < len(stringSlice)-3; i++ {
		if stringSlice[i] != stringSlice[i+1] {
			if stringSlice[i] == stringSlice[i+3] {
				if stringSlice[i+1] == stringSlice[i+2] {
					return true
				}
			}
		}
	}
	return false
}

func getAbas(s string) []string {
	var out []string
	stringSlice := strings.Split(s, "")
	for i := 0; i < len(stringSlice)-2; i++ {
		if stringSlice[i] != stringSlice[i+1] {
			if stringSlice[i] == stringSlice[i+2] {
				out = append(out, strings.Join(stringSlice[i:i+3], ""))
			}
		}
	}
	return out
}

func main() {
	f, err := os.Open("./input")
	check(err)
	defer f.Close()
	var ipAddresses []ip
	scanner := bufio.NewScanner(f)
	var nets []string
	var addrs []string
	for scanner.Scan() {
		var netflag = false
		var curstring []byte
		for _, x := range scanner.Text() {
			switch x {
			case '[':
				if netflag == false {
					addrs = append(addrs, string(curstring))
					curstring = []byte{}
					netflag = true
				}
			case ']':
				if netflag == true {
					nets = append(nets, string(curstring))
					curstring = []byte{}
					netflag = false
				}
			default:
				curstring = append(curstring, byte(x))
			}

		}
		if netflag == true {
			nets = append(nets, string(curstring))
		} else {
			addrs = append(addrs, string(curstring))
		}
		curstring = []byte{}
		ipAddresses = append(ipAddresses, ip{addrs, nets})
		addrs = []string{}
		nets = []string{}
	}

	sumTLS := 0
	sumSSL := 0
	for _, x := range ipAddresses {
		if x.verifyTLS() {
			sumTLS++
		}
		if x.verifySSL() {
			sumSSL++
		}
	}

	fmt.Println(sumTLS)
	fmt.Println(sumSSL)

}
