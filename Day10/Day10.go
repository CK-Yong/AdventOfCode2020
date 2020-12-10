package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parsed := strings.Split(PuzzleInput, "\n")

	adapterSet := AdapterSet{map[int]int{1: 0, 2: 0, 3: 0}}
	adapterSet.Traverse(parsed)

	fmt.Printf("Part 1: 1 jolt * 3 jolt counts: %v", adapterSet.Values[1]*adapterSet.Values[3])
}

type AdapterSet struct {
	Values map[int]int
}

func (set AdapterSet) Traverse(input []string) {
	currentValue := 0
	for {
		nextStep := FindNextStep(input, currentValue)
		if nextStep < 0 {
			set.Values[3]++ // Add own device
			return
		}

		set.Values[nextStep - currentValue]++
		currentValue = nextStep
	}
}

func FindNextStep(input []string, currentValue int) int {
	for i := 1; i < 4; i++ {
		target := currentValue + i
		if Contains(input, target) {
			return target
		}
	}
	return -1
}

func Contains(input []string, target int) bool {
	for _, val := range input {
		number, _ := strconv.Atoi(val)
		if target == number {
			return true
		}
	}
	return false
}

func (set AdapterSet) CountPossibleArrangements(parsed []string) int64 {
	
}