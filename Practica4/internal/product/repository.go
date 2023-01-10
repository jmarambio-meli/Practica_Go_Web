package product

import (
	"errors"
	"regexp"

	"github.com/jmarambio/prueba/internal/domain"
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
func NewRepository(productos []domain.Producto) ProductRepository {
	return &sliceRepository{productos}
}

type sliceRepository struct {
	products []domain.Producto
}

func (repo *sliceRepository) GetProducts() ([]domain.Producto, error) {
	if len(repo.products) == 0 {
		return repo.products, errors.New("no existen productos")
	}
	return repo.products, nil
}

func (repo *sliceRepository) GetProductById(id int) (domain.Producto, error) {
	for _, v := range repo.products {
		if v.Id == id {
			return v, nil
		}
	}
	return domain.Producto{}, errors.New("producto no encontrado")
}

func (repo *sliceRepository) GetProductByFilter(valor float64) ([]domain.Producto, error) {
	var productos []domain.Producto
	for _, v := range repo.products {
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

	_, err = codeValueRepeated(producto, repo.products)
	if err != nil {
		return domain.Producto{}, err
	}

	producto.Id = len(domain.Productos) + 1
	repo.products = append(repo.products, producto)
	//domain.Productos = append(domain.Productos, producto)

	return producto, nil
}

func (repo *sliceRepository) EditProduct(producto domain.Producto, id int) (domain.Producto, error) {

	err := validaciones(producto)
	if err != nil {
		return domain.Producto{}, err
	}

	_, err = codeValueRepeated(producto, repo.products)
	if err != nil {
		return domain.Producto{}, err
	}

	isExist := false
	for _, v := range repo.products {
		if v.Id == id {
			v = producto
			v.Id = id
			isExist = true
		}
	}

	if !isExist {
		return producto, errors.New("producto a modificar no encontrado")
	} else {
		p := &repo.products[id-1]
		p.Id = id
		p.Code_value = producto.Code_value
		p.Expiration = producto.Expiration
		p.Is_published = producto.Is_published
		p.Name = producto.Name
		p.Price = producto.Price
		p.Quantity = producto.Quantity
		return repo.products[id-1], nil
	}

}

func (repo *sliceRepository) PatchProduct(producto domain.Producto, id int) (domain.Producto, error) {
	_, err := codeValueRepeated(producto, repo.products)
	if err != nil {
		return domain.Producto{}, err
	}

	isExist := false
	for _, v := range repo.products {
		if v.Id == id {
			isExist = true
		}
	}

	if !isExist {
		return producto, errors.New("producto a modificar no encontrado")
	} else {
		patchProduct, err := validarPatch(repo.products[id-1], producto)
		if err != nil {
			return producto, errors.New("error al modificar el producto")
		}
		p := &repo.products[id-1]
		p.Id = id
		p.Code_value = patchProduct.Code_value
		p.Expiration = patchProduct.Expiration
		p.Is_published = patchProduct.Is_published
		p.Name = patchProduct.Name
		p.Price = patchProduct.Price
		p.Quantity = patchProduct.Quantity
		return repo.products[id-1], nil
	}

}

func (repo *sliceRepository) DeleteProduct(id int) error {
	isExist := false
	var index int
	for i, v := range repo.products {
		if v.Id == id {
			index = i
			isExist = true
		}
	}

	if !isExist {
		return errors.New("producto no encontrado")
	}
	repo.products = append(repo.products[:index], repo.products[index+1:]...)
	return nil

}

func codeValueRepeated(p domain.Producto, products []domain.Producto) (bool, error) {
	for _, v := range products {
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
