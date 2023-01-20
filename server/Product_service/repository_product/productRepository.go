package repository_product

import (
	"context"

	"github.com/IRSHIT033/E-comm-GO-/server/Product_service/domain_product"
	"gorm.io/gorm"
)

type productRepository struct {
	database *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain_product.ProductRepository {
	return &productRepository{
		database: db,
	}
}

func (pr *productRepository) Create(c context.Context, product domain_product.Product) error {
	err := pr.database.Create(&product).Error
	return err
}
func (pr *productRepository) Fetch(c context.Context) ([]domain_product.Product, error) {
	var products []domain_product.Product
	err := pr.database.Find(&products).Error

	return products, err
}
func (pr *productRepository) FetchbyId(c context.Context, productID uint) (domain_product.Product, error) {
	var product domain_product.Product
	err := pr.database.Find(&product, "id = ?", productID).Error
	return product, err
}
func (pr *productRepository) FetchbyCategory(c context.Context, categories []string) ([]domain_product.Product, error) {
	var products []domain_product.Product
	err := pr.database.Find(&products, "category IN ?", categories).Error
	return products, err
}
