package main

import (
	"strings"
	"testing"
)

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

func Test_should_result_in_21_cubes_after_2_cycle(test *testing.T) {
	input := ".#.\n..#\n###"

	split := strings.Split(input, "\n")
	parsed := make([][]string, len(split))
	for i, line := range split {
		parsed[i] = strings.Split(line, "")
	}

	grid := CreateGrid(parsed)

	grid.Update(2)
	count := grid.CountActiveCubes()

	if count != 21 {
		test.Errorf("Expected 21 cubes to be active after 1 cycle. Got: %v", count)
	}
}

func Test_should_result_in_112_cubes_after_3_cycles(test *testing.T) {
	input := ".#.\n..#\n###"

	split := strings.Split(input, "\n")
	parsed := make([][]string, len(split))
	for i, line := range split {
		parsed[i] = strings.Split(line, "")
	}

	grid := CreateGrid(parsed)

	grid.Update(6)
	count := grid.CountActiveCubes()

	if count != 112 {
		test.Errorf("Expected 112 cubes to be active after 3 cycle. Got: %v", count)
	}
}
