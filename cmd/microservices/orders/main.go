package main

import (
	"net/http"
	"log"
	"os"
	"fmt"

)

func main(){
	log.Println("Starting the orders microservice")

	ctx := cmd.Context()

	r, closeFn := createOrderMicroservice()

	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r}

	go func(){
		if err := server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}()

		<-ctx.Done()

		log.Println("Closing microservice")

		if err := server.Close(); err != nil {
			panic(err)
		}
	}
}

func createOrderMicroservice()(router *chi.Mux, closeFn func()){
	cmd.WaitForservice(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r := cmd.CreateRouter()

	orders_public_http.ADDRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRputes(r, ordersService, ordersRepo)

	return r, func() {
		
	}
}