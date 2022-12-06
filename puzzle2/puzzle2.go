package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"haydenbr/adventofcode-2022/util"
)

/*
rules:
rock > scissors
paper > rock
scissors > paper

key:
rock: A X
paper: B Y
scissors: C Z

X: lose
Y: draw
Z: win

score:
rock: 1
paper: 2
scissors: 3

loss: 0
tie: 3
win: 6
*/

var part1Cases = map[string]int{
	"A X": 1 + 3,
	"A Y": 2 + 6,
	"A Z": 3 + 0,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 1 + 6,
	"C Y": 2 + 0,
	"C Z": 3 + 3,
}

var part2Cases = map[string]int{
	"A X": 3 + 0,
	"A Y": 1 + 3,
	"A Z": 2 + 6,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 2 + 0,
	"C Y": 3 + 3,
	"C Z": 1 + 6,
}

// what do I choose? did I win, draw, or loose?

func main() {
	input, inputErr := os.ReadFile("puzzle2.txt")

	if inputErr != nil {
		log.Fatalln(inputErr)
	}

	rounds := strings.Split(strings.TrimSpace(string(input)), "\n")
	part1Score := util.Fold(0, rounds, func(score int, round string) int {
		return score + part1Cases[round]
	})
	part2Score := util.Fold(0, rounds, func(score int, round string) int {
		return score + part2Cases[round]
	})

	fmt.Println("part 1", part1Score)
	fmt.Println("part 2", part2Score)
}
