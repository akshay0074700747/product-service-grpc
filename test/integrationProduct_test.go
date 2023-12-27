package test

import (
	"context"
	"os"
	"testing"

	"github.com/akshay0074700747/products-service/db"
	"github.com/akshay0074700747/products-service/initializer"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestAddProduct(t *testing.T) {

	if err := godotenv.Load("../cmd/.env"); err != nil {
		t.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	db, err := db.InitDB(addr)
	if err != nil {
		t.Fatal(err.Error())
	}

	productSrv := initializer.Initialize(db)
	tests := []struct {
		name      string
		request   *pb.AddProductRequest
		wantError bool
		result    *pb.AddProductResponce
	}{
		{
			name:      "Success",
			request:   &pb.AddProductRequest{Name: "phone", Price: 10000, Stock: 50},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 1, Name: "phone", Price: 10000, Stock: 50},
		},
		{
			name:      "Failure",
			request:   &pb.AddProductRequest{Name: "", Price: 10000, Stock: 50},
			wantError: true,
			result:    nil,
		},
		{
			name:      "Success",
			request:   &pb.AddProductRequest{Name: "tv", Price: 100000, Stock: 45},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 2, Name: "tv", Price: 100000, Stock: 45},
		},
	}

	for _, test := range tests {

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

func TestGetProducts(t *testing.T) {

	if err := godotenv.Load("../cmd/.env"); err != nil {
		t.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	db, err := db.InitDB(addr)
	if err != nil {
		t.Fatal(err.Error())
	}

	productSrv := initializer.Initialize(db)
	tests := []struct {
		name      string
		request   *pb.GetProductByID
		wantError bool
		result    *pb.AddProductResponce
	}{
		{
			name:      "Success",
			request:   &pb.GetProductByID{Id: 1},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 1, Name: "phone", Price: 10000, Stock: 50},
		},
		{
			name:      "Failure",
			request:   &pb.GetProductByID{Id: 100},
			wantError: true,
			result:    nil,
		},
		{
			name:      "Success",
			request:   &pb.GetProductByID{Id: 2},
			wantError: false,
			result:    &pb.AddProductResponce{Id: 2, Name: "tv", Price: 100000, Stock: 45},
		},
	}

	for _, test := range tests {

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

func TestGetAllProduct(t *testing.T) {

	if err := godotenv.Load("../cmd/.env"); err != nil {
		t.Fatal(err.Error())
	}

	addr := os.Getenv("TEST_DATABASE_ADDR")

	db, err := db.InitDB(addr)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer func() {
		db.Exec("drop table products")
	}()

	productSrv := initializer.Initialize(db)
	tests := []struct {
		name      string
		request   *empty.Empty
		wantError bool
		result    *pb.GetAllProductsResponce
	}{
		{
			name:      "Success",
			request:   &emptypb.Empty{},
			wantError: false,
			result: &pb.GetAllProductsResponce{
				Products: []*pb.AddProductResponce{
					{Id: 1, Name: "phone", Price: 10000, Stock: 50},
					{Id: 2, Name: "tv", Price: 100000, Stock: 45},
				},
			},
		},
		// {
		// 	name: "Failure",
		// 	request:   &emptypb.Empty{},
		// 	wantError: true,
		// 	result:    nil,
		// },
	}

	for _, test := range tests {

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
