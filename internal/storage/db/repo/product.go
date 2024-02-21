package repo

import (
	"database/sql"
	"errors"
	"shop/internal/models"
)

type ProductRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product
	rows, err := r.DB.Query("SELECT id, title, description, price, category_id, created_at, updated_at FROM product")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product models.Product
		err = rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return &products, nil
}

func (r *ProductRepo) GetProductByID(id uint64) (*models.Product, error) {
	var product models.Product
	err := r.DB.QueryRow("SELECT id, title, description, price, category_id, created_at, updated_at FROM product WHERE id = ?", id).
		Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepo) CreateProduct(product *models.Product) (uint64, error) {
	if product == nil {
		return 0, errors.New("no product provided")
	}
	if r.TX != nil {
		stmt, err := r.TX.Prepare("INSERT INTO product(title, description, price, category_id) VALUES(?, ?, ?, ?)")
		if err != nil {
			return 0, err
		}
		_, err = stmt.Exec(
			&product.Title,
			&product.Description,
			&product.Price,
			&product.CategoryID,
		)
		if err != nil {
			return 0, err
		}
		return product.ID, nil
	}
	stmt, err := r.DB.Prepare("INSERT INTO product(title, description, price, category_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	_, err = stmt.Exec(
		&product.Title,
		&product.Description,
		&product.Price,
		&product.CategoryID,
	)
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (r *ProductRepo) UpdateProduct(product *models.Product) (uint64, error) {
	if product == nil {
		return 0, errors.New("no product provided")
	}
	if r.TX != nil {
		stmt, err := r.TX.Prepare("UPDATE product SET title = ?, description = ?, price = ?, category_id = ? WHERE id = ?")
		if err != nil {
			return 0, err
		}
		_, err = stmt.Exec(product.Title, product.Description, product.Price, product.CategoryID, product.ID)
		if err != nil {
			return 0, err
		}
		return product.ID, nil
	}
	stmt, err := r.DB.Prepare("UPDATE product SET title = ?, description = ?, price = ?, category_id = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	_, err = stmt.Exec(product.Title, product.Description, product.Price, product.CategoryID, product.ID)
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (r *ProductRepo) DeleteProduct(id uint64) (uint64, error) {
	if r.TX != nil {
		_, err := r.TX.Exec("DELETE FROM products WHERE id = ?", id)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	_, err := r.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *ProductRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Commit()
	}
	return nil
}

func (r *ProductRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.TX.Rollback()
	}
	return nil
}

func (r *ProductRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *ProductRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}
