package repo

type UserRequest struct {
	PhoneNumber string
}

type UserResponse struct {
	Id          string
	PhoneNumber string
	CreatedAt   string
	UpdatedAt   string
}

type UserId struct {
	Id string
}

type UserUpdateReq struct {
	Id          string
	PhoneNumber string
}

type AllUsersParams struct {
	Page   int64
	Limit  int64
	Search string
}

type AllUsers struct {
	Users []*UserResponse
}

type UserStorageI interface {
	CreateUser(u *UserRequest) (*UserResponse, error)
	UpdateUser(u *UserUpdateReq) (*UserResponse, error)
	GetUserById(id string) (*UserResponse, error)
	GetAllUser(params *AllUsersParams) (*AllUsers, error)
	DeleteUser(id string) (*Empty, error)
}
