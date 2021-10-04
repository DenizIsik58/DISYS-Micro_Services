package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var courses = make([]*Course, 0)
var students = make([]*Student, 0)
var teachers = make([]*Teacher, 0)

func coursePutTeacherStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	courseId, error := strconv.Atoi(vars["courseId"])

	if error != nil {

		http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
		return
	}

	course := getCourseById(courseId)

	if course == nil {
		http.Error(w, "{\"error\": \"COURSE_NOT_FOUND\"}", http.StatusNotFound)
		return
	}

	r.ParseMultipartForm(math.MaxInt)

	if r.FormValue("type") == "teacher" {
		teacherId, error := strconv.Atoi(vars["id"])

		if error != nil {
			http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
			return
		}

		teacher := getTeacherById(teacherId)

		if teacher == nil {
			http.Error(w, "{\"error\": \"TEACHER_NOT_FOUND\"}", http.StatusNotFound)
			return
		}

		course.Teachers = append(course.Teachers, teacher)
	} else {
		studentId, error := strconv.Atoi(vars["id"])

		if error != nil {
			http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
			return
		}

		student := getStudentById(studentId)

		if student == nil {
			http.Error(w, "{\"error\": \"STUDENT_NOT_FOUND\"}", http.StatusNotFound)
			return
		}

		course.Students = append(course.Students, student)
	}

	w.Write([]byte("{\"status\": true}"))
	w.WriteHeader(200)
}

func courseDeleteTeacherStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	courseId, error := strconv.Atoi(vars["courseId"])

	if error != nil {

		http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
		return
	}

	course := getCourseById(courseId)

	if course == nil {
		http.Error(w, "{\"error\": \"COURSE_NOT_FOUND\"}", http.StatusNotFound)
		return
	}

	r.ParseMultipartForm(math.MaxInt)

	removed := false
	searchType := strings.ToUpper(r.FormValue("type"))

	if searchType == "TEACHER" {
		teacherId, error := strconv.Atoi(vars["id"])

		if error != nil {
			http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
			return
		}

		teacher := getTeacherById(teacherId)

		if teacher == nil {
			http.Error(w, "{\"error\": \"TEACHER_NOT_FOUND\"}", http.StatusNotFound)
			return
		}

		for i, element := range course.Teachers {
			if element.Id == teacher.Id {
				course.Teachers = append(course.Teachers[:i], course.Teachers[i+1:]...)
				removed = true
				break
			}
		}
	} else {
		studentId, error := strconv.Atoi(vars["id"])

		if error != nil {
			http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
			return
		}

		student := getStudentById(studentId)

		if student == nil {
			http.Error(w, "{\"error\": \"STUDENT_NOT_FOUND\"}", http.StatusNotFound)
			return
		}

		for i, element := range course.Students {
			if element.Id == student.Id {
				course.Students = append(course.Students[:i], course.Students[i+1:]...)
				removed = true
				break
			}
		}
	}

	if !removed {
		http.Error(w, "{\"error\": \""+searchType+"_NOT_IN_COURSE\"}", http.StatusNotFound)
		return
	}

	w.Write([]byte("{\"status\": true}"))
	w.WriteHeader(200)
}

func courseDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	courseId, error := strconv.Atoi(vars["courseId"])

	if error != nil {
		http.Error(w, "{\"error\": \"INVALID_ID\"}", http.StatusForbidden)
		return
	}

	course := getCourseById(courseId)

	if course == nil {
		http.Error(w, "{\"error\": \"COURSE_NOT_FOUND\"}", http.StatusNotFound)
		return
	}

	for i, element := range courses {
		if element.Id == course.Id {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}

	w.Write([]byte("{\"status\": true}"))
	w.WriteHeader(200)
}

func coursePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(math.MaxInt)

	name := r.FormValue("name")
	course := &Course{Id: len(courses) + 1, Name: name, Students: make([]*Student, 0), Teachers: make([]*Teacher, 0), SatisfactionRatingAvg: 0}
	courses = append(courses, course)

	w.Write([]byte("{\"status\": true}"))
	w.WriteHeader(200)
}

func courseGetHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courses)
}

func main() {
	fmt.Println("Starting now...")

	courses = append(courses, &Course{Students: make([]*Student, 0), Name: "DISYS", Teachers: make([]*Teacher, 0), SatisfactionRatingAvg: 10})

	r := mux.NewRouter()
	r.HandleFunc("/courses/{courseId}/{id}", coursePutTeacherStudentHandler).Methods("PUT")
	r.HandleFunc("/courses/{courseId}/{id}", courseDeleteTeacherStudentHandler).Methods("DELETE")
	r.HandleFunc("/courses", courseGetHandler).Methods("GET")
	r.HandleFunc("/courses", coursePostHandler).Methods("POST")
	r.HandleFunc("/courses/{courseId}", courseDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":7000", r))
}

func getStudentById(id int) *Student {
	for i := 0; i < len(students); i++ {
		if students[i].Id == id {
			return students[i]
		}
	}
	return nil
}

func getCourseById(id int) *Course {
	for _, course := range courses {
		if course.Id == id {
			return course
		}
	}

	return nil
}

func getTeacherById(id int) *Teacher {
	for i := 0; i < len(teachers); i++ {
		if teachers[i].Id == id {
			return teachers[i]
		}
	}

	return nil
}
