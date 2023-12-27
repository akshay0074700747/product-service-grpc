package adapters

import "github.com/akshay0074700747/products-service/entities"

type AdapterInterface interface {
	AddProduct(req entities.Products) (entities.Products, error)
	GetProduct(id uint) (entities.Products, error)
	GetAllProducts() ([]entities.Products, error)
	IncrementStock(id uint, stock int) (entities.Products, error)
	DecrementStock(id uint, stock int) (entities.Products, error)
}
