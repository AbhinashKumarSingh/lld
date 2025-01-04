package coupon_discount

type ShoppingCart struct {
	ProductList []Product
}

func (s *ShoppingCart) AddToCartPercent(product Product) {
	percentDiscount := PercentCouponDecorator{Product: product}
	percentDiscount.GetDiscountPrice(10)
	s.ProductList = append(s.ProductList, percentDiscount.Product)

}

func (s *ShoppingCart) AddToCartValue(product Product) {
	valueDiscount := ValueCouponDecorator{Product: product}
	valueDiscount.GetDiscountPrice(5)
	s.ProductList = append(s.ProductList, valueDiscount.Product)

}
func (s *ShoppingCart) GetTotalPrice() float64 {
	totalPrice := 0.0
	for _, v := range s.ProductList {
		totalPrice += v.Price
	}
	return totalPrice
}
