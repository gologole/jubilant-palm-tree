package grpcserver

import (
	"cmd/main.go/internal/service"
	pb "cmd/main.go/internal/transport/grpc/rpc/protogen"
	"cmd/main.go/models"
	mylogger "cmd/main.go/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
)

type server struct {
	pb.UnimplementedOrderServiceServer // Embed the unimplemented server
	service                            service.Service
}

// NewGrpcServer returns an instance of OrderServiceServer
func NewGrpcServer(service service.Service) pb.OrderServiceServer {
	return &server{service: service}
}

func (s *server) CreateOrder(ctx context.Context, newOrder *pb.NewOrder) (*emptypb.Empty, error) {
	order := models.Order{
		ID:          int64(rand.Int()),
		OverPrice:   newOrder.Overprice,
		Description: newOrder.Description,
	}
	err := s.service.AddOrder(order)

	if err != nil {
		mylogger.GlobalLogger.Error(fmt.Sprintf("Error adding new order %s", err.Error()))
		return nil, err
	}

	return nil, nil
}

func (s *server) GetOrder(ctx context.Context, _ *emptypb.Empty) (*pb.NewOrder, error) {
	order, err := s.service.GetOrder()
	if err != nil {
		mylogger.GlobalLogger.Error(fmt.Sprintf("Error getting order %s", err.Error()))
		return nil, status.Errorf(codes.Internal, "error getting order: %v", err)
	}

	// Преобразование models.Order в pb.NewOrder
	pbOrder := &pb.NewOrder{
		Overprice:   order.OverPrice,   // Пример поля
		Description: order.Description, // Пример поля
	}

	return pbOrder, nil
}
func (s *server) UploadDocument(ctx context.Context, resp *pb.Response) (*emptypb.Empty, error) {
	err := s.service.ReturnResult(resp)

	if err != nil {
		mylogger.GlobalLogger.Error(fmt.Sprintf("Error upload document %s", err.Error()))
	}
	return &emptypb.Empty{}, nil
}
