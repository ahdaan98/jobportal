package storer

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type JOBstorer struct {
	db *sqlx.DB
}

func NewJOBStorer(db *sqlx.DB) *JOBstorer {
	return &JOBstorer{
		db: db,
	}
}

func (js *JOBstorer) CreateJob(ctx context.Context, j *JobReq) (*JobRes, error) {
	var id int64
	query := `
        INSERT INTO jobs (employer_id, title, employment_type, description, location, salary, experience_level)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
	if err := js.db.QueryRowContext(ctx, query,
		j.Title, j.EmploymentType, j.Description, j.Location,
		j.Salary, j.ExperienceLevel).Scan(&id); err != nil {
		return nil, fmt.Errorf("error creating job: %w", err)
	}

	var job JobRes
	err := js.db.GetContext(ctx, &job, "SELECT * FROM jobs WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting job detail: %w", err)
	}

	return &job, nil
}

func (js *JOBstorer) GetJob(ctx context.Context, id int64) (*JobRes, error) {
	var j JobRes
	err := js.db.GetContext(ctx, &j, "SELECT * FROM jobs WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting job detail: %w", err)
	}

	return &j, nil
}

func (js *JOBstorer) ListJobs(ctx context.Context) ([]*JobRes, error) {
	var jobs []*JobRes
	err := js.db.SelectContext(ctx, &jobs, "SELECT * FROM jobs")
	if err != nil {
		return nil, fmt.Errorf("error listing jobs: %w", err)
	}

	return jobs, nil
}

func (js *JOBstorer) UpdateJob(ctx context.Context, id int64, j *JobReq) (*JobRes, error) {
	var job JobRes
	log.Println(j)
	query := `
        UPDATE jobs 
        SET title = :title, 
            employment_type = :employment_type, 
            description = :description, 
            location = :location, 
            salary = :salary, 
            experience_level = :experience_level
        WHERE id = :id`

	params := map[string]interface{}{
		"title":            j.Title,
		"employment_type":  j.EmploymentType,
		"description":      j.Description,
		"location":         j.Location,
		"salary":           j.Salary,
		"experience_level": j.ExperienceLevel,
		"id":               id,
	}

	_, err := js.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return nil, fmt.Errorf("error updating job: %w", err)
	}

	err = js.db.GetContext(ctx, &job, "SELECT * FROM jobs WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting updated job detail: %w", err)
	}

	return &job, nil
}

func (js *JOBstorer) AddApplyJob(ctx context.Context, aj *ApplyJobReq) (error) {
	res, err := js.db.NamedExecContext(ctx, "INSERT INTO applications (job_id, jobseeker_id, resume) VALUES (:job_id, :jobseeker_id, :resume)", aj)
	if err!=nil{
		return fmt.Errorf("error insert apply job detail: %w", err)
	}

	fmt.Println(res)
	return nil
}

func (js *JOBstorer) GetApplyJob(ctx context.Context, aj *ApplyJobReq) (*ApplyJobRes, error) {
	var j ApplyJobRes
	err := js.db.GetContext(ctx, &j, "SELECT * FROM applications WHERE jobseeker_id=$1 AND job_id=$2 ORDER BY id DESC LIMIT 1", aj.JobseekerID, aj.JobID)

	if err!=nil{
		return nil,fmt.Errorf("error select apply job detail: %w", err)
	}

	return &j, err
}

func (js *JOBstorer) ListApplyJobbyID(ctx context.Context, jobid int64) ([]*ApplyJobRes, error) {
	var j []*ApplyJobRes
	err := js.db.SelectContext(ctx, &j, "SELECT * FROM applications WHERE job_id=$1", jobid)

	if err!=nil{
		return nil,fmt.Errorf("error select apply job detail: %w", err)
	}

	return j, err
}

func (js *JOBstorer) IsApplyJobexist(ctx context.Context, jsid, jobid int64) (bool, error) {
	var count int64
	err := js.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM applications WHERE job_id=$1 AND jobseeker_id=$2", jobid, jsid)
	if err != nil {
		return false, fmt.Errorf("failed to get count of jobid and jobseeker id from apply jobs: %w", err)
	}

	return count > 0, nil
}

func (js *JOBstorer) UpdateApplicants(ctx context.Context, appl *UpdateApplicants) (error) {
	row, err := js.db.NamedExecContext(ctx, "UPDATE applications SET status=:status WHERE job_id=:job_id AND jobseeker_id=:jobseeker_id", appl)
	if err!=nil {
		return fmt.Errorf("error updating applicant status - %v : %w", appl.Status, err)
	}

	fmt.Println(row)
	return nil
}

func (js *JOBstorer) GetApplicantsByStatus(ctx context.Context, a *ApplicantsReq) ([]*ApplyJobRes, error) {
	var j []*ApplyJobRes
	err := js.db.SelectContext(ctx, &j, "SELECT * FROM applications WHERE job_id=$1 AND status=$2", a.JobID, a.Status)

	if err!=nil{
		return nil,fmt.Errorf("error select applicants detail: %w", err)
	}

	return j, err
}