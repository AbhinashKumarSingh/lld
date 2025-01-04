package coupon_discount

type Product struct {
	Name  string
	Price float64
	Type  string
}

func AddProduct(name string, price float64, typeof string) *Product {
	return &Product{Name: name, Price: price, Type: typeof}
}

func (p *Product) GetPrice() float64 { return p.Price }
