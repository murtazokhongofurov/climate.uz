package models

type UserReq struct {
	PhoneNumber string `json:"phone_number" example:"(998)"`
}

type UserRes struct {
	Id          string `json:"id" example:"98fe8389-8d29-4e4c-838e-782818e6b966"`
	PhoneNumber string `json:"phone_number" example:"(998)91 900-89-99"`
	CreatedAt   string `json:"created_at" example:"Mon, 15 Jan 2023 07:41:56 UTC"`
	UpdateAt    string `json:"updated_at" example:"Mon, 15 Jan 2023 07:41:56 UTC"`
}

type UpdateUserReq struct {
	Id          string `json:"id" example:"uuid"`
	PhoneNumber string `json:"phone_number"`
}
