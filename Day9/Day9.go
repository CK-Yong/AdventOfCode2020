package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    parsed := ParseInput(PuzzleInput)
    sequence := Sequence{
        input:          parsed,
        preambleLength: 25,
    }

    weakness := sequence.FindWeakness()
    fmt.Printf("Part 1: Weak number is %v.", weakness)
}

func ParseInput(input string) []int {
    parsed := strings.Split(input, "\n")
    converted := make([]int, 0)
    for _, entry := range parsed {
        number, _ := strconv.Atoi(entry)
        converted = append(converted, number)
    }
    return converted
}

type Sequence struct {
    input          []int
    preambleLength int
}

func (s Sequence) FindWeakness() int {
    startIndex := s.preambleLength
    weaknessFound := false
    for i := startIndex; i < len(s.input); i++ {
        preamble := s.input[i-s.preambleLength : i]
        currentNumber := s.input[i]
        for _, p := range preamble {
            diff := currentNumber - p
            if Contains(preamble, diff) {
                weaknessFound = false
                break
            } else {
                weaknessFound = true
            }
        }
        if weaknessFound {
            return currentNumber
        }
    }

    return -1
}

func Contains(input []int, value int) bool {
    for _, x := range input {
        if x == value {
            return true
        }
    }
    return false
}
