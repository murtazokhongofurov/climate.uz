package models

type ProductReq struct {
	CategoryId  int64   `json:"category_id"`
	Title       string  `json:"title"`
	MediaLink   string  `json:"media_link"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type UpdateProductReq struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	MediaLink   string  `json:"media_link"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductRes struct {
	Id          int64   `json:"id"`
	CategoryId  int64   `json:"category_id"`
	Title       string  `json:"title"`
	MediaLink   string  `json:"media_link"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type AllProducts struct {
	Products []ProductRes `json:"products"`
}
