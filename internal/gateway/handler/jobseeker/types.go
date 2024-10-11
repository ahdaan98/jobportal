package jobseeker

import (
	"errors"
	"time"
)

type CreateJobseekerReq struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Gender      string    `json:"gender"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type CreateJobseekerRes struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}

type JobSeekerProfileReq struct {
	JobseekerID int64  `json:"jobseeker_id"`
	Summary     string `json:"summary"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
}

type JobSeekerProfileRes struct {
	CreateJobseekerRes
	Summary    string `json:"summary"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Education  string `json:"education"`
	Experience string `json:"experience"`
}

type JSLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FollowEmployerReq struct {
	JobseekerID int64 `json:"jobseeker_id"`
	EmployerID  int64 `json:"employer_id"`
}

type EmployerRes struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Country string `json:"country"`
	Website string `json:"website"`
}

func (req *CreateJobseekerReq) Validate() error {
	if req.FirstName == "" {
		return errors.New("first_name cannot be empty")
	}
	if req.LastName == "" {
		return errors.New("last_name cannot be empty")
	}
	if req.Email == "" {
		return errors.New("email cannot be empty")
	}
	if req.Password == "" {
		return errors.New("password cannot be empty")
	}
	if req.Gender == "" {
		return errors.New("gender cannot be empty")
	}
	if req.Phone == "" {
		return errors.New("phone cannot be empty")
	}
	if req.DateOfBirth.IsZero() {
		return errors.New("date_of_birth cannot be empty")
	}
	return nil
}