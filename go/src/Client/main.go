package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {

	fmt.Println("Client successfully started....")
	fmt.Println("Please simply pick one of these 5 options")
	fmt.Println()
	fmt.Println("Press 1 to add a new course")
	fmt.Println("Press 2 to delete an existing course")
	fmt.Println("Press 3 to delete a student or a teacher from a course")
	fmt.Println("Press 4 to change the name of an existing course")
	fmt.Println("Press 5 to get an overview of current courses")
	var k string
	fmt.Scanln(&k)

	switch k {
	case "1":
		fmt.Println("To add a course, simply write a name in the terminal and hit enter")

		var courseName string
		fmt.Scanln(&courseName)
		addCourse(courseName)
		break
	case "2":
		fmt.Println("To delete an existing course, simply write the id in the terminal and hit enter")

		var courseId int
		fmt.Scanln(&courseId)
		deleteCourse(courseId)
		break
	case "3":
		fmt.Println("To delete a teacher or a student from a course, please type TEACHER or STUDENT to delete either of these followed by the courseid and student/teacherid, id int, courseId int")

		var searchType, typeId, courseid string
		fmt.Scan(&searchType)
		fmt.Scan(&typeId)
		fmt.Scan(&courseid)
		break

	case "4":
		fmt.Println("To change the name of an existing course, please type a name in the terminal followed by the id of the course")

		var courseName string
		var courseId int
		fmt.Scan(&courseName)
		fmt.Scan(&courseId)
		changeNameOfCourse(courseName, courseId)

		break

	case "5":
		fmt.Println("You requested to get an overview of all existing courses! \n")
		getAllTeachers()
		break
	}







}


func addCourse(input string) {

	_, err := http.PostForm("http://localhost:7000/courses", url.Values{"name": {input}})

	if err != nil {
		log.Fatal(err)
	}

}

func deleteCourse(courseID int) {
	_, err := http.NewRequest("DELETE", "http://localhost:7000/courses/1", nil)

	if err != nil {
		fmt.Println("Error occured")
	}

}

func deleteTeacherStudent(enrollmentType string, id int, courseId int){

	courseIdToString := strconv.Itoa(courseId)
	IdToString := strconv.Itoa(id)

	_, err := http.NewRequest("DELETE", "http://localhost:7000/courses" + courseIdToString + "/" + IdToString + "/" + enrollmentType,nil)

	if err != nil {
		fmt.Println(err)
	}
}

func changeNameOfCourse(newCourseName string, courseId int) {

	courseIdToString := strconv.Itoa(courseId)

	jsonStr := "{title:" + newCourseName + "}"


	_, err := http.NewRequest("PUT","http://localhost:7000/courses" + courseIdToString, bytes.NewReader([]byte(jsonStr)) )

	if err != nil {
		fmt.Println(err)
	}

}

func getAllTeachers(){
	resp, err := http.Get("http://localhost:7000/courses")

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
	log.Printf(sb)
}


