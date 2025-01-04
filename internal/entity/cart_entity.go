package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID    `json:"user_id" gorm:"type:uuid;not null"`
	User      User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	ProductID uuid.UUID    `json:"product_id" gorm:"type:uuid;not null"`
	Product   Product      `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int          `json:"quantity" gorm:"type:integer;not null"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty" gorm:"index"`
}

type CartResponse struct {
	ID         string       `json:"id"`
	UserID     uuid.UUID    `json:"user_id"`
	ProductID  uuid.UUID    `json:"product_id"`
	Quantity   int          `json:"quantity"`
	TotalPrice float64      `json:"total_price"`
	CreatedAt  time.Time    `json:"created_at,omitempty"`
	UpdatedAt  time.Time    `json:"updated_at,omitempty"`
	DeletedAt  sql.NullTime `json:"deleted_at,omitempty"`
}
