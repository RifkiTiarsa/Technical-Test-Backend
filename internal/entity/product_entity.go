package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string       `json:"name" gorm:"type:varchar(255);not null"`
	Description string       `json:"description" gorm:"type:text;not null"`
	CategoryID  uuid.UUID    `json:"-" gorm:"type:uuid;not null"`
	Category    Category     `json:"category" gorm:"-"`
	Price       float64      `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int          `json:"stock" gorm:"type:integer;not null"`
	Rating      float64      `json:"rating" gorm:"type:decimal(1,1);default:0"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
	DeletedAt   sql.NullTime `json:"deleted_at,omitempty" gorm:"index"`
}

type Category struct {
	ID        uuid.UUID    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string       `json:"name" gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
