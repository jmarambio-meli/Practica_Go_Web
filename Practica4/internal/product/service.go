package product

import "github.com/jmarambio/prueba/internal/domain"

// Service proporciona funcionalidades para trabajar con productos
type Service struct {
	repo ProductRepository
}

// NewService crea una nueva instancia de Service
func NewService(repository ProductRepository) *Service {
	return &Service{
		repo: repository,
		//repo: NewRepository(),
	}
}

// List recupera todos los productos
func (s *Service) List() ([]domain.Producto, error) {
	return s.repo.GetProducts()
}

// ListId recupera un producto con el id indicado
func (s *Service) ListId(id int) (domain.Producto, error) {
	return s.repo.GetProductById(id)
}

// Listfilter recupera todos los productos con el filtro
func (s *Service) ListFilter(valor float64) ([]domain.Producto, error) {
	return s.repo.GetProductByFilter(valor)
}

// AddToList agrega un producto a la lista
func (s *Service) AddList(producto domain.Producto) (domain.Producto, error) {
	return s.repo.AddProduct(producto)
}

// EditToList modifica un producto de la lista con la peticion PUT
func (s *Service) EditList(producto domain.Producto, id int) (domain.Producto, error) {
	return s.repo.EditProduct(producto, id)
}

// PatchToList modifica un producto de la lista con la peticion PATCH
func (s *Service) PatchList(producto domain.Producto, id int) (domain.Producto, error) {
	return s.repo.PatchProduct(producto, id)
}

// DeleteToList borra un producto de la lista
func (s *Service) DeleteList(id int) error {
	return s.repo.DeleteProduct(id)
}
