package models

type Order struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	CustomerID uint    `json:"customer_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
