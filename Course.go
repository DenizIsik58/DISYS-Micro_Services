package Micro_Services


type Course struct {
	teachers []Teacher `json:"teachers"`
	satisfaction_rating_AVG float64 `json:"satisfaction_rate_AVG"`
	student_enrollment int `json:"student_enrollment"`

}


