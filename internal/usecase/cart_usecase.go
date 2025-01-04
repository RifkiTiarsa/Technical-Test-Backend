package usecase

import (
	"technical-test/internal/entity"
	"technical-test/internal/repository"
)

type cartUsecase struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
	userRepo    repository.UserRepository
}

// AddCart implements CartUsecase.
func (c *cartUsecase) AddCart(cart entity.Cart) (entity.CartResponse, error) {
	return c.cartRepo.CreateCart(cart)
}

type CartUsecase interface {
	AddCart(cart entity.Cart) (entity.CartResponse, error)
}

func NewCartUsecase(cartRepo repository.CartRepository) CartUsecase {
	return &cartUsecase{cartRepo: cartRepo}
}
