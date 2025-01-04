package repository

import (
	"fmt"
	"technical-test/internal/entity"
	"technical-test/internal/shared/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type checkoutRepository struct {
	db *gorm.DB
}

// CreateCheckout implements CheckoutRepository.
func (c *checkoutRepository) CreateCheckout(checkout entity.Checkout) (entity.CheckoutResponse, error) {
	var newCheckout entity.CheckoutResponse

	err := c.db.Transaction(func(tx *gorm.DB) error {
		// Get cart
		var cart entity.Cart
		if err := tx.Where("id =?", checkout.CartID).First(&cart).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("cart not found: %w", err)
			}
			return fmt.Errorf("failed to retrieve cart: %w", err)
		}

		// locking
		var product entity.Product
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id =?", cart.ProductID).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("product not found: %w", err)
			}
			return fmt.Errorf("failed to retrieve product: %w", err)
		}

		// Validation quantity
		if cart.Quantity > product.Stock {
			return common.ErrStockAvailability
		}

		// Update stock product
		product.Stock -= cart.Quantity
		if err := tx.Save(&product).Error; err != nil {
			return fmt.Errorf("failed to update product stock: %w", err)
		}

		// Get amount
		checkout.Amount = float64(cart.Quantity) * product.Price

		// Set payment status
		checkout.PaymentStatus = "pending"

		// Require field
		if checkout.PaymentMethod == "" || checkout.Address == "" || checkout.LogisticProvider == "" {
			return common.ErrInvalidInput
		}

		// Set shipping status
		checkout.ShippingStatus = "menunggu pembayaran"

		// Create checkout
		if err := tx.Create(&checkout).Error; err != nil {
			return fmt.Errorf("failed to create checkout: %w", err)
		}

		newCheckout = entity.CheckoutResponse{
			ID:               checkout.ID.String(),
			CartID:           checkout.CartID,
			Amount:           checkout.Amount,
			PaymentStatus:    checkout.PaymentStatus,
			PaymentMethod:    checkout.PaymentMethod,
			Address:          checkout.Address,
			LogisticProvider: checkout.LogisticProvider,
			ShippingStatus:   checkout.ShippingStatus,
			CreatedAt:        checkout.CreatedAt,
		}

		return nil
	})

	if err != nil {
		return entity.CheckoutResponse{}, err
	}

	return newCheckout, nil
}

type CheckoutRepository interface {
	CreateCheckout(checkout entity.Checkout) (entity.CheckoutResponse, error)
}

func NewCheckoutRepository(db *gorm.DB) CheckoutRepository {
	return &checkoutRepository{db: db}
}
