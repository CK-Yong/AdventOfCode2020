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

	fmt.Printf("Part 1: 2020th number is %v", result)
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

func (game *Game) PlayUntil(input int) int {
	for i := 0; i < input; i++ {
		if i >= len(game.Numbers) {
			game.Numbers = append(game.Numbers, game.GetAge(game.Numbers[i-1], i))
		}
	}

	return game.Numbers[input-1]
}

var cache = make(map[int]int, 0)

func (game *Game) GetAge(value int, index int) int {
	for i := len(game.Numbers) - 2; i >= 0; i-- {
		if game.Numbers[i] == value {
			age := len(game.Numbers) - i - 1
			return age
		}
	}
	return 0
}

