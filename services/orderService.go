package services

import "hacktiv8_assignment_3/repositories"

type OrderService interface {
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) orderService {
	return orderService{
		orderRepository: orderRepo,
	}
}
