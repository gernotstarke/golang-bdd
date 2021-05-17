package basket

func NewBasket() *Basket {
	return &Basket{
		products: make(map[string]float64),
	}
}

type Basket struct {
	products map[string]float64
}

func (b *Basket) AddItem(productName string, price float64) {
	b.products[productName] = price
}