package main

import (
	"strings"
	"testing"
)

func Test_should_result_in_10th_number(test *testing.T) {
	input := "0,3,6"
	parsed := strings.Split(input, ",")

	game := CreateGame(parsed)

	result := game.PlayUntil(10)

	if result != 0 {
		test.Errorf("Expected 10th number to be 0. Got: %v", result)
	}
}

func Test_should_result_in_2020th_number(test *testing.T) {
	input := "2,1,3"
	parsed := strings.Split(input, ",")

	game := CreateGame(parsed)

	result := game.PlayUntil(2020)

	if result != 10 {
		test.Errorf("Expected 2020th number to be 10. Got: %v", result)
	}
}

func Test_should_result_in_30000000th_number(test *testing.T) {
	input := "2,1,3"
	parsed := strings.Split(input, ",")

	game := CreateGame(parsed)

	result := game.PlayUntil(30000000)

	if result != 3544142 {
		test.Errorf("Expected 10th number to be 0. Got: %v", result)
	}
}
