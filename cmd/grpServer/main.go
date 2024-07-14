package main

import (
	"database/sql"
	"net"

	"github.com/thiagohmm/gRPCEstudo/internal/database"
	"github.com/thiagohmm/gRPCEstudo/internal/pb"
	"github.com/thiagohmm/gRPCEstudo/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("Sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	CategoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*CategoryDB)

	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryService)
	reflection.Register(server)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
