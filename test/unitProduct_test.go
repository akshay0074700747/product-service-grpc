package test

import (
	"context"
	"errors"
	"testing"

	mock_adapters "github.com/akshay0074700747/products-service/adapters/mocks"
	"github.com/akshay0074700747/products-service/entities"
	"github.com/akshay0074700747/products-service/service"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestAddProducts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapters.NewMockAdapterInterface(ctrl)
	tests := []struct {
		name      string
		mockFunc  func(req entities.Products) (entities.Products, error)
		request   *pb.AddProductRequest
		wantError bool
		result    *pb.AddProductResponce
	}{
		{
			name: "Success",
			mockFunc: func(req entities.Products) (entities.Products, error) {
				return entities.Products{ID: 1, Name: req.Name, Price: req.Price, Stock: req.Stock}, nil
			},
			request:   &pb.AddProductRequest{Name: "phone", Price: 10000, Stock: 50},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 1, Name: "phone", Price: 10000, Stock: 50},
		},
		{
			name: "Failure",
			mockFunc: func(req entities.Products) (entities.Products, error) {
				return entities.Products{}, errors.New("this is a failure case")
			},
			request:   &pb.AddProductRequest{Name: "phone", Price: 10000, Stock: 50},
			wantError: true,
			result:    nil,
		},
		{
			name: "Success",
			mockFunc: func(req entities.Products) (entities.Products, error) {
				return entities.Products{ID: 2, Name: req.Name, Price: req.Price, Stock: req.Stock}, nil
			},
			request:   &pb.AddProductRequest{Name: "tv", Price: 100000, Stock: 45},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 2, Name: "tv", Price: 100000, Stock: 45},
		},
	}

	for _, test := range tests {

		adapter.EXPECT().AddProduct(gomock.Any()).DoAndReturn(test.mockFunc).AnyTimes().Times(1)
		productSrv := service.NewProductService(adapter)

		res, err := productSrv.AddProducts(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, res)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, test.result, res)
		}
	}

}

func TestGetProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapters.NewMockAdapterInterface(ctrl)
	tests := []struct {
		name      string
		mockFunc  func(id uint) (entities.Products, error)
		request   *pb.GetProductByID
		wantError bool
		result    *pb.AddProductResponce
	}{
		{
			name: "Success",
			mockFunc: func(id uint) (entities.Products, error) {
				return entities.Products{ID: id, Name: "phone", Price: 10000, Stock: 50}, nil
			},
			request:   &pb.GetProductByID{Id: 1},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 1, Name: "phone", Price: 10000, Stock: 50},
		},
		{
			name: "Failure",
			mockFunc: func(id uint) (entities.Products, error) {
				return entities.Products{}, errors.New("this is a failure")
			},
			request:   &pb.GetProductByID{Id: 100},
			wantError: true,
			result:    nil,
		},
		{
			name: "Success",
			mockFunc: func(id uint) (entities.Products, error) {
				return entities.Products{ID: id, Name: "tv", Price: 100000, Stock: 45}, nil
			},
			request:   &pb.GetProductByID{Id: 2},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 2, Name: "tv", Price: 100000, Stock: 45},
		},
	}

	for _, test := range tests {

		adapter.EXPECT().GetProduct(gomock.Any()).DoAndReturn(test.mockFunc).AnyTimes().Times(1)
		productSrv := service.NewProductService(adapter)

		res, err := productSrv.GetProduct(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, res)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, test.result, res)
		}
	}
}

func TestGetAllProducts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapters.NewMockAdapterInterface(ctrl)
	tests := []struct {
		name      string
		mockFunc  func() ([]entities.Products, error)
		request   *empty.Empty
		wantError bool
		result    *pb.GetAllProductsResponce
	}{
		{
			name: "Success",
			mockFunc: func() ([]entities.Products, error) {
				return []entities.Products{
					{ID: 1, Name: "phone", Price: 10000, Stock: 50},
					{ID: 2, Name: "tv", Price: 100000, Stock: 45},
				}, nil
			},
			request:   &emptypb.Empty{},
			wantError: false,
			result: &pb.GetAllProductsResponce{
				Products: []*pb.AddProductResponce{
					{Id: 1, Name: "phone", Price: 10000, Stock: 50},
					{Id: 2, Name: "tv", Price: 100000, Stock: 45},
				},
			},
		},
		{
			name: "Failure",
			mockFunc: func() ([]entities.Products, error) {
				return []entities.Products{}, errors.New("this is a failed case")
			},
			request:   &emptypb.Empty{},
			wantError: true,
			result:    nil,
		},
	}

	for _, test := range tests {

		adapter.EXPECT().GetAllProducts().DoAndReturn(test.mockFunc).AnyTimes().Times(1)
		productSrv := service.NewProductService(adapter)

		res, err := productSrv.GetAllProducts(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, res)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, test.result, res)
		}
	}
}

func TestUpdateStock(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	adapter := mock_adapters.NewMockAdapterInterface(ctrl)
	tests := []struct {
		name      string
		mockFunc  func(id uint, stock int) (entities.Products, error)
		request   *pb.UpdateStockRequest
		wantError bool
		result    *pb.AddProductResponce
	}{
		{
			name: "Success",
			mockFunc: func(id uint, stock int) (entities.Products, error) {
				return entities.Products{ID: id, Name: "phone", Price: 10000, Stock: 55}, nil
			},
			request:   &pb.UpdateStockRequest{Id: 1, Stock: 5, Increase: true},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 1, Name: "phone", Price: 10000, Stock: 55},
		},
		{
			name: "Failure",
			mockFunc: func(id uint, stock int) (entities.Products, error) {
				return entities.Products{}, errors.New("the stock cannot go below zero")
			},
			request:   &pb.UpdateStockRequest{Id: 1, Stock: 60, Increase: false},
			wantError: true,
			result:    nil,
		},
		{
			name: "Success",
			mockFunc: func(id uint, stock int) (entities.Products, error) {
				return entities.Products{ID: id, Name: "tv", Price: 100000, Stock: 40}, nil
			},
			request:   &pb.UpdateStockRequest{Id: 2, Stock: 5, Increase: false},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 2, Name: "tv", Price: 100000, Stock: 40},
		},
	}

	for _, test := range tests {

		if test.request.Increase {
			adapter.EXPECT().IncrementStock(gomock.Any(), gomock.Any()).DoAndReturn(test.mockFunc).AnyTimes().Times(1)
		} else {
			adapter.EXPECT().DecrementStock(gomock.Any(), gomock.Any()).DoAndReturn(test.mockFunc).AnyTimes().Times(1)
		}
		productSrv := service.NewProductService(adapter)

		res, err := productSrv.UpdateStock(context.TODO(), test.request)
		if test.wantError {
			assert.Error(t, err)
			assert.Nil(t, res)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, test.result, res)
		}
	}
}
