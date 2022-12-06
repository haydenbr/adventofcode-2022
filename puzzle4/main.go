package main

import (
	"fmt"
	"haydenbr/adventofcode-2022/util"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, inputErr := util.GetInputLines("4")

	if inputErr != nil {
		log.Fatalln(inputErr)
	}

	subsetCounts := util.Fold(0, input, func(sum int, pair string) int {
		if arePairsSubsets(pair) {
			return sum + 1
		} else {
			return sum
		}
	})

	overlapCounts := util.Fold(0, input, func(sum int, pair string) int {
		if doPairsOverlap(pair) {
			return sum + 1
		} else {
			return sum
		}
	})

	fmt.Println(subsetCounts)
	fmt.Println(overlapCounts)
}

func arePairsSubsets(pair string) bool {
	split := strings.Split(pair, ",")

	range1 := strings.Split(split[0], "-")
	range2 := strings.Split(split[1], "-")

	a, _ := strconv.Atoi(range1[0])
	b, _ := strconv.Atoi(range1[1])
	c, _ := strconv.Atoi(range2[0])
	d, _ := strconv.Atoi(range2[1])

	return (a <= c && b >= d) || (a >= c && b <= d)
}

func doPairsOverlap(pair string) bool {
	split := strings.Split(pair, ",")

	range1 := strings.Split(split[0], "-")
	range2 := strings.Split(split[1], "-")

	a, _ := strconv.Atoi(range1[0])
	b, _ := strconv.Atoi(range1[1])
	c, _ := strconv.Atoi(range2[0])
	d, _ := strconv.Atoi(range2[1])

	return (a <= d && b >= c) || (a >= d && b <= c)
}
