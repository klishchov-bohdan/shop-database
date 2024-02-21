package repo

import (
	"database/sql"
	"shop/internal/models"
)

type OrderProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewOrderProductRepo(db *sql.DB) *OrderProductRepo {
	return &OrderProductRepo{DB: db}
}

func (r *OrderProductRepo) GetAllOrderProducts() (*[]models.OrderProduct, error) {
	var OrderProducts []models.OrderProduct
	query := "SELECT DISTINCT `order_product`.id, `order`.id, `product`.id, `product`.title, `product`.description, `category`.title, `order_product`.product_quantity " +
		"FROM `order` " +
		"JOIN `order_product` ON `order`.id = `order_product`.order_id " +
		"JOIN `product` ON `order_product`.product_id = `product`.id " +
		"JOIN `category` ON `product`.category_id = `category`.id"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var OrderProduct models.OrderProduct
		err = rows.Scan(
			&OrderProduct.ID,
			&OrderProduct.OrderID,
			&OrderProduct.ProductID,
			&OrderProduct.ProductTitle,
			&OrderProduct.ProductDescription,
			&OrderProduct.ProductCategory,
			&OrderProduct.ProductQuantity,
		)
		if err != nil {
			return nil, err
		}
		OrderProducts = append(OrderProducts, OrderProduct)
	}
	return &OrderProducts, nil
}

func (r *OrderProductRepo) GetManyOrderProductsByIds(ids ...string) (*[]models.OrderProduct, error) {
	var OrderProducts []models.OrderProduct
	whereStr := " WHERE "
	for idx, id := range ids {
		whereStr += "`order`.id = " + id
		if idx < len(ids)-1 {
			whereStr += " || "
		}
	}
	query := "SELECT DISTINCT `order_product`.id, `order`.id, `product`.id, `product`.title, `product`.description, `category`.title, `order_product`.product_quantity " +
		"FROM `order` " +
		"JOIN `order_product` ON `order`.id = `order_product`.order_id " +
		"JOIN `product` ON `order_product`.product_id = `product`.id " +
		"JOIN `category` ON `product`.category_id = `category`.id" + whereStr
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var OrderProduct models.OrderProduct
		err = rows.Scan(
			&OrderProduct.ID,
			&OrderProduct.OrderID,
			&OrderProduct.ProductID,
			&OrderProduct.ProductTitle,
			&OrderProduct.ProductDescription,
			&OrderProduct.ProductCategory,
			&OrderProduct.ProductQuantity,
		)
		if err != nil {
			return nil, err
		}
		OrderProducts = append(OrderProducts, OrderProduct)
	}
	return &OrderProducts, nil
}
