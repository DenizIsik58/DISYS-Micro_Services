package main

type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}
