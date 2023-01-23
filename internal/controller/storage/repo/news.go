package repo

type NewsProductRequest struct {
	CategoryId  int64
	Title       string
	MediaLink   string
	Description string
	Price       float64
}

type NewsProductResponse struct {
	Id          int64
	CategoryId  int64
	Title       string
	MediaLink   string
	Description string
	Price       float64
	CreatedAt   string
	UpdatedAt   string
}

type NewsProductId struct {
	Id int64
}

type NewsProductUpdateReq struct {
	Id          int64
	Title       string
	MediaLink   string
	Description string
	Price       float64
}

type AllNewsProductParams struct {
	Page   int64
	Limit  int64
	Search string
}

type AllNewsProducts struct {
	NewProducts []*NewsProductResponse
}

type NewsStorageI interface {
	CreateNews(n *NewsProductRequest) (*NewsProductResponse, error)
	UpdateNews(n *NewsProductUpdateReq) (*NewsProductResponse, error)
	GetNewsById(id int64) (*NewsProductResponse, error)
	GetNewsByCategoryId(id int64) (*AllNewsProducts, error)
	GetAllNews(params *AllNewsProductParams) (*AllNewsProducts, error)
	DeleteNewsById(id int64) (*Empty, error)
	DeleteNewsByCategoryId(id int64) (*Empty, error)
}
