package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string       `json:"username" gorm:"type:varchar(255);not null"`
	Email     string       `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string       `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty" gorm:"index"`
}
