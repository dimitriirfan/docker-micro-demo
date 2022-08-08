package product

type Service interface {
	GetAllProduct() ([]Product, error)
}

type service struct {
	productRepository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAllProduct() ([]Product, error) {
	products, err := s.productRepository.GetAll()

	if err != nil {
		return products, err
	}

	return products, nil
}
