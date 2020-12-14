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

	fmt.Printf("Part 1: Earliest bus ID * number of minutes to wait: %v\n", earliestBus.Interval*(earliestBus.FirstDeparture-PuzzleInputTimestamp))

	buses = CreateBuses(-1, service)
	earliestTimestamp := GetSyncedTimestamp(buses)

	fmt.Printf("Part 2: Earliest timestamp: %v", earliestTimestamp)
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

func GetSyncedTimestamp(buses []Bus) int64 {
	ts := searchTimeStamp(0, 0, buses)
	commonMultiplier := GetCommonMultiplier(buses)

	for ts-int64(commonMultiplier) > 0 {
		ts -= int64(commonMultiplier)
	}
	return ts
}

func GetCommonMultiplier(buses []Bus) int {
	aggregator := 1
	for _, bus := range buses {
		if bus.Interval < 0 {
			continue
		}
		aggregator *= bus.Interval
	}
	return aggregator
}

func searchTimeStamp(index int, timestamp int64, buses []Bus) int64 {
	if index == len(buses) {
		return 0
	}

	bus := buses[index]
	interval := int64(bus.Interval)

	if bus.Interval < 0 {
		return searchTimeStamp(index+1, timestamp, buses)
	}

	newTimestamp := int64(1)
	for i, bus := range buses {
		if i == index || bus.Interval < 0 {
			continue
		}
		newTimestamp *= int64(bus.Interval)
	}

	i := int64(0)
	for ((newTimestamp*i)+int64(index))%interval != 0 {
		i++
	}

	return (newTimestamp * i) + searchTimeStamp(index+1, newTimestamp, buses)
}
