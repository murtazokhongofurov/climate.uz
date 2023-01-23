package repo

type ProductRequest struct {
	CategoryId  int64
	Title       string
	MeidaLink   string
	Description string
	Price       float64
}

type ProductResponse struct {
	Id          int64
	CategoryId  int64
	Title       string
	MediaLink   string
	Description string
	Price       float64
	CreatedAt   string
	UpdatedAt   string
}

type ProductId struct {
	Id int64
}

type AllProductsParams struct {
	Page   int64
	Limit  int64
	Search string
}

type AllProducts struct {
	Products []*ProductResponse
}
type ProductUpdateReq struct {
	Id          int64
	Title       string
	MediaLink   string
	Description string
	Price       float64
}

type ProductStorageI interface {
	CreateProduct(p *ProductRequest) (*ProductResponse, error)
	UpdateProduct(p *ProductUpdateReq) (*ProductResponse, error)
	GetProductById(id int64) (*ProductResponse, error)
	GetAllProducts(params *AllProductsParams) (*AllProducts, error)
	DeleteProductById(id int64) (*Empty, error)
	DeleteProductByCategoryId(id int64) (*Empty, error)
}
