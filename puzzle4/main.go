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
	fmt.Println(subsetCounts)
}

/*
2-3
1-4

a.first >= b.first?
	1 >= 4?

1-4
4-7

4-7
1-4

1234.....
.23......
*/

func arePairsSubsets(pair string) bool {
	split := strings.Split(pair, ",")

	range1 := strings.Split(split[0], "-")
	range2 := strings.Split(split[1], "-")

	a, _ := strconv.Atoi(range1[0])
	b, _ := strconv.Atoi(range1[1])
	c, _ := strconv.Atoi(range2[0])
	d, _ := strconv.Atoi(range2[1])

	fmt.Println(a, b, c, d)

	return (a <= c && b >= d) || (a >= c && b <= d)
}
