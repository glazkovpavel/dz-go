package product

import (
	"go/order-api/pkg/db"
	"gorm.io/gorm/clause"
)

type Repository struct {
	Database *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{
		Database: db,
	}
}

func (repo *Repository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *Repository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *Repository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *Repository) FindById(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *Repository) FindAll() ([]*Product, error) {
	var products []*Product
	result := repo.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *Repository) FindWithPagination(limit, offset int) ([]*Product, error) {
	var products []*Product
	result := repo.Database.DB.Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
