package usecase

import (
	"technical-test/internal/entity"
	"technical-test/internal/repository"
	"technical-test/internal/shared/common"
)

type productUsecase struct {
	productRepo repository.ProductRepository
}

// ListProduct implements ProductUsecase.
func (p *productUsecase) ListProduct() ([]entity.Product, error) {
	return p.productRepo.ListProduct()
}

// GetProductByID implements ProductUsecase.
func (p *productUsecase) GetProductByID(id string) (entity.Product, error) {
	return p.productRepo.GetProductByID(id)
}

// GetProductByName implements ProductUsecase.
func (p *productUsecase) GetProductByName(name string) (entity.Product, error) {
	product, err := p.productRepo.GetProductByName(name)
	if err != nil {
		return entity.Product{}, common.ErrNotFound
	}

	return product, nil
}

type ProductUsecase interface {
	ListProduct() ([]entity.Product, error)
	GetProductByID(id string) (entity.Product, error)
	GetProductByName(name string) (entity.Product, error)
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: productRepo}
}
