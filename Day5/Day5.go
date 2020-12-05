package main

import (
    "fmt"
    "sort"
    "strings"
)

type BoardingPass struct {
    value           string
    Id, Row, Column int
}

func (b *BoardingPass) Init(value string) {
    b.Row = calculateRow(value)
    b.Column = calculateColumn(value)
    b.Id = b.Row*8 + b.Column
}

func calculateColumn(value string) int {
    lowerBoundary := 0
    upperBoundary := 7
    for i := 7; i < len(value); i++ {
        if value[i] == 'L' {
            upperBoundary = (upperBoundary-lowerBoundary)/2 + lowerBoundary
        } else {
            lowerBoundary = ((upperBoundary-lowerBoundary)/2 + lowerBoundary) + 1
        }
    }
    return lowerBoundary
}

func calculateRow(value string) int {
    lowerBoundary := 0
    upperBoundary := 127
    for i := 0; i < 7; i++ {
        if value[i] == 'F' {
            upperBoundary = (upperBoundary-lowerBoundary)/2 + lowerBoundary
        } else {
            lowerBoundary = ((upperBoundary-lowerBoundary)/2 + lowerBoundary) + 1
        }
    }
    return lowerBoundary
}

func main() {
    inputs := strings.Split(PuzzleInput, "\n")
    var currentId int
    var listOfIds []int
    for _, input := range inputs {
        pass := new(BoardingPass)
        pass.Init(input)
        if pass.Id > currentId {
            currentId = pass.Id
        }
        listOfIds = append(listOfIds, pass.Id)
    }

    fmt.Printf("Part 1: Highest ID is %v\n", currentId)

    sort.Ints(listOfIds)

    prevId := listOfIds[0]
    for _, currentId := range listOfIds {
        if prevId + 1 == currentId {
            break
        }
        prevId++
    }

    fmt.Printf("Part 2: My seat ID is %v\n", prevId)
}
