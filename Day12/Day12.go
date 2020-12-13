package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	parsed := strings.Split(PuzzleInput, "\n")

	ship := CreateShip()
	ship.Traverse(parsed)

	manhattan := ship.GetManhattanDistance()

	fmt.Printf("Part 1: Manhattan distance of %v\n", manhattan)

	ship = CreateShip()
	ship.TraverseUsingWaypoint(parsed)

	manhattan = ship.GetManhattanDistance()

	fmt.Printf("Part 2: Manhattan distance of %v", manhattan)
}

type Ship struct {
	PosY int
	PosX int
	Deg  int
	Waypoint Waypoint
}

type Waypoint struct {
	PosY int
	PosX int
}

func (ship Ship) GetManhattanDistance() float64 {
	return math.Abs(float64(ship.PosX)) + math.Abs(float64(ship.PosY))
}

func (ship *Ship) Traverse(instructions []string) {
	for _, val := range instructions {
		dir := val[0]
		num, _ := strconv.Atoi(val[1:])
		ship.Execute(rune(dir), num)
	}
}

func (ship *Ship) Execute(instruction rune, amount int) {
	switch instruction {
	case 'N':
		ship.PosY += amount
	case 'S':
		ship.PosY -= amount
	case 'E':
		ship.PosX += amount
	case 'W':
		ship.PosX -= amount
	case 'L':
		ship.Deg -= amount
	case 'R':
		ship.Deg += amount
	case 'F':
		for ship.Deg >= 360 {
			ship.Deg -= 360
		}
		for ship.Deg <= -360 {
			ship.Deg += 360
		}
		switch ship.Deg {
		case 0:
			ship.PosX += amount
		case 90, -270:
			ship.PosY -= amount
		case 180, -180:
			ship.PosX -= amount
		case 270, -90:
			ship.PosY += amount
		}
	}
}

func (ship *Ship) TraverseUsingWaypoint(instructions []string) {
	for _, val := range instructions {
		dir := val[0]
		num, _ := strconv.Atoi(val[1:])
		ship.ExecuteUsingWaypoint(rune(dir), num)
	}
}

func (ship *Ship) ExecuteUsingWaypoint(instruction rune, amount int) {
	switch instruction {
	case 'N':
		ship.Waypoint.PosY += amount
	case 'S':
		ship.Waypoint.PosY -= amount
	case 'E':
		ship.Waypoint.PosX += amount
	case 'W':
		ship.Waypoint.PosX -= amount
	case 'L':
		for i := 0; i < amount; i += 90{
			posX := ship.Waypoint.PosX
			ship.Waypoint.PosX = -ship.Waypoint.PosY
			ship.Waypoint.PosY = posX
		}
	case 'R':
		for i := 0; i < amount; i += 90 {
			posX := ship.Waypoint.PosX
			ship.Waypoint.PosX = ship.Waypoint.PosY
			ship.Waypoint.PosY = -posX
		}
	case 'F':
		for i := 0; i < amount; i++ {
			ship.PosX += ship.Waypoint.PosX
			ship.PosY += ship.Waypoint.PosY
		}
	}
}

func CreateShip() Ship {
	return Ship{Waypoint: Waypoint{PosY: 1, PosX: 10}}
}