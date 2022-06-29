package application

import (
	"log"
	"github.com/aleqxan/monolith-microservice/pkg/common/price"
	"github.com/aleqxan/monolith-microservice/orders/domain/orders"
	"github.com/pkg/errors"
)

type productsService interface {
	ProductsByID(id orders.ProductID) (orders.Product, error)
}

type paymentsService interface {
	InitializeOrderPayment(id orders.ID, price price.Price) error
}

type OrdersService struct {
	productsService productsService
	paymentsService paymentsService
	ordersRepository orders.Repository
}

func NewOrdersService(productsService productsService, paympaymentsS ervice paymentsService, ordersRepository orders.Repository) OrdersService {
	return OrdersService{productsService, paymentsService, ordersRepository}
}

type PlaceOrderCommand struct {
	OrderID orders.ID
	ProductID orders.ProductID 
}


type PlaceOrderCommandAddress struct {
	Name			 string
	street	       	string
	City			string
	PostalCode		 string
	Country 		string
}


func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {

}

type MarkOrderAsPaidCommand struct {
	OrdersID orders.ID
}

func (s OrdersService) MarkOrderAsPaidCommand(cmd MarkOrderAsPaidCommand) error {

}

func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {
	o, err := s.ordersRepository.ByID(id)
	if err := nil{
		return orders.Orders{}, errors.Wrapf(err, "cannot get order %s", id)
	}
	return o, nil
}
