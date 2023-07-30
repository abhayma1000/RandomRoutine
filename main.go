package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type WorkoutsCsv struct {
	Recent string `csv:"RecentWorkouts"`
}

type WorkoutsJson struct {
	AllWorkouts   []string
	BaseUrl       string
	SlidingWindow int
}

func main() {

	workoutHandler := func(w http.ResponseWriter, req *http.Request) {

		jsonData := readJson("config.json")
		allWorkouts := jsonData.AllWorkouts
		baseUrl := jsonData.BaseUrl
		slidingWindow := jsonData.SlidingWindow
		latest := readCsvAdvanced("test.csv")

		workout, err := getRandom(latest, allWorkouts, slidingWindow)
		if err != nil {
			panic(err)
		}

		io.WriteString(w, fmt.Sprintf("%s/%s", baseUrl, workout))

		AppendCsvAdvanced("test.csv", []string{workout})
	}

	http.HandleFunc("/workout", workoutHandler)
	log.Println("Listing for requests at http://localhost:8000/workout")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getRandom(latest []WorkoutsCsv, allWorkouts []string, slidingWindow int) (string, error) {

	for 1 < 2 {
		not_found := false
		workout := allWorkouts[rand.Intn(len(allWorkouts))]

		if len(latest) > 0 {
			if slidingWindow > len(latest) {
				slidingWindow = len(latest)
			}

			for i := 0; i < slidingWindow; i++ {
				if latest[len(latest)-1-i].Recent == workout {
					// If workout has been executed recently...
					not_found = true
				}
			}
		}

		if !not_found {
			return workout, nil
		}
	}

	return "", fmt.Errorf("no new workout to be done found")
}

func readJson(filename string) WorkoutsJson {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var allWorkouts WorkoutsJson
	err = json.Unmarshal(byteValue, &allWorkouts)
	if err != nil {
		panic(err)
	}

	return allWorkouts
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

func readCsvAdvanced(filename string) []WorkoutsCsv {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	var workouts []WorkoutsCsv
	if err := gocsv.UnmarshalFile(file, &workouts); err != nil {
		panic(err)
	}

	return workouts
}

func AppendCsvAdvanced(filename string, data []string) {
	// Open the CSV file for appending
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write a new row to the CSV file
	err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}
