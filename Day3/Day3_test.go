package main

import (
    "testing"
)

var testInput = []string{
    "..##.......",
    "#...#...#..",
    ".#....#..#.",
    "..#.#...#.#",
    ".#...##..#.",
    "..#.##.....",
    ".#.#.#....#",
    ".#........#",
    "#.##...#...",
    "#...##....#",
    ".#..#...#.#",
}

var forestWithTree = []string{
    "....",
    "...#",
}

func Test_should_encounter_tree(test *testing.T) {
    var traveller = Toboggan{StepsRight: 3, StepsDown: 1, Forest: forestWithTree}

    traveller.Traverse()

    if traveller.Count != 1 {
        test.Errorf("Expected 7 trees to be encountered, got %v", traveller.Count)
    }
}

var forestWithoutTree = []string{
    "####",
    "###.",
}

func Test_should_not_encounter_tree(test *testing.T) {
    var traveller = Toboggan{StepsRight: 3, StepsDown: 1, Forest: forestWithoutTree}

    traveller.Traverse()

    if traveller.Count != 0 {
        test.Errorf("Expected 7 trees to be encountered, got %v", traveller.Count)
    }
}

func Test_should_encounter_7_trees(test *testing.T) {
    var traveller = Toboggan{StepsRight: 3, StepsDown: 1, Forest: testInput}

    traveller.Traverse()

    if traveller.Count != 7 {
        test.Errorf("Expected 7 trees to be encountered, got %v", traveller.Count)
    }
}
