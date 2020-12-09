package main

import (
    "fmt"
    "math"
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
    fmt.Printf("Part 1: Weak number is %v.\n", weakness)

    slice := sequence.FindContiguousSet(weakness)
    min,max := FindMinMax(slice)
    fmt.Printf("Part 2: Sum of min-max is %v.", min + max)
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

func (s Sequence) FindContiguousSet(value int) []int {
    var slice []int
    for i := range s.input {
        for j := i; j < len(s.input); j++ {
            slice = s.input[i:j]
            sum := Sum(slice)
            if sum == value {
                return slice
            }
        }
    }

    return nil
}

func Sum(input []int) int {
    sum := 0
    for _, number := range input {
        sum += number
    }
    return sum
}

func Contains(input []int, value int) bool {
    for _, x := range input {
        if x == value {
            return true
        }
    }
    return false
}

func FindMinMax(slice []int) (int, int) {
    max := 0
    min := math.MaxInt32

    for _, number := range slice {
        if min > number {
            min = number
        }
        if max < number {
            max = number
        }
    }

    return min, max
}
