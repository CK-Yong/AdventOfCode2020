package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

func main() {
    parsed := strings.Split(PuzzleInput, "\n")

    adapterSet := AdapterSet{map[int]int{1: 0, 2: 0, 3: 0}}
    adapterSet.Traverse(parsed)

    fmt.Printf("Part 1: 1 jolt * 3 jolt counts: %v\n", adapterSet.Values[1]*adapterSet.Values[3])

    arrangementCount := adapterSet.CountPossibleArrangements(parsed)
    fmt.Printf("Part 2: total number of arrangements: %f", arrangementCount)
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

        set.Values[nextStep-currentValue]++
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

func (set AdapterSet) CountPossibleArrangements(input []string) float64 {
    adapters := make([]int, 0)
    for _, val := range input {
        number, _ := strconv.Atoi(val)
        adapters = append(adapters, number)
    }

    sort.Ints(adapters)
    // Add initial value and my own device
    adapters = append([]int{0}, adapters...)
    adapters = append(adapters, adapters[len(adapters)-1]+3)

    // We take slices of 3, so stop when there are two values left
    omittable := make([]int, 0)
    for i := 0; i < len(adapters)-2; i++ {
        slice := adapters[i : i+3]
        if slice[2]-slice[0] < 4 {
            omittable = append(omittable, slice[1])
        }
    }

    // Get all sequences
    prev := omittable[0] - 1
    omitSequences := make([][]int, 1)
    currentRow := 0
    for _, value := range omittable {
        if value != prev+1 {
            currentRow++
            omitSequences = append(omitSequences, make([]int, 0))
        }
        omitSequences[currentRow] = append(omitSequences[currentRow], value)
        prev = value
    }

    // Loop over sequences and search for valid permutations
    count := float64(1)
    for _, sequence := range omitSequences {
        count *= CountPermutations(sequence)
    }
    return count
}

func CountPermutations(sequence []int) float64 {
    // Because all these are omitabble, we know that start = [0] - 1, and end = [length - 1] + 1
    sequence = append([]int{sequence[0] - 1}, sequence...)
    sequence = append(sequence, sequence[len(sequence)-1]+1)

    return GetPossibleBranches(sequence[0], sequence)
}

func GetPossibleBranches(value int, sequence []int) float64 {
    last := sequence[len(sequence)-1]
    if value == last {
        return 1
    }

    count := float64(0)
    for i := 1; i < 4; i++ {
        next := value + i
        if ContainsInt(sequence, next) {
            count += GetPossibleBranches(next, sequence)
        }
    }
    return count
}

func ContainsInt(sequence []int, value int) bool {
    for _, val := range sequence {
        if val == value {
            return true
        }
    }
    return false
}
