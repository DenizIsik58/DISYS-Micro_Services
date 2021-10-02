package main

type Course struct {
	Id 						int 	`json:"course_id"`
	Name                  string    `json:"name"`
	Students              []*Student `json:"students"`
	Teachers              []*Teacher `json:"teachers"`
	SatisfactionRatingAvg int       `json:"satisfaction_rating_avg"`
}

