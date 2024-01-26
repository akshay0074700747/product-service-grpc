package service

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/akshay0074700747/products-service/adapters"
	"github.com/akshay0074700747/products-service/entities"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

var (
	Tracer opentracing.Tracer
)

func RetrieveTreacer(tr opentracing.Tracer) {
	Tracer = tr
}

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

	span := Tracer.StartSpan("add products grpc")
	defer span.Finish()

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

	span := Tracer.StartSpan("get product grpc")
	defer span.Finish()

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

func (product *ProductService) GetAllProducts(em *empty.Empty, srv pb.ProductService_GetAllProductsServer) error {

	span := Tracer.StartSpan("get all products grpc")
	defer span.Finish()

	fmt.Println("GetAll products called third...")
	products, err := product.Adapter.GetAllProducts()
	if err != nil {
		return err
	}
	for _, prod := range products {
		if err = srv.Send(&pb.AddProductResponce{
			Id:    uint32(prod.ID),
			Name:  prod.Name,
			Price: int32(prod.Price),
			Stock: int32(prod.Stock),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (product *ProductService) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.AddProductResponce, error) {

	span := Tracer.StartSpan("update stock of product grpc")
	defer span.Finish()

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

func (product *ProductService) GetArrayofProducts(srv pb.ProductService_GetArrayofProductsServer) error {

	span := Tracer.StartSpan("get array of products grpc")
	defer span.Finish()
	for {
		ID, err := srv.Recv()
		if err == io.EOF {
			log.Println(err)
			break
		}
		if err != nil {
			return err
		}
		prod, err := product.Adapter.GetProduct(uint(ID.GetId()))
		if err = srv.Send(&pb.AddProductResponce{
			Id:    uint32(prod.ID),
			Name:  prod.Name,
			Price: int32(prod.Price),
			Stock: int32(prod.Stock),
		}); err != nil {
			return err
		}
	}

	return nil
}

type HealthChecker struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (s *HealthChecker) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Println("check called")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthChecker) Watch(in *grpc_health_v1.HealthCheckRequest, srv grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watching is not supported")
}
