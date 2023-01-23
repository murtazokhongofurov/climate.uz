package postgres

import (
	"context"
	"database/sql"
	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/db"
)

type CategoryRepo struct {
	db *db.Postgres
}

func NewCategory(db *db.Postgres) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (c *CategoryRepo) CreateCategory(category *repo.CategoryRequest) (*repo.CategoryResponse, error) {
	var (
		res = repo.CategoryResponse{}
	)
	query, _, err := c.db.Builder.Insert("categories").
		Columns("category_name").Values(category.CategoryName).
		Suffix("RETURNING id, category_name, created_at, updated_at").ToSql()
	if err != nil {
		return &repo.CategoryResponse{}, err
	}
	err = c.db.Pool.QueryRow(context.Background(),
		query, category.CategoryName).Scan(
		&res.Id, &res.CatergoryName,
		&res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return &repo.CategoryResponse{}, err
	}
	return &res, nil
}

func (c *CategoryRepo) GetCategoryById(id int) (*repo.CategoryResponse, error) {
	res := repo.CategoryResponse{}
	query, _, err := c.db.Builder.
		Select(
			"id",
			"category_name",
			"created_at",
			"updated_at").
		From("categories").Where("id=$1 AND deleted_at IS NULL", id).ToSql()
	if err != nil {
		return &repo.CategoryResponse{}, err
	}
	err = c.db.Pool.QueryRow(context.Background(), query, id).
		Scan(&res.Id, &res.CatergoryName, &res.CreatedAt, &res.UpdatedAt)
	if err == sql.ErrNoRows {
		return &repo.CategoryResponse{}, nil
	}
	return &res, nil
}

func (c *CategoryRepo) UpdateCategory(category *repo.CategoryUpdateReq) (*repo.CategoryResponse, error) {
	res := repo.CategoryResponse{}
	query := `UPDATE 
				categories SET category_name=$1, 
				updated_at=NOW() WHERE id=$2
				RETURNING id, category_name, created_at, updated_at`

	err := c.db.Pool.QueryRow(context.Background(), query, category.CategoryName, category.Id).
		Scan(&res.Id, &res.CatergoryName, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return &repo.CategoryResponse{}, err
	}
	return &res, nil
}

func (c *CategoryRepo) GetAllCategories(param *repo.AllCategoriesParams) (*repo.AllCategory, error) {
	res := repo.AllCategory{}
	query, _, err := c.db.Builder.Select(
		"id",
		"category_name",
		"created_at",
		"updated_at",
	).From("categories").Where("deleted_at IS NULL").
		Where("category_name ILIKE $1 LIMIT $2 OFFSET $3",
			"%"+param.Search+"%", param.Limit, (param.Page-1)*param.Limit).ToSql()
	if err != nil {
		return &repo.AllCategory{}, err
	}
	rows, err := c.db.Pool.Query(context.Background(), query,
		"%"+param.Search+"%",
		param.Limit,
		(param.Page-1)*param.Limit)
	if err == sql.ErrNoRows {
		return &repo.AllCategory{}, nil
	}
	for rows.Next() {
		temp := repo.CategoryResponse{}
		err = rows.Scan(&temp.Id, &temp.CatergoryName, &temp.CreatedAt, &temp.UpdatedAt)
		if err != nil {
			return &repo.AllCategory{}, err
		}
		res.Categories = append(res.Categories, &temp)
	}
	return &res, nil
}

func (c *CategoryRepo) DeleteCategoryById(id int) (*repo.Empty, error) {
	res := repo.Empty{}
	_, err := c.db.Pool.Exec(context.Background(),
		`UPDATE categories SET deleted_at=NOW() 
	WHERE id=$1 and deleted_at IS NULL`, id)
	if err != nil {
		return &repo.Empty{}, err
	}
	return &res, nil
}
