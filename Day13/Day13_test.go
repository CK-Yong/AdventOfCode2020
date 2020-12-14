package main

import (
	"strings"
	"testing"
)

func Test_should_result_in_properly_parsed_buses(test *testing.T) {
	timestamp := 939
	service := strings.Split("7,13,x,x,59,x,31,19", ",")

	buses := CreateBuses(timestamp, service)

	earliestBus := GetBusById(buses, 59)

	if earliestBus.FirstDeparture != 944 {
		test.Errorf("Expected bus 59 to depart at 944. Got: %v", earliestBus.FirstDeparture)
	}
}

func Test_should_result_in_earliest_bus(test *testing.T) {
	timestamp := 939
	service := strings.Split("7,13,x,x,59,x,31,19", ",")

	buses := CreateBuses(timestamp, service)

	earliestBus := GetEarliestBus(buses)

	if earliestBus.FirstDeparture != 944 || earliestBus.Interval != 59 {
		test.Errorf("Expected bus 59 to depart at 944. Got: ID %v, Departure %v", earliestBus.Interval, earliestBus.FirstDeparture)
	}
}

func Test_should_get_earliest_timestamp_at_which_buses_sync(test *testing.T) {
	parsed := strings.Split("17,x,13,19", ",")

	buses := CreateBuses(-1, parsed)

	earliestTimestamp := GetSyncedTimestamp(buses)

	if earliestTimestamp != 3417 {
		test.Errorf("Expected earliest timestamp to be 3417, got %v", earliestTimestamp)
	}
}

