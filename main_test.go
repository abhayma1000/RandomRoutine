package main

import (
	"fmt"
	"testing"
)

func TestReadCsvBasic(t *testing.T) {
	data := readCsvBasic("test.csv")

	// Print the CSV data
	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}

func TestReadCsvAdvanced(t *testing.T) {
	workouts := readCsvAdvanced("test.csv")
	// Print the data
	for _, record := range workouts {
		fmt.Printf("Recent: %s, All: %s\n", record.Recent, record.All)
	}
}
