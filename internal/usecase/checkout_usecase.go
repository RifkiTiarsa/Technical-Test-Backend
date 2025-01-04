package usecase

import (
	"technical-test/internal/entity"
	"technical-test/internal/repository"
)

type checkoutUsecase struct {
	checkoutRepo repository.CheckoutRepository
}

// CreateCheckout implements CheckoutUsecase.
func (c *checkoutUsecase) CreateCheckout(checkout entity.Checkout) (entity.CheckoutResponse, error) {
	return c.checkoutRepo.CreateCheckout(checkout)
}

type CheckoutUsecase interface {
	CreateCheckout(checkout entity.Checkout) (entity.CheckoutResponse, error)
}

func NewCheckoutUsecase(checkoutRepo repository.CheckoutRepository) CheckoutUsecase {
	return &checkoutUsecase{checkoutRepo: checkoutRepo}
}
