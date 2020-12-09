package main

import (
    "testing"
)

func Test_should_find_first_invalid_number(test *testing.T) {
    input := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
    converted := ParseInput(input)

    sequence := Sequence{converted, 5}
    weakness, _, _ := sequence.FindWeakness()

    if weakness != 127 {
        test.Errorf("Expected weakness number to be 127. Got %v", weakness)
    }
}

func Test_should_find_the_two_parts_of_the_weak_number(test *testing.T) {
    input := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
    converted := ParseInput(input)

    sequence := Sequence{converted, 5}
    weakness, _, _ := sequence.FindWeakness()

    if weakness != 127 {
        test.Errorf("Expected weakness number to be 127. Got %v", weakness)
    }
}
