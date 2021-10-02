package main

type Course struct {
	Student               []Student `json:"students"`
	Teachers              []Teacher `json:"teachers"`
	SatisfactionRatingAvg float64   `json:"satisfaction_rating_avg"`
}
