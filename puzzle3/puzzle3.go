package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"haydenbr/adventofcode-2022/util"
)

/*
- all rucksacks have exactly two compartments for items
- items denoted by single upper case or lower case letter
- first half of line input is in first compartment, second half is in the second compartment
- each rucksack has an item type that occurs in both compartments
- each item type has a priority:
	- a-z: 1-26
	- A-Z: 27-52
- sum priority of item type that appears in both compartments for each rucksack
*/

func main() {
	rucksacks, inputErr := util.GetInputLines("3")

	if inputErr != nil {
		log.Fatalln(inputErr)
	}

	fmt.Println(part1(rucksacks))
	fmt.Println(part2(rucksacks))
}

func part1(rucksacks []string) int {
	return util.Fold(0, rucksacks, func(sum int, rucksack string) int {
		container1, container2 := util.PartitionHalf(util.ToRuneSlice(rucksack))
		commonItem, _ := util.NewSetFromSlice(container1).Intersect(util.NewSetFromSlice(container2)).Pop()
		itemPriority, priorityErr := getItemPriority(commonItem)

		if priorityErr != nil {
			log.Fatalln(priorityErr)
		}

		return sum + itemPriority
	})
}

func part2(rucksacks []string) int {
	return util.Fold(0, util.PartitionEveryN(rucksacks, 3), func(sum int, group []string) int {
		elf1 := util.NewSetFromSlice(util.ToRuneSlice(group[0]))
		elf2 := util.NewSetFromSlice(util.ToRuneSlice(group[1]))
		elf3 := util.NewSetFromSlice(util.ToRuneSlice(group[2]))

		commonItem, _ := elf1.Intersect(elf2).Intersect(elf3).Pop()
		itemPriority, priorityErr := getItemPriority(commonItem)

		if priorityErr != nil {
			log.Fatalln(priorityErr)
		}

		return sum + itemPriority
	})
}

const (
	lowerCaseShift = 'a' - 1
	upperCaseShift = 'A' - 27
)

func getItemPriority(r rune) (int, error) {
	if r <= 'Z' {
		return int(r - upperCaseShift), nil
	} else if r >= 'a' {
		return int(r - lowerCaseShift), nil
	} else {
		return 0, errors.New("invalid input: " + strconv.Itoa(int(r)))
	}
}
