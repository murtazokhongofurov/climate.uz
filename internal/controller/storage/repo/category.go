package repo

import "time"

type CategoryRequest struct {
	CategoryName string
}

type CategoryResponse struct {
	Id            int
	CatergoryName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CategoryId struct {
	Id int
}
type CategoryUpdateReq struct {
	Id           int
	CategoryName string
}
type AllCategoriesParams struct {
	Page   int64
	Limit  int64
	Search string
}
type AllCategory struct {
	Categories []*CategoryResponse
}
type Empty struct{}

type CategoryStorageI interface {
	CreateCategory(c *CategoryRequest) (*CategoryResponse, error)
	GetCategoryById(id int) (*CategoryResponse, error)
	GetAllCategories(params *AllCategoriesParams) (*AllCategory, error)
	UpdateCategory(c *CategoryUpdateReq) (*CategoryResponse, error)
	DeleteCategoryById(id int) (*Empty, error)
}
