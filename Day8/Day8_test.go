package main

import (
    "testing"
)

func Test_should_return_accumulator_5(test *testing.T) {
    input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
    commands := ParseCommands(input)

    traverser := Traverser{BootCode: commands}

    traverser.Traverse()

    if traverser.Accumulator != 5 {
        test.Errorf("Expected accumulator to be 5. Got: %v", traverser.Accumulator)
    }
}

func Test_should_fix_their_own_program(test *testing.T) {
    input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
    commands := ParseCommands(input)

    traverser := GetTraverserForTraversibleRoute(commands)

    if traverser.Accumulator != 8 {
        test.Errorf("Expected accumulator to be 8. Got: %v", traverser.Accumulator)
    }
}

