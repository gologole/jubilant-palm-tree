package service

import (
	pb "cmd/main.go/internal/transport/grpc/rpc/protogen"
	"cmd/main.go/models"
	"cmd/main.go/pkg/logger"
	"errors"
	"fmt"
)

type Service interface {
	AddOrder(order models.Order) error    // для создания заказа
	GetOrder() (*models.Order, error)     // при запросе из расчета приоритета
	ReturnResult(resp *pb.Response) error // для возврата клиенту и записи в бд
}
type service struct {
	Queue *Queue
}

func NewService() Service {
	return service{
		NewQueue(),
	}
}

func (s service) AddOrder(order models.Order) error {
	go s.Queue.Enqueue(&order)
	mylogger.GlobalLogger.Info(fmt.Sprintf("Added order %s", order))
	return nil
}

func (s service) GetOrder() (*models.Order, error) {
	order, ok := s.Queue.Get()
	mylogger.GlobalLogger.Info(fmt.Sprintf("Getting order %s", order))
	if !ok {
		mylogger.GlobalLogger.Error("Error getting order")
		return nil, errors.New("Error getting order")
	}
	return order.(*models.Order), nil
}

func (s service) ReturnResult(resp *pb.Response) error {
	//grpc.UploadDocument()
	mylogger.GlobalLogger.Info(fmt.Sprintf("Upload response %s", resp))
	return nil
}
