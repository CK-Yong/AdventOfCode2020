package main

import (
    "testing"
)

func Test_should_calculate_the_row(test *testing.T) {
    pass := new(BoardingPass)
    pass.Init("FBFBBFFRLR")

    if pass.Row != 44 {
        test.Errorf("Expected boarding pass to have row 44. Got: %v", pass.Row)
    }
}

func Test_should_calculate_the_column(test *testing.T) {
    pass := new(BoardingPass)
    pass.Init("FBFBBFFRLR")

    if pass.Column != 5 {
        test.Errorf("Expected boarding pass to have row 5. Got: %v", pass.Column)
    }
}

func Test_should_calculate_the_passId(test *testing.T) {
    pass := new(BoardingPass)
    pass.Init("FBFBBFFRLR")

    if pass.Id != 357 {
        test.Errorf("Expected boarding pass to have pass ID 357. Got: %v", pass.Id)
    }
}
