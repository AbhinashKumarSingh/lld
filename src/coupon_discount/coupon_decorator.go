package coupon_discount

type CouponDecorator interface {
	Product
	GetDiscountPrice() float64
}

type PercentCouponDecorator struct {
	Product
}

func (c *PercentCouponDecorator) GetDiscountPrice(percent float64) {
	c.Price = c.GetPrice() - (c.GetPrice()*percent)/100
}

type ValueCouponDecorator struct {
	Product
}

func (v ValueCouponDecorator) GetDiscountPrice(value float64) {
	v.Price = (v.GetPrice() - value)
}
