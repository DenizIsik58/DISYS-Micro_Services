package main

type Teacher struct {
	Id               int      `json:"teacher_id"`
	Name             string   `json:"name"`
	Courses         []Course `json:"courses"`
	PopularityScore float64  `json:"popularity_score"`
}
