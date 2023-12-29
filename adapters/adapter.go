package adapters

import (
	"fmt"

	"github.com/akshay0074700747/products-service/entities"
)

func (product *ProductAdapter) AddProduct(req entities.Products) (entities.Products, error) {

	var res entities.Products
	query := "INSERT INTO products (name,price,stock) VALUES($1,$2,$3) RETURNING id,name,price,stock"

	return res, product.DB.Raw(query, req.Name, req.Price, req.Stock).Scan(&res).Error
}

func (product *ProductAdapter) GetProduct(id uint) (entities.Products, error) {

	var res entities.Products
	query := "SELECT * FROM products WHERE id = $1"

	return res, product.DB.Raw(query, id).Scan(&res).Error
}

func (product *ProductAdapter) GetAllProducts() ([]entities.Products, error) {

	var res []entities.Products
	query := "SELECT * FROM products"

	return res, product.DB.Raw(query).Scan(&res).Error
}

func (product *ProductAdapter) IncrementStock(id uint, stock int) (entities.Products, error) {

	var res entities.Products
	query := "UPDATE products SET stock = stock + $1 WHERE id = $2"

	return res, product.DB.Raw(query, stock, id).Scan(&res).Error
}

func (product *ProductAdapter) DecrementStock(id uint, stock int) (entities.Products, error) {

	fmt.Println("jhsafkjdfj")
	var res entities.Products
	query := "UPDATE products SET stock = stock - $1 WHERE id = $2"

	tx := product.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Raw(query, stock, id).Scan(&res).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	if res.Stock < 0 {
		tx.Rollback()
		return res, fmt.Errorf("updated stock would be negative")
	}

	if err := tx.Commit().Error; err != nil {
		return res, err
	}

	return res, nil
}
