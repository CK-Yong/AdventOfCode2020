package main

import "fmt"

type Toboggan struct {
    Count, StepsRight, StepsDown int
    Forest                       []string
    Position                     Position
}

type Position struct {
    X, Y int
}

func (toboggan *Toboggan) Traverse() {
    toboggan.Position.X += toboggan.StepsRight
    toboggan.Position.Y += toboggan.StepsDown

    if toboggan.Position.Y >= len(toboggan.Forest) {
        return
    }

    currentRow := toboggan.Forest[toboggan.Position.Y]
    for toboggan.Position.X >= len(currentRow) {
        currentRow += currentRow
    }

    var nextSquare = rune(currentRow[toboggan.Position.X])
    if nextSquare == '#' {
        toboggan.Count++
    }

    toboggan.Traverse()
}

func main() {
    traveller := Toboggan{StepsRight: 3, StepsDown: 1, Forest: PuzzleInput}

    traveller.Traverse()

    fmt.Printf("Part 1: %v trees encountered\n", traveller.Count)

    travellers := []Toboggan{
        {StepsRight: 1, StepsDown: 1, Forest: PuzzleInput},
        {StepsRight: 3, StepsDown: 1, Forest: PuzzleInput},
        {StepsRight: 5, StepsDown: 1, Forest: PuzzleInput},
        {StepsRight: 7, StepsDown: 1, Forest: PuzzleInput},
        {StepsRight: 1, StepsDown: 2, Forest: PuzzleInput},
    }

    for i := 0; i < len(travellers); i++{
        travellers[i].Traverse()
    }

    count := 1
    for i := 0; i < len(travellers); i++ {
        count *= travellers[i].Count
    }

    fmt.Printf("Part 2: %v trees encountered\n", count)
}
