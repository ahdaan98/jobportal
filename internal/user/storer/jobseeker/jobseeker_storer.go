package jobseeker

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	 _ "github.com/lib/pq"
)

type JOBSEEKERstorer struct {
	db *sqlx.DB
}

func NewJOBSEEKERStorer(db *sqlx.DB) *JOBSEEKERstorer {
	return &JOBSEEKERstorer{
		db: db,
	}
}

func (js *JOBSEEKERstorer) CreateJobseeker(ctx context.Context, j *CreateJobseekerReq) (*CreateJobseekerRes, error) {
	var id int64
	query := `
        INSERT INTO jobseekers (first_name, last_name, email, password, gender, phone, date_of_birth)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	if err := js.db.QueryRowContext(ctx, query,
		j.FirstName, j.LastName, j.Email, j.Password,
		j.Gender, j.Phone, j.DateOfBirth).Scan(&id); err != nil {
		return nil, fmt.Errorf("error creating jobseeker: %w", err)
	}

	var jobseeker CreateJobseekerRes
	err := js.db.GetContext(ctx, &jobseeker, "SELECT * FROM jobseekers WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting jobseeker detail: %w", err)
	}

	return &jobseeker, nil
}

func (js *JOBSEEKERstorer) CreateJobseekerProfile(ctx context.Context, j *JobSeekerProfileReq) error {
	query := `
        INSERT INTO profiles (jobseeker_id, summary, city, country, education, experience)
		VALUES (:jobseeker_id, :summary, :city, :country, :education, :experience)`
	_, err := js.db.NamedExecContext(ctx, query, j)

	if err != nil {
		return fmt.Errorf("error creating jobseeker profile: %w", err)
	}

	return nil
}

func (js *JOBSEEKERstorer) GetJobSeeker(ctx context.Context, id int64) (*JobSeekerProfileRes, error) {
	var j CreateJobseekerRes
	err := js.db.GetContext(ctx, &j, "SELECT * FROM jobseekers WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting job seeker detail: %w", err)
	}

	jq := `
	  SELECT 
            js.id, js.first_name, js.last_name, js.email, js.gender, js.phone, 
            js.date_of_birth, js.created_at, js.updated_at, 
            jp.summary, jp.city, jp.country, jp.education, jp.experience
        FROM jobseekers js
        JOIN profiles jp ON js.id = jp.jobseeker_id
        WHERE js.id = $1`

	var jobseekerProfile JobSeekerProfileRes
	err = js.db.GetContext(ctx, &jobseekerProfile, jq, id)
	if err != nil {
		return nil, fmt.Errorf("error getting jobseeker detail: %w", err)
	}

	return &jobseekerProfile, nil
}

func (js *JOBSEEKERstorer) IsJSexist(ctx context.Context, email string) (bool, error) {
	var count int64
	err := js.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM jobseekers WHERE email=$1", email)
	if err != nil {
		return false, fmt.Errorf("failed to get jobseeker count: %w", err)
	}

	return count > 0, nil
}

func (js *JOBSEEKERstorer) GetJSpass(ctx context.Context, email string) (*GetEPI, error) {
	var d GetEPI
	err := js.db.GetContext(ctx, &d, "SELECT id, password FROM jobseekers WHERE email=$1", email)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobseeker detail: %w", err)
	}

	d.Email = email

	return &d, nil
}

func (js *JOBSEEKERstorer) GetBasicJSProfile(ctx context.Context, email string) (*CreateJobseekerRes, error) {
	var jobseeker CreateJobseekerRes
	err := js.db.GetContext(ctx, &jobseeker, "SELECT * FROM jobseekers WHERE email=$1", email)
	if err != nil {
		return nil, fmt.Errorf("error getting jobseeker detail: %w", err)
	}

	return &jobseeker, nil
}

func (js *JOBSEEKERstorer) GetBasicJSProfilebyID(ctx context.Context, id int64) (*CreateJobseekerRes, error) {
	var jobseeker CreateJobseekerRes
	err := js.db.GetContext(ctx, &jobseeker, "SELECT * FROM jobseekers WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting jobseeker detail: %w", err)
	}

	return &jobseeker, nil
}

func (js *JOBSEEKERstorer) FollowEmployer(ctx context.Context, f *FollowEmployerReq) (error) {
	query := `
	INSERT INTO follows(jobseeker_ID, employer_id) VALUES (:jobseeker_id, :employer_id)
	`

	_, err := js.db.NamedExecContext(ctx, query, f)
	if err!=nil {
		return fmt.Errorf("error inserting into follows: %w",err)
	}

	return nil
}

func (js *JOBSEEKERstorer) UnFollowEmployer(ctx context.Context, f *FollowEmployerReq) (error) {
	query := `
	DELETE FROM follows WHERE jobseeker_id=$1 AND employer_id=$2
	`

	_, err := js.db.ExecContext(ctx, query, f.JobseekerID, f.EmployerID)
	if err!=nil {
		return fmt.Errorf("error deleting follows: %w",err)
	}

	return nil
}