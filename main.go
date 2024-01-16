package main

import (
	"context"
	"github.com/Filtronic/Minio/app/models"
	"github.com/Filtronic/Minio/gapi/core/grpc_handlers"
	pb "github.com/Filtronic/Minio/gapi/pb/mutation_gen"
	"github.com/Filtronic/Minio/pkg/configs"
	"github.com/Filtronic/Minio/pkg/routes"
	"github.com/Filtronic/Minio/pkg/utils"
	"github.com/Filtronic/Minio/util"
	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
)

var db *gorm.DB

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)
	db = initDB()
	routes.PublicRoutes(app, db)
	utils.StartServer(app)

	configGrpc, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	runGrpcServer(configGrpc)
	go runGatewayServer(configGrpc, db)
}

func initDB() *gorm.DB {
	dsn := os.Getenv("CONNECTION_STRING") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.FileMetadata{}, &models.BucketMetadata{})
	return db

}

func runGrpcServer(config util.Config) {
	grpcServer := grpc.NewServer()
	server, err := grpc_handlers.NewServer(config)
	if err != nil {
		log.Fatal("cannot create grpc server:", err)
	}
	pb.RegisterMutationServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Unable to start the listener: ", err)
	}
	log.Printf("Starting gRPC server on %s\n", listener.Addr().String())

	//start the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Unable to start the gRPC server: ", err)
	}
}

func runGatewayServer(config util.Config, db *gorm.DB) {

	server, err := grpc_handlers.NewServer(config)
	if err != nil {
		log.Fatal("cannot create grpc server:", err)
	}
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterMutationServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Unable to create the listener :", err)
	}
	log.Printf("Starting http server on %s\n", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Unable to start the http server: ", err)
	}
}
