package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	baseUrl := "https://musclewiki.com/exercises/male"

	workouts := []string{"chest", "shoulders", "traps", "biceps"}

	workoutHandler := func(w http.ResponseWriter, req *http.Request) {

		num := rand.Intn(len(workouts))

		io.WriteString(w, fmt.Sprintf("%s/%s", baseUrl, workouts[num]))
	}

	http.HandleFunc("/workout", workoutHandler)
	log.Println("Listing for requests at http://localhost:8000/workout")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
