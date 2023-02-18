package repo

type BrandRequst struct {
	BrandName string
	Logo      string
}
type BrandResponse struct {
	Id        int64
	BrandName string
	Logo      string
}

type GetBrandResponse struct {
	Id         int64
	BrandName  string
	Logo       string
	Categories []*CategoryResponse
}

type BrandId struct {
	Id int64
}

type AllBrands struct {
	Brands []*BrandResponse
}

type ParamBrands struct {
	Page    int64
	Limit   int64
	Keyword string
}

type BrandStorageI interface {
	CreateBrand(b *BrandRequst) (*BrandResponse, error)
	GetBrandById(b *BrandId) (*GetBrandResponse, error)
	GetBrandAll(p *ParamBrands) (*AllBrands, error)
	DeleteBrand(b *BrandId) (*Empty, error)
}
