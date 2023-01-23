package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/db"
)

type ProductRepo struct {
	db *db.Postgres
}

func NewProduct(db *db.Postgres) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) CreateProduct(product *repo.ProductRequest) (*repo.ProductResponse, error) {
	var (
		create, update time.Time
		res            = repo.ProductResponse{}
	)
	query := `INSERT INTO products(
			category_id, 
			title, 
			media_link, 
			description, 
			price) 
			VALUES($1, $2, $3, $4, $5) 
			RETURNING 
			id, category_id, title, media_link, 
			description, price, created_at, updated_at`

	err := p.db.Pool.QueryRow(context.Background(), query,
		product.CategoryId,
		product.Title,
		product.MeidaLink,
		product.Description,
		product.Price).Scan(
		&res.Id,
		&res.CategoryId,
		&res.Title,
		&res.MediaLink,
		&res.Description,
		&res.Price, &create, &update,
	)
	res.CreatedAt = create.Format(time.RFC1123)
	res.UpdatedAt = update.Format(time.RFC1123)
	if err != nil {
		fmt.Println("error >> ", err)
		return &repo.ProductResponse{}, err
	}
	return &res, nil
}

func (p *ProductRepo) GetProductById(id int64) (*repo.ProductResponse, error) {
	var (
		create, update time.Time
		res            = repo.ProductResponse{}
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
		FROM products 
		WHERE id=$1 AND deleted_at IS NULL`
	err := p.db.Pool.QueryRow(context.Background(), query, id).Scan(
		&res.Id, &res.CategoryId,
		&res.Title, &res.MediaLink,
		&res.Description, &res.Price, &create, &update,
	)
	res.CreatedAt = create.Format(time.RFC1123)
	res.UpdatedAt = update.Format(time.RFC1123)
	if err == sql.ErrNoRows {
		return &repo.ProductResponse{}, nil
	}
	return &res, nil
}

func (p *ProductRepo) UpdateProduct(product *repo.ProductUpdateReq) (*repo.ProductResponse, error) {
	var (
		create_time, update_time time.Time
		res                      = repo.ProductResponse{}
	)
	query := `UPDATE 
		products SET 
		title=$1, 
		media_link=$2, 
		description=$3, 
		price=$4, 
		updated_at=NOW() 
		WHERE id=$5 AND deleted_at IS NULL
		RETURNING 
		id, category_id, title, media_link, 
		description, price, created_at, updated_at`
	err := p.db.Pool.QueryRow(context.Background(), query,
		product.Title,
		product.MediaLink,
		product.Description,
		product.Price, product.Id).Scan(
		&res.Id,
		&res.CategoryId,
		&res.Title,
		&res.MediaLink,
		&res.Description,
		&res.Price,
		&create_time,
		&update_time,
	)
	res.CreatedAt = create_time.Format(time.RFC1123)
	res.UpdatedAt = update_time.Format(time.RFC1123)
	if err != nil {
		return &repo.ProductResponse{}, err
	}

	return &res, nil
}

func (p *ProductRepo) GetAllProducts(product *repo.AllProductsParams) (*repo.AllProducts, error) {
	var (
		create, update time.Time
		res            = repo.AllProducts{}
	)
	query := `SELECT 
			id, category_id, title, 
			media_link, description, price, 
			created_at, updated_at FROM products
			WHERE deleted_at IS NULL AND 
			title ILIKE $1 OR description ILIKE $2 LIMIT $3 OFFSET $4
			`
	rows, err := p.db.Pool.Query(context.Background(), query,
		"%"+product.Search+"%", "%"+product.Search+"%",
		product.Limit, (product.Page-1)*product.Limit)
	if err != nil {
		fmt.Println("error while getting all products -> ", err)
		return &repo.AllProducts{}, err
	}
	for rows.Next() {
		temp := repo.ProductResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.CategoryId,
			&temp.Title,
			&temp.MediaLink,
			&temp.Description,
			&temp.Price, &create, &update)
		temp.CreatedAt = create.Format(time.RFC1123)
		temp.UpdatedAt = update.Format(time.RFC1123)
		if err != nil {
			return &repo.AllProducts{}, err
		}
		res.Products = append(res.Products, &temp)
	}
	return &res, nil
}

func (p *ProductRepo) DeleteProductById(id int64) (*repo.Empty, error) {
	_, err := p.db.Pool.Exec(context.Background(),
		`UPDATE products SET deleted_at=NOW() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return &repo.Empty{}, err
	}
	return &repo.Empty{}, nil
}

func (p *ProductRepo) DeleteProductByCategoryId(category_id int64) (*repo.Empty, error) {
	_, err := p.db.Pool.Exec(context.Background(),
		`UPDATE products SET deleted_at=NOW() WHERE category_id=$1 AND deleted_at IS NULL`, category_id)
	if err != nil {
		return &repo.Empty{}, err
	}
	return &repo.Empty{}, nil
}
