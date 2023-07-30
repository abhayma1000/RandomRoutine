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

const doInit string = "init"
const doExec string = "exec"

type OptionsCsv struct {
	Recent string `csv:"Recent"`
}

type OptionsJson struct {
	AllOptions    []string
	BaseUrl       string
	SlidingWindow int
}

func main() {
	// go run main.go init workouts
	// go run main.go serve workouts

	// If incorrect go run statement
	if len(os.Args) != 3 || (os.Args[1] != doInit && os.Args[1] != doExec) {
		panic(fmt.Errorf("run statement is invalid. Follow schema specified. Input: %v", os.Args[1:]))
	}
	action := os.Args[1]
	upon := os.Args[2]

	if action == doInit {
		initCsv := []string{"Recent"}
		writeCsvBasic(fmt.Sprintf("%s.csv", upon), initCsv)

		initJson := OptionsJson{
			[]string{}, "", 0,
		}
		createJson(fmt.Sprintf("%s.json", upon), initJson)

	}

	if action == doExec {
		// If init has not been run
		_, err := os.Open(fmt.Sprintf("%s.csv", upon))
		if err != nil {
			panic(fmt.Errorf("run init before exec: %s", err))
		}
		_, err = os.Open(fmt.Sprintf("%s.json", upon))
		if err != nil {
			panic(fmt.Errorf("run init before exec: %s", err))
		}

		optionHandler := func(w http.ResponseWriter, req *http.Request) {

			jsonData := readJson(fmt.Sprintf("%s.json", upon))
			allOptions := jsonData.AllOptions
			baseUrl := jsonData.BaseUrl
			slidingWindow := jsonData.SlidingWindow
			if slidingWindow == 0 {
				panic(fmt.Errorf("set SlidingWindow to non-zero value: %d", slidingWindow))
			}

			latest := readCsvAdvanced(fmt.Sprintf("%s.csv", upon))

			option, err := getRandom(latest, allOptions, slidingWindow)
			if err != nil {
				panic(err)
			}

			io.WriteString(w, fmt.Sprintf("%s/%s", baseUrl, option))

			AppendCsvAdvanced(fmt.Sprintf("%s.csv", upon), []string{option})
		}

		http.HandleFunc(fmt.Sprintf("/%s", upon), optionHandler)
		log.Printf("Listing for requests at http://localhost:8000/%s", upon)
		log.Fatal(http.ListenAndServe(":8000", nil))
	}
}

func getRandom(latest []OptionsCsv, allOptions []string, slidingWindow int) (string, error) {

	for 1 < 2 {
		not_found := false

		option := allOptions[rand.Intn(len(allOptions))]

		if len(latest) > 0 {
			if slidingWindow > len(latest) {
				slidingWindow = len(latest)
			}

			for i := 0; i < slidingWindow; i++ {
				if latest[len(latest)-1-i].Recent == option {
					// If option has been executed recently...
					not_found = true
				}
			}
		}

		if !not_found {
			return option, nil
		}
	}

	return "", fmt.Errorf("no new option to be done found")
}

func readJson(filename string) OptionsJson {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var allOptions OptionsJson
	err = json.Unmarshal(byteValue, &allOptions)
	if err != nil {
		panic(err)
	}

	return allOptions
}

func createJson(filename string, data OptionsJson) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func writeCsvBasic(filename string, data []string) {
	// Write the CSV data
	file2, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	defer writer.Flush()

	writer.Write(data)
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

func readCsvAdvanced(filename string) []OptionsCsv {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	var options []OptionsCsv
	if err := gocsv.UnmarshalFile(file, &options); err != nil {
		panic(err)
	}

	return options
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

// func createCsvAdvanced(filename string, data *[]OptionsCsv) string {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	if err := gocsv.MarshalFile(&data, file); err != nil {
// 		panic(err)
// 	}

// 	return filename
// }
