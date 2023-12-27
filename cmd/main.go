package main

import (
	"log"
	"net"
	"os"

	"github.com/akshay0074700747/products-service/db"
	"github.com/akshay0074700747/products-service/initializer"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}

	addr := os.Getenv("DATABASE_ADDR")

	DB, err := db.InitDB(addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	servicee := initializer.Initialize(DB)

	server := grpc.NewServer()

	pb.RegisterProductServiceServer(server, servicee)

	listener, err := net.Listen("tcp", ":50004")

	if err != nil {
		log.Fatalf("Failed to listen on port 50004: %v", err)
	}

	log.Printf("Product Server is listening on port")

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to listen on port 50004: %v", err)
	}
}
