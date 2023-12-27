package service

import (
	"context"
	"fmt"

	"github.com/akshay0074700747/products-service/adapters"
	"github.com/akshay0074700747/products-service/entities"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type ProductService struct {
	Adapter adapters.AdapterInterface
	pb.UnimplementedProductServiceServer
}

func NewProductService(adapter adapters.AdapterInterface) *ProductService {
	return &ProductService{
		Adapter: adapter,
	}
}

func (product *ProductService) AddProducts(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponce, error) {

	if req.Name == "" {
		return nil, fmt.Errorf("the name of the product cant be empty")
	}
	fmt.Println("Add Product called first...")
	reqEntity := entities.Products{
		Name:  req.Name,
		Price: int(req.Price),
		Stock: int(req.Stock),
	}

	res, err := product.Adapter.AddProduct(reqEntity)
	if err != nil {
		return nil, err
	}

	return &pb.AddProductResponce{
		Id:    uint32(res.ID),
		Name:  res.Name,
		Price: int32(res.Price),
		Stock: int32(res.Stock),
	}, nil
}

func (product *ProductService) GetProduct(ctx context.Context, req *pb.GetProductByID) (*pb.AddProductResponce, error) {
	fmt.Println("Get Product called second...")
	res, err := product.Adapter.GetProduct(uint(req.Id))
	if err != nil {
		return nil, err
	}

	if res.Name == "" {
		return nil, fmt.Errorf("the product with the given id doesnt exist")
	}

	return &pb.AddProductResponce{
		Id:    uint32(res.ID),
		Name:  res.Name,
		Price: int32(res.Price),
		Stock: int32(res.Stock),
	}, nil
}

func (product *ProductService) GetAllProducts(context.Context, *empty.Empty) (*pb.GetAllProductsResponce, error) {

	fmt.Println("GetAll products called third...")
	products, err := product.Adapter.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var res []*pb.AddProductResponce

	for _, prod := range products {
		res = append(res, &pb.AddProductResponce{
			Id:    uint32(prod.ID),
			Name:  prod.Name,
			Price: int32(prod.Price),
			Stock: int32(prod.Stock),
		})
	}

	return &pb.GetAllProductsResponce{
		Products: res,
	}, nil
}

func (product *ProductService) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.AddProductResponce, error) {

	var res *pb.AddProductResponce

	if req.Increase {

		result, err := product.Adapter.IncrementStock(uint(req.Id), int(req.Stock))
		if err != nil {
			return nil, err
		}

		res = &pb.AddProductResponce{
			Id:    uint32(result.ID),
			Name:  result.Name,
			Price: int32(result.Price),
			Stock: int32(result.Stock),
		}
	} else {

		result, err := product.Adapter.DecrementStock(uint(req.Id), int(req.Stock))
		if err != nil {
			return nil, err
		}

		res = &pb.AddProductResponce{
			Id:    uint32(result.ID),
			Name:  result.Name,
			Price: int32(result.Price),
			Stock: int32(result.Stock),
		}
	}

	return res, nil
}
