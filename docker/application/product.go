package application

import "errors"

// vamos setar uma interface especificando as caracteristicas do produto
type ProductInterface interface {
	IsValid() (bool, error) // um método que retorna um booleano ou um erro
	Enable() error // se o erro retorna vazio, tá tudo bem
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	// essas constantes vão setar o estado do produto
	DISABLED = "disabled"
	ENABLED = "enabled"
)

// o struct poderia ser uma analogia as classes do POO,
// porém ela já implementa uma interface implicitamente
type Product struct {
	Id string 
	Name string
	Price float64
	Status string
}

// func (object *Product) IsValid() (bool, error) {
// }

// Enable() com "E" maiusculo pode ser acessado externamente (não é um metodo privado)
// enable() com "e" minusculo não pode ser acessado externamente (é um metodo privado)
func (object *Product) Enable() error {
	// para um produto ser habilitado ele precisa ter um preço maior que zero
	if object.Price > 0 {
		object.Status = ENABLED
		return nil
	}
	return errors.New("The price must be greater then zero to enable the product")
}

// func (object *Product) Disable() error {
// }

func (object *Product) GetId() string {
	return object.Id
	
}

func (object *Product) GetName() string {
	return object.Name
}

func (object *Product) GetStatus() string {
	return object.Status
}

func (object *Product) GetPrice() float64 {
	return object.Price
}