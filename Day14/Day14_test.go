package main

import (
	"strings"
	"testing"
)

func Test_should_execute_instructions(test *testing.T) {
	input := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0"
	split := strings.Split(input, "\n")

	executer := CreateExecuter(split)
	executer.Execute()

	if executer.ReadInt(8) != 64 {
		test.Errorf("Expected memory address 8 to have value 64. Got: %v", executer.ReadInt(8))
	}

	if executer.ReadInt(7) != 101 {
		test.Errorf("Expected memory address 7 to have value 101. Got: %v", executer.ReadInt(7))
	}
}

func Test_should_sum_all_number_in_memory(test *testing.T){
	input := "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0"
	split := strings.Split(input, "\n")

	executer := CreateExecuter(split)
	executer.Execute()

	if executer.GetSum() != 165 {
		test.Errorf("Expected memory address 8 to have value 64. Got: %v", executer.GetSum())
	}
}

func Test_should_sum_all_number_in_memory_V2(test *testing.T){
	input := "mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1"
	split := strings.Split(input, "\n")

	executer := CreateExecuter(split)
	executer.Execute()

	if executer.GetSumV2() != 208 {
		test.Errorf("Expected memory address 8 to have value 208. Got: %v", executer.GetSumV2())
	}
}
