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
	sums := util.ProjectTo(groups, func(group string) int {
		return util.Fold(0, strings.Split(group, "\n"), func(result int, s string) int {
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
