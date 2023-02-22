package main

import (
	"fmt"
	"io/ioutil"
)

const STARTING_PACKET_SIZE = 4

func hasUniqueBytes(s string) bool {
	hasUnique := true

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			hasUnique = hasUnique && s[i] != s[j]
		}

		if !hasUnique {
			break
		}
	}

	return hasUnique
}

func findFirstNUnique(input string, n int) (string, int) {
	for i := n; i < len(input); i++ {
		window := input[i-n : i]
		if hasUniqueBytes(window) {
			return window, i
		}
	}

	return "", -1
}

func main() {
	inputRaw, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := string(inputRaw)
	fmt.Println(findFirstNUnique(input, 4))
	fmt.Println(findFirstNUnique(input, 14))
}
