package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type Workouts struct {
	Recent string `csv:"RecentWorkouts"`
	All    string `csv:"AllWorkouts"`
}

func main() {

	// data := readCsv("test.csv")

	stretchesBaseUrl := "https://musclewiki.com/stretches/male"
	baseUrl := "https://musclewiki.com/exercises/male"

	workouts := []string{"chest", "shoulders", "traps", "biceps"}

	workoutHandler := func(w http.ResponseWriter, req *http.Request) {

		num := rand.Intn(len(workouts))

		io.WriteString(w, fmt.Sprintf("%s/%s", stretchesBaseUrl, workouts[num]))

		io.WriteString(w, "\n")

		io.WriteString(w, fmt.Sprintf("%s/%s", baseUrl, workouts[num]))
	}

	http.HandleFunc("/workout", workoutHandler)
	log.Println("Listing for requests at http://localhost:8000/workout")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func readCsvBasic(filename string) [][]string {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return data
}

func readCsvAdvanced(filename string) []Workouts {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	var workouts []Workouts
	if err := gocsv.UnmarshalFile(file, &workouts); err != nil {
		panic(err)
	}

	return workouts
}
