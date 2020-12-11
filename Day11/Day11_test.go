package main

import (
	"strings"
	"testing"
)

func Test_should_all_be_occupied_after_1_round(test *testing.T) {
	var input = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed)

	room.Update(1)

	for _, row := range room.CurrentState {
		for _, seat := range row {
			if !seat.IsSeat {
				print(".")
				continue
			}

			if seat.Occupied {
				print("#")
				continue
			}

			if !seat.Occupied {
				print("L")
				continue
			}
		}
		println()
	}
}

func Test_should_partially_be_occupied_after_2_rounds(test *testing.T) {
	var input = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed)

	room.Update(2)

	for _, row := range room.CurrentState {
		for _, seat := range row {
			if !seat.IsSeat {
				print(".")
				continue
			}

			if seat.Occupied {
				print("#")
				continue
			}

			if !seat.Occupied {
				print("L")
				continue
			}
		}
		println()
	}
}
