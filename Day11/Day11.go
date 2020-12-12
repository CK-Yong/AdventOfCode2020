package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
)

func main()  {
	parsed := strings.Split(PuzzleInput, "\n")

	room := ParseRoom(parsed)

	room.Update(1)

	for !room.HasStabilized() {
		room.Update(1)
	}

	occupiedSeats := room.CountOccupiedSeats()
	fmt.Printf("Part 1: Found %v occupied seats.", occupiedSeats)
}

type Cell struct {
	Occupied, IsSeat bool
	Row, Column      int
}

// Adjacent cell needs to be added from top left corner, going clockwise
func (cell *Cell) Update(adjacentCells []Cell) {
	if !cell.IsSeat {
		return
	}

	if !cell.Occupied && CountOccupiedSeats(adjacentCells) == 0 {
		cell.Occupied = true
		return
	}

	if cell.Occupied && CountOccupiedSeats(adjacentCells) >= 4 {
		cell.Occupied = false
	}
}

func (cell *Cell) LookAround(hall Hall){

}

type Hall struct {
	CurrentState  [][]Cell
	previousState [][]Cell
}

func (hall *Hall) Update(times int) {
	for i := 0; i < times; i++ {
		hall.previousState = Clone(hall.CurrentState)
		for i := range hall.CurrentState {
			for j := range hall.CurrentState[i] {
				cells := hall.GetAdjacentCellsFor(i, j)
				hall.CurrentState[i][j].Update(cells)
			}
		}
	}
}

func (hall Hall) CountOccupiedSeats() int {
	count := 0
	for _, row := range hall.CurrentState {
		for _, cell := range row {
			if cell.Occupied {
				count++
			}
		}
	}
	return count
}

func (hall Hall) GetAdjacentCellsFor(row int, column int) []Cell {
	adjacent := make([]Cell, 0)
	for i := row - 1; i < row+2; i++ {
		for j := column - 1; j < column+2; j++ {
			if i == row && j == column {
				continue
			}
			if i < 0 || i >= len(hall.previousState) {
				continue
			}
			if j < 0 || j >= len(hall.previousState[i]) {
				continue
			}
			adjacent = append(adjacent, hall.previousState[i][j])
		}
	}
	return adjacent
}

func (hall Hall) HasStabilized() bool {
	for i := range hall.CurrentState{
		for j := range hall.CurrentState[i]{
			if !cmp.Equal(hall.CurrentState[i][j], hall.previousState[i][j]){
				return false
			}
		}
	}
	return true
}

func CreateCell(input rune, row, column int) Cell {
	if input == '.' {
		return Cell{IsSeat: false, Row: row, Column: column}
	} else {
		return Cell{IsSeat: true, Row: row, Column: column}
	}
}

func ParseRoom(parsed []string) Hall {
	cells := make([][]Cell, 0)
	for i, rowInput := range parsed {
		cells = append(cells, make([]Cell, 0))
		for j, cellInput := range rowInput {
			cells[i] = append(cells[i], CreateCell(cellInput, i, j))
		}
	}

	return Hall{Clone(cells), Clone(cells)}
}

func CountOccupiedSeats(adjacentCells []Cell) int {
	count := 0
	for _, val := range adjacentCells {
		if val.IsSeat && val.Occupied {
			count++
		}
	}
	return count
}

func Clone(cells [][]Cell) [][]Cell {
	clone := make([][]Cell, len(cells))
	for i := range cells {
		clone[i] = make([]Cell, len(cells[i]))
		copy(clone[i], cells[i])
	}
	return clone
}
