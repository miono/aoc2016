package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func getMD5(in string) string {
	hasher := md5.New()
	hasher.Write([]byte(in))
	return hex.EncodeToString(hasher.Sum(nil))

}

func checkIfDone(pw []string) bool {
	for _, x := range pw {
		if string(x) == "_" {
			return false
		}
	}
	return true
}

func a(doorID string) string {
	var password string
	for j := 0; ; j++ {
		stringToHash := doorID + strconv.Itoa(j)
		if getMD5(stringToHash)[:5] == "00000" {
			password = password + string(getMD5(stringToHash)[5])
			if len(password) == 8 {
				break
			}

		}

	}
	return password
}

func b(doorID string) string {
	password := []string{"_", "_", "_", "_", "_", "_", "_", "_"}
	for j := 0; ; j++ {
		stringTohash := doorID + strconv.Itoa(j)
		hash := getMD5(stringTohash)
		if hash[:5] == "00000" {
			position, err := strconv.Atoi(string(hash[5]))
			value := string(hash[6])
			if err != nil {
				continue
			}
			if position < 8 && password[position] == "_" {
				password[position] = value
				fmt.Println(password)
				if checkIfDone(password) {
					break
				}
			}
		}

	}
	return strings.Join(password, "")
}

func main() {
	fmt.Printf("The first password is: %s\n", a("wtnhxymk"))
	fmt.Println("Starting to decrypt second password:")
	fmt.Printf("The second password is: %s\n", b("wtnhxymk"))
}
