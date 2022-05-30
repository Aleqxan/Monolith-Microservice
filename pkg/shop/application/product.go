package application

type ProductReadModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductsService struct {
	repo      products.Repository
	readModel ProductReadModel
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() {

}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {

	price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)

	products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)

	s.repo.Save
}
