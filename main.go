package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

var courses []Course

func courseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(courses)
		return
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(math.MaxInt)
		name := r.FormValue("name")
		courses = append(courses, Course{Students: make([]Student, 0), Name: name, Teachers: make([]Teacher, 0), SatisfactionRatingAvg: 0})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	fmt.Println("Starting now...")

	courses = append(courses, Course{Students: make([]Student, 0), Name: "DISYS", Teachers: make([]Teacher, 0), SatisfactionRatingAvg: 10})

	http.HandleFunc("/courses", courseHandler)

	log.Fatal(http.ListenAndServe(":7000", nil))
}
