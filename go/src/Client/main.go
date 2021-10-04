package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const URL = "http://localhost:7000"

func main() {

	fmt.Println("Client successfully started....")

	for {

		fmt.Println("Please simply pick one of these 5 options")
		fmt.Println()
		fmt.Println("Press 1 to add a new course")
		fmt.Println("Press 2 to delete an existing course")
		fmt.Println("Press 3 to delete a student or a teacher from a course")
		fmt.Println("Press 4 to change the name of an existing course")
		fmt.Println("Press 5 to get an overview of current courses")
		var selection string
		fmt.Scanln(&selection)

		switch selection {
		case "1":
			fmt.Println("To add a course, simply write a name in the terminal and hit enter")

			var courseName string
			fmt.Scanln(&courseName)
			addCourse(courseName)
		case "2":
			fmt.Println("To delete an existing course, simply write the id in the terminal and hit enter")

			var courseId string
			fmt.Scanln(&courseId)
			deleteCourse(courseId)
		case "3":
			fmt.Println("To delete a teacher or a student from a course, please type TEACHER or STUDENT to delete either of these followed by the courseid and student/teacherid, id int, courseId int")

			var searchType, typeId, courseid string
			fmt.Scan(&searchType)
			fmt.Scan(&typeId)
			fmt.Scan(&courseid)
		case "4":
			fmt.Println("To change the name of an existing course, please type a name in the terminal followed by the id of the course")

			var courseName string
			var courseId string
			fmt.Scan(&courseName)
			fmt.Scan(&courseId)
			changeNameOfCourse(courseName, courseId)
		case "5":
			fmt.Println("You requested to get an overview of all existing courses!")
			fmt.Println(" ")
			getAllTeachers()
		}
	}
}

func addCourse(input string) {
	_, err := http.PostForm(URL+"/courses", url.Values{"name": {input}})

	if err != nil {
		log.Fatal(err)
	}

}

func deleteCourse(courseID string) {
	_, err := http.NewRequest("DELETE", URL+"/courses/"+courseID, nil)

	if err != nil {
		log.Fatalln(err)
	}

}

func deleteTeacherStudent(enrollmentType string, id string, courseId string) {
	_, err := http.NewRequest("DELETE", URL+"/courses"+courseId+"/"+id+"/"+enrollmentType, nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func changeNameOfCourse(newCourseName string, courseId string) {
	jsonStr := "{title:" + newCourseName + "}"

	_, err := http.NewRequest("PUT", URL+"/courses"+courseId, bytes.NewReader([]byte(jsonStr)))

	if err != nil {
		log.Fatalln(err)
	}
}

func getAllTeachers() {
	resp, err := http.Get(URL + "/courses")

	if err != nil {
		log.Fatalln(err)
		return
	}

	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(body), &result)

	if err != nil {
		log.Fatalln(err)
		return
	}

	sb := string(body)
	log.Println(sb)
}
