package job

import "time"

type JobReq struct {
	EmployerID      int64   `json:"employer_id"`
	Title           string  `json:"title"`
	EmploymentType  string  `json:"employment_type"`
	Description     string  `json:"description"`
	Location        string  `json:"location"`
	Salary          float64 `json:"salary"`
	ExperienceLevel string  `json:"experience_level"`
}

type JobRes struct {
	ID              int64     `json:"id"`
	EmployerID      int64     `json:"employer_id"`
	Title           string    `json:"title"`
	EmploymentType  string    `json:"employment_type"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	Salary          float64   `json:"salary"`
	ExperienceLevel string    `json:"experience_level"`
	PostedDate      time.Time `json:"posted_date"`
}

type ApplyJobReq struct {
	JobID       int64  `json:"job_id"`
	JobseekerID int64  `json:"jobseeker_id"`
	Resume      string `json:"resume"`
}

type ApplyJobRes struct {
	ID          int64     `json:"id"`
	JobID       int64     `json:"job_id"`
	JobseekerID int64     `json:"jobseeker_id"`
	Resume      string    `json:"resume"`
	Status      string    `json:"status"`
	AppliedAt   time.Time `json:"applied_at"`
}

type UpdateApplicants struct {
	JobID       int64  `json:"job_id"`
	JobseekerID int64  `json:"jobseeker_id"`
	Status      string `json:"status"`
}
