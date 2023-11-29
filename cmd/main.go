package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"golang-project-template/internal/bootstrap"
	"golang-project-template/internal/deliver/grpc/controller"
	pb "golang-project-template/internal/deliver/grpc/stub"
	"golang-project-template/internal/deliver/rest/api/router"
	"golang-project-template/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	// Load environment variables and set up the database
	bootstrap.LoadEnv()
	db := bootstrap.SetupDB()

	// Set up the HTTP server
	r := chi.NewRouter()
	router.Setup(r, db)

	// Start HTTP server in a goroutine
	go func() {
		err := http.ListenAndServe(":5005", r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Set up the gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterImageServiceServer(grpcServer, &controller.ImageServerImpl{UseCase: usecase.NewImageUseCase(db)})

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":5006")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started on :5006")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
