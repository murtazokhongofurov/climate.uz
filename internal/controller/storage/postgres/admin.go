package postgres

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/db"
)

type AdminRepo struct {
	db *db.Postgres
}

func NewAdmin(db *db.Postgres) *AdminRepo {
	return &AdminRepo{
		db: db,
	}
}

func (a *AdminRepo) GetAdmin(username string) (*repo.AdminResponse, error) {
	var (
		create, update time.Time
		res            = repo.AdminResponse{}
	)
	query := `
	SELECT 
		id, username, password, created_at, updated_at
	FROM 
		admins
	WHERE
		username=$1 AND deleted_at IS NULL	
		`
	err := a.db.Pool.QueryRow(context.Background(), query, username).
		Scan(&res.Id, &res.UserName, &res.Password, &create, &update)
	if err != nil {
		fmt.Println("error while getting admin info>> ", err)
		return &repo.AdminResponse{}, err
	}
	if res.UserName == "" {
		res.CreatedAt = ""
		res.UpdatedAt = ""
	} else {
		res.CreatedAt = create.Format(time.RFC1123)
		res.UpdatedAt = update.Format(time.RFC1123)
	}
	return &res, nil
}
