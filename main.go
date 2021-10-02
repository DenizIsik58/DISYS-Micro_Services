package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var courses []*Course
var students = make([]*Student, 0)
var teachers = make([]*Teacher, 0)

func courseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(courses)
		return
	}

	if r.Method == "POST" {
		r.ParseMultipartForm(math.MaxInt)
		name := r.FormValue("name")
		rating := r.FormValue("satisfaction_rating_avg")
		satisFactionrating, _ := strconv.Atoi(rating)

		courses = append(courses, &Course{Students: make([]Student, 0), Name: name, Teachers: make([]Teacher, 0), SatisfactionRatingAvg: satisFactionrating})
		return
	}

	if r.Method == "PUT" {
		// Add a student to a course
		studentId := r.Form.Get("student_id")
		courseName := r.Form.Get("course_name")
		teacherId := r.Form.Get("teacher_id")

		if !courseExists(courseName) {
			fmt.Println("This course doesn't exsits")
			return
		}
		if studentExistsInCourse(courseName) || teacherExistsInCourse(courseName) {
			return
		}
		addStudentToCourse(studentId, courseName)

		// Add a teacher to a course

		addTeacherToCourse(teacherId, courseName)



	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(courses)
		return
	}

	if r.Method == "POST" {
		name := r.Form.Get("name")
		append(students, &Student{Id: len(students) + 1, Name: name, Courses: make([]Course, 0)})
	}

	if r.Method == "POST" {
		name := r.Form.Get("name")
		append(students, &Student{Id: len(students) + 1, Name: name, Courses: make([]Course, 0)})


	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	fmt.Println("Starting now...")

	courses = append(courses, &Course{Students: make([]Student, 0), Name: "DISYS", Teachers: make([]Teacher, 0), SatisfactionRatingAvg: 10})
	http.HandleFunc("/courses", courseHandler)
	http.HandleFunc("/students", studentHandler)

	log.Fatal(http.ListenAndServe(":7000", nil))

}

func addStudent(Id int, name string, courses []Course) {
	students = append(students, &Student{Id, name, courses})
}

func getStudent(index int) *Student {
	return students[index]
}

func studentExists(student *Student) bool {
	for i := 0; i < len(students); i++ {
		if students[i] == student {
			return false
		}
	}
	return true
}

func teacherExists(teacher *Teacher) bool {
	for i := 0; i < len(students); i++ {
		if teachers[i] == teacher {
			return false
		}
	}
	return true
}

func courseExists(courseName string) bool{
	for i:=0; i < len(courses); i++ {
		if courses[i].Name == courseName {
			return true
		}
	}
	return false
}

func studentExistsInCourse(courseName string) bool{
	for i:=0; i < len(students); i++ {
		for j := 0; j < len(students[i].Courses); j++ {
			if students[i].Courses[j].Name == courseName {
				return false
			}
		}
	}
	return true
}

func teacherExistsInCourse(courseName string) bool{
	for i:=0; i < len(teachers); i++ {
		for j := 0; j < len(teachers[i].Courses); j++ {
			if teachers[i].Courses[j].Name == courseName {
				return false
			}
		}
	}
	return true
}

func addStudentToCourse(studentId int, courseName string)  {

	if !studentExists(students[studentId]) {
		errors.New("this student doesn't exists")
	}

	for i := 0; i < len(courses); i++ {
		if courses[i].Name == courseName {
			append(students[studentId].Courses, *courses[i])
		}
	}
}

func addTeacherToCourse(teacherId int, courseName string) (*Course, error) {

	if !teacherExists(teachers[teacherId]) {
		errors.New("this teacher doesn't exists")
	}

	for i := 0; i < len(courses); i++ {
		if courses[i].Name == courseName {
			append(teachers[teacherId].Courses, *courses[i])
			return courses[i], errors.New("")
		}
	}
	return nil, errors.New("this course does not Exists")
}
