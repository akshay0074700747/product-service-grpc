package adapters

import "gorm.io/gorm"

type ProductAdapter struct {
	DB *gorm.DB
}

func NewProductAdapter(db *gorm.DB) *ProductAdapter {
	return &ProductAdapter{
		DB: db,
	}
}
