package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	service := strings.Split(PuzzleInput, ",")

	buses := CreateBuses(PuzzleInputTimestamp, service)
	earliestBus := GetEarliestBus(buses)

	fmt.Printf("Part 1: Earliest bus ID * number of minutes to wait: %v", earliestBus.Interval*(earliestBus.FirstDeparture-PuzzleInputTimestamp))
}

type Bus struct {
	Interval       int
	FirstDeparture int
}

func GetEarliestBus(buses []Bus) Bus {
	earliestDeparture := math.MaxInt32
	earliestId := 0
	for _, bus := range buses {
		if bus.Interval < 0 {
			continue
		}

		if bus.FirstDeparture < earliestDeparture {
			earliestDeparture = bus.FirstDeparture
			earliestId = bus.Interval
		}
	}

	return GetBusById(buses, earliestId)
}

func GetBusById(buses []Bus, id int) Bus {
	for _, bus := range buses {
		if bus.Interval == id {
			return bus
		}
	}
	return Bus{
		Interval:       -1,
		FirstDeparture: -1,
	}
}

func CreateBuses(timestamp int, busInputs []string) []Bus {
	result := make([]Bus, 0)
	for _, input := range busInputs {
		if input == "x" {
			result = append(result, Bus{Interval: -1, FirstDeparture: -1})
			continue
		}

		interval, _ := strconv.Atoi(input)
		departure := GetEarliestDeparture(timestamp, interval)
		result = append(result, Bus{Interval: interval, FirstDeparture: departure})
	}
	return result
}

func GetEarliestDeparture(timestamp int, interval int) int {
	currentDeparture := 0
	for {
		currentDeparture += interval
		if currentDeparture > timestamp {
			return currentDeparture
		}
	}
}
