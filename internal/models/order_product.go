package models

type OrderProduct struct {
	ID                 uint64
	OrderID            uint64
	ProductID          uint64
	ProductTitle       string
	ProductDescription string
	ProductCategory    string
	ProductQuantity    uint64
}
