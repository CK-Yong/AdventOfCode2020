package main

import (
	"fmt"
	"strconv"
	"strings"
)

var PuzzleInput = "13,16,0,12,15,1"

func main() {
	parsed := strings.Split(PuzzleInput, ",")
	game := CreateGame(parsed)

	result := game.PlayUntil(2020)
	fmt.Printf("Part 1: 2020th number is %v\n", result)

	result2 := game.PlayUntil(30000000)
	fmt.Printf("Part 2: 2020th number is %v", result2)
}

type Game struct {
	Numbers []int
}

func CreateGame(input []string) Game {
	game := Game{make([]int, 0)}
	for _, val := range input {
		num, _ := strconv.Atoi(val)
		game.Numbers = append(game.Numbers, num)
	}
	return game
}


func (game Game) PlayUntil(input int) int {
	cache := make(map[int]int, 0)

	for i, val := range game.Numbers {
		cache[val] = i
	}

	lastVal := game.Numbers[len(game.Numbers) - 1]
	for i := len(game.Numbers); i < input; i++ {
		lastSeen, ok := cache[lastVal]

		value := 0
		if ok {
			value = i - 1 - lastSeen
		}

		cache[lastVal] = i - 1
		lastVal = value
	}

	return lastVal
}
