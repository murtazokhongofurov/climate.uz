package postgres

import (
	"context"

	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/db"
)

type BrandRepo struct {
	db *db.Postgres
}

func NewBrand(db *db.Postgres) *BrandRepo {
	return &BrandRepo{
		db: db,
	}
}

func (b *BrandRepo) CreateBrand(brand *repo.BrandRequst) (*repo.BrandResponse, error) {
	var (
		res repo.BrandResponse
	)
	query := `
	INSERT INTO 
		brands(brand_name, logo)
	VALUES
		($1, $2)
	RETURNING 
		id, brand_name, logo
	`
	err := b.db.Pool.QueryRow(context.Background(), query, brand.BrandName, brand.Logo).
		Scan(&res.Id, &res.BrandName, &res.Logo)
	if err != nil {
		return &repo.BrandResponse{}, err
	}

	query2 := `
	INSERT INTO 
		brand_category(brand_id)
	VALUES
		($1)`
	_, err = b.db.Pool.Exec(context.Background(), query2, res.Id)
	if err != nil {
		return &repo.BrandResponse{}, err
	}
	return &res, nil
}

func (b *BrandRepo) GetBrandById(brand *repo.BrandId) (*repo.GetBrandResponse, error) {
	var (
		res repo.GetBrandResponse
	)
	query := `
	SELECT 
		id, brand_name, logo
	FROM 
		brands
	WHERE 
		id=$1
	`
	err := b.db.Pool.QueryRow(context.Background(), query, brand.Id).
		Scan(&res.Id, &res.BrandName, &res.Logo)
	if err != nil {
		return &repo.GetBrandResponse{}, err
	}

	query2 := `
	SELECT 
		c.id, c.category_name
	FROM 
		categories c
	INNER JOIN 
		brand_category bc ON bc.category_id=c.id
	WHERE 
		bc.brand_id=$1`

	rows, err := b.db.Pool.Query(context.Background(), query2, res.Id)
	if err != nil {
		return &repo.GetBrandResponse{}, err
	}
	for rows.Next() {
		temp := repo.CategoryResponse{}
		err = rows.Scan(&temp.Id, &temp.CatergoryName)
		if err != nil {
			return &repo.GetBrandResponse{}, err
		}
		res.Categories = append(res.Categories, &temp)
	}
	return &res, nil

}

func (b *BrandRepo) GetBrandAll(brand *repo.ParamBrands) (*repo.AllBrands, error) {
	var (
		res repo.AllBrands
	)
	query := `
	SELECT 
		id, brand_name
	FROM 
		brands
	WHERE
		brand_name ILIKE $1 
	ORDER BY
		brand_name ASC
	LIMIT 
		$2 OFFSET $3 
	`
	rows, err := b.db.Pool.Query(context.Background(),
		query, "%"+brand.Keyword+"%", brand.Limit, (brand.Page-1)*brand.Limit)
	if err != nil {
		return &repo.AllBrands{}, err
	}
	for rows.Next() {
		var temp repo.BrandResponse
		err = rows.Scan(&temp.Id, &temp.BrandName, &temp.Logo)
		if err != nil {
			return &repo.AllBrands{}, err
		}
		res.Brands = append(res.Brands, &temp)
	}

	return &res, nil

}

func (b *BrandRepo) DeleteBrand(brand *repo.BrandId) (*repo.Empty, error) {
	query := `
	DELETE FROM 
		brands
	WHERE 
		id=$1
	`
	_, err := b.db.Pool.Exec(context.Background(), query, brand.Id)
	if err != nil {
		return &repo.Empty{}, err
	}
	query2 := `
	DELETE FROM 
		brand_category
	WHERE 
		brand_id=$1`
	_, err = b.db.Pool.Exec(context.Background(), query2, brand.Id)
	if err != nil {
		return &repo.Empty{}, err
	}
	return &repo.Empty{}, nil
}
