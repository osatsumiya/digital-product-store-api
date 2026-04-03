package models

type License struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ProductID  uint   `json:"product_id"`
	CustomerID uint   `json:"customer_id"`
	LicenseKey string `json:"license_key"`
	IsActive   bool   `json:"is_active"`
}
