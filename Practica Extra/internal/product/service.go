package product

import "github.com/jmarambio/prueba/internal/domain"

// Service proporciona funcionalidades para trabajar con productos

type Service interface {
	GetProducts() ([]domain.Producto, error)
	GetProductById(id int) (domain.Producto, error)
	GetProductByFilter(valor float64) ([]domain.Producto, error)
	GetTotalByConsumer(id []int) ([]domain.Producto, float64, error)
	AddProduct(producto domain.Producto) (domain.Producto, error)
	EditProduct(producto domain.Producto, id int) (domain.Producto, error)
	PatchProduct(producto domain.Producto, id int) (domain.Producto, error)
	DeleteProduct(id int) error
}

type service struct {
	repo ProductRepository
}

// NewService crea una nueva instancia de Service
func NewService(repository ProductRepository) Service {
	return &service{
		repo: repository,
	}
}

// List recupera todos los productos
func (s *service) GetProducts() ([]domain.Producto, error) {
	return s.repo.GetProducts()
}

// ListId recupera un producto con el id indicado
func (s *service) GetProductById(id int) (domain.Producto, error) {
	return s.repo.GetProductById(id)
}

// Listfilter recupera todos los productos con el filtro
func (s *service) GetProductByFilter(valor float64) ([]domain.Producto, error) {
	return s.repo.GetProductByFilter(valor)
}
func (s *service) GetTotalByConsumer(id []int) ([]domain.Producto, float64, error) {
	return s.repo.GetTotalByConsumer(id)
}

// AddToList agrega un producto a la lista
func (s *service) AddProduct(producto domain.Producto) (domain.Producto, error) {
	return s.repo.AddProduct(producto)
}

// EditToList modifica un producto de la lista con la peticion PUT
func (s *service) EditProduct(producto domain.Producto, id int) (domain.Producto, error) {
	return s.repo.EditProduct(producto, id)
}

// PatchToList modifica un producto de la lista con la peticion PATCH
func (s *service) PatchProduct(producto domain.Producto, id int) (domain.Producto, error) {
	return s.repo.PatchProduct(producto, id)
}

// DeleteToList borra un producto de la lista
func (s *service) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
