package main

import (
	"strings"
	"testing"
)

func Test_should_all_be_occupied_after_1_round(test *testing.T) {
	var input = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed,4)

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

	room := ParseRoom(parsed, 4)

	room.Update(2)
	actual := ToString(room)

	expected := "#.LL.L#.##\n#LLLLLL.L#\nL.L.L..L..\n#LLL.LL.L#\n#.LL.LL.LL\n#.LLLL#.##\n..L.L.....\n#LLLLLLLL#\n#.LLLLLL.L\n#.#LLLL.##\n"

	if actual != expected {
		test.Errorf("Expected value to equal: \n%v\n\n Got:\n%v", expected, actual)
	}
}

func Test_should_find_8_seats_when_looking_around(test *testing.T){
	var input = ".......#.\n...#.....\n.#.......\n.........\n..#L....#\n....#....\n.........\n#........\n...#....."
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed,4)

	seat := room.CurrentState[4][3]
	cells := seat.LookAround(room)

	if len(cells) != 8{
		test.Errorf("Expected 8 cells to be found. Got %v", len(cells))
	}

	count := 0
	for _, cell := range cells {
		if cell.IsSeat {
			count++
		}
	}

	if count != 8 {
		test.Errorf("Expected 8 cells to be seats. Got %v", count)
	}
}

func Test_should_find_0_seats_when_looking_around(test *testing.T) {
	var input = ".##.##.\n#.#.#.#\n##...##\n...L...\n##...##\n#.#.#.#\n.##.##."
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed, 4)

	seat := room.CurrentState[3][3]
	cells := seat.LookAround(room)

	if len(cells) != 0 {
		test.Errorf("Expected 0 cells to be found. Got %v", len(cells))
	}

	count := 0
	for _, cell := range cells {
		if cell.IsSeat {
			count++
		}
	}

	if count != 0 {
		test.Errorf("Expected 0 cells to be seats. Got %v", count)
	}
}

func Test_should_be_partially_occupied_after_1_round(test *testing.T) {
	var input = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed,4)

	room.UpdateWithLookaround(1)
	actual := ToString(room)

	expected := "#.##.##.##\n#######.##\n#.#.#..#..\n####.##.##\n#.##.##.##\n#.#####.##\n..#.#.....\n##########\n#.######.#\n#.#####.##\n"

	if actual != expected {
		test.Errorf("Expected value to equal: \n%v\n\n Got:\n%v", expected, actual)
	}
}

func Test_should_be_partially_occupied_after_2_rounds(test *testing.T) {
	var input = "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"
	parsed := strings.Split(input, "\n")

	room := ParseRoom(parsed,4)

	room.UpdateWithLookaround(2)
	actual := ToString(room)

	expected := "#.LL.LL.L#\n#LLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLL#\n#.LLLLLL.L\n#.LLLLL.L#\n"

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
