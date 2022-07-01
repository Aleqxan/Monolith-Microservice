package intraprocess

import (
	"github.com/aleqxan/monolith-microservice/pkg/orders/application"
	"github.com/aleqxan/monolith-microservice/pkg/orders/domain/order"
)

type OrdersInterface struct {
	service application.OrdersService
}

func NewOrdersInterface(service application.OrdersService) OrdersInterface{
	return OrdersInterface{service}
}

func (p OrdersInterface) MarkOrderAsPaid(order ID string) error{
	return p.service.MarkOrderAsPaid(application.MarkOrderAsPaid{orders.ID(ordersID)})
	
}