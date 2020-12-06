package main

import (
	"fmt"
	"strings"
)

type Group struct {
	People  int
	Answers map[rune]int
}

type YesCounter struct {
	raw    string
	Groups []Group
}

func (y *YesCounter) Init(input string) {
	rawSplit := strings.Split(input, "\n")

	currentGroup := 0
	// Initial group
	y.Groups = append(y.Groups, Group{0, make(map[rune]int)})

	for _, value := range rawSplit {
		if len(value) == 0 {
			y.Groups = append(y.Groups, Group{0, make(map[rune]int)})
			currentGroup++
			continue
		}

		y.Groups[currentGroup].People++
		for _, question := range value {
			y.Groups[currentGroup].Answers[question]++
		}
	}
}

func (y YesCounter) CountAnswers() int {
	count := 0
	for _, group := range y.Groups {
		count += len(group.Answers)
	}
	return count
}

func (y YesCounter) CountCommonAnswers() int {
	count := 0
	for _, currentGroup := range y.Groups {
		for _, answerCount := range currentGroup.Answers {
			if answerCount == currentGroup.People{
				count++
			}
		}
	}
	return count
}

func main() {
	counter := new(YesCounter)
	counter.Init(PuzzleInput)
	yesAnswers := counter.CountAnswers()

	fmt.Printf("Part 1: Counted %v yes-answers.\n", yesAnswers)

	commonYesAnswers := counter.CountCommonAnswers()
	fmt.Printf("Part 2: Counted %v yes-answers.", commonYesAnswers)
}
