package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/jmarambio/prueba/internal/domain"
)

type Store interface {
	GetProducts() ([]domain.Producto, error)
	GetProductById(id int) (domain.Producto, error)
	//GetProductByFilter(valor float64) ([]domain.Producto, error)
	AddProducts(products []domain.Producto) error
	AddProduct(product domain.Producto) (domain.Producto, error)
	EditProduct(product domain.Producto) error
	//PatchProduct(p domain.Producto, id int) (domain.Producto, error)
	DeleteProduct(id int) error
	CargarJson() ([]domain.Producto, error)
}

type jsonStore struct {
	pathToFile string
}

func (s *jsonStore) CargarJson() ([]domain.Producto, error) {
	var productos []domain.Producto
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &productos)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

// AddProducts guarda los productos en un archivo json
func (s *jsonStore) AddProducts(products []domain.Producto) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de products
func NewStore(path string) Store {
	return &jsonStore{
		pathToFile: path,
	}
}

// GetProducts devuelve todos los productos
func (s *jsonStore) GetProducts() ([]domain.Producto, error) {
	products, err := s.CargarJson()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductById entrega un producto por su id
func (s *jsonStore) GetProductById(id int) (domain.Producto, error) {
	products, err := s.CargarJson()
	if err != nil {
		return domain.Producto{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Producto{}, errors.New("producto no encontrado")
}

// AddProduct agrega un producto
func (s *jsonStore) AddProduct(product domain.Producto) (domain.Producto, error) {
	products, err := s.CargarJson()
	if err != nil {
		return domain.Producto{}, err
	}
	product.Id = len(products) + 1
	products = append(products, product)
	return product, s.AddProducts(products)
}

// EditProduct actualiza un producto
func (s *jsonStore) EditProduct(product domain.Producto) error {
	products, err := s.CargarJson()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return s.AddProducts(products)
		}
	}
	return errors.New("product not found")
}

// DeleteProduct elimina un producto
func (s *jsonStore) DeleteProduct(id int) error {
	products, err := s.CargarJson()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.AddProducts(products)
		}
	}
	return errors.New("product not found")
}
