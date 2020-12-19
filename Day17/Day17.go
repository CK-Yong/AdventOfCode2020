package main

import (
	"fmt"
	"math"
	"strings"
)

var PuzzleInput = "#.#..###\n.#....##\n.###...#\n..####..\n....###.\n##.#.#.#\n..#..##.\n#.....##"

func main(){
	split := strings.Split(PuzzleInput, "\n")
	parsed := make([][]string, len(split))
	for i, line := range split {
		parsed[i] = strings.Split(line, "")
	}

	grid := CreateGrid(parsed)

	grid.Update(6)
	count := grid.CountActiveCubes()

	fmt.Printf("Part 1: %v active cubes after 6 cycles", count)
}

type Cube struct {
	X, Y, Z int
}

type Grid struct {
	ActiveCubes map[Cube]struct{}
	prevState   map[Cube]struct{}
}

func CreateGrid(input [][]string) Grid {
	// coordinates will be in z, y, x
	cubes := make(map[Cube]struct{}, 0)
	for i, row := range input {
		for j, col := range row {
			if col == "#" {
				cube := Cube{j, i, 0}
				cubes[cube] = struct{}{}
			}
		}
	}

	return Grid{cubes, clone(cubes)}
}

func clone(cubes map[Cube]struct{}) map[Cube]struct{} {
	prevState := make(map[Cube]struct{}, len(cubes))
	for k, v := range cubes {
		prevState[k] = v
	}
	return prevState
}

func (grid Grid) Update(times int) {
	for i := 0; i < times; i++ {
		boundaries := grid.GetBoundaries()

		for z := boundaries.Z.Min - 1; z < boundaries.Z.Max+2; z++ {
			for y := boundaries.Y.Min - 1; y < boundaries.Y.Max+2; y++ {
				for x := boundaries.X.Min - 1; x < boundaries.X.Max+2; x++ {
					currentCube := Cube{x, y, z}
					_, currentCubeIsActive := grid.ActiveCubes[currentCube]
					neighbours := grid.GetActiveNeighbours(z, y, x)

					if currentCubeIsActive && (len(neighbours) < 2 || len(neighbours) > 3) {
						delete(grid.ActiveCubes, currentCube)
						continue
					}

					if !currentCubeIsActive && (len(neighbours) == 3) {
						grid.ActiveCubes[currentCube] = struct{}{}
					}
				}
			}
		}
		
		grid.prevState = clone(grid.ActiveCubes)
	}
}

func (grid Grid) GetActiveNeighbours(z int, y int, x int) []Cube {
	neighbours := make([]Cube, 0)
	for i := z - 1; i < z+2; i++ {
		for j := y - 1; j < y+2; j++ {
			for k := x - 1; k < x+2; k++ {
				if i == z && j == y && k == x {
					// This is the cube being evaluated
					continue
				}

				cube := Cube{k, j, i}
				_, isActive := grid.prevState[cube]
				if isActive {
					neighbours = append(neighbours, cube)
				}
			}
		}
	}
	return neighbours
}

func (grid Grid) GetBoundaries() Boundaries {
	zMin := math.MaxInt32
	zMax := math.MinInt32
	yMin := math.MaxInt32
	yMax := math.MinInt32
	xMin := math.MaxInt32
	xMax := math.MinInt32
	for cube := range grid.ActiveCubes {
		zMin = int(math.Min(float64(zMin), float64(cube.Z)))
		zMax = int(math.Max(float64(zMax), float64(cube.Z)))
		yMin = int(math.Min(float64(yMin), float64(cube.Y)))
		yMax = int(math.Max(float64(yMax), float64(cube.Y)))
		xMin = int(math.Min(float64(xMin), float64(cube.X)))
		xMax = int(math.Max(float64(xMax), float64(cube.X)))
	}
	return Boundaries{Limit{zMin, zMax}, Limit{yMin, yMax}, Limit{xMin, xMax}}
}

type Limit struct {
	Min, Max int
}

type Boundaries struct {
	Z, Y, X Limit
}

func (grid Grid) CountActiveCubes() int {
	return len(grid.ActiveCubes)
}
