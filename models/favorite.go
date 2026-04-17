package models

import "time"

type FavoriteProduct struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"uniqueIndex:idx_user_product" json:"user_id"`
	ProductID uint      `gorm:"uniqueIndex:idx_user_product" json:"product_id"`
	CreatedAt time.Time `json:"created_at"`

	User    User    `gorm:"foreignKey:UserID" json:"user"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}
