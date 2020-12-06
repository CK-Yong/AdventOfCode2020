package main

import (
	"testing"
)

func Test_should_split_value_into_groups(test *testing.T) {
	var input = "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	var expectations = []int{1, 3, 2, 4, 1} // number of people in groups, there are 5 groups with x people

	var yesCounter = new(YesCounter)
	yesCounter.Init(input)

	for i, numberOfPeople := range expectations {
		actual := yesCounter.Groups[i].People
		if actual != numberOfPeople {
			test.Errorf("Expected group %v to have %v people. Got: %v", i, numberOfPeople, actual)
		}
	}
}

func Test_should_count_the_total_number_of_yes_answers(test *testing.T) {
	var input = "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	var expectation = 11

	var yesCounter = new(YesCounter)
	yesCounter.Init(input)

	var result = yesCounter.CountAnswers()
	if result != expectation {
		test.Errorf("Expected total answers to be 11. Got: %v", result)
	}
}

func Test_should_count_the_total_number_of_common_yes_answers(test *testing.T) {
	var input = "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	var expectation = 6

	var yesCounter = new(YesCounter)
	yesCounter.Init(input)

	var result = yesCounter.CountCommonAnswers()
	if result != expectation {
		test.Errorf("Expected total answers to be 6. Got: %v", result)
	}
}
