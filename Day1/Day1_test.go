package main

import (
    "testing"
)

func Test_should_find_numbers_that_sum_to_2020(test *testing.T) {
    input := []int{1721, 979, 366, 299, 675, 1456}

    result := Find2020Sum(input)

    if result.first != 1721 {
        test.Errorf("First result in pair was incorrect, actual: %d, expected: %d", result.first, 1721)
    }

    if result.second != 299 {
        test.Errorf("Second result in pair was incorrect, actual: %d, expected: %d", result.second, 299)
    }
}
