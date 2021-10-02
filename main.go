package main

import (
	"encoding/json"
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

		courses = append(courses, &Course{Students: make([]*Student, 0), Name: name, Teachers: make([]*Teacher, 0), SatisfactionRatingAvg: satisFactionrating})
		return
	}

	if r.Method == "PUT" {
		// Add a student to a course
		studentId := r.FormValue("student_id")
		studentIdToInt, _ := strconv.Atoi(studentId)

		teacherId := r.FormValue("teacher_id")
		teacherIdToInt, _ := strconv.Atoi(teacherId)

		studentName := r.FormValue("student_name")
		courseName := r.FormValue("course_name")

		teacher := getTeacherById(teacherIdToInt)

		student := getStudentById(studentIdToInt)

		course := getCourseByName(courseName)

		if student == nil {
			student = addStudent(studentName)
		}
			course.Students = append (course.Students, student)


		if teacher == nil {
			teacher = addTeacher(studentName)
		}
			course.Teachers = append(course.Teachers, teacher)

		return
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
		students = append(students, &Student{Id: len(students) + 1, Name: name, Courses: make([]Course, 0)})
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	fmt.Println("Starting now...")

	courses = append(courses, &Course{Students: make([]*Student, 0), Name: "DISYS", Teachers: make([]*Teacher, 0), SatisfactionRatingAvg: 10})
	http.HandleFunc("/courses", courseHandler)
	http.HandleFunc("/students", studentHandler)

	log.Fatal(http.ListenAndServe(":7000", nil))

}

func getStudentById(id int) *Student {
	for i:=0; i < len(students); i++ {
		if students[i].Id == id {
			return students[i]
		}
	}
	return nil
}

func getCourseByName(name string) *Course{
	for _, course := range courses {
		if course.Name == name {
			return course
		}
	}
	return nil
}

func getTeacherById(id int) *Teacher {
	for i:=0; i < len(teachers); i++ {
		if teachers[i].Id == id {
			return teachers[i]
		}
	}
	return nil
}

func addStudent(name string) *Student{
	student := Student{Id: len(students) + 1, Name: name, Courses: make([]Course, 0)}
	students = append(students, &student)
	return &student
}

func addTeacher(name string) *Teacher{
	teacher := Teacher{Id: len(students) + 1, Name: name, Courses: make([]Course, 0)}
	teachers = append(teachers, &teacher)
	return &teacher
}
