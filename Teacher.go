package main

type Teacher struct {
	id               int      `json:"teacher_id"`
	name             string   `json:"name"`
	courses          []Course `json:"courses"`
	popularity_score float64  `json:"popularity_score"`
}
