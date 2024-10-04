package jobseeker

import "time"

type FollowEmployerReq struct {
	JobseekerID int64 `db:"jobseeker_id"`
	EmployerID  int64 `db:"employer_id"`
}

type CreateJobseekerReq struct {
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	Gender      string    `db:"gender"`
	Phone       string    `db:"phone"`
	DateOfBirth time.Time `db:"date_of_birth"`
}

type CreateJobseekerRes struct {
	ID          int64     `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	Gender      string    `db:"gender"`
	Phone       string    `db:"phone"`
	DateOfBirth time.Time `db:"date_of_birth"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type JobSeekerProfileReq struct {
	JobseekerID int64  `db:"jobseeker_id"`
	Summary     string `db:"summary"`
	City        string `db:"city"`
	Country     string `db:"country"`
	Education   string `db:"education"`
	Experience  string `db:"experience"`
}

type JobSeekerProfileRes struct {
	CreateJobseekerRes
	Summary    string `db:"summary"`
	City       string `db:"city"`
	Country    string `db:"country"`
	Education  string `db:"education"`
	Experience string `db:"experience"`
}

type GetEPI struct {
	Id    int64  `db:"id"`
	Email string `db:"email"`
	Pass  string `db:"password"`
}