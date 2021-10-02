package Micro_Services



type Student struct {
	id int `json:"student_id"`
	name string `json:"name"`
	course_workload []Course `json:"courses"`

}
