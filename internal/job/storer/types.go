package storer

import "time"

type JobReq struct {
	EmployerID      int64   `db:"employer_id"`
	Title           string  `db:"title"`
	EmploymentType  string  `db:"employment_type"`
	Description     string  `db:"description"`
	Location        string  `db:"location"`
	Salary          float64 `db:"salary"`
	ExperienceLevel string  `db:"experience_level"`
}

type JobRes struct {
	ID              int64     `db:"id"`
	EmployerID      int64     `db:"employer_id"`
	Title           string    `db:"title"`
	EmploymentType  string    `db:"employment_type"`
	Description     string    `db:"description"`
	Location        string    `db:"location"`
	Salary          float64   `db:"salary"`
	ExperienceLevel string    `db:"experience_level"`
	PostedDate      time.Time `db:"posted_date"`
}

type ApplyJobReq struct {
	JobID       int64 `db:"job_id"`
	JobseekerID int64 `db:"jobseeker_id"`
	Resumedata  byte
	Resume      string `db:"resume"`
}

type ApplyJobRes struct {
	ID          int64     `db:"id"`
	JobID       int64     `db:"job_id"`
	JobseekerID int64     `db:"jobseeker_id"`
	Resume      string    `db:"resume"`
	Status      string    `db:"status"`
	AppliedAt   time.Time `db:"applied_at"`
}

type UpdateApplicants struct {
	JobID       int64  `db:"job_id"`
	JobseekerID int64  `db:"jobseeker_id"`
	Status      string `db:"status"`
}

type ApplicantsReq struct {
	JobID  int64  `db:"job_id"`
	Status string `db:"status"`
}
