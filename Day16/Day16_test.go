package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_should_result_in_error_rate_of_71(test *testing.T) {
	input := "7,3,47\n40,4,50\n55,2,20\n38,6,12"
	tickets := ParseTickets(strings.Split(input, "\n"))

	ruleInput := "class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50"
	rules := ParseRules(strings.Split(ruleInput, "\n"))

	result := ScanForErrors(tickets, rules)

	if result != 71 {
		test.Errorf("Expected error rate to be 71. Got %v", result)
	}
}

func Test_should_result_in_row_class_seat(test *testing.T) {
	input := "11,12,13\n3,9,18\n15,1,5\n5,14,9"
	tickets := ParseTickets(strings.Split(input, "\n"))

	ruleInput := "class: 0-1 or 4-19\nrow: 0-5 or 8-19\nseat: 0-13 or 16-19"
	rules := ParseRules(strings.Split(ruleInput, "\n"))

	names := SortLabels(tickets, rules, len(tickets[0].Numbers))

	expected := []string{"row", "class", "seat"}
	if !reflect.DeepEqual(names, expected) {
		test.Errorf("Expected error rate to be %v. Got %v", expected, names)
	}
}

