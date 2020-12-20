package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	split := strings.Split(PuzzleInput, "\n")
	count := 0
	for _, line := range split {
		expression := CreateExpression(line)
		count += expression.Solve()
	}
	fmt.Printf("Part 1: Sum of all math problems is %v.\n", count)

	countV2 := 0
	for _, line := range split {
		expression := CreateExpression(line)
		countV2 += expression.SolveV2()
	}
	fmt.Printf("Part 2: Sum of all math problems is %v.\n", countV2)
}

type Expression struct {
	originalValue, currentValue string
}

func (expression *Expression) Solve() int {
	if !strings.ContainsAny(expression.currentValue, "+*") {
		result, _ := strconv.Atoi(expression.currentValue)
		return result
	}

	currentProblem := expression.GetCurrentProblem()

	solution := 0

	if len(GetOperators(currentProblem)) > 1 {
		subExpression := CreateExpression(strings.Trim(currentProblem, "()"))
		solution = subExpression.Solve()
		expression.currentValue = strings.Replace(expression.currentValue, currentProblem, fmt.Sprint(solution), 1)
		return expression.Solve()
	}

	operator := GetOperator(currentProblem)
	operands := strings.Split(currentProblem, operator)
	op1, _ := strconv.Atoi(strings.Trim(operands[0], " ()"))
	op2, _ := strconv.Atoi(strings.Trim(operands[1], " ()"))

	if operator == "+" {
		solution = op1 + op2
	} else {
		solution = op1 * op2
	}
	expression.currentValue = strings.Replace(expression.currentValue, currentProblem, fmt.Sprint(solution), 1)

	return expression.Solve()
}

func (expression *Expression) GetCurrentProblem() string {
	if strings.ContainsAny(expression.currentValue, "()") {
		start := -1
		end := -1
		for i, val := range expression.currentValue {
			if val == '(' {
				start = i
			}
			if val == ')' {
				end = i
				return expression.currentValue[start : end+1]
			}
		}
	}

	operators := GetOperators(expression.currentValue)
	currentProblem := ""

	if len(operators) > 1 {
		currentProblem = expression.currentValue[0:operators[1]]
	} else {
		currentProblem = expression.currentValue
	}
	return currentProblem
}

func GetOperators(value string) []int {
	operators := make([]int, 0)
	for i, r := range value {
		if r == '*' || r == '+' {
			operators = append(operators, i)
		}
	}
	return operators
}

func GetOperator(value string) string {
	for _, r := range value {
		if r == '*' || r == '+' {
			return string(r)
		}
	}
	return ""
}

func CreateExpression(input string) Expression {
	return Expression{currentValue: input, originalValue: input}
}

func (expression *Expression) SolveV2() int {
	if !strings.ContainsAny(expression.currentValue, "+*") {
		result, _ := strconv.Atoi(expression.currentValue)
		return result
	}

	currentProblem := expression.GetCurrentProblemV2()

	solution := 0

	if len(GetOperators(currentProblem)) > 1 {
		subExpression := CreateExpression(strings.Trim(currentProblem, "()"))
		solution = subExpression.SolveV2()
		expression.currentValue = strings.Replace(expression.currentValue, currentProblem, fmt.Sprint(solution), 1)
		return expression.SolveV2()
	}

	operator := GetOperator(currentProblem)
	operands := strings.Split(currentProblem, operator)
	op1, _ := strconv.Atoi(strings.Trim(operands[0], " ()"))
	op2, _ := strconv.Atoi(strings.Trim(operands[1], " ()"))

	if operator == "+" {
		solution = op1 + op2
	} else {
		solution = op1 * op2
	}
	expression.currentValue = strings.Replace(expression.currentValue, currentProblem, fmt.Sprint(solution), 1)

	return expression.SolveV2()
}

func (expression *Expression) GetCurrentProblemV2() string {
	if strings.ContainsAny(expression.currentValue, "()") {
		start := -1
		end := -1
		for i, val := range expression.currentValue {
			if val == '(' {
				start = i
			}
			if val == ')' {
				end = i
				return expression.currentValue[start : end+1]
			}
		}
	}

	operatorsMap := GetOperatorsV2(expression.currentValue)

	if len(operatorsMap) > 1 {
		firstAdditionIndex := GetFirstAddition(operatorsMap)
		operatorPositions := GetOperators(expression.currentValue)
		if firstAdditionIndex < 0 {
			return expression.currentValue[0:operatorPositions[1]]
		}

		start, end := GetIndices(operatorPositions, firstAdditionIndex)
		if start < 0 && end < 0 {

		}
		if start < 0 {
			return expression.currentValue[:end]
		}
		if end < 0 {
			return expression.currentValue[start+1:]
		}
		return expression.currentValue[start+1 : end]
	}
	return expression.currentValue
}

func GetIndices(positions []int, index int) (int, int) {
	for i, position := range positions {
		if position == index {
			if i == 0 {
				return -1, positions[i+1]
			}
			if i == len(positions) - 1 {
				return positions[i-1], -1
			}
			return positions[i-1], positions[i+1]
		}
	}
	return -1, -1
}

func GetFirstAddition(operators map[int]rune) int {
	for index, operator := range operators {
		if operator == '+' {
			return index
		}
	}
	return -1
}

func GetOperatorsV2(value string) map[int]rune {
	operators := make(map[int]rune, 0)
	for i, r := range value {
		if r == '*' || r == '+' {
			operators[i] = r
		}
	}
	return operators
}
