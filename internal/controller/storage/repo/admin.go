package repo

type AdminRequest struct {
	Id       string
	UserName string
	Password string
}

type AdminResponse struct {
	Id          string
	UserName    string
	Password    string
	AccessToken string
	CreatedAt   string
	UpdatedAt   string
}

type UpdateAdminReq struct {
	Id       string
	UserName string
	Password string
}

type AllAdminParams struct {
	Page   int64
	Limit  int64
	Search string
}

type AllAdmins struct {
	Admins []AdminResponse
}

type CheckFieldReq struct {
	Field string
	Value string
}

type CheckFieldRes struct {
	Exists bool
}

type AdminStorageI interface {
	// CreateAdmin(a *AdminRequest) (*AdminResponse, error)
	// GetAdminById(id string) (*AdminResponse, error)
	// GetAllAdminRequest(a *AllAdminParams) (*AllAdmins, error)
	// UpdateAdmin(a *UpdateAdminReq) (*AdminResponse, error)
	// DeleteAdmin(id string) (*Empty, error)
	// CheckField(a *CheckFieldReq) (*CheckFieldRes, error)
	GetAdmin(username string) (*AdminResponse, error)
}
