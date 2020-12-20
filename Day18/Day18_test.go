package main

import (
	"testing"
)

func Test_should_evaluate_expression(test *testing.T) {
	var input = "1 + 2 * 3"
	expression := CreateExpression(input)

	result := expression.Solve()

	if result != 9 {
		test.Errorf("Expected solution to be 9. Got %v", result)
	}
}

func Test_should_consider_parentheses(test *testing.T) {
	var input = "1 + (2 * 3)"
	expression := CreateExpression(input)

	result := expression.Solve()

	if result != 7 {
		test.Errorf("Expected solution to be 7. Got %v", result)
	}
}

func Test_should_solve_a_bit_more_difficult_cases(test *testing.T) {
	var input = "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	expression := CreateExpression(input)

	result := expression.Solve()

	if result != 12240 {
		test.Errorf("Expected solution to be 12240. Got %v", result)
	}
}

func Test_should_solve_with_addition_priority(test *testing.T) {
	var input = "2 * 3 + (4 * 5)"
	expression := CreateExpression(input)

	result := expression.SolveV2()

	if result != 46 {
		test.Errorf("Expected solution to be 46. Got %v", result)
	}
}

func Test_should_solve_with_addition_priority_2(test *testing.T) {
	var input = "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	expression := CreateExpression(input)

	result := expression.SolveV2()

	if result != 669060 {
		test.Errorf("Expected solution to be 669060. Got %v", result)
	}
}
