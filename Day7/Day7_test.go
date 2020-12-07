package main

import (
    "strings"
    "testing"
)

var input = "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags."
var parsed = strings.Split(input, "\n")
var bags = ParseBags(parsed)

func Test_should_parse_9_bags(test *testing.T) {
    if len(bags) != 9 {
        test.Errorf("Expected 9 bags to be parsed. Got %v.", len(bags))
    }
}

func Test_should_count_4_bags(test *testing.T) {
    noOfBags := CalculateBagsThatCanHold(bags, "shiny gold")

    if noOfBags != 4 {
        test.Errorf("Expected 4 bags to be counted. Got %v.", noOfBags)
    }
}

func Test_should_require_32_bags(test *testing.T) {
    noOfBags := CalculateTotalBagsRequired(bags, "shiny gold")

    if noOfBags - 1 != 32 {
        test.Errorf("Expected 32 bags to be required. Got: %v.", noOfBags)
    }
}

func Test_should_require_126_bags(test *testing.T) {
    input := "shiny gold bags contain 2 dark red bags.\ndark red bags contain 2 dark orange bags.\ndark orange bags contain 2 dark yellow bags.\ndark yellow bags contain 2 dark green bags.\ndark green bags contain 2 dark blue bags.\ndark blue bags contain 2 dark violet bags.\ndark violet bags contain no other bags."
    parsed := strings.Split(input, "\n")
    bags := ParseBags(parsed)

    noOfBags := CalculateTotalBagsRequired(bags, "shiny gold")

    if noOfBags - 1 != 126 {
        test.Errorf("Expected 126 bags to be required. Got: %v.", noOfBags)
    }
}