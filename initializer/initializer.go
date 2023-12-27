package initializer

import (
	"github.com/akshay0074700747/products-service/adapters"
	"github.com/akshay0074700747/products-service/service"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) *service.ProductService {

	adapter := adapters.NewProductAdapter(db)
	service := service.NewProductService(adapter)

	return service
}
