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
		fmt.Printf("Recent: %s\n", record.Recent)
	}
}

func TestAppendCsvAdvanced(t *testing.T) {
	data := []string{"test1"}

	AppendCsvAdvanced("test.csv", data)

	TestReadCsvAdvanced(t)
}

func TestReadJson(t *testing.T) {
	workouts := readJson("config.json")

	fmt.Printf("All workouts: %v \n", workouts.AllWorkouts)
	fmt.Printf("Workouts Base URL: %s \n", workouts.BaseUrl)
}
