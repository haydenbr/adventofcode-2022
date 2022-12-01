package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"haydenbr/adventofcode-2022/util"
)

func main() {
	input, inputErr := os.ReadFile("puzzle1.txt")

	if inputErr != nil {
		log.Fatalln(inputErr)
	}

	groups := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	sums := util.Map(groups, func(group string, _ int) int {
		return util.Fold(strings.Split(group, "\n"), 0, func(result int, s string, _ int) int {
			parsedInt, parsedIntError := strconv.Atoi(s)

			if parsedIntError != nil {
				log.Fatalln(parsedIntError)
			}

			return result + parsedInt
		})
	})

	sort.Ints(sums)
	fmt.Println(util.Last(sums))
	fmt.Println(util.Sum(util.Lastn(sums, 3)))
}
