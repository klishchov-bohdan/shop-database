package storage

import (
	"database/sql"
	"shop/internal/storage/db/repo"
)

type Storage struct {
	Products      *repo.ProductRepo
	Orders        *repo.OrderRepo
	OrderProducts *repo.OrderProductRepo
}

func NewStorage(db *sql.DB) *Storage {
	pr := repo.NewProductRepo(db)
	or := repo.NewOrderRepo(db)
	opr := repo.NewOrderProductRepo(db)

	return &Storage{
		Products:      pr,
		Orders:        or,
		OrderProducts: opr,
	}
}
