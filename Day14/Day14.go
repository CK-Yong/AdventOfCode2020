package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	split := strings.Split(PuzzleInput, "\n")

	executer := CreateExecuter(split)
	executer.Execute()

	fmt.Printf("Part 1: Total sum is %v\n", executer.GetSum())
	fmt.Printf("Part 2: Total sum is %v", executer.GetSumV2())
}

type Executer struct {
	Commands    []Command
	Memory      map[int][]byte
	MemoryV2    map[int]int
	CurrentMask Mask
}

type Command interface {
	execute(*Executer)
	executeV2(*Executer)
}

type Mask struct {
	Value string
}

type SetAddress struct {
	Address int
	Value   int
}

func (mask Mask) execute(executer *Executer) {
	executer.CurrentMask = mask
}

func (mask Mask) executeV2(executer *Executer) {
	mask.execute(executer)
}

func (mask Mask) Apply(value int) []byte {
	result := ToByteArray(value)
	for i, bit := range mask.Value {
		if bit == 'X' {
			continue
		}
		if bit == '1' {
			result[i] = 1
		}
		if bit == '0' {
			result[i] = 0
		}
	}
	return result
}

func ToByteArray(value int) []byte {
	result := make([]byte, 36)
	binary := strconv.FormatInt(int64(value), 2)

	startIndex := len(result) - len(binary)
	for i := range binary {
		if binary[i] == '1' {
			result[startIndex+i] = 1
		} else {
			result[startIndex+i] = 0
		}
	}
	return result
}

func (command SetAddress) execute(executer *Executer) {
	bytes := executer.CurrentMask.Apply(command.Value)
	executer.Memory[command.Address] = bytes
}

func (executer *Executer) Execute() {
	for _, command := range executer.Commands {
		command.execute(executer)
	}
	for _, command := range executer.Commands {
		command.executeV2(executer)
	}
}

func CreateExecuter(input []string) Executer {
	executer := Executer{Commands: make([]Command, 0), Memory: make(map[int][]byte, 0), MemoryV2: make(map[int]int, 0)}
	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			maskInput := strings.Split(line, " = ")[1]
			executer.Commands = append(executer.Commands, Mask{Value: maskInput})
		} else {
			addrEnd := strings.Index(line, "]")
			address, _ := strconv.Atoi(line[4:addrEnd])
			valString := strings.Split(line, " = ")[1]
			val, _ := strconv.Atoi(valString)
			executer.Commands = append(executer.Commands, SetAddress{Address: address, Value: val})
		}
	}
	return executer
}

func (executer Executer) ReadInt(index int) int {
	bytes := executer.Memory[index]
	return ToInt(bytes)
}

func ToInt(bytes []byte) int {
	count := float64(0)
	for i := 0; i < len(bytes); i++ {
		if bytes[len(bytes)-1-i] == 1 {
			count += math.Pow(float64(2), float64(i))
		}
	}
	return int(count)
}

func (executer *Executer) GetSum() int {
	count := 0
	for key := range executer.Memory {
		count += executer.ReadInt(key)
	}
	return count
}

/// Part 2 logic.. It's gonna be big??
func (command SetAddress) executeV2(executer *Executer) {
	addressMask := executer.CurrentMask.ApplyV2(command.Address)
	addresses := GetAddresses(addressMask)
	for _, address := range addresses {
		executer.MemoryV2[address] = command.Value
	}
}

func (mask Mask) ApplyV2(value int) []byte {
	result := ToByteArray(value)
	for i, bit := range mask.Value {
		switch bit {
		case 'X':
			result[i] = 'X'
		case '1':
			result[i] = 1
		case '0':
			continue
		}
	}
	return result
}

func GetAddresses(addressBytes []byte) []int {
	result := make([]int, 2)
	addr1 := make([]byte, 36)
	copy(addr1, addressBytes)
	addr2 := make([]byte, 36)
	copy(addr2, addressBytes)

	for i := 0; i < len(addressBytes); i++ {
		if addressBytes[i] == 'X' {
			addr1[i] = 1
			addr2[i] = 0
			result = append(result, GetAddresses(addr1)...)
			result = append(result, GetAddresses(addr2)...)
			return result
		}
	}

	int1 := ToInt(addr1)
	return []int{int1}
}

func (executer *Executer) GetSumV2() int {
	count := 0
	for key, val := range executer.MemoryV2 {
		if key == 0 {
			continue
		}
		count += val
	}
	return count
}
