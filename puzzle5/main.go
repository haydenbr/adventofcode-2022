package main

import (
	"bufio"
	"fmt"
	"haydenbr/adventofcode-2022/util"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const CHAR_INDEX_OFFSET = len("[X] ")

var MOVEMENT_REGEX = regexp.MustCompile("(\\d+)")

func main() {
	file, fileErr := os.Open("puzzle5.txt")
	defer file.Close()

	if fileErr != nil {
		panic(fileErr)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	stackDiagram := make([]string, 0)

	for !util.IsWhiteSpace(fileScanner.Text()) {
		stackDiagram = append(stackDiagram, fileScanner.Text())
		fileScanner.Scan()
	}

	stackIndices := strings.Split(strings.TrimSpace(stackDiagram[len(stackDiagram)-1]), "   ")
	stackDiagram = util.TrimLast(stackDiagram)

	stacks := util.SliceToMap(stackIndices, func(i string) string { return i }, func(i string) []string { return make([]string, 0) })

	util.ForEachReverse(stackDiagram, func(row string) {
		for i, stackIndex := range stackIndices {
			rowCharIndex := (i * CHAR_INDEX_OFFSET) + 1
			char := row[rowCharIndex : rowCharIndex+1]

			if !util.IsWhiteSpace(char) {
				stack := stacks[stackIndex]
				stacks[stackIndex] = util.Push(stack, char)
			}
		}
	})

	for fileScanner.Scan() {
		movement := fileScanner.Text()
		parsedMovement := MOVEMENT_REGEX.FindAllString(movement, 3)
		moveAmount, err := strconv.Atoi(parsedMovement[0])

		if err != nil {
			panic(err)
		}

		sourceIndex := parsedMovement[1]
		destinationIndex := parsedMovement[2]
		source := stacks[sourceIndex]
		destination := stacks[destinationIndex]

		_source, popped := util.CutN(source, moveAmount)
		source = _source
		destination = util.Push(destination, popped...)

		stacks[sourceIndex] = source
		stacks[destinationIndex] = destination
	}

	result := util.Fold("", stackIndices, func(result string, i string) string {
		stack := stacks[i]
		return result + util.Peek(stack)
	})

	fmt.Println(result)
}
