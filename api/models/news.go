package models

type NewsProductReq struct {
	CategoryId  int64   `json:"category_id"`
	Title       string  `json:"title"`
	MediaLink   string  `json:"media_link"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type NewsProductRes struct {
	Id          int32   `json:"id"`
	CategoryId  int64   `json:"category_id"`
	Title       string  `json:"title"`
	MediaLink   string  `json:"media_link"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type AllNewsProducts struct {
	NewsProducts []NewsProductRes `json:"news_products"`
}

type UpdateNewsProductReq struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	MediaLink   string  `json:"media_link"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
