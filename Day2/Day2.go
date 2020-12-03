package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Policy struct {
    Min, Max int
    Value    string
}

func (policy Policy) Validate(password string) bool {
    count := strings.Count(password, policy.Value)
    if count >= policy.Min && count <= policy.Max {
        return true
    }
    return false
}


type PolicyV2 struct {
    Position1, Position2 int
    Value    rune
}

func (policy PolicyV2) Validate(password string) bool {
    firstPositionMatches := rune(password[policy.Position1 - 1]) == policy.Value
    secondPositionMatches := rune(password[policy.Position2 - 1]) == policy.Value

    if firstPositionMatches && secondPositionMatches {
        return false
    }

    if !firstPositionMatches && !secondPositionMatches {
        return false
    }

    return true
}

func IsDashOrWhitespace(input rune) bool {
    return input == '-' || input == ' '
}

func main() {
    count := 0

    for _, entry := range PuzzleInput {
        parsed := strings.Split(entry, ": ")

        policyInput := strings.FieldsFunc(parsed[0], IsDashOrWhitespace)

        min, _ := strconv.Atoi(policyInput[0])
        max, _ := strconv.Atoi(policyInput[1])

        policy := Policy{min, max, policyInput[2]}

        isValid := policy.Validate(parsed[1])

        if isValid { count++ }
    }

    fmt.Printf("Part 1: Counted valid passwords: %v\n", count)
    // Result 1: 467

    count = 0
    for _, entry := range PuzzleInput {
        parsed := strings.Split(entry, ": ")

        policyInput := strings.FieldsFunc(parsed[0], IsDashOrWhitespace)

        first, _ := strconv.Atoi(policyInput[0])
        second, _ := strconv.Atoi(policyInput[1])

        policy := PolicyV2{first, second, []rune(policyInput[2])[0]}

        isValid := policy.Validate(parsed[1])

        if isValid { count++ }
    }

    fmt.Printf("Part 2: Counted valid passwords: %v", count)
    // Result 2: 441
}
