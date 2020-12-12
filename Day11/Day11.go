package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
)

func main() {
	parsed := strings.Split(PuzzleInput, "\n")

	room := ParseRoom(parsed, 4)

	room.Update(1)
	for !room.HasStabilized() {
		room.Update(1)
	}

	occupiedSeats := room.CountOccupiedSeats()
	fmt.Printf("Part 1: Found %v occupied seats.\n", occupiedSeats)

	room = ParseRoom(parsed, 5)
	room.UpdateWithLookaround(1)

	for !room.HasStabilized(){
		room.UpdateWithLookaround(1)
	}
	occupiedSeats = room.CountOccupiedSeats()
	fmt.Printf("Part 2: Found %v occupied seats.", occupiedSeats)
}

type Cell struct {
	Occupied, IsSeat bool
	Row, Column      int
}

// Adjacent cell needs to be added from top left corner, going clockwise
func (cell *Cell) Update(adjacentCells []Cell, leniency int) {
	if !cell.IsSeat {
		return
	}

	if !cell.Occupied && CountOccupiedSeats(adjacentCells) == 0 {
		cell.Occupied = true
		return
	}

	if cell.Occupied && CountOccupiedSeats(adjacentCells) >= leniency {
		cell.Occupied = false
	}
}

func (cell *Cell) LookAround(hall [][]Cell) []Cell {
	cells := make([]Cell, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j ==0 {
				continue
			}
			result := cell.LookAt(hall, i, j)
			if result.IsSeat {
				cells = append(cells, result)
			}
		}
	}
	return cells
}

func (cell *Cell) LookAt(hall [][]Cell, rowOffset int, colOffset int) Cell {
	if cell.Row+rowOffset < 0 || cell.Row+rowOffset >= len(hall) {
		return Cell{Row: -1, Column: -1}
	}

	if cell.Column+colOffset < 0 || cell.Column+colOffset >= len(hall[cell.Row]) {
		return Cell{Row: -1, Column: -1}
	}

	target := hall[cell.Row+rowOffset][cell.Column+colOffset]
	if target.IsSeat {
		return target
	}
	return target.LookAt(hall, rowOffset, colOffset)
}

type Hall struct {
	CurrentState  [][]Cell
	previousState [][]Cell
	Leniency      int
}

func (hall *Hall) Update(times int) {
	for i := 0; i < times; i++ {
		hall.previousState = Clone(hall.CurrentState)
		for i := range hall.CurrentState {
			for j := range hall.CurrentState[i] {
				cells := hall.GetAdjacentCellsFor(i, j)
				hall.CurrentState[i][j].Update(cells, hall.Leniency)
			}
		}
	}
}

func (hall *Hall) UpdateWithLookaround(times int) {
	for i := 0; i < times; i++ {
		hall.previousState = Clone(hall.CurrentState)
		for i := range hall.CurrentState {
			for j := range hall.CurrentState[i] {
				cells := hall.CurrentState[i][j].LookAround(hall.previousState)
				hall.CurrentState[i][j].Update(cells, hall.Leniency)
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
	for i := range hall.CurrentState {
		for j := range hall.CurrentState[i] {
			if !cmp.Equal(hall.CurrentState[i][j], hall.previousState[i][j]) {
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

func ToString(room Hall) string {
	str := ""
	for _, row := range room.CurrentState {
		for _, seat := range row {
			if !seat.IsSeat {
				str += "."
				continue
			}

			if seat.Occupied {
				str += "#"
				continue
			}

			if !seat.Occupied {
				str += "L"
				continue
			}
		}
		str += "\n"
	}
	return str
}

func PrevToString(room Hall) string {
	str := ""
	for _, row := range room.previousState {
		for _, seat := range row {
			if !seat.IsSeat {
				str += "."
				continue
			}

			if seat.Occupied {
				str += "#"
				continue
			}

			if !seat.Occupied {
				str += "L"
				continue
			}
		}
		str += "\n"
	}
	return str
}


func ParseRoom(parsed []string, leniency int) Hall {
	cells := make([][]Cell, 0)
	for i, rowInput := range parsed {
		cells = append(cells, make([]Cell, 0))
		for j, cellInput := range rowInput {
			cells[i] = append(cells[i], CreateCell(cellInput, i, j))
		}
	}

	return Hall{Clone(cells), Clone(cells), leniency}
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
