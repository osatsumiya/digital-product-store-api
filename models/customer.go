package models

type Customer struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `json:"full_name"`
	Email    string `gorm:"unique" json:"email"`
}
