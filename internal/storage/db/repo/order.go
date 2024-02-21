package repo

import (
	"database/sql"
	"shop/internal/models"
	"strconv"
)

type OrderRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func (r *OrderRepo) GetAllOrders() (*[]models.Order, error) {
	var Orders []models.Order
	rows, err := r.DB.Query("SELECT id, total_price, created_at, updated_at FROM `order`")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Order models.Order
		err = rows.Scan(
			&Order.ID,
			&Order.TotalPrice,
			&Order.CreatedAt,
			&Order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Orders = append(Orders, Order)
	}
	return &Orders, nil
}

func (r *OrderRepo) GetManyOrdersByIds(ids ...int) (*[]models.Order, error) {
	var Orders []models.Order
	whereStr := " WHERE "
	for idx, id := range ids {
		whereStr += "id = " + strconv.Itoa(id)
		if idx < len(ids)-1 {
			whereStr += " || "
		}
	}
	rows, err := r.DB.Query("SELECT id, total_price, created_at, updated_at FROM `order`" + whereStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Order models.Order
		err = rows.Scan(
			&Order.ID,
			&Order.TotalPrice,
			&Order.CreatedAt,
			&Order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		Orders = append(Orders, Order)
	}
	return &Orders, nil
}
