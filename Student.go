package main

type Student struct {
	Id      int      `json:"student_id"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}



