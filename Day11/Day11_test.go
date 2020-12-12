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
	actual := ToString(room)

	expected := "#.##.##.##\n#######.##\n#.#.#..#..\n####.##.##\n#.##.##.##\n#.#####.##\n..#.#.....\n##########\n#.######.#\n#.#####.##\n"

	if actual != expected {
		test.Errorf("Expected value to equal: \n%v\n\n Got:\n%v", expected, actual)
	}
}

func Test_should_partially_be_occupied_after_2_rounds(test *testing.T) {
	var input = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed)

	room.Update(2)
	actual := ToString(room)

	expected := "#.LL.L#.##\n#LLLLLL.L#\nL.L.L..L..\n#LLL.LL.L#\n#.LL.LL.LL\n#.LLLL#.##\n..L.L.....\n#LLLLLLLL#\n#.LLLLLL.L\n#.#LLLL.##\n"

	if actual != expected {
		test.Errorf("Expected value to equal: \n%v\n\n Got:\n%v", expected, actual)
	}
}

func ToString(room Hall) string {
	str := ""
	for _, row := range room.CurrentState {
		for _, seat := range row {
			if !seat.IsSeat {
				str += "."
				continue
			}

			if seat.Occupied {
				str += "#"
				continue
			}

			if !seat.Occupied {
				str += "L"
				continue
			}
		}
		str += "\n"
	}
	return str
}
