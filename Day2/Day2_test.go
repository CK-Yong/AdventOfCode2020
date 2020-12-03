package main

import (
    "testing"
)

var validPasswords = []struct {
    Policy   Policy
    Password string
}{
    {Policy{1, 3, "a"}, "abcde"},
    {Policy{2, 9, "c"}, "ccccccccc"},
}

func Test_should_be_valid(test *testing.T) {
    for _, testCase := range validPasswords {
        isValid := testCase.Policy.Validate(testCase.Password)

        if !isValid {
            test.Errorf("Expected password to be valid but got %v", isValid)
        }
    }
}

func Test_should_be_invalid(test *testing.T) {
    testCase := struct{ Policy Policy; Password string }{Policy{1, 3, "b"}, "cdefg"}
    isValid := testCase.Policy.Validate(testCase.Password)

    if isValid {
        test.Errorf("Expected password to be invalid but got %v", isValid)
    }
}

