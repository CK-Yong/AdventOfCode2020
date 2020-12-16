package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tickets := ParseTickets(strings.Split(TicketInputs, "\n"))
	rules := ParseRules(strings.Split(RuleInputs, "\n"))

	result := ScanForErrors(tickets, rules)
	fmt.Printf("Part 1: Ticket error rate is %v\n", result)

	myTicket := ParseTickets([]string{YourTicket})[0]
	numberOfIds := len(tickets[0].Numbers)
	validTickets := RemoveErrors(tickets, rules)
	validTickets = append(validTickets, myTicket)
	names := SortLabels(validTickets, rules, numberOfIds)

	departureFields := make([]int, 0)
	for i, str := range names {
		if strings.HasPrefix(str, "departure"){
			departureFields = append(departureFields, i)
		}
	}

	multiple := 1
	for _, index := range departureFields {
		multiple *= myTicket.Numbers[index]
	}

	fmt.Printf("Part 2: Multiple of six values is %v", multiple)

}

type Ticket struct {
	Numbers []int
}

type Boundary struct {
	Min, Max int
}

type Rule struct {
	name       string
	boundaries []Boundary
}

func ScanForErrors(tickets []Ticket, rules []Rule) int {
	errors := 0
	for _, ticket := range tickets {
		errors += ApplyRule(ticket, rules)
	}
	return errors
}

func RemoveErrors(tickets []Ticket, rules []Rule) []Ticket {
	validTickets := make([]Ticket, 0)
	for _, ticket := range tickets {
		isValid := isTicketValid(ticket, rules)
		if isValid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func isTicketValid(ticket Ticket, rules []Rule) bool {
	for _, num := range ticket.Numbers {
		numIsValid := false
		for _, rule := range rules {
			if numIsValid {
				break
			}

			if isValidNumber(num, rule) {
				numIsValid = true
				break
			}
		}
		if !numIsValid {
			return false
		}
	}
	return true
}

func ApplyRule(ticket Ticket, rules []Rule) int {
	errorRate := 0

	for _, num := range ticket.Numbers {
		numIsValid := false
		for _, rule := range rules {
			if numIsValid {
				break
			}

			for _, boundary := range rule.boundaries {
				if boundary.Min <= num && boundary.Max >= num {
					numIsValid = true
					break
				}
			}
		}
		if !numIsValid {
			errorRate += num
		}
	}

	return errorRate
}

func isValidNumber(number int, rule Rule) bool {
    for _, boundary := range rule.boundaries {
        if boundary.Min <= number && number <= boundary.Max {
            return true
        }
    }
    return false
}

func ParseRules(input []string) []Rule {
	rules := make([]Rule, 0)
	for _, line := range input {
		nameValues := strings.Split(line, ": ")
		boundaryInputs := strings.Split(nameValues[1], " or ")
		boundaries := make([]Boundary, 2)

		for i, boundaryStr := range boundaryInputs {
			numStrs := strings.Split(boundaryStr, "-")
			min, _ := strconv.Atoi(numStrs[0])
			max, _ := strconv.Atoi(numStrs[1])
			boundaries[i] = Boundary{min, max}
		}
		rules = append(rules, Rule{name: nameValues[0], boundaries: boundaries})
	}
	return rules
}

func ParseTickets(input []string) []Ticket {
	tickets := make([]Ticket, 0)

	for _, line := range input {
		split := strings.Split(line, ",")
		numbers := make([]int, 0)
		for _, str := range split {
			number, _ := strconv.Atoi(str)
			numbers = append(numbers, number)
		}
		tickets = append(tickets, Ticket{Numbers: numbers})
	}
	return tickets
}

func SortLabels(tickets []Ticket, rules []Rule, numbersPerTicket int) []string {
	possbilityMap := GetAllPossibilities(tickets, rules, numbersPerTicket)

	for !finishedFiltering(possbilityMap){
		confirmedFields := GetConfirmedFields(possbilityMap)
		possbilityMap = Filter(possbilityMap, confirmedFields)
	}

	result := make([]string, 0)
	for _, possibilities := range possbilityMap {
		result = append(result, GetSingle(possibilities))
	}
	return result
}

func Filter(possbilityMap []map[string]bool, fields []string) []map[string]bool {
	for _, possibilities := range possbilityMap {
		for _, field := range fields {
			if len(possibilities) == 1 {
				continue
			}
			delete(possibilities, field)
		}
	}
	return possbilityMap
}

func GetConfirmedFields(possbilityMap []map[string]bool) []string {
	result := make([]string, 0)
	for _, possibilities := range possbilityMap {
		if len(possibilities) == 1 {
			result = append(result, GetSingle(possibilities))
		}
	}
	return result
}

func finishedFiltering(possbilityMap []map[string]bool) bool {
	for _, possibilities := range possbilityMap {
		if len(possibilities) != 1 {
			return false
		}
	}
	return true
}

func GetSingle(possibilities map[string]bool) string {
	for field := range possibilities {
		return field
	}
	return ""
}

func GetAllPossibilities(tickets []Ticket, rules []Rule, numbersPerTicket int) []map[string]bool {
	possibilities := make([]map[string]bool, len(rules))
	for i := range possibilities {
		possibilities[i] = make(map[string]bool,0)
	}

	for _, rule := range rules {
		for j := 0; j < numbersPerTicket; j++ {
			isValidRule := false
			for i := range tickets {
				number := tickets[i].Numbers[j]
				if !isValidNumber(number, rule) {
					isValidRule = false
					break
				}
				isValidRule = true
			}

			if isValidRule {
				possibilities[j][rule.name] = true
			}
		}
	}
	return possibilities
}
