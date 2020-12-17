package main

import (
	"strings"
	"testing"
)

type Cube struct {
	X, Y, Z  int
	isActive bool
}

type Grid struct {
	Cubes     [][][]Cube
	prevState [][][]Cube
}

func CreateGrid(input [][]string) Grid {
	// coordinates will be in z, y, x
	cubes := make([][][]Cube, 1)
	cubes[0] = make([][]Cube, len(input))
	for i, row := range input {
		cubes[0][i] = make([]Cube, len(row))
		for j, col := range row {
			cubes[0][i][j] = Cube{j, i, 0, col == "#"}
		}
	}

	return Grid{cubes, clone(cubes)}
}

func clone(cubes [][][]Cube) [][][]Cube {
	prevState := make([][][]Cube, len(cubes))
	prevState[0] = make([][]Cube, len(cubes[0][0]))

	for z, zPlane := range cubes {
		prevState[z] = make([][]Cube, len(cubes[0][0]))
		for y := range zPlane {
			prevState[z][y] = make([]Cube, len(cubes[0][0]))
			copy(prevState[z][y], cubes[z][y])
		}
	}
	return prevState
}

func (grid *Grid) Update(times int) {
	grid.Expand()
	grid.prevState = clone(grid.Cubes)


}

func (grid Grid) CountActiveCubes() int {
	count := 0
	for _, z := range grid.Cubes {
		for _, y := range z {
			for _, cube := range y {
				if cube.isActive {
					count++
				}
			}
		}
	}
	return count
}

func (grid *Grid) Expand() {
	length := len(grid.Cubes[0])
	// Expand z-plane
	grid.Cubes = append(make([][][]Cube, 1), grid.Cubes...)
	grid.Cubes[0] = make([][]Cube, length+2)
	for i := range grid.Cubes[0] {
		grid.Cubes[0][i] = make([]Cube, length+2)
	}
	grid.Cubes = append(grid.Cubes, make([][]Cube, length+2))
	for i := range grid.Cubes[len(grid.Cubes)-1] {
		grid.Cubes[len(grid.Cubes)-1][i] = make([]Cube, length+2)
	}

	// Expand xy-plane in all z-planes
	for z := range grid.Cubes {
		if len(grid.Cubes[z]) == length {
			grid.Cubes[z] = append(make([][]Cube, 1), grid.Cubes[z]...)
			grid.Cubes[z][0] = make([]Cube, length+2)
			grid.Cubes[z] = append(grid.Cubes[z], make([]Cube, 0))
			grid.Cubes[z][len(grid.Cubes[z])-1] = make([]Cube, length+2)
		}

		// Expand X planes where still short
		for _, zDim := range grid.Cubes {
			for y, row := range zDim {
				if len(row) == length {
					zDim[y] = append(make([]Cube, 1), zDim[y]...)
					zDim[y] = append(zDim[y], Cube{0,0,0, false})
				}
			}
		}
	}

	centerZIndex := len(grid.Cubes)/2
	// Populate all cubes with correct XYZ dimensions
	for z, plane := range grid.Cubes {
		for y, row := range plane {
			for x, cube := range row {
				// Check if this is a default cube and populate coordinates. Do check for the center cube
				if cube.X == 0 && cube.Y == 0 && cube.Z == 0 {
					grid.Cubes[z][y][x] = Cube{x, y, z - centerZIndex, false}
				}
			}
		}
	}
}

func Test_should_result_in_11_cubes_after_1_cycle(test *testing.T) {
	input := ".#.\n..#\n###"

	split := strings.Split(input, "\n")
	parsed := make([][]string, len(split))
	for i, line := range split {
		parsed[i] = strings.Split(line, "")
	}

	grid := CreateGrid(parsed)

	grid.Update(1)
	count := grid.CountActiveCubes()

	if count != 11 {
		test.Errorf("Expected 11 cubes to be active after 1 cycle. Got: %v", count)
	}
}
