package main

import (
	"strings"
	"testing"
)

func Test_should_find_manhattan_distance(test *testing.T){
	var input = "F10\nN3\nF7\nR90\nF11"
	parsed := strings.Split(input, "\n")

	ship := Ship{}
	ship.Traverse(parsed)

	manhattan := ship.GetManhattanDistance()

	if manhattan != 25 {
		test.Errorf("Expected Manhattan distance of 25. Got: %v", manhattan)
	}
}

func Test_should_find_manhattan_distance_after_waypoint_navigation(test *testing.T){
	var input = "F10\nN3\nF7\nR90\nF11"
	parsed := strings.Split(input, "\n")

	ship := CreateShip()
	ship.TraverseUsingWaypoint(parsed)

	manhattan := ship.GetManhattanDistance()

	if manhattan != 286 {
		test.Errorf("Expected Manhattan distance of 286. Got: %v", manhattan)
	}
}