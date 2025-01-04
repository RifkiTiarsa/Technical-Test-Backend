package repository

import (
	"technical-test/internal/entity"
	"technical-test/internal/shared/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cartRepository struct {
	db          *gorm.DB
	userRepo    UserRepository
	productRepo ProductRepository
}

// // CreateCart implements CartRepository.
func (c *cartRepository) CreateCart(cart entity.Cart) (entity.CartResponse, error) {
	var finalCart entity.CartResponse

	c.db.Transaction(func(tx *gorm.DB) error {

		// Get product by id
		var product entity.Product
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, "id = ?", cart.ProductID).Error; err != nil {
			return err
		}

		// Validation quantity
		if cart.Quantity <= 0 {
			return common.ErrInvalidInput
		}

		// Check product stock
		if cart.Quantity > product.Stock {
			return common.ErrStockAvailability
		}

		// Check existing cart
		existingCart, err := c.GetCartByUserIdAndProductId(cart.UserID.String(), cart.ProductID.String())
		if err == nil {
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&existingCart, "id = ?", existingCart.ID).Error; err != nil {
				return err
			}

			if existingCart.Quantity > product.Stock {
				return err
			}
			existingCart.Quantity += cart.Quantity
			existingTotalPrice := float64(existingCart.Quantity) * product.Price
			if err := tx.Save(&existingCart).Error; err != nil {
				return err
			}
			finalCart = entity.CartResponse{
				ID:         existingCart.ID.String(),
				UserID:     cart.UserID,
				ProductID:  product.ID,
				Quantity:   existingCart.Quantity,
				TotalPrice: existingTotalPrice,
				CreatedAt:  existingCart.CreatedAt,
				UpdatedAt:  existingCart.UpdatedAt,
				DeletedAt:  existingCart.DeletedAt,
			}
		} else if err == gorm.ErrRecordNotFound {
			// Create new cart
			cart.ProductID = product.ID
			totalPrice := float64(cart.Quantity) * product.Price
			if err := tx.Create(&cart).Error; err != nil {
				return err
			}
			finalCart = entity.CartResponse{
				ID:         cart.ID.String(),
				UserID:     cart.UserID,
				ProductID:  product.ID,
				Quantity:   cart.Quantity,
				TotalPrice: totalPrice,
				CreatedAt:  cart.CreatedAt,
				UpdatedAt:  cart.UpdatedAt,
				DeletedAt:  cart.DeletedAt,
			}
		} else {
			return err
		}

		return nil
	})

	return finalCart, nil

}

// GetCartByUserIdAndProductId implements CartRepository.
func (c *cartRepository) GetCartByUserIdAndProductId(userId string, productId string) (entity.Cart, error) {
	var cart entity.Cart

	if err := c.db.Where("user_id = ? AND product_id = ?", userId, productId).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return cart, gorm.ErrRecordNotFound
		}
		return cart, err
	}

	return cart, nil
}

type CartRepository interface {
	CreateCart(cart entity.Cart) (entity.CartResponse, error)
	GetCartByUserIdAndProductId(userId, productId string) (entity.Cart, error)
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}
