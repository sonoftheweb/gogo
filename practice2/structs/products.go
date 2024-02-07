package structs

type Product struct {
	id    int
	title string
	price float64
}

type Products []Product

func New(id int, title string, price float64) Product {
	return Product{id, title, price}
}

func NewProducts() *Products {
	return &Products{}
}

func (products *Products) Add(product Product) {
	*products = append(*products, product)
}
