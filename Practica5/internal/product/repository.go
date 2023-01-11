package product

import (
	"errors"
	"regexp"

	"github.com/jmarambio/prueba/internal/domain"
	"github.com/jmarambio/prueba/pkg/store"
)

type ProductRepository interface {
	GetProducts() ([]domain.Producto, error)
	GetProductById(id int) (domain.Producto, error)
	GetProductByFilter(valor float64) ([]domain.Producto, error)
	AddProduct(p domain.Producto) (domain.Producto, error)
	EditProduct(p domain.Producto, id int) (domain.Producto, error)
	PatchProduct(p domain.Producto, id int) (domain.Producto, error)
	DeleteProduct(id int) error
}

// NewRepository crea una nueva instancia de Repository
func NewRepository(storage store.Store) ProductRepository {
	return &sliceRepository{storage}
}

type sliceRepository struct {
	storage store.Store
}

func (repo *sliceRepository) GetProducts() ([]domain.Producto, error) {
	productos, err := repo.storage.GetProducts()
	if err != nil {
		return []domain.Producto{}, err
	}
	if len(productos) == 0 {
		return []domain.Producto{}, err
	}
	return productos, nil
}

func (repo *sliceRepository) GetProductById(id int) (domain.Producto, error) {
	product, err := repo.storage.GetProductById(id)
	if err != nil {
		return domain.Producto{}, errors.New("producto no encontrado")
	}
	return product, nil
}

func (repo *sliceRepository) GetProductByFilter(valor float64) ([]domain.Producto, error) {

	var productos []domain.Producto
	products, err := repo.storage.GetProducts()
	if err != nil {
		return productos, err
	}
	for _, v := range products {
		if v.Price > valor {
			productos = append(productos, v)
		}
	}
	if len(productos) == 0 {
		return []domain.Producto{}, errors.New("productos no encontrado")
	}
	return productos, nil
}

func (repo *sliceRepository) AddProduct(producto domain.Producto) (domain.Producto, error) {

	err := validaciones(producto)
	if err != nil {
		return domain.Producto{}, err
	}

	_, err = repo.codeValueRepeated(producto)
	if err != nil {
		return domain.Producto{}, err
	}

	addProduct, err := repo.storage.AddProduct(producto)
	if err != nil {
		return domain.Producto{}, err
	}
	return addProduct, nil
}

func (repo *sliceRepository) EditProduct(producto domain.Producto, id int) (p domain.Producto, err error) {

	err = validaciones(producto)
	if err != nil {
		return p, err
	}

	_, err = repo.codeValueRepeated(producto)
	if err != nil {
		return p, err
	}

	err = repo.storage.EditProduct(producto)
	if err != nil {
		return domain.Producto{}, errors.New("producto no editado")
	}
	return p, nil

}

func (repo *sliceRepository) PatchProduct(producto domain.Producto, id int) (patchProduct domain.Producto, err error) {
	_, err = repo.codeValueRepeated(producto)
	if err != nil {
		return patchProduct, err
	}

	foundProduct, err := repo.storage.GetProductById(id)
	if err != nil {
		return patchProduct, err
	}
	patchProduct, err = validarPatch(foundProduct, producto)
	/*
		data, err := json.Marshal(&producto)
		if err != nil {
			log.Fatal(err)
		}
		reader := bytes.NewReader(data)

		err = json.NewDecoder(reader).Decode(&repo.products[id-1])
	*/
	if err != nil {
		return producto, errors.New("error al modificar el producto")
	}
	patchProduct.Id = id
	err = repo.storage.EditProduct(patchProduct)
	if err != nil {
		return domain.Producto{}, err
	}
	return patchProduct, nil

}

func (repo *sliceRepository) DeleteProduct(id int) error {
	err := repo.storage.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil

}

func (repo *sliceRepository) codeValueRepeated(p domain.Producto) (bool, error) {
	productos, _ := repo.storage.GetProducts()
	for _, v := range productos {
		if v.Code_value == p.Code_value {
			return true, errors.New("el code value ya existe")
		}
	}
	return false, nil
}

func validaciones(p domain.Producto) error {
	if p.Name == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	if p.Expiration == "" {
		return errors.New("la fecha de expiracion no puede estar vacio")
	}
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if !re.MatchString(p.Expiration) {
		return errors.New("fecha incorrecta o Fomato incorrecto de expiración, el formato es : dd/mm/yyyy")
	}
	if p.Code_value == "" {
		return errors.New("el code value no puede ser vacio")
	}
	if p.Price <= 0 {
		return errors.New("el precio no puede ser igual o menor a 0")
	}
	if p.Quantity <= 0 {
		return errors.New("la cantidad no puede ser igual o menor a 0")
	}
	return nil
}

func validarPatch(productoModificar domain.Producto, productoPatch domain.Producto) (product domain.Producto, err error) {
	product = productoPatch
	if productoPatch.Name == "" {
		product.Name = productoModificar.Name
	}
	if productoPatch.Expiration == "" {
		product.Expiration = productoModificar.Expiration
	}
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if !re.MatchString(product.Expiration) {
		return product, errors.New("fecha incorrecta o Fomato incorrecto de expiración, el formato es : dd/mm/yyyy")
	}
	if productoPatch.Code_value == "" {
		product.Code_value = productoModificar.Code_value
	}
	if productoPatch.Price <= 0 {
		product.Price = productoModificar.Price
	}
	if productoPatch.Quantity <= 0 {
		product.Quantity = productoModificar.Quantity
	}
	return product, nil
}
