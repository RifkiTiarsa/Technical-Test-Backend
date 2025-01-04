package repository

import (
	"technical-test/internal/entity"
	"technical-test/internal/shared/common"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// ListProduct implements ProductRepository.
func (p *productRepository) ListProduct() ([]entity.Product, error) {
	var products []entity.Product

	rows, err := p.db.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON p.category_id = c.id").
		Select("p.id, p.name, p.description, c.id, c.name AS category_name, p.price, p.stock, p.rating, p.created_at, p.updated_at, p.deleted_at").
		Where("p.deleted_at IS NULL").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		var category entity.Category
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &category.ID, &category.Name, &product.Price, &product.Stock, &product.Rating, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}

		product.CategoryID = category.ID
		product.Category = category
		products = append(products, product)
	}

	return products, nil

}

// GetProductByID implements ProductRepository.
func (p *productRepository) GetProductByID(id string) (entity.Product, error) {
	var product entity.Product
	var category entity.Category

	row := p.db.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON p.category_id = c.id").
		Select("p.id, p.name, p.description, c.id, c.name AS category_name, p.price, p.stock, p.rating, p.created_at, p.updated_at, p.deleted_at").
		Where("p.id = ? AND p.deleted_at IS NULL", id).
		Row()

	if err := row.Scan(&product.ID, &product.Name, &product.Description, &category.ID, &category.Name, &product.Price, &product.Stock, &product.Rating, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
		return entity.Product{}, common.ErrNotFound
	}

	product.Category = entity.Category{
		ID:   category.ID,
		Name: category.Name,
	}

	return product, nil
}

// GetProductByName implements ProductRepository.
func (p *productRepository) GetProductByName(name string) (entity.Product, error) {
	var product entity.Product
	var category entity.Category

	row := p.db.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON p.category_id = c.id").
		Select("p.id, p.name, p.description, c.id, c.name AS category_name, p.price, p.stock, p.rating, p.created_at, p.updated_at, p.deleted_at").
		Where("p.name ILIKE ? AND p.deleted_at IS NULL", "%"+name+"%").
		Row()

	if err := row.Scan(&product.ID, &product.Name, &product.Description, &category.ID, &category.Name, &product.Price, &product.Stock, &product.Rating, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
		return entity.Product{}, common.ErrNotFound
	}

	product.Category = entity.Category{
		ID:   category.ID,
		Name: category.Name,
	}

	return product, nil
}

type ProductRepository interface {
	ListProduct() ([]entity.Product, error)
	GetProductByID(id string) (entity.Product, error)
	GetProductByName(name string) (entity.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
