package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main(){
    tickets := ParseTickets(strings.Split(TicketInputs, "\n"))
    rules := ParseRules(strings.Split(RuleInputs, "\n"))

    result := ScanForErrors(tickets, rules)
    fmt.Printf("Part 1: Ticket error rate is %v\n", result)

    validTickets := RemoveErrors(tickets, rules)
    
}

type Ticket struct {
    Numbers []int
    isValid bool
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
        errors := ApplyRule(ticket, rules)
        if errors == 0 {
            validTickets = append(validTickets, ticket)
        }
    }
    return validTickets
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
