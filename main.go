package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var courses []Course

func studentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(courses)
		return
	}

	if r.Method == "POST" {

		name := r.Form.Get("name")
		students = append(students, Student{Id: len(students) + 1, Name: name, Courses: make([]Course, 0)})
	}
}

func main() {
	fmt.Println("Starting now...")
	courses = append(courses, Course{Id: 1, Name: "Deniz", Courses: make([]Course, 0)})

	http.HandleFunc("/courses", studentHandler)

	log.Fatal(http.ListenAndServe(":7000", nil))
}
