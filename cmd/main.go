package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/akshay0074700747/products-service/db"
	"github.com/akshay0074700747/products-service/initializer"
	"github.com/akshay0074700747/products-service/service"
	servicediscoveryconsul "github.com/akshay0074700747/products-service/servicediscovery_consul"
	"github.com/akshay0074700747/proto-files-for-microservices/pb"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
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

	go func() {
		time.Sleep(5 * time.Second)

		servicediscoveryconsul.RegisterService()
	}()

	healthService := &service.HealthChecker{}

	grpc_health_v1.RegisterHealthServer(server, healthService)

	tracer, closer := initTracer()

	defer closer.Close()

	service.RetrieveTreacer(tracer)

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to listen on port 50004: %v", err)
	}
}

func initTracer() (tracer opentracing.Tracer, closer io.Closer) {
	jaegerEndpoint := "http://localhost:14268/api/traces"

	cfg := &config.Configuration{
		ServiceName: "product-service",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: jaegerEndpoint,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("updated")

	return
}
