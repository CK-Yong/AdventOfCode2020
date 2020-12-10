package main

import (
	"reflect"
	"strings"
	"testing"
)


func Test_should_traverse_all_adapters_and_give_back_difference(test *testing.T) {
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	parsed := strings.Split(input, "\n")

	adapterSet := AdapterSet{map[int]int{1: 0, 2: 0, 3: 0}}
	adapterSet.Traverse(parsed)

	actual := adapterSet.Values
	expected := map[int]int{1: 7, 2: 0, 3: 5}
	if !reflect.DeepEqual(actual, expected) {
		test.Errorf("Expected adapterset to be equal to {1: 7, 2: 0, 3: 3}. Got: %v", actual)
	}
}

func Test_should_find_all_possible_arrangements(test *testing.T) {
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	parsed := strings.Split(input, "\n")

	adapterSet := AdapterSet{map[int]int{1: 0, 2: 0, 3: 0}}
	count := adapterSet.CountPossibleArrangements(parsed)

	if count != 8 {
		test.Errorf("Expected adapterset to find 8 arrangements. Got: %v", count)
	}
}
