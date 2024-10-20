package employer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type EMPLOYERstorer struct {
	db *sqlx.DB
}

func NewEMPLOYERstorer(db *sqlx.DB) *EMPLOYERstorer {
	return &EMPLOYERstorer{
		db: db,
	}
}

func (es *EMPLOYERstorer) CreateEmployer(ctx context.Context, e *CreateEmployerReq) (int64, error) {
    stmt, err := es.db.PrepareNamedContext(ctx, `
        INSERT INTO employers (name, email, password, phone, address, country, website)
        VALUES (:name, :email, :password, :phone, :address, :country, :website)
        RETURNING id`)
    if err != nil {
        return 0, fmt.Errorf("error preparing employer insert statement: %w", err)
    }

    var id int64
    err = stmt.Get(&id, e)
    if err != nil {
        return 0, fmt.Errorf("error executing employer insert: %w", err)
    }

    return id, nil
}

func (es *EMPLOYERstorer) GetEmployer(ctx context.Context, id int64) (*EmployerRes, error) {
	query := `
	SELECT id, name, email, phone, address, country, website FROM employers WHERE id=$1; 
	`
	var e EmployerRes
	err := es.db.GetContext(ctx, &e, query, id)
	if err != nil {
		return nil, fmt.Errorf("error inserting into follows: %w", err)
	}

	return &e, nil
}

func (es *EMPLOYERstorer) GetEmployers(ctx context.Context) ([]*EmployerRes, error) {
	var emps []*EmployerRes
	err := es.db.SelectContext(ctx, &emps, "SELECT id, name, email, phone, address, country, website FROM employers")
	if err != nil {
		return nil, fmt.Errorf("error listing employers: %w", err)
	}

	return emps, nil
}

func (es *EMPLOYERstorer) GetFollowersCount(ctx context.Context, id int64) (int64, error) {
	var count int64
	err := es.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM follows WHERE employer_id=$1",id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("error getting employer followeres count: %w", err)
	}

	return count, nil
}

func (es *EMPLOYERstorer) GetJobCounts(ctx context.Context, id int64) (int64, error) {
	var count int64
	err := es.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM jobs WHERE employer_id=$1",id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("error getting employer job counts: %w", err)
	}

	return count, nil
}

func (es *EMPLOYERstorer) IsEmployerExist(ctx context.Context, email string) (bool, error) {
	var count int64
	err := es.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM employers WHERE email=$1", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to get employer count: %w", err)
	}

	return count > 0, nil
}

func (es *EMPLOYERstorer) GetEmployerPass(ctx context.Context, email string) (*GetEPI, error) {
	var d GetEPI
	err := es.db.GetContext(ctx, &d, "SELECT id, password FROM employers WHERE email=$1", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no employer found with email: %s", email) 
		}
		return nil, fmt.Errorf("failed to get employer details: %w", err)
	}

	d.Email = email
	return &d, nil
}