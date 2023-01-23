package storage

import (
	"gitlab.com/climate.uz/internal/controller/storage/postgres"
	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/db"
)

type StorageI interface {
	Admin() repo.AdminStorageI
	Category() repo.CategoryStorageI
	Product() repo.ProductStorageI
	NewProduct() repo.NewsStorageI
	User() repo.UserStorageI
}

type StoragePg struct {
	Adminrepo       repo.AdminStorageI
	Categoryrepo    repo.CategoryStorageI
	Productrepo     repo.ProductStorageI
	Userrepo        repo.UserStorageI
	NewsProductrepo repo.NewsStorageI
}

func NewStoragePg(db *db.Postgres) StorageI {
	return &StoragePg{
		Adminrepo:       postgres.NewAdmin(db),
		Categoryrepo:    postgres.NewCategory(db),
		Productrepo:     postgres.NewProduct(db),
		NewsProductrepo: postgres.NewNewsProduct(db),
		Userrepo:        postgres.NewUser(db),
	}
}

func (a *StoragePg) Admin() repo.AdminStorageI {
	return a.Adminrepo
}

func (s *StoragePg) Category() repo.CategoryStorageI {
	return s.Categoryrepo
}

func (s *StoragePg) Product() repo.ProductStorageI {
	return s.Productrepo
}

func (s *StoragePg) NewProduct() repo.NewsStorageI {
	return s.NewsProductrepo
}

func (s *StoragePg) User() repo.UserStorageI {
	return s.Userrepo
}
