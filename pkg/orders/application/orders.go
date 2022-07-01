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
	Name			    string
	street	 	      	string
	City				string
	PostalCode			string
	Country 			string
}


func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := order.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street,
		cmd.Address.City,
		cmd.Address.PostCode,
		cmd.Address.Country,
	)

	if err != nil {
		return errors.Wrap(err, "Invalid address")
	}

	// 1. Gettting the product by id

	product, err := s.productsService.ProductsByID(cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "Cannot get the product")
	}
	// 2. Create a new PlaceOrder

	mewOrder, err := orders.NewOrder(cmd.OrderID, product, address)
	if err != nil {
		return errors.Wrap(err, "Cannot create order")
	}
	// 3. Save the order

	if err := s.ordersRepository.Save(newOrder); err != nil{
		return errors.Wrap(err, "cannot save order")
	}
	// 4. Initialize the payment

	if err := s.paymentsService.InitializeOrderPayment(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.Wrap(err, "Cannot initialize payment")
	}
	log.Printf("order %s placed", &cmd.OrderID)
	return nil
}

type MarkOrderAsPaidCommand struct {
	OrdersID orders.ID
}

func (s OrdersService) MarkOrderAsPaidCommand(cmd MarkOrderAsPaidCommand) error {
	o, err := s.ByID(cmd.OrdersID)
	if err != nil {
		return errors.Wrap(err, "Cannot get order %s", cmd.OrdersID)
	}
	o.MarkAsPaid()

	if err := s.ordersRepository.Save(o); err != nil{
		return errors.Wrap(err, "cannot save order")
	}

	log.Printf("marked order %s, as paid", cmd.OrdersID)
	return nil 
}

func (s OrdersService) OrderByID(id orders.ID) (orders.Order, error) {
	o, err := s.ordersRepository.ByID(id)
	if err := nil{
		return orders.Orders{}, errors.Wrapf(err, "cannot get order %s", id)
	}
	return o, nil
}
