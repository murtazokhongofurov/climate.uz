package models

type AdminLoginReq struct {
	UserName string `json:"AdminName"`
	Password string `json:"password"`
}
