package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/db"
)

type NewsProductRepo struct {
	db *db.Postgres
}

func NewNewsProduct(db *db.Postgres) *NewsProductRepo {
	return &NewsProductRepo{
		db: db,
	}
}

func (n *NewsProductRepo) CreateNews(new *repo.NewsProductRequest) (*repo.NewsProductResponse, error) {
	var (
		create, update time.Time
		res            = repo.NewsProductResponse{}
	)
	query := `INSERT INTO new_products(
		category_id, 
		title, 
		media_link, 
		description, 
		price) VALUES($1, $2, $3, $4, $5)
		RETURNING 
		id, 
		category_id, 
		title, 
		media_link, 
		description, 
		price, created_at, updated_at`
	err := n.db.Pool.QueryRow(context.Background(), query,
		new.CategoryId,
		new.Title,
		new.MediaLink,
		new.Description,
		new.Price,
	).Scan(
		&res.Id,
		&res.CategoryId,
		&res.Title,
		&res.MediaLink,
		&res.Description,
		&res.Price,
		&create, &update)
	res.CreatedAt = create.Format(time.RFC1123)
	res.UpdatedAt = update.Format(time.RFC1123)
	if err != nil {
		return &repo.NewsProductResponse{}, err
	}

	return &res, nil
}

func (n *NewsProductRepo) GetNewsById(id int64) (*repo.NewsProductResponse, error) {
	var (
		create, update time.Time
		res            = repo.NewsProductResponse{}
	)
	query := `SELECT 
		id, 
		category_id, 
		title, 
		media_link, 
		description, 
		price, 
		created_at, 
		updated_at 
		FROM new_products WHERE deleted_at IS NULL AND id=$1`
	err := n.db.Pool.QueryRow(context.Background(), query, id).
		Scan(
			&res.Id,
			&res.CategoryId,
			&res.Title,
			&res.MediaLink,
			&res.Description,
			&res.Price,
			&create, &update,
		)
	if res.Title == "" {
		res.CreatedAt, res.UpdatedAt = "", ""
	} else {
		res.CreatedAt = create.Format(time.RFC1123)
		res.UpdatedAt = update.Format(time.RFC1123)
	}
	if err == sql.ErrNoRows {
		return &repo.NewsProductResponse{}, nil
	}
	return &res, nil
}

func (n *NewsProductRepo) GetNewsByCategoryId(id int64) (*repo.AllNewsProducts, error) {
	var (
		create, update time.Time
		res            = repo.AllNewsProducts{}
	)
	query := `
	SELECT 
		id, 
		category_id, 
		title, 
		media_link, 
		description, 
		price, 
		created_at, 
		updated_at 
	FROM new_products
		WHERE category_id=$1 AND deleted_at IS NULL`
	rows, err := n.db.Pool.Query(context.Background(), query, id)
	if err != nil {
		fmt.Println("error while getting news by categoryId", err)
		return &repo.AllNewsProducts{}, err
	}
	for rows.Next() {
		temp := repo.NewsProductResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.CategoryId,
			&temp.Title,
			&temp.MediaLink,
			&temp.Description,
			&temp.Price,
			&create, &update,
		)
		if temp.Title == "" {
			temp.CreatedAt, temp.UpdatedAt = "", ""
		} else {
			temp.CreatedAt, temp.UpdatedAt = create.Format(time.RFC1123), update.Format(time.RFC1123)
		}
		if err != nil {
			return &repo.AllNewsProducts{}, err
		}
		res.NewProducts = append(res.NewProducts, &temp)

	}
	return &res, nil
}

func (n *NewsProductRepo) GetAllNews(params *repo.AllNewsProductParams) (*repo.AllNewsProducts, error) {
	var (
		create, update time.Time
		res            = repo.AllNewsProducts{}
	)
	query := `
	SELECT 
		id, 
		category_id, 
		title, 
		media_link, 
		description, 
		price, 
		created_at, 
		updated_at
	FROM new_products 
		WHERE deleted_at IS NULL AND 
		title ILIKE $1 or description ILIKE $2 
		LIMIT $3 OFFSET $4`
	rows, err := n.db.Pool.Query(context.Background(), query,
		"%"+params.Search+"%",
		"%"+params.Search+"%",
		params.Limit, params.Page,
	)
	if err != nil {
		return &repo.AllNewsProducts{}, err
	}
	for rows.Next() {
		temp := repo.NewsProductResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.CategoryId,
			&temp.Title,
			&temp.MediaLink,
			&temp.Description,
			&temp.Price,
			&create, &update,
		)
		temp.CreatedAt = create.Format(time.RFC1123)
		temp.UpdatedAt = update.Format(time.RFC1123)
		if err != nil {
			return &repo.AllNewsProducts{}, err
		}
		res.NewProducts = append(res.NewProducts, &temp)
	}
	return &res, nil
}

func (n *NewsProductRepo) UpdateNews(new *repo.NewsProductUpdateReq) (*repo.NewsProductResponse, error) {
	var (
		create, update time.Time
		res            = repo.NewsProductResponse{}
	)
	query := `
	UPDATE new_products SET 
		title=$1, 
		media_link=$2, 
		description=$3,
		updated_at=NOW(), 
		price=$4 WHERE 
		id=$5 AND deleted_at IS NULL
	RETURNING 
		id, category_id, 
		title, media_link, 
		description, price, 
		created_at, updated_at`
	err := n.db.Pool.QueryRow(context.Background(), query,
		new.Title, new.MediaLink, new.Description, new.Price, new.Id,
	).Scan(
		&res.Id,
		&res.CategoryId,
		&res.Title,
		&res.MediaLink,
		&res.Description,
		&res.Price,
		&create, &update)
	res.CreatedAt = create.Format(time.RFC1123)
	res.UpdatedAt = update.Format(time.RFC1123)
	if err != nil {
		return &repo.NewsProductResponse{}, err
	}

	return &res, nil
}

func (n *NewsProductRepo) DeleteNewsById(id int64) (*repo.Empty, error) {
	_, err := n.db.Pool.Exec(context.Background(),
		`UPDATE new_products SET deleted_at=NOW() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		fmt.Println("error while delete news ", err)
		return &repo.Empty{}, err
	}
	return &repo.Empty{}, nil
}

func (n *NewsProductRepo) DeleteNewsByCategoryId(id int64) (*repo.Empty, error) {
	_, err := n.db.Pool.Exec(context.Background(),
		`UPDATE new_products SET deleted_at=NOW() WHERE category_id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return &repo.Empty{}, err
	}
	return &repo.Empty{}, nil
}
