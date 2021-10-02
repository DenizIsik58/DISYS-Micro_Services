package main

type Course struct {
	Name                  string    `json:"name"`
	Students              []Student `json:"students"`
	Teachers              []Teacher `json:"teachers"`
	SatisfactionRatingAvg int       `json:"satisfaction_rating_avg"`
}

