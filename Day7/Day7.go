package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    parsed := strings.Split(PuzzleInput, "\n")

    bags := ParseBags(parsed)
    noOfBags := CalculateBagsThatCanHold(bags, "shiny gold")

    fmt.Printf("Part 1: Found %v bags.\n", noOfBags)

    noOfBags = CalculateTotalBagsRequired(bags, "shiny gold")
    fmt.Printf("Part 2: Found %v bags.", noOfBags-1) // Includes the 1 shiny gold bag
}

type Bag struct {
    Color      string
    Capacity   map[string]int
    wasCounted bool
}

func ParseBags(input []string) []Bag {
    var bags = make([]Bag, 0)

    for _, entry := range input {
        commasRemoved := strings.Replace(entry, ",", "", -1)
        dotsRemoved := strings.Replace(commasRemoved, ",", "", -1)
        split := strings.Split(dotsRemoved, " ")
        bag := Bag{strings.Join(split[0:2], " "), make(map[string]int), false}

        if strings.Contains(entry, "contain no other bags") {
            bags = append(bags, bag)
            continue
        }

        for i := 4; i < len(split); i += 4 {
            color := strings.Join(split[i+1:i+3], " ")
            bag.Capacity[color], _ = strconv.Atoi(split[i])
        }

        bags = append(bags, bag)
    }

    return bags
}

func CalculateBagsThatCanHold(input []Bag, color string) int {
    count := 0
    for i, bag := range input {
        if bag.CanHold(color) {
            if !bag.wasCounted {
                count++
                input[i].wasCounted = true
            }
            count += CalculateBagsThatCanHold(input, bag.Color)
        }
    }

    return count
}

func (b *Bag) CanHold(color string) bool {
    if number, ok := b.Capacity[color]; ok {
        return number > 0
    }
    return false
}

func CalculateTotalBagsRequired(input []Bag, color string) int {
    count := 1
    for _, bag := range input {
        if bag.Color == color {
            for holdsColor, holdsNumber := range bag.Capacity {
                count += holdsNumber * CalculateTotalBagsRequired(input, holdsColor)
            }
        }
    }

    return count
}
