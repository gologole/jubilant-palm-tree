package main

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/service"
	pb "cmd/main.go/internal/transport/grpc/rpc/protogen"
	mylogger "cmd/main.go/pkg/logger"
	"cmd/main.go/server/grpcserver"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

func main() {
	cfg, err := config.InitConfig("./config/config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	mylogger.NewLogger()

	// Создание сервиса
	svc := service.NewService()

	// Создание gRPC сервера
	grpcServer := grpc.NewServer()

	// Создание экземпляра вашего gRPC сервера
	orderServiceServer := grpcserver.NewGrpcServer(svc)

	// Регистрация вашего gRPC сервера
	pb.RegisterOrderServiceServer(grpcServer, orderServiceServer)

	reflection.Register(grpcServer)
	// Прослушивание входящих соединений
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.GrpcServer.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	mylogger.GlobalLogger.Info(fmt.Sprintf("Server is listening on port %s", cfg.GrpcServer.Port))

	// Запуск gRPC сервера
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
