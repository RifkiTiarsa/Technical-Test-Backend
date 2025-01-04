package entity

import (
	"time"

	"github.com/google/uuid"
)

type Checkout struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CartID           uuid.UUID `json:"cart_id" gorm:"type:uuid;not null"`
	Cart             Cart      `json:"cart" gorm:"foreignKey:CartID;references:ID"`
	Amount           float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	PaymentStatus    string    `json:"payment_status" gorm:"type:varchar(20);default:'pending';not null"`
	PaymentMethod    string    `json:"payment_method" gorm:"type:varchar(50);not null"`
	Address          string    `json:"address" gorm:"type:text;not null"`
	LogisticProvider string    `json:"logistic_provider" gorm:"type:varchar(20);not null"`
	ShippingStatus   string    `json:"shipping_status" gorm:"type:varchar(20);default:'menunggu pembayaran';not null"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}

type CheckoutResponse struct {
	ID               string    `json:"id"`
	CartID           uuid.UUID `json:"cart_id"`
	Amount           float64   `json:"amount"`
	PaymentStatus    string    `json:"payment_status"`
	PaymentMethod    string    `json:"payment_method"`
	Address          string    `json:"address"`
	LogisticProvider string    `json:"logistic_provider"`
	ShippingStatus   string    `json:"shipping_status"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
}
