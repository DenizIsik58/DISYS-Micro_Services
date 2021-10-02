package main

type Student struct {
	Id      int      `json:"student_id"`
	Name    string   `json:"student_name"`
	Courses []Course `json:"courses"`
}



